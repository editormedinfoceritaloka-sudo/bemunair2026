import { actionError, apiRequest, formValue, nullable, tokenFrom } from '$lib/server/api';
import type { Cabinet, OrganizationUnit, WorkProgram } from '$lib/types';
import { fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch, cookies, parent, url }) => {
  const { user } = await parent();
  const token = tokenFrom(cookies);
  const cabinets = (await apiRequest<Cabinet[]>(fetch, token, '/admin/cabinet-terms?page=1&per_page=50')).data;
  const cabinet = cabinets[0];
  const allUnits = cabinet
    ? (await apiRequest<OrganizationUnit[]>(fetch, token, `/admin/organizations?cabinet_id=${cabinet.id}&unit_type=KEMENTERIAN`)).data
    : [];
  const units = user?.role === 'ADMIN_MEDINFO'
    ? allUnits
    : allUnits.filter((unit) => unit.id === user?.ministry_id);
  const requested = Number(url.searchParams.get('unit_id') || 0);
  const unit = units.find((item) => item.id === requested) ?? units[0];
  const programs = unit
    ? (await apiRequest<WorkProgram[]>(fetch, token, `/admin/work-programs?unit_id=${unit.id}&page=1&per_page=50`)).data
    : [];
  return { units, programs, unitId: unit?.id ?? null };
};

export const actions: Actions = {
  create: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), '/admin/work-programs', {
        method: 'POST',
        body: JSON.stringify({
          ministry_id: Number(formValue(form, 'ministry_id')),
          name: formValue(form, 'name'),
          short_description: nullable(form, 'short_description'),
          description: nullable(form, 'description'),
          start_date: nullable(form, 'start_date'),
          end_date: nullable(form, 'end_date'),
          status: formValue(form, 'status'),
          is_featured: form.has('is_featured')
        })
      });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  },
  delete: async ({ request, fetch, cookies }) => {
    const form = await request.formData();
    try {
      await apiRequest(fetch, tokenFrom(cookies), `/admin/work-programs/${formValue(form, 'id')}`, { method: 'DELETE' });
      return { success: true };
    } catch (error) {
      const result = actionError(error);
      return fail(result.status, { error: result.message });
    }
  }
};
