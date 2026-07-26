import { apiRequest, tokenFrom } from '$lib/server/api';
import type { LetterSubmission } from '$lib/types';
import type { PageServerLoad } from './$types';
export const load: PageServerLoad = async ({ fetch, cookies }) => ({ letters: (await apiRequest<LetterSubmission[]>(fetch, tokenFrom(cookies), '/letter-submissions')).data });
