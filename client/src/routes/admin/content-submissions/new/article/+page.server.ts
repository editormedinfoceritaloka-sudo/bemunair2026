import { apiRequest, tokenFrom } from '$lib/server/api';
import type { MediaSubmissionSetting } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies, parent }) => {
  const { user } = await parent();
  const { data: setting } = await apiRequest<MediaSubmissionSetting>(
    fetch, tokenFrom(cookies), '/media-submission-settings/ARTICLE'
  );
  return { user, setting };
};
