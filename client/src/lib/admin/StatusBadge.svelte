<script lang="ts">
  import { Badge } from '$lib/components/ui/badge';
  let { status }: { status: string } = $props();
  const labels: Record<string,string> = { PENDING:'Menunggu', IN_REVIEW:'Ditinjau', DRAFT:'Draft', SUBMITTED:'Dikirim', PENDING_REVIEW:'Menunggu peninjauan', REVISION_REQUIRED:'Perlu revisi', REVISION_SUBMITTED:'Revisi dikirim', APPROVED:'Disetujui', SCHEDULED:'Dijadwalkan', PUBLISHED:'Dipublikasikan', COMPLETED:'Selesai', REJECTED:'Ditolak' };
  const colors: Record<string,string> = { PENDING:'bg-warning/10 text-warning border-warning/25', IN_REVIEW:'bg-info/10 text-info border-info/25', DRAFT:'bg-muted text-muted-foreground border-border', SUBMITTED:'bg-blue-50 text-blue-700 border-blue-200', PENDING_REVIEW:'bg-warning/10 text-warning border-warning/25', REVISION_REQUIRED:'bg-orange-50 text-orange-700 border-orange-200', REVISION_SUBMITTED:'bg-blue-50 text-blue-700 border-blue-200', APPROVED:'bg-success/10 text-success border-success/25', SCHEDULED:'bg-blue-50 text-blue-800 border-blue-200', PUBLISHED:'bg-success/10 text-success border-success/25', COMPLETED:'bg-success/10 text-success border-success/25', REJECTED:'bg-destructive/10 text-destructive border-destructive/25' };
</script>
<Badge variant="outline" class={colors[status] || ''}>{labels[status] || status}</Badge>
