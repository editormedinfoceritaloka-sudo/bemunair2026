import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';
import type { User } from '$lib/types';
import { apiRequest, tokenFrom } from '$lib/server/api';
import { ImageUploadError, parsePurpose, uploadImage, validateImage } from '$lib/server/imagekit';

export const POST: RequestHandler = async ({ request, cookies, fetch }) => {
  const token = tokenFrom(cookies);
  if (!token) {
    return json({ status: false, message: 'Sesi admin diperlukan.', error: { code: 'UNAUTHENTICATED' } }, { status: 401 });
  }

  try {
    const session = await apiRequest<User>(fetch, token, '/auth/me');
    if (!['ADMIN', 'ADMIN_MEDINFO'].includes(session.data?.role ?? '')) {
      return json({ status: false, message: 'Akses upload hanya untuk admin.', error: { code: 'FORBIDDEN' } }, { status: 403 });
    }

    let data: FormData;
    try {
      data = await request.formData();
    } catch {
      throw new ImageUploadError('Request harus menggunakan multipart/form-data.', 400, 'INVALID_MULTIPART_FORM');
    }
    const file = data.get('file');
    if (!(file instanceof File)) {
      throw new ImageUploadError('Field file wajib berisi gambar.', 400, 'FILE_REQUIRED');
    }

    const purpose = parsePurpose(data.get('purpose'));
    await validateImage(file);
    const uploaded = await uploadImage(file, purpose);

    return json({
      status: true,
      message: 'Gambar berhasil diunggah.',
      data: {
        file_id: uploaded.fileId,
        name: uploaded.name,
        file_path: uploaded.filePath,
        url: uploaded.url,
        thumbnail_url: uploaded.thumbnailUrl,
        width: uploaded.width,
        height: uploaded.height,
        size: uploaded.size,
        file_type: uploaded.fileType
      }
    }, { status: 201 });
  } catch (error) {
    if (error instanceof ImageUploadError) {
      return json({ status: false, message: error.message, error: { code: error.code } }, { status: error.status });
    }

    const status = typeof error === 'object' && error && 'status' in error && typeof error.status === 'number'
      ? error.status
      : 502;
    console.error('ImageKit upload failed', {
      name: error instanceof Error ? error.name : 'UnknownError',
      status
    });
    return json({
      status: false,
      message: 'Gagal mengunggah gambar ke media storage.',
      error: { code: 'IMAGEKIT_UPLOAD_FAILED' }
    }, { status: status >= 400 && status < 600 ? status : 502 });
  }
};
