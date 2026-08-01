import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { Cabinet, OrganizationUnit } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies }) => { const token = tokenFrom(cookies); const cabinets = (await apiRequest<Cabinet[]>(fetch, token, '/admin/cabinet-terms?page=1&per_page=50')).data; const cabinet = cabinets[0]; const units = cabinet ? (await apiRequest<OrganizationUnit[]>(fetch, token, `/admin/organizations?cabinet_id=${cabinet.id}`)).data : []; return { cabinets, units, cabinetId: cabinet?.id || null }; };

export const actions: Actions = { create: async ({ request, fetch, cookies }) => { const form = await request.formData(); try { await apiRequest(fetch, tokenFrom(cookies), '/admin/organizations', { method: 'POST', body: JSON.stringify({ cabinet_term_id: Number(formValue(form, 'cabinet_term_id')), parent_id: formValue(form, 'parent_id') ? Number(formValue(form, 'parent_id')) : null, code: formValue(form, 'code'), name: formValue(form, 'name'), unit_type: 'KEMENTERIAN', slug: formValue(form, 'slug'), short_name: nullable(form, 'short_name'), description: nullable(form, 'description'), is_active: form.has('is_active'), is_published: form.has('is_published') }) }); return { success: true }; } catch (error) { const result = actionError(error); return fail(result.status, { error: result.message }); } } };
