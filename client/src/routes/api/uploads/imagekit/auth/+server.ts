import ImageKit from '@imagekit/nodejs';
import { env } from '$env/dynamic/private';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import type { User } from '$lib/types';
import { apiRequest, tokenFrom } from '$lib/server/api';

const PURPOSES = {
  article: { folder: 'articles', medinfoOnly: true },
  cover: { folder: 'article-covers', medinfoOnly: true },
  profile: { folder: 'profiles', medinfoOnly: true },
  submission_media: { folder: 'submissions/media', medinfoOnly: false },
  submission_brief: { folder: 'submissions/briefs', medinfoOnly: false }
} as const;

type Purpose = keyof typeof PURPOSES;

export const GET: RequestHandler = async ({ url, cookies, fetch }) => {
  const token = tokenFrom(cookies);
  if (!token) {
    return json({ status: false, message: 'Sesi admin diperlukan.' }, { status: 401 });
  }

  const purpose = (url.searchParams.get('purpose') || '') as Purpose;
  if (!(purpose in PURPOSES)) {
    return json({ status: false, message: 'Tujuan upload tidak valid.' }, { status: 400 });
  }

  try {
    const { data: user } = await apiRequest<User>(fetch, token, '/auth/me');
    if (!['ADMIN', 'ADMIN_MEDINFO'].includes(user.role)) {
      return json({ status: false, message: 'Akses upload ditolak.' }, { status: 403 });
    }
    if (PURPOSES[purpose].medinfoOnly && user.role !== 'ADMIN_MEDINFO') {
      return json({ status: false, message: 'Upload ini hanya untuk Admin Medinfo.' }, { status: 403 });
    }

    const privateKey = env.IMAGEKIT_PRIVATE_KEY?.trim();
    const publicKey = env.IMAGEKIT_PUBLIC_KEY?.trim();
    if (!privateKey || !publicKey) {
      return json({ status: false, message: 'ImageKit belum dikonfigurasi.' }, { status: 503 });
    }

    const imagekit = new ImageKit({ privateKey });
    const authentication = imagekit.helper.getAuthenticationParameters(undefined, 5 * 60);
    const root = (env.IMAGEKIT_UPLOAD_FOLDER || '/bemunair').trim().replace(/^\/+|\/+$/g, '');
    const folder = `/${[root, PURPOSES[purpose].folder].filter(Boolean).join('/')}`;

    return json({
      status: true,
      message: 'Autentikasi upload dibuat.',
      data: { ...authentication, public_key: publicKey, folder }
    });
  } catch {
    return json({ status: false, message: 'Gagal membuat autentikasi upload.' }, { status: 502 });
  }
};
