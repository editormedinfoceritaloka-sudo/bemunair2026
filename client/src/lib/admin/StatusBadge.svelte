<script lang="ts">
  import { Badge } from '$lib/components/ui/badge';
  let { status }: { status: string } = $props();
  const labels: Record<string,string> = { PENDING:'Menunggu', IN_REVIEW:'Ditinjau', APPROVED:'Disetujui', REJECTED:'Ditolak', DRAFT:'Draft', PUBLISHED:'Terbit' };
  const colors: Record<string,string> = { PENDING:'bg-warning/10 text-warning border-warning/25', IN_REVIEW:'bg-info/10 text-info border-info/25', APPROVED:'bg-success/10 text-success border-success/25', REJECTED:'bg-destructive/10 text-destructive border-destructive/25', DRAFT:'bg-muted text-muted-foreground border-border', PUBLISHED:'bg-success/10 text-success border-success/25' };
</script>
<Badge variant="outline" class={colors[status] || ''}>{labels[status] || status}</Badge>
