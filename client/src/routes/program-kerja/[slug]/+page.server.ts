import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import type { WorkProgram } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
	try {
		console.log('slug:', params.slug);

		const response = await apiRequest<WorkProgram>(
			fetch,
			undefined,
			`/cabinet/programs/${params.slug}`
		);

		console.log('program response:', response);
		console.log('program data:', response.data);

		return {
			program: response.data
		};
	} catch (err) {
		console.error('failed to fetch program:', err);
		error(404, 'Program kerja tidak ditemukan');
	}
};