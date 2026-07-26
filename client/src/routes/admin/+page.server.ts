import { apiRequest, tokenFrom } from '$lib/server/api';
import type { Article, ContentSubmission, LetterSubmission, LetterTemplate, QueueItem, User } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => {
  const token = tokenFrom(cookies)!;
  const settled = await Promise.allSettled([
    apiRequest<User[]>(fetch, token, '/users'), apiRequest<ContentSubmission[]>(fetch, token, '/content-submissions'),
    apiRequest<LetterSubmission[]>(fetch, token, '/letter-submissions'), apiRequest<QueueItem[]>(fetch, token, '/medinfo-pj/queue'),
    apiRequest<LetterTemplate[]>(fetch, token, '/letter-templates'), apiRequest<Article[]>(fetch, token, '/admin/articles?page=1&per_page=50')
  ]);
  const value = <T>(index: number, fallback: T): T => settled[index].status === 'fulfilled' ? (settled[index] as PromiseFulfilledResult<any>).value.data : fallback;
  const users = value<User[]>(0, []), content = value<ContentSubmission[]>(1, []), letters = value<LetterSubmission[]>(2, []), queue = value<QueueItem[]>(3, []), templates = value<LetterTemplate[]>(4, []), articles = value<Article[]>(5, []);
  return { users, content, letters, queue, templates, articles, partialFailure: settled.some((item) => item.status === 'rejected') };
};
