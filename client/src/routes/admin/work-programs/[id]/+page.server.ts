import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { MediaAsset, WorkProgram } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies, params }) => ({ program: (await apiRequest<WorkProgram>(fetch, tokenFrom(cookies), `/admin/work-programs/${params.id}`)).data });

export const actions: Actions = {
  milestone: async ({ request, fetch, cookies, params }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), "/admin/work-programs/" + params.id + "/milestones", {
        method: "POST",
        body: JSON.stringify({
          title: formValue(form, "title"),
          description: nullable(form, "description"),
          status: formValue(form, "status"),
          display_order: Number(formValue(form, "display_order") || 0)
        })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },
  documentation: async ({ request, fetch, cookies, params }) => {
    const form = await request.formData();
    try {
      const media = await apiRequest<MediaAsset>(fetch, tokenFrom(cookies), "/admin/media-assets", {
        method: "POST",
        body: JSON.stringify({
          file_id: formValue(form, "file_id"),
          file_path: nullable(form, "file_path"),
          url: formValue(form, "url"),
          thumbnail_url: nullable(form, "thumbnail_url"),
          name: formValue(form, "name"),
          alt_text: formValue(form, "alt_text"),
          caption: nullable(form, "caption"),
          mime_type: formValue(form, "mime_type"),
          size_bytes: Number(formValue(form, "size_bytes") || 0),
          width: formValue(form, "width") ? Number(formValue(form, "width")) : null,
          height: formValue(form, "height") ? Number(formValue(form, "height")) : null,
          purpose: "work_program_documentation"
        })
      });
      await apiRequest(fetch, tokenFrom(cookies), "/admin/work-programs/" + params.id + "/documentations", {
        method: "POST",
        body: JSON.stringify({
          media_asset_id: media.data.id,
          title: nullable(form, "title"),
          caption: nullable(form, "caption"),
          display_order: Number(formValue(form, "display_order") || 0),
          is_cover: form.has("is_cover")
        })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};