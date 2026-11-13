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

  const body = new FormData();
  body.set('file', file);
  body.set('purpose', purpose);

  return await new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/uploads/imagekit');
    xhr.upload.onprogress = (event) => {
      if (event.lengthComputable) onProgress?.(Math.round((event.loaded / event.total) * 100));
    };
    xhr.onerror = () => reject(new Error('Koneksi upload terputus.'));
    xhr.onload = () => {
      let result: Record<string, any> = {};
      try { result = JSON.parse(xhr.responseText); } catch { /* handled below */ }
      const data = result?.data;
      if (xhr.status < 200 || xhr.status >= 300 || !data) {
        reject(new Error(String(result?.message || 'Upload file gagal.')));
        return;
      }
      resolve({
        file_id: String(data.file_id),
        name: String(data.name || file.name),
        url: String(data.url),
        thumbnail_url: data.thumbnail_url ? String(data.thumbnail_url) : undefined,
        file_type: String(data.file_type || file.type),
        size: Number(data.size || file.size),
        mime_type: file.type
      });
    };
    xhr.send(body);
  });
}
