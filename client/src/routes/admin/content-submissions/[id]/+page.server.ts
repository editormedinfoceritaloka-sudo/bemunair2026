import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { ContentSubmission } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, fetch, cookies }) => ({
  submission: (await apiRequest<ContentSubmission>(fetch, tokenFrom(cookies), `/content-submissions/${params.id}`)).data
});

export const actions: Actions = {
  status: async ({ params, request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}/status`, {
        method: 'PUT', body: JSON.stringify({ status: formValue(form, 'status'), notes: nullable(form, 'notes') })
      });
      return { success: true };
    } catch (error) { const x = actionError(error); return fail(x.status, { error: x.message }); }
  },
  delete: async ({ params, fetch, cookies }) => {
    try { await apiRequest(fetch, tokenFrom(cookies), `/content-submissions/${params.id}`, { method: 'DELETE' }); }
    catch (error) { const x = actionError(error); return fail(x.status, { error: x.message }); }
    redirect(303, '/admin/content-submissions');
  }
};
