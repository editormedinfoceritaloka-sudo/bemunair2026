import { actionError, apiRequest, formValue, tokenFrom } from '$lib/server/api';
import type { LetterSubmission, LetterTemplate } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => ({
  templates: (await apiRequest<LetterTemplate[]>(fetch, tokenFrom(cookies), '/letter-templates')).data
});

export const actions: Actions = {
  default: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    const deadline = formValue(form, 'deadline');
    const body = {
      ministry: formValue(form, 'ministry'),
      letter_type: formValue(form, 'letter_type'),
      subject: formValue(form, 'subject'),
      body: formValue(form, 'body'),
      deadline: `${deadline}:00+07:00`
    };
    let id: number;
    try {
      const result = await apiRequest<LetterSubmission>(fetch, tokenFrom(cookies), '/letter-submissions', { method: 'POST', body: JSON.stringify(body) });
      id = result.data.id;
    } catch (error) {
      const x = actionError(error);
      return fail(x.status, { error: x.message });
    }
    redirect(303, `/admin/letter-submissions/${id}`);
  }
};
