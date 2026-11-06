import { redirect } from '@sveltejs/kit';

import { getPublishedArticles } from '$lib/server/articles';

import type { PageServerLoad } from './$types';

function positiveInteger(
  value: string | null,
  fallback: number
): number {
  if (!value) {
    return fallback;
  }

  const parsed = Number.parseInt(value, 10);

  if (Number.isNaN(parsed) || parsed < 1) {
    return fallback;
  }

  return parsed;
}

export const load: PageServerLoad = async ({ fetch, url }) => {
  const page = positiveInteger(url.searchParams.get('page'), 1);

  const perPage = Math.min(
    positiveInteger(url.searchParams.get('per_page'), 6),
    50
  );

  const search = url.searchParams.get('search')?.trim() ?? '';

  const result = await getPublishedArticles(fetch, {
    page,
    perPage,
    search
  });

  if (result.meta.total_pages > 0 && page > result.meta.total_pages) {
    const params = new URLSearchParams();

    if (search) {
      params.set('search', search);
    }

    params.set('page', String(result.meta.total_pages));
    params.set('per_page', String(perPage));

    redirect(302, `/news?${params.toString()}`);
  }

  return {
    articles: result.articles,
    meta: result.meta,
    filters: {
      search,
      page,
      per_page: perPage
    },
    seo: {
      title: search
        ? `Pencarian "${search}" | News BEM UNAIR 2026`
        : 'News | BEM Universitas Airlangga 2026',
      description: search
        ? `Hasil pencarian berita "${search}" dari Kabinet Cerita Loka BEM Universitas Airlangga 2026.`
        : 'Berita, kegiatan, program, dan cerita terbaru Kabinet Cerita Loka BEM Universitas Airlangga 2026.',
      image: '/og-image.png'
    }
  };
};
