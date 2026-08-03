import { env } from '$env/dynamic/private';
import type { ApiEnvelope, ApiErrorBody } from '$lib/types';
import {
  error as kitError,
  type Cookies,
  type RequestEvent
} from '@sveltejs/kit';

export const SESSION_COOKIE = 'bemunair_admin_session';

const API_BASE = (
  env.API_INTERNAL_URL ||
  'http://127.0.0.1:8080/api/v1'
).replace(/\/$/, '');

export class ApiError extends Error {
  constructor(
    message: string,
    public status: number,
    public code = 'API_ERROR',
    public details?: unknown,
    public cause?: unknown
  ) {
    super(message);

    this.name = 'ApiError';
  }
}

export interface NormalizedApiError {
  message: string;
  code: string;
  status: number;
  details?: unknown;
}

export type SafeApiResult<T> =
  | {
      ok: true;
      data: T;
      error: null;
    }
  | {
      ok: false;
      data: null;
      error: NormalizedApiError;
    };

function normalizePath(path: string): string {
  return path.startsWith('/') ? path : `/${path}`;
}

function getRequestHandler(fetcher: typeof fetch): typeof fetch {
  return API_BASE.startsWith('http://server')
    ? globalThis.fetch
    : fetcher;
}

function isApiEnvelope<T>(value: unknown): value is ApiEnvelope<T> {
  return typeof value === 'object' && value !== null;
}

async function parseApiResponse<T>(
  response: Response
): Promise<ApiEnvelope<T>> {
  const contentType = response.headers.get('content-type') ?? '';
  const rawBody = await response.text();

  if (!rawBody.trim()) {
    throw new ApiError(
      `Respons API kosong (${response.status})`,
      response.ok ? 502 : response.status,
      'EMPTY_API_RESPONSE'
    );
  }

  if (!contentType.includes('application/json')) {
    throw new ApiError(
      `Respons API bukan JSON (${response.status})`,
      response.ok ? 502 : response.status,
      'INVALID_CONTENT_TYPE',
      {
        contentType,
        bodyPreview: rawBody.slice(0, 300)
      }
    );
  }

  let payload: unknown;

  try {
    payload = JSON.parse(rawBody);
  } catch (cause) {
    throw new ApiError(
      `Respons API tidak valid (${response.status})`,
      response.ok ? 502 : response.status,
      'INVALID_JSON_RESPONSE',
      {
        bodyPreview: rawBody.slice(0, 300)
      },
      cause
    );
  }

  if (!isApiEnvelope<T>(payload)) {
    throw new ApiError(
      `Format respons API tidak sesuai (${response.status})`,
      response.ok ? 502 : response.status,
      'INVALID_API_ENVELOPE',
      payload
    );
  }

  return payload;
}

export function normalizeApiError(
  error: unknown
): NormalizedApiError {
  if (error instanceof ApiError) {
    return {
      message: error.message,
      code: error.code,
      status: error.status,
      details: error.details
    };
  }

  if (error instanceof Error) {
    return {
      message: error.message || 'Terjadi kesalahan',
      code: 'UNKNOWN_ERROR',
      status: 500
    };
  }

  return {
    message: 'Terjadi kesalahan yang tidak diketahui',
    code: 'UNKNOWN_ERROR',
    status: 500,
    details: error
  };
}

export async function apiRequest<T>(
  fetcher: typeof fetch,
  token: string | undefined,
  path: string,
  init: RequestInit = {}
): Promise<ApiEnvelope<T>> {
  const headers = new Headers(init.headers);

  if (token) {
    headers.set('Authorization', `Bearer ${token}`);
  }

  if (
    init.body &&
    !(init.body instanceof FormData) &&
    !headers.has('Content-Type')
  ) {
    headers.set('Content-Type', 'application/json');
  }

  headers.set('Accept', 'application/json');

  const request = getRequestHandler(fetcher);
  const url = `${API_BASE}${normalizePath(path)}`;

  let response: Response;

  try {
    response = await request(url, {
      ...init,
      headers
    });
  } catch (cause) {
    throw new ApiError(
      'Server API tidak dapat dihubungi',
      503,
      'API_UNREACHABLE',
      {
        path: normalizePath(path)
      },
      cause
    );
  }

  const payload = await parseApiResponse<T>(response);
  const succeeded = payload.status ?? payload.success ?? false;

  if (!response.ok || !succeeded) {
    const objectError =
      typeof payload.error === 'object' &&
      payload.error !== null
        ? (payload.error as ApiErrorBody)
        : undefined;

    const stringError =
      typeof payload.error === 'string'
        ? payload.error
        : undefined;

    throw new ApiError(
      payload.message ||
        stringError ||
        `Request gagal (${response.status})`,
      response.status,
      objectError?.code ||
        stringError ||
        'API_REQUEST_FAILED',
      objectError?.details
    );
  }

  return payload;
}

export async function apiData<T>(
  fetcher: typeof fetch,
  token: string | undefined,
  path: string,
  init: RequestInit = {}
): Promise<T> {
  const response = await apiRequest<T>(
    fetcher,
    token,
    path,
    init
  );

  return response.data;
}

export async function safeApiData<T>(
  fetcher: typeof fetch,
  token: string | undefined,
  path: string,
  init: RequestInit = {}
): Promise<SafeApiResult<T>> {
  try {
    const data = await apiData<T>(
      fetcher,
      token,
      path,
      init
    );

    return {
      ok: true,
      data,
      error: null
    };
  } catch (error) {
    return {
      ok: false,
      data: null,
      error: normalizeApiError(error)
    };
  }
}

export async function requireApiData<T>(
  fetcher: typeof fetch,
  token: string | undefined,
  path: string,
  init: RequestInit = {}
): Promise<T> {
  try {
    return await apiData<T>(
      fetcher,
      token,
      path,
      init
    );
  } catch (error) {
    const normalized = normalizeApiError(error);

    console.error('[API request failed]', {
      path,
      status: normalized.status,
      code: normalized.code,
      message: normalized.message,
      details: normalized.details
    });

    kitError(
      normalizeHttpStatus(normalized.status),
      normalized.message
    );
  }
}

function normalizeHttpStatus(status: number): number {
  if (status >= 400 && status <= 599) {
    return status;
  }

  return 500;
}

export function tokenFrom(
  cookies: Cookies
): string | undefined {
  return cookies.get(SESSION_COOKIE);
}

export function clearSession(cookies: Cookies): void {
  cookies.delete(SESSION_COOKIE, {
    path: '/'
  });
}

export function setSession(
  cookies: Cookies,
  token: string,
  secure: boolean
): void {
  cookies.set(SESSION_COOKIE, token, {
    path: '/',
    httpOnly: true,
    sameSite: 'lax',
    secure,
    maxAge: 60 * 60 * 24
  });
}

export function formValue(
  data: FormData,
  name: string
): string {
  return String(data.get(name) ?? '').trim();
}

export function nullable(
  data: FormData,
  name: string
): string | null {
  const value = formValue(data, name);

  return value || null;
}

export function actionError(
  error: unknown
): NormalizedApiError {
  return normalizeApiError(error);
}

export type AdminEvent = RequestEvent & {
  locals: App.Locals;
};