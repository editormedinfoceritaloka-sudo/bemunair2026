import { requireApiData } from '$lib/server/api';

import type {
  ArticleListItem,
  Cabinet
} from '$lib/types';

import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
  const [cabinet, articles] = await Promise.all([
    requireApiData<Cabinet>(
      fetch,
      undefined,
      '/cabinet'
    ),

    requireApiData<ArticleListItem[]>(
      fetch,
      undefined,
      '/articles?page=1&per_page=4'
    )
  ]);

  return {
    cabinet,
    articles,
    seo: {
      title: 'BEM Universitas Airlangga 2026 | Kabinet Cerita Loka',
      description:
        'Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka. Temukan informasi kabinet, kementerian, program kerja, agenda, dan berita terbaru BEM UNAIR.',
      image: '/og-image.png'
    }
  };
};