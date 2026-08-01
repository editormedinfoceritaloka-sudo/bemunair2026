import { apiRequest, tokenFrom } from '$lib/server/api';
import type { LetterTemplate } from '$lib/types';
import type { PageServerLoad } from './$types';
export const load: PageServerLoad = async ({ fetch, cookies }) => ({ templates: (await apiRequest<LetterTemplate[]>(fetch, tokenFrom(cookies), '/letter-templates')).data });
