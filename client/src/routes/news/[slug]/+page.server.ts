import { error } from '@sveltejs/kit';

import { ApiError } from '$lib/server/api';
import {
  getPublishedArticleBySlug,
  getRelatedArticles
} from '$lib/server/articles';

import type { PageServerLoad } from './$types';

const RELATED_LIMIT = 3;

function calculateReadTime(html: string): number {
  const text = html
    .replace(/<[^>]*>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();

  if (!text) {
    return 1;
  }

  const words = text.split(' ').length;

  return Math.max(1, Math.ceil(words / 200));
}

export const load: PageServerLoad = async ({ fetch, params }) => {
  const [articleResult, relatedResult] = await Promise.allSettled([
    getPublishedArticleBySlug(fetch, params.slug),
    getRelatedArticles(fetch, params.slug, RELATED_LIMIT)
  ]);

  if (articleResult.status === 'rejected') {
    const reason = articleResult.reason;

    // Hanya 404 bila API memang tidak menemukan artikel; gangguan server
    // tidak boleh tampil sebagai "tidak ditemukan".
    if (reason instanceof ApiError && reason.status >= 500) {
      error(503, 'Berita belum dapat dimuat. Coba lagi beberapa saat.');
    }

    error(404, 'Artikel tidak ditemukan');
  }

  const article = articleResult.value;

  return {
    article,
    readTime: calculateReadTime(article.body ?? ''),
    // Berita lainnya bersifat pelengkap: kegagalan tidak menggagalkan halaman.
    related:
      relatedResult.status === 'fulfilled' ? relatedResult.value : [],
    seo: {
      title: `${article.title} | BEM Universitas Airlangga 2026`,
      description:
        article.excerpt ??
        'Berita terbaru Kabinet Cerita Loka BEM Universitas Airlangga 2026.',
      image: article.cover_image ?? '/og-image.png'
    }
  };
};
