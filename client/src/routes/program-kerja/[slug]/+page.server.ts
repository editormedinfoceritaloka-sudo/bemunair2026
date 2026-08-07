 import { error } from '@sveltejs/kit';
 import { apiRequest } from '$lib/server/api';
 import type { WorkProgram } from '$lib/types';
 import type { PageServerLoad } from './$types';

  export const load: PageServerLoad = async ({ fetch, params }) => {
    try { return { program: (await apiRequest<WorkProgram>(fetch, undefined, `/cabinet/programs/${params.slug}`)).data }; } catch { error(404, 'Program kerja tidak ditemukan'); }
  };
