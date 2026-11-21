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
          start_date: nullable(form, "start_date"),
          end_date: nullable(form, "end_date"),
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
  },
  update: async ({ request, fetch, cookies, params }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), "/admin/work-programs/" + params.id, {
        method: "PUT",
        body: JSON.stringify({
          ministry_id: Number(formValue(form, "ministry_id")),
          name: formValue(form, "name"),
          short_description: nullable(form, "short_description"),
          description: nullable(form, "description"),
          objectives: nullable(form, "objectives"),
          target_audience: nullable(form, "target_audience"),
          start_date: nullable(form, "start_date"),
          end_date: nullable(form, "end_date"),
          status: formValue(form, "status"),
          cover_media_id: formValue(form, "cover_media_id") ? Number(formValue(form, "cover_media_id")) : null,
          display_order: Number(formValue(form, "display_order") || 0),
          is_featured: form.has("is_featured"),
          is_published: form.has("is_published")
        })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
