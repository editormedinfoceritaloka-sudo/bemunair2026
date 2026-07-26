import { actionError, apiRequest, clearSession, formValue, setSession } from '$lib/server/api';
import type { LoginResult, User } from '$lib/types';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ parent }) => {
  const data = await parent();
  if (data.user?.role === 'ADMIN') throw redirect(303, '/admin');
  return {};
};

export const actions: Actions = {
  default: async ({ request, fetch, cookies, url }) => {
    const form = await request.formData();
    const email = formValue(form, 'email');
    const password = formValue(form, 'password');
    if (!email || !password) return fail(422, { error: 'Email dan password wajib diisi', email });
    try {
      const { data } = await apiRequest<LoginResult>(fetch, undefined, '/auth/login', { method: 'POST', body: JSON.stringify({ email, password }) });
      if (data.user.role !== 'ADMIN') { clearSession(cookies); return fail(403, { error: 'Akun ini tidak memiliki akses admin', email }); }
      setSession(cookies, data.token, url.protocol === 'https:');
      await apiRequest<User>(fetch, data.token, '/auth/me');
    } catch (error) {
      return fail(actionError(error).status, { error: actionError(error).message, email });
    }
    throw redirect(303, '/admin');
  }
};
