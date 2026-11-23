import { requireApiData, tokenFrom } from '$lib/server/api';
import type { Article } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, fetch, cookies }) => ({
  article: await requireApiData<Article>(fetch, tokenFrom(cookies), `/admin/articles/${params.id}`)
});
