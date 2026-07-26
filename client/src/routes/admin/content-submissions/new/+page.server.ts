import { actionError, apiRequest, tokenFrom } from '$lib/server/api';
import type { ContentSubmission } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
  default: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    let id: number;
    try {
      const result = await apiRequest<ContentSubmission>(fetch, tokenFrom(cookies), '/content-submissions', { method: 'POST', body: form });
      id = result.data.id;
    } catch (error) {
      const x = actionError(error);
      return fail(x.status, { error: x.message });
    }
    redirect(303, `/admin/content-submissions/${id}`);
  }
};
