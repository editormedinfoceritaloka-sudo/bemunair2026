import { requireApiData } from '$lib/server/api';

import type { Cabinet } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
  const cabinet = await requireApiData<Cabinet>(
    fetch,
    undefined,
    '/cabinet'
  );

  return {
    cabinet,
    seo: {
      title: 'BEM Universitas Airlangga 2026 | Kabinet Cerita Loka',
      description:
        'Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka. Temukan informasi kabinet, kementerian, program kerja, agenda, dan berita terbaru BEM UNAIR.',
      image: '/og-image.png'
    }
  };
};