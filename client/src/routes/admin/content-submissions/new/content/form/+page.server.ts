import { actionError, apiRequest, tokenFrom } from '$lib/server/api';
import type { ContentSubmission, MediaSubmissionSetting } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies, parent }) => {
  const user = (await parent()).user!;
  const { data: setting } = await apiRequest<MediaSubmissionSetting>(
    fetch, tokenFrom(cookies), '/media-submission-settings/CONTENT'
  );
  return { user, setting };
};

export const actions: Actions = {
  default: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    let submission: ContentSubmission;
    try {
      submission = (await apiRequest<ContentSubmission>(
        fetch, tokenFrom(cookies), '/content-submissions', { method: 'POST', body: form }
      )).data;
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
    redirect(303, `/admin/content-submissions/success?id=${submission.id}`);
  }
};
