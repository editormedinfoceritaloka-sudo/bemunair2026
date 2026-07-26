import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { MediaSubmissionSetting } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => {
  const token = tokenFrom(cookies);
  const [content, article] = await Promise.all([
    apiRequest<MediaSubmissionSetting>(fetch, token, '/media-submission-settings/CONTENT'),
    apiRequest<MediaSubmissionSetting>(fetch, token, '/media-submission-settings/ARTICLE')
  ]);
  return { content: content.data, article: article.data };
};

export const actions: Actions = {
  save: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    const serviceType = formValue(form, 'service_type');
    const terms = formValue(form, 'terms').split('\n').map((item) => item.trim()).filter(Boolean);
    const payload = {
      sop_url: nullable(form, 'sop_url'),
      ministry_template_url: nullable(form, 'ministry_template_url'),
      brief_template_url: nullable(form, 'brief_template_url'),
      caption_template_url: nullable(form, 'caption_template_url'),
      pic_name: nullable(form, 'pic_name'),
      pic_whatsapp: nullable(form, 'pic_whatsapp'),
      terms,
      minimum_lead_days: Number(formValue(form, 'minimum_lead_days')),
      publish_time_start: formValue(form, 'publish_time_start'),
      publish_time_end: formValue(form, 'publish_time_end'),
      slot_interval_minutes: Number(formValue(form, 'slot_interval_minutes')),
      daily_capacity: nullable(form, 'daily_capacity') ? Number(formValue(form, 'daily_capacity')) : null
    };
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/media-submission-settings/${serviceType}`, {
        method: 'PUT', body: JSON.stringify(payload)
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
