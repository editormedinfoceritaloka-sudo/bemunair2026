import { apiRequest, tokenFrom } from '$lib/server/api';
import type { ContentSubmission } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => ({
  submissions: (await apiRequest<ContentSubmission[]>(fetch, tokenFrom(cookies), '/content-submissions')).data
});
