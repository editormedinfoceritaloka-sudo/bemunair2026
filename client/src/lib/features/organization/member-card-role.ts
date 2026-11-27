import type { OrganizationMember } from '$lib/types';

const fallbackRoles: Partial<Record<OrganizationMember['position_type'], string>> = {
  MENKO: 'Menteri Koordinator',
  MENTERI: 'MENTERI',
  MINISTER: 'MENTERI',
  KABIRO: 'KEPALA BIRO',
  DEPUTI: 'DEPUTI',
  KEPALA: 'KEPALA PAK'
};

export function memberCardRole({ position, position_type }: OrganizationMember): string {
  if (position_type === 'DIRJEN' || position_type === 'DIRECTOR_GENERAL') {
    return 'Direktur Jenderal';
  }
  return position?.trim() || fallbackRoles[position_type] || position_type;
}
