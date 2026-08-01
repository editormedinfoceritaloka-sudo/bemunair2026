import { apiRequest } from '$lib/server/api';
import type { Cabinet } from '$lib/types';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ fetch }) => {
	try {
		return { cabinet: (await apiRequest<Cabinet>(fetch, undefined, '/cabinet')).data };
	} catch {
		return { cabinet: null };
	}
};
