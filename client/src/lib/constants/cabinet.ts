export const organizationUnitTypes = [
  { value: 'KEMENKOAN', label: 'Kemenkoan' },
  { value: 'KEMENTERIAN', label: 'Kementerian' }
] as const;

export const managedOrganizationUnitTypes = [{ value: 'KEMENTERIAN', label: 'Kementerian' }] as const;

export const workProgramStatuses = [
  { value: 'DRAFT', label: 'Draft' },
  { value: 'PLANNED', label: 'Direncanakan' },
  { value: 'ONGOING', label: 'Berlangsung' },
  { value: 'COMPLETED', label: 'Selesai' },
  { value: 'POSTPONED', label: 'Ditunda' },
  { value: 'CANCELLED', label: 'Dibatalkan' },
  { value: 'ARCHIVED', label: 'Diarsipkan' }
] as const;
