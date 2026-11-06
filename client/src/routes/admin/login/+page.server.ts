import { actionError, apiRequest, clearSession, formValue, setSession } from '$lib/server/api';
import type { LoginResult, User } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Cookies } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

const RATE_LIMIT_MAX = 5;
const RATE_LIMIT_WINDOW_MS = 15 * 60 * 1000;
type AttemptRecord = { count: number; resetAt: number };
const loginAttempts = new Map<string, AttemptRecord>();

function getClientIp(request: Request, cookies: Cookies): string {
  return (
    request.headers.get('x-forwarded-for')?.split(',')[0]?.trim() ||
    request.headers.get('x-real-ip')?.trim() ||
    cookies.get('client_ip')?.toString() ||
    'unknown'
  );
}

function checkRateLimit(ip: string): { allowed: boolean; retryAfter: number } {
  const now = Date.now();
  const record = loginAttempts.get(ip);

  if (!record || record.resetAt < now) {
    loginAttempts.set(ip, { count: 1, resetAt: now + RATE_LIMIT_WINDOW_MS });
    return { allowed: true, retryAfter: 0 };
  }

  if (record.count >= RATE_LIMIT_MAX) {
    return { allowed: false, retryAfter: Math.ceil((record.resetAt - now) / 1000) };
  }

  record.count += 1;
  return { allowed: true, retryAfter: 0 };
}

function clearRateLimit(ip: string): void {
  loginAttempts.delete(ip);
}

export const load: PageServerLoad = async ({ parent }) => {
  const data = await parent();
  if (['ADMIN', 'ADMIN_MEDINFO'].includes(data.user?.role ?? '')) throw redirect(303, '/admin');
  return {};
};

export const actions: Actions = {
  default: async ({ request, fetch, cookies, url }) => {
    const form = await request.formData();
    const email = formValue(form, 'email');
    const password = formValue(form, 'password');
    const ip = getClientIp(request, cookies);
    const limit = checkRateLimit(ip);
    if (!limit.allowed) {
      return fail(429, { error: `Terlalu banyak percobaan login. Coba lagi dalam ${Math.ceil(limit.retryAfter / 60)} menit.`, email });
    }
    if (!email || !password) return fail(422, { error: 'Email dan password wajib diisi', email });
    try {
      const { data } = await apiRequest<LoginResult>(fetch, undefined, '/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) });
      if (!['ADMIN', 'ADMIN_MEDINFO'].includes(data.user.role)) { clearSession(cookies); return fail(403, { error: 'Akun ini tidak memiliki akses admin', email }); }
      clearRateLimit(ip);
      setSession(cookies, data.token, url.protocol === 'https:');
      await apiRequest<User>(fetch, data.token, '/auth/me');
    } catch (error) {
      return fail(actionError(error).status, { error: actionError(error).message, email });
    }
    throw redirect(303, '/admin');
  }
};
