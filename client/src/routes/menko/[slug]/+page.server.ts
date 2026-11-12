import { error } from '@sveltejs/kit';
import { apiRequest } from '$lib/server/api';
import type { OrganizationUnit } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, params }) => {
  try {
    const response = await apiRequest<OrganizationUnit>(
      fetch,
      undefined,
      `/cabinet/units/${params.slug}`
    );

    return {
      unit: response.data
    };
  } catch (err) {


    error(404, 'Kemenkoan atau kementerian tidak ditemukan');
  }
};