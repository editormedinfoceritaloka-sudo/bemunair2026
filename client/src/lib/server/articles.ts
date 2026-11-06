import { apiRequest } from '$lib/server/api';

import type {
  Article,
  ArticleListItem,
  Meta
} from '$lib/types';

type FetchLike = typeof globalThis.fetch;

export type ArticleListResult = {
  articles: ArticleListItem[];
  meta: Meta;
};

function fallbackMeta(
  page: number,
  perPage: number,
  count: number
): Meta {
  return {
    page,
    per_page: perPage,
    total: count,
    total_pages: count > 0 ? 1 : 0
  };
}

async function fetchPublishedArticlePage(
  fetch: FetchLike,
  page: number,
  perPage: number
): Promise<ArticleListResult> {
  const params = new URLSearchParams({
    page: String(page),
    per_page: String(perPage)
  });

  const response = await apiRequest<ArticleListItem[]>(
    fetch,
    undefined,
    `/articles?${params.toString()}`
  );

  return {
    articles: response.data,
    meta:
      response.meta ??
      fallbackMeta(page, perPage, response.data.length)
  };
}

const MAX_FANOUT_PAGES = 5;

async function fetchAllPublishedArticles(
  fetch: FetchLike
): Promise<ArticleListItem[]> {
  const perPage = 50;

  const first = await fetchPublishedArticlePage(fetch, 1, perPage);

  const totalPages = Math.min(first.meta.total_pages, MAX_FANOUT_PAGES);

  if (totalPages <= 1) {
    return first.articles;
  }

  const requests = Array.from(
    { length: totalPages - 1 },
    (_, index) => fetchPublishedArticlePage(fetch, index + 2, perPage)
  );

  const rest = await Promise.all(requests);

  return [...first.articles, ...rest.flatMap((result) => result.articles)];
}

function matchesSearch(
  article: ArticleListItem,
  search: string
): boolean {
  const query = search.trim().toLocaleLowerCase('id-ID');

  if (!query) {
    return true;
  }

  const source = [
    article.title,
    article.excerpt ?? '',
    article.author?.name ?? ''
  ]
    .join(' ')
    .toLocaleLowerCase('id-ID');

  return source.includes(query);
}

export async function getPublishedArticles(
  fetch: FetchLike,
  options: {
    page: number;
    perPage: number;
    search?: string;
  }
): Promise<ArticleListResult> {
  const search = options.search?.trim() ?? '';

  if (!search) {
    return fetchPublishedArticlePage(fetch, options.page, options.perPage);
  }

  const allArticles = await fetchAllPublishedArticles(fetch);

  const filtered = allArticles.filter((article) =>
    matchesSearch(article, search)
  );

  const total = filtered.length;

  const totalPages = total > 0 ? Math.ceil(total / options.perPage) : 0;

  const offset = (options.page - 1) * options.perPage;

  return {
    articles: filtered.slice(offset, offset + options.perPage),
    meta: {
      page: options.page,
      per_page: options.perPage,
      total,
      total_pages: totalPages
    }
  };
}

export async function getPublishedArticleBySlug(
  fetch: FetchLike,
  slug: string
): Promise<Article> {
  const response = await apiRequest<Article>(
    fetch,
    undefined,
    `/articles/${encodeURIComponent(slug)}`
  );

  return response.data;
}
