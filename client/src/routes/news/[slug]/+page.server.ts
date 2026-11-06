import { error } from '@sveltejs/kit';

import { getPublishedArticleBySlug } from '$lib/server/articles';

import type { PageServerLoad } from './$types';

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
  try {
    const article = await getPublishedArticleBySlug(fetch, params.slug);

    return {
      article,
      readTime: calculateReadTime(article.body ?? ''),
      seo: {
        title: `${article.title} | BEM Universitas Airlangga 2026`,
        description:
          article.excerpt ??
          'Berita terbaru Kabinet Cerita Loka BEM Universitas Airlangga 2026.',
        image: article.cover_image ?? '/og-image.png'
      }
    };
  } catch {
    error(404, 'Artikel tidak ditemukan');
  }
};
