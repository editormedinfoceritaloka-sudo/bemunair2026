import { apiRequest, tokenFrom } from '$lib/server/api';
import type { Article, ContentSubmission, LetterSubmission, LetterTemplate, QueueItem, User } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies, parent }) => {
  const user = (await parent()).user!;
  const token = tokenFrom(cookies)!;
  const common = await Promise.allSettled([
    apiRequest<ContentSubmission[]>(fetch, token, '/content-submissions'),
    apiRequest<LetterSubmission[]>(fetch, token, '/letter-submissions')
  ]);
  const content = common[0].status === 'fulfilled' ? common[0].value.data : [];
  const letters = common[1].status === 'fulfilled' ? common[1].value.data : [];

  if (user.role !== 'ADMIN_MEDINFO') {
    return {
      user, content, letters, users: [] as User[], queue: [] as QueueItem[],
      templates: [] as LetterTemplate[], articles: [] as Article[],
      partialFailure: common.some((item) => item.status === 'rejected')
    };
  }

  const operations = await Promise.allSettled([
    apiRequest<User[]>(fetch, token, '/users'),
    apiRequest<QueueItem[]>(fetch, token, '/medinfo-pj/queue'),
    apiRequest<LetterTemplate[]>(fetch, token, '/letter-templates'),
    apiRequest<Article[]>(fetch, token, '/admin/articles?page=1&per_page=50')
  ]);
  const value = <T>(index: number, fallback: T): T =>
    operations[index].status === 'fulfilled'
      ? (operations[index] as unknown as PromiseFulfilledResult<{ data: T }>).value.data
      : fallback;

  return {
    user, content, letters,
    users: value<User[]>(0, []), queue: value<QueueItem[]>(1, []),
    templates: value<LetterTemplate[]>(2, []), articles: value<Article[]>(3, []),
    partialFailure: [...common, ...operations].some((item) => item.status === 'rejected')
  };
};
