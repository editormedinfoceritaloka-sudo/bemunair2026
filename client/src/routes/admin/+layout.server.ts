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
    if (!['ADMIN', 'ADMIN_MEDINFO'].includes(user.role)) throw new Error('Role tidak diizinkan');
    const medinfoOnly = ['/admin/users', '/admin/medinfo-queue', '/admin/articles', '/admin/system', '/admin/ministries', '/admin/settings', '/admin/letter-templates'];
    const adminOnly = ['/admin/letter-submissions'];
    if (user.role === 'ADMIN_MEDINFO' && adminOnly.some((path) => url.pathname.startsWith(path))) throw redirect(303, '/admin');
    if (user.role === 'ADMIN' && medinfoOnly.some((path) => url.pathname.startsWith(path))) throw redirect(303, '/admin');
    if (user.role === 'ADMIN' && url.pathname.startsWith('/admin/letter-submissions/')) throw redirect(303, '/admin/letter-submissions');
    return { user };
  } catch (error) {
    if (typeof error === 'object' && error && 'status' in error && Number(error.status) >= 300 && Number(error.status) < 400) throw error;
    clearSession(cookies);
    throw redirect(303, '/admin/login?expired=1');
  }
};
