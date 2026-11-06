import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { Cabinet, MediaAsset, OrganizationUnit } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

type UploadedMedia = {
  file_id: string;
  file_path?: string;
  url: string;
  thumbnail_url?: string;
  name: string;
  file_type: string;
  size: number;
  width?: number;
  height?: number;
};

const numeric = (form: FormData, key: string) => {
  const value = formValue(form, key);
  if (!value) return null;
  const parsed = Number(value);
  return Number.isFinite(parsed) && parsed > 0 ? parsed : null;
};

const uploaded = (form: FormData, key: string): UploadedMedia | null => {
  const value = formValue(form, key);
  if (!value) return null;
  try {
    return JSON.parse(value) as UploadedMedia;
  } catch {
    return null;
  }
};

async function mediaID(
  fetch: typeof globalThis.fetch,
  token: string | undefined,
  form: FormData,
  key: string,
  altText: string,
  purpose: string
) {
  const value = uploaded(form, `${key}_upload`);
  if (!value) return numeric(form, `${key}_media_id`);
  const result = await apiRequest<MediaAsset>(fetch, token, '/admin/media-assets', {
    method: 'POST',
    body: JSON.stringify({
      file_id: value.file_id,
      file_path: value.file_path || null,
      url: value.url,
      thumbnail_url: value.thumbnail_url || null,
      name: value.name,
      alt_text: altText,
      mime_type: value.file_type,
      size_bytes: value.size,
      width: value.width || null,
      height: value.height || null,
      purpose
    })
  });
  return result.data.id;
}

async function unitBody(fetch: typeof globalThis.fetch, token: string | undefined, form: FormData) {
  const name = formValue(form, 'name');
  return {
    cabinet_term_id: numeric(form, 'cabinet_term_id'),
    parent_id: formValue(form, 'unit_type') === 'KEMENTERIAN' ? numeric(form, 'parent_id') : null,
    code: formValue(form, 'code'),
    name,
    unit_type: formValue(form, 'unit_type'),
    slug: formValue(form, 'slug'),
    short_name: nullable(form, 'short_name'),
    description: nullable(form, 'description'),
    vision: nullable(form, 'vision'),
    mission: nullable(form, 'mission'),
    logo_media_id: await mediaID(fetch, token, form, 'logo', `Logo ${name}`, 'ORGANIZATION_LOGO'),
    cover_media_id: await mediaID(fetch, token, form, 'cover', `Cover ${name}`, 'ORGANIZATION_COVER'),
    display_order: Number(formValue(form, 'display_order') || 0),
    is_active: form.has('is_active'),
    is_published: form.has('is_published')
  };
}

export const load: PageServerLoad = async ({ fetch, cookies, url }) => {
  const token = tokenFrom(cookies);
  const cabinets = (await apiRequest<Cabinet[]>(fetch, token, '/admin/cabinet-terms?page=1&per_page=50')).data;
  const requestedID = Number(url.searchParams.get('cabinet') || 0);
  const cabinet = cabinets.find((item) => item.id === requestedID)
    ?? cabinets.find((item) => item.is_active)
    ?? cabinets[0];
  const units = cabinet
    ? (await apiRequest<OrganizationUnit[]>(fetch, token, `/admin/organizations?cabinet_id=${cabinet.id}`)).data
    : [];

  return {
    cabinets,
    cabinetId: cabinet?.id ?? null,
    units,
    filters: {
      q: url.searchParams.get('q') ?? '',
      type: url.searchParams.get('type') ?? 'ALL',
      status: url.searchParams.get('status') ?? 'ALL'
    }
  };
};

export const actions: Actions = {
  saveUnit: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    const id = numeric(form, 'id');
    const token = tokenFrom(cookies);
    try {
      await apiRequest(fetch, token, id ? `/admin/organizations/${id}` : '/admin/organizations', {
        method: id ? 'PUT' : 'POST',
        body: JSON.stringify(await unitBody(fetch, token, form))
      });
      return { success: true, message: id ? 'Unit organisasi diperbarui' : 'Unit organisasi dibuat' };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },

  deleteUnit: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/admin/organizations/${formValue(form, 'id')}`, {
        method: 'DELETE'
      });
      return { success: true, message: 'Unit organisasi dihapus' };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },

  saveMember: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    const id = numeric(form, 'member_id');
    const unitID = numeric(form, 'ministry_id');
    if (!unitID) return fail(422, { error: 'Unit organisasi wajib dipilih' });
    const token = tokenFrom(cookies);
    try {
      const name = formValue(form, 'member_name');
      await apiRequest(fetch, token, id ? `/admin/organizations/members/${id}` : `/admin/organizations/${unitID}/members`, {
        method: id ? 'PUT' : 'POST',
        body: JSON.stringify({
          ministry_id: unitID,
          name,
          position: formValue(form, 'position'),
          position_type: formValue(form, 'position_type'),
          biography: nullable(form, 'biography'),
          quote: nullable(form, 'quote'),
          photo_media_id: await mediaID(fetch, token, form, 'photo', `Foto ${name}`, 'ORGANIZATION_MEMBER'),
          display_order: Number(formValue(form, 'member_display_order') || 0),
          is_leader: form.has('is_leader'),
          is_active: form.has('member_is_active')
        })
      });
      return { success: true, message: id ? 'Anggota diperbarui' : 'Anggota ditambahkan' };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
