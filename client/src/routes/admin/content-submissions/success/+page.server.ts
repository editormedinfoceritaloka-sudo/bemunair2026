import { apiRequest, tokenFrom } from '$lib/server/api';
import type { ContentSubmission } from '$lib/types';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ url, fetch, cookies }) => {
  const id = Number(url.searchParams.get('id'));
  if (!Number.isInteger(id) || id < 1) throw error(404, 'Pengajuan tidak ditemukan');
  const { data: submission } = await apiRequest<ContentSubmission>(
    fetch, tokenFrom(cookies), `/content-submissions/${id}`
  );
  return { submission };
};
