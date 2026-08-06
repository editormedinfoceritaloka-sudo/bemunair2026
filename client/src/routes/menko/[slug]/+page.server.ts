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

    console.log('Slug:', params.slug);
    console.log('Organization unit:', response.data);
    console.dir(response.data, {
      depth: null,
      colors: true
    });

    return {
      unit: response.data
    };
  } catch (err) {
    console.error(
      `Gagal mengambil unit dengan slug "${params.slug}":`,
      err
    );

    error(404, 'Kemenkoan atau kementerian tidak ditemukan');
  }
};