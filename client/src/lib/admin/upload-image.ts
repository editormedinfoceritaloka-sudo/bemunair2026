export type UploadPurpose = 'article' | 'cover' | 'profile' | 'submission';

export interface UploadedImage {
  file_id: string;
  name: string;
  file_path: string;
  url: string;
  thumbnail_url?: string;
  width?: number;
  height?: number;
  size: number;
  file_type: string;
}

interface UploadResponse {
  status: boolean;
  message: string;
  data?: UploadedImage;
  error?: { code?: string };
}

const ACCEPTED_TYPES = new Set(['image/jpeg', 'image/png', 'image/webp', 'image/gif', 'image/avif']);
const CLIENT_MAX_BYTES = 10 * 1024 * 1024;

export async function uploadImageFile(file: File, purpose: UploadPurpose): Promise<UploadedImage> {
  if (!ACCEPTED_TYPES.has(file.type)) {
    throw new Error('Format gambar harus JPEG, PNG, WebP, GIF, atau AVIF.');
  }
  if (file.size > CLIENT_MAX_BYTES) throw new Error('Ukuran gambar maksimal 10 MB.');

  const body = new FormData();
  body.set('file', file);
  body.set('purpose', purpose);

  const response = await fetch('/api/uploads/imagekit', {
    method: 'POST',
    body
  });
  const payload = await response.json().catch(() => null) as UploadResponse | null;
  if (!response.ok || !payload?.status || !payload.data) {
    throw new Error(payload?.message || 'Gagal mengunggah gambar.');
  }
  return payload.data;
}
