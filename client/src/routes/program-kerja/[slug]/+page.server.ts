import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import type { WorkProgram } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
	try {
		const response = await apiRequest<WorkProgram>(
			fetch,
			undefined,
			`/cabinet/programs/${params.slug}`
		);

		return {
			program: response.data
		};
	} catch {
		error(404, 'Program kerja tidak ditemukan');
	}
};