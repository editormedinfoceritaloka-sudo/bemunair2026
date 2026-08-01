import { env } from '$env/dynamic/private';
import type { ApiEnvelope, ApiErrorBody } from '$lib/types';
import type { Cookies, RequestEvent } from '@sveltejs/kit';

export const SESSION_COOKIE = 'bemunair_admin_session';
const API_BASE = (env.API_INTERNAL_URL || 'http://server:8080/api/v1').replace(/\/$/, '');

export class ApiError extends Error {
  constructor(message: string, public status: number, public code = 'API_ERROR', public details?: unknown) { super(message); }
}

export async function apiRequest<T>(fetcher: typeof fetch, token: string | undefined, path: string, init: RequestInit = {}): Promise<ApiEnvelope<T>> {
  const headers = new Headers(init.headers);
  if (token) headers.set('Authorization', `Bearer ${token}`);
  if (init.body && !(init.body instanceof FormData) && !headers.has('Content-Type')) headers.set('Content-Type', 'application/json');
  headers.set('Accept', 'application/json');
  const request = API_BASE.startsWith("http://server") ? globalThis.fetch : fetcher;
  const response = await request(API_BASE + path, { ...init, headers });
  let payload: ApiEnvelope<T> | undefined;
  try { payload = await response.json(); } catch { /* handled below */ }
  if (!payload) throw new ApiError("Respons API tidak valid (" + response.status + ")", response.status);
  const succeeded = payload.status ?? payload.success ?? false;
  if (!response.ok || !succeeded) {
    const apiError = typeof payload.error === 'object' ? payload.error as ApiErrorBody : undefined;
    const fallback = typeof payload.error === 'string' ? payload.error : undefined;
    throw new ApiError(payload.message || fallback || `Request gagal (${response.status})`, response.status, apiError?.code || fallback || 'API_ERROR', apiError?.details);
  }
  return payload;
}

export function tokenFrom(cookies: Cookies) { return cookies.get(SESSION_COOKIE); }
export function clearSession(cookies: Cookies) { cookies.delete(SESSION_COOKIE, { path: '/' }); }
export function setSession(cookies: Cookies, token: string, secure: boolean) {
  cookies.set(SESSION_COOKIE, token, { path: '/', httpOnly: true, sameSite: 'lax', secure, maxAge: 60 * 60 * 24 });
}
export function formValue(data: FormData, name: string) { return String(data.get(name) || '').trim(); }
export function nullable(data: FormData, name: string) { const value = formValue(data, name); return value || null; }
export function actionError(error: unknown) {
  if (error instanceof ApiError) return { message: error.message, code: error.code, status: error.status };
  return { message: error instanceof Error ? error.message : 'Terjadi kesalahan', code: 'UNKNOWN', status: 500 };
}
export type AdminEvent = RequestEvent & { locals: App.Locals };
