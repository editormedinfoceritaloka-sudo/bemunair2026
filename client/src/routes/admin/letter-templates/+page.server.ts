import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { LetterTemplate, MediaAsset } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => ({
  templates: (await apiRequest<LetterTemplate[]>(fetch, tokenFrom(cookies), '/letter-templates')).data
});

export const actions: Actions = {
  create: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      const media = await apiRequest<MediaAsset>(fetch, tokenFrom(cookies), '/admin/media-assets', {
        method: 'POST',
        body: JSON.stringify({
          file_id: formValue(form, 'file_id'),
          file_path: nullable(form, 'file_path'),
          url: formValue(form, 'url'),
          name: formValue(form, 'file_name'),
          alt_text: formValue(form, 'alt_text'),
          mime_type: formValue(form, 'mime_type'),
          size_bytes: Number(formValue(form, 'size_bytes') || 0),
          purpose: 'letter_template'
        })
      });
      await apiRequest<LetterTemplate>(fetch, tokenFrom(cookies), '/letter-templates', {
        method: 'POST',
        body: JSON.stringify({
          name: formValue(form, 'name'),
          type: formValue(form, 'type'),
          subject: nullable(form, 'subject') || '',
          body: '',
          media_asset_id: media.data.id,
          is_active: true,
          display_order: Number(formValue(form, 'display_order') || 0)
        })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },
  delete: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), '/letter-templates/' + formValue(form, 'id'), { method: 'DELETE' });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
