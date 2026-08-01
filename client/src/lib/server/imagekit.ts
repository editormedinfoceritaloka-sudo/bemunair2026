import ImageKit from '@imagekit/nodejs';
import { env } from '$env/dynamic/private';

const DEFAULT_MAX_BYTES = 10 * 1024 * 1024;
const ALLOWED_TYPES = new Set(['image/jpeg', 'image/png', 'image/webp', 'image/gif', 'image/avif']);
const ALLOWED_PDF_TYPES = new Set(['application/pdf']);
const PURPOSE_FOLDERS = {
  article: 'articles',
  cover: 'article-covers',
  profile: 'profiles',
  submission: 'submissions',
  documentation: 'program-documentations',
  letter_template: 'letter-templates'
} as const;

export type ImagePurpose = keyof typeof PURPOSE_FOLDERS;

export class ImageUploadError extends Error {
  constructor(message: string, public status: number, public code: string) {
    super(message);
  }
}

function configuredMaxBytes() {
  const configured = Number(env.IMAGEKIT_MAX_UPLOAD_BYTES);
  return Number.isFinite(configured) && configured > 0 ? configured : DEFAULT_MAX_BYTES;
}

function hasImageSignature(bytes: Uint8Array, type: string) {
  const startsWith = (...signature: number[]) => signature.every((byte, index) => bytes[index] === byte);
  if (type === 'image/jpeg') return startsWith(0xff, 0xd8, 0xff);
  if (type === 'image/png') return startsWith(0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a);
  if (type === 'image/gif') return new TextDecoder().decode(bytes.slice(0, 6)) === 'GIF87a' || new TextDecoder().decode(bytes.slice(0, 6)) === 'GIF89a';
  if (type === 'image/webp') return new TextDecoder().decode(bytes.slice(0, 4)) === 'RIFF' && new TextDecoder().decode(bytes.slice(8, 12)) === 'WEBP';
  if (type === 'image/avif') return new TextDecoder().decode(bytes.slice(4, 12)).includes('ftypavif') || new TextDecoder().decode(bytes.slice(4, 12)).includes('ftypavis');
  return false;
}

function safeFileName(original: string) {
  const dot = original.lastIndexOf('.');
  const extension = dot >= 0 ? original.slice(dot + 1).toLowerCase().replace(/[^a-z0-9]/g, '') : '';
  const base = (dot >= 0 ? original.slice(0, dot) : original)
    .normalize('NFKD')
    .replace(/[\u0300-\u036f]/g, '')
    .toLowerCase()
    .replace(/[^a-z0-9_-]+/g, '-')
    .replace(/^-+|-+$/g, '')
    .slice(0, 80) || 'image';
  return extension ? `${base}.${extension}` : base;
}

export async function validateUpload(file: File, purpose: ImagePurpose) {
  if (purpose === 'letter_template') {
    if (!ALLOWED_PDF_TYPES.has(file.type)) throw new ImageUploadError('Template surat harus berupa file PDF.', 415, 'UNSUPPORTED_PDF_TYPE');
    if (file.size === 0) throw new ImageUploadError('File PDF kosong.', 400, 'EMPTY_FILE');
    if (file.size > configuredMaxBytes()) throw new ImageUploadError('Ukuran PDF maksimal ' + Math.floor(configuredMaxBytes() / 1024 / 1024) + ' MB.', 413, 'FILE_TOO_LARGE');
    const header = new TextDecoder().decode(new Uint8Array(await file.slice(0, 5).arrayBuffer()));
    if (header !== '%PDF-') throw new ImageUploadError('Isi file tidak sesuai dengan PDF.', 415, 'INVALID_PDF_SIGNATURE');
    return;
  }
  return validateImage(file);
}

export async function validateImage(file: File) {
  if (!ALLOWED_TYPES.has(file.type)) {
    throw new ImageUploadError('Format gambar harus JPEG, PNG, WebP, GIF, atau AVIF.', 415, 'UNSUPPORTED_IMAGE_TYPE');
  }
  if (file.size === 0) throw new ImageUploadError('File gambar kosong.', 400, 'EMPTY_FILE');
  if (file.size > configuredMaxBytes()) {
    throw new ImageUploadError(`Ukuran gambar maksimal ${Math.floor(configuredMaxBytes() / 1024 / 1024)} MB.`, 413, 'FILE_TOO_LARGE');
  }

  const header = new Uint8Array(await file.slice(0, 16).arrayBuffer());
  if (!hasImageSignature(header, file.type)) {
    throw new ImageUploadError('Isi file tidak sesuai dengan format gambar yang dikirim.', 415, 'INVALID_IMAGE_SIGNATURE');
  }
}

export function parsePurpose(value: FormDataEntryValue | null): ImagePurpose {
  const purpose = String(value || 'article') as ImagePurpose;
  if (!(purpose in PURPOSE_FOLDERS)) throw new ImageUploadError('Tujuan upload tidak valid.', 400, 'INVALID_UPLOAD_PURPOSE');
  return purpose;
}

export async function uploadImage(file: File, purpose: ImagePurpose) {
  const privateKey = env.IMAGEKIT_PRIVATE_KEY?.trim();
  if (!privateKey) throw new ImageUploadError('ImageKit belum dikonfigurasi pada server.', 503, 'IMAGEKIT_NOT_CONFIGURED');

  const root = (env.IMAGEKIT_UPLOAD_FOLDER || '/bemunair').trim().replace(/^\/+|\/+$/g, '');
  const folder = `/${[root, PURPOSE_FOLDERS[purpose]].filter(Boolean).join('/')}`;
  const imagekit = new ImageKit({ privateKey, timeout: 30_000, maxRetries: 1 });

  return imagekit.files.upload({
    file,
    fileName: safeFileName(file.name),
    folder,
    useUniqueFileName: true,
    tags: ['bemunair', purpose]
  });
}
