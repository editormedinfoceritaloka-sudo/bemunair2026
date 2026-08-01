import { apiRequest } from '$lib/server/api';
import type { Cabinet } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => ({ cabinet: (await apiRequest<Cabinet>(fetch, undefined, '/cabinet')).data });
