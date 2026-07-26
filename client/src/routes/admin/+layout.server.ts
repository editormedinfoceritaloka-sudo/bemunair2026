import { apiRequest, clearSession, tokenFrom } from '$lib/server/api';
import type { User } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ url, cookies, fetch }) => {
  if (url.pathname === '/admin/login') return { user: null };
  const token = tokenFrom(cookies);
  if (!token) throw redirect(303, '/admin/login');
  try {
    const { data: user } = await apiRequest<User>(fetch, token, '/auth/me');
    if (user.role !== 'ADMIN') throw new Error('Role bukan ADMIN');
    return { user };
  } catch {
    clearSession(cookies);
    throw redirect(303, '/admin/login?expired=1');
  }
};
