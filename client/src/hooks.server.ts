import { redirect, type Handle } from '@sveltejs/kit';
import { SESSION_COOKIE } from '$lib/server/api';

export const handle: Handle = async ({ event, resolve }) => {
  const path = event.url.pathname;
  if (path.startsWith('/admin') && path !== '/admin/login' && !event.cookies.get(SESSION_COOKIE)) {
    throw redirect(303, `/admin/login?next=${encodeURIComponent(path)}`);
  }
  return resolve(event);
};
