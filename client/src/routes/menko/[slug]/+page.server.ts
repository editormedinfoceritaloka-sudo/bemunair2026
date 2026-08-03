// import { error } from '@sveltejs/kit';
// import { apiRequest } from '$lib/server/api';
// import type { OrganizationUnit } from '$lib/types';
// import type { PageServerLoad } from './$types';

// export const load: PageServerLoad = async ({ fetch, params }) => {
//   try { return { unit: (await apiRequest<OrganizationUnit>(fetch, undefined, `/cabinet/units/${params.slug}`)).data }; } catch { error(404, 'Kemenkoan tidak ditemukan'); }
// };
