export type SubmissionUploadPurpose = 'submission_media' | 'submission_brief';

export interface UploadedSubmissionFile {
  file_id: string;
  name: string;
  url: string;
  thumbnail_url?: string;
  file_type: string;
  size: number;
  mime_type: string;
}

interface UploadAuthentication {
  token: string;
  expire: number;
  signature: string;
  public_key: string;
  folder: string;
}

const MEDIA_TYPES = new Set([
  'image/jpeg', 'image/png', 'image/webp'
]);
const BRIEF_TYPES = new Set([
  'application/pdf',
  'application/msword',
  'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
]);
const MAX_MEDIA_BYTES = 20 * 1024 * 1024;
const MAX_BRIEF_BYTES = 15 * 1024 * 1024;

function validate(file: File, purpose: SubmissionUploadPurpose) {
  const allowed = purpose === 'submission_media' ? MEDIA_TYPES : BRIEF_TYPES;
  const max = purpose === 'submission_media' ? MAX_MEDIA_BYTES : MAX_BRIEF_BYTES;
  if (!allowed.has(file.type)) {
    throw new Error(purpose === 'submission_media'
      ? 'Format materi harus JPG, PNG, atau WebP.'
      : 'Format brief harus PDF, DOC, atau DOCX.');
  }
  if (!file.size) throw new Error('File tidak boleh kosong.');
  if (file.size > max) throw new Error(`Ukuran file maksimal ${Math.floor(max / 1024 / 1024)} MB.`);
}

export async function uploadSubmissionFile(
  file: File,
  purpose: SubmissionUploadPurpose,
  onProgress?: (percentage: number) => void
): Promise<UploadedSubmissionFile> {
  validate(file, purpose);
  const authResponse = await fetch(`/api/uploads/imagekit/auth?purpose=${purpose}`);
  const authPayload = await authResponse.json() as { status: boolean; message: string; data?: UploadAuthentication };
  if (!authResponse.ok || !authPayload.data) throw new Error(authPayload.message || 'Gagal menyiapkan upload.');

  const auth = authPayload.data;
  const body = new FormData();
  body.set('file', file);
  body.set('fileName', file.name);
  body.set('publicKey', auth.public_key);
  body.set('token', auth.token);
  body.set('expire', String(auth.expire));
  body.set('signature', auth.signature);
  body.set('folder', auth.folder);
  body.set('useUniqueFileName', 'true');
  body.set('tags', `bemunair,${purpose}`);

  return await new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', 'https://upload.imagekit.io/api/v1/files/upload');
    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable) onProgress?.(Math.round((event.loaded / event.total) * 100));
    };
    xhr.onerror = () => reject(new Error('Koneksi upload terputus.'));
    xhr.onload = () => {
      let result: Record<string, unknown> = {};
      try { result = JSON.parse(xhr.responseText); } catch { /* handled below */ }
      if (xhr.status < 200 || xhr.status >= 300) {
        reject(new Error(String(result.message || 'Upload ke ImageKit gagal.')));
        return;
      }
      resolve({
        file_id: String(result.fileId),
        name: String(result.name || file.name),
        url: String(result.url),
        thumbnail_url: result.thumbnailUrl ? String(result.thumbnailUrl) : undefined,
        file_type: String(result.fileType || ''),
        size: Number(result.size || file.size),
        mime_type: file.type
      });
    };
    xhr.send(body);
  });
}
