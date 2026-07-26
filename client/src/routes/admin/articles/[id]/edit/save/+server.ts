import { actionError, apiRequest, tokenFrom } from '$lib/server/api';
import type { Article } from '$lib/types';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const PUT: RequestHandler = async ({ params, request, fetch, cookies }) => {
  try {
    const body = await request.json();
    const result = await apiRequest<Article>(fetch, tokenFrom(cookies), `/admin/articles/${params.id}`, { method: 'PUT', body: JSON.stringify(body) });
    return json(result.data);
  } catch (error) {
    const x = actionError(error);
    return json({ message: x.message }, { status: x.status });
  }
};

export const POST: RequestHandler = async ({ params, fetch, cookies }) => {
  try {
    const current = (await apiRequest<Article>(fetch, tokenFrom(cookies), `/admin/articles/${params.id}`)).data;
    const result = await apiRequest<Article>(fetch, tokenFrom(cookies), `/admin/articles/${params.id}/publish`, {
      method: 'PUT', body: JSON.stringify({ published: current.status !== 'PUBLISHED' })
    });
    return json(result.data);
  } catch (error) {
    const x = actionError(error);
    return json({ message: x.message }, { status: x.status });
  }
};
