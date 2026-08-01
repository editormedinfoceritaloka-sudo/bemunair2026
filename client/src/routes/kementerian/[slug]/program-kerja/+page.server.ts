import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import { pageFromSearch } from '$lib/hooks/use-pagination.svelte';
import type { OrganizationUnit, WorkProgram, Meta } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params, url }) => {
  try {
    const unit = (await apiRequest<OrganizationUnit>(fetch, undefined, `/cabinet/units/${params.slug}`)).data;
    const result = await apiRequest<WorkProgram[]>(fetch, undefined, `/cabinet/units/${params.slug}/programs?page=${pageFromSearch(url.searchParams)}`);
    return { unit, programs: result.data, meta: result.meta as Meta };
  } catch { error(404, 'Kementerian tidak ditemukan'); }
};
