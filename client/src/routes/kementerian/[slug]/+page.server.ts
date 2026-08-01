import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import type { OrganizationUnit, WorkProgram } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
  try {
    const unit = (await apiRequest<OrganizationUnit>(fetch, undefined, `/cabinet/units/${params.slug}`)).data;
    const programs = (await apiRequest<WorkProgram[]>(fetch, undefined, `/cabinet/units/${params.slug}/programs?page=1`)).data;
    return { unit, programs };
  } catch { error(404, 'Kementerian tidak ditemukan'); }
};
