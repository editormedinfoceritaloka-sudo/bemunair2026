export function formatDate(value?: string, withTime = false) {
  if (!value) return '—';
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return value;
  return new Intl.DateTimeFormat('id-ID', withTime
    ? { dateStyle: 'medium', timeStyle: 'short', timeZone: 'Asia/Jakarta' }
    : { dateStyle: 'medium', timeZone: 'Asia/Jakarta' }).format(date);
}

export function deadlineTone(value?: string) {
  if (!value) return 'text-muted-foreground';
  const hours = (new Date(value).getTime() - Date.now()) / 3_600_000;
  if (hours < 0) return 'text-red-700';
  if (hours < 48) return 'text-amber-700';
  return 'text-black-300';
}
