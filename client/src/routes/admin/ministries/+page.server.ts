import { actionError, apiRequest, formValue, tokenFrom } from '$lib/server/api';
import type { Ministry } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => ({
  ministries: (await apiRequest<Ministry[]>(fetch, tokenFrom(cookies), '/ministries')).data
});

export const actions: Actions = {
  create: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), '/ministries', {
        method: 'POST',
        body: JSON.stringify({ code: formValue(form, 'code'), name: formValue(form, 'name'), is_active: true })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },
  toggle: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/ministries/${formValue(form, 'id')}`, {
        method: 'PUT', body: JSON.stringify({ is_active: formValue(form, 'active') !== 'true' })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
