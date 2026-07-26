import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { ContentSubmission, SubmissionHistory, QueueItem } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, fetch, cookies, parent }) => {
  const user = (await parent()).user!;
  const token = tokenFrom(cookies);
  const [submission, timeline, queue] = await Promise.all([
    apiRequest<ContentSubmission>(fetch, token, `/content-submissions/${params.id}`),
    apiRequest<SubmissionHistory[]>(fetch, token, `/content-submissions/${params.id}/timeline`),
    user.role === 'ADMIN_MEDINFO'
      ? apiRequest<QueueItem[]>(fetch, token, '/medinfo-pj/queue')
      : Promise.resolve({ data: [] as QueueItem[] })
  ]);
  return { submission: submission.data, timeline: timeline.data, queue: queue.data, user };
};

export const actions: Actions = {
  assign: async ({ params, request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}/assignee`, {
        method: 'PUT',
        body: JSON.stringify({ assigned_pj_id: Number(formValue(form, 'assigned_pj_id')) })
      });
      return { success: true };
    } catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); }
  },
  status: async ({ params, request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}/status`, {
        method: 'PUT',
        body: JSON.stringify({ status: formValue(form, 'status'), notes: nullable(form, 'notes') })
      });
      return { success: true };
    } catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); }
  },
  revision: async ({ params, request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}/revision`, {
        method: 'POST', body: JSON.stringify({ notes: nullable(form, 'notes') })
      });
      return { success: true };
    } catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); }
  },
  articleDraft: async ({ params, fetch, cookies }) => {
    let articleID: number;
    try {
      const result = await apiRequest<{ id: number }>(fetch, tokenFrom(cookies), `/content-submissions/${params.id}/article-draft`, { method: 'POST' });
      articleID = result.data.id;
    } catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); }
    redirect(303, `/admin/articles/${articleID}/edit`);
  },
  delete: async ({ params, fetch, cookies }) => {
    try { await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}`, { method: 'DELETE' }); }
    catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); }
    redirect(303, '/admin/content-submissions');
  }
};
