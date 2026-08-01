<script lang="ts">
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Card, CardContent } from '$lib/components/ui/card';
  import { CheckCircle2, ArrowRight, LayoutDashboard } from '@lucide/svelte';
  let { data } = $props();
  const row = $derived(data.submission);
</script>

<svelte:head><title>Pengajuan Berhasil · BEM UNAIR</title></svelte:head>
<div class="mx-auto max-w-2xl py-8 md:py-16">
  <Card class="overflow-hidden border-green-200 shadow-lg">
    <div class="h-2 bg-green-500"></div>
    <CardContent class="p-7 text-center md:p-10">
      <div class="mx-auto grid size-16 place-items-center rounded-full bg-green-50 text-green-600"><CheckCircle2 class="size-9" /></div>
      <h1 class="mt-5 font-serif text-3xl font-bold text-blue-900">Pengajuan berhasil dikirim</h1>
      <p class="mt-3 text-sm leading-6 text-muted-foreground">Tim Media dan Informasi akan memeriksa materi dan menghubungi nomor WhatsApp Anda jika diperlukan revisi.</p>

      <div class="mt-7 rounded-xl border bg-muted/30 p-5 text-left">
        <div class="grid gap-4 sm:grid-cols-2">
          <div><p class="text-xs text-muted-foreground">Kode pengajuan</p><p class="font-mono font-bold text-blue-800">{row.request_code || `MED-${row.id}`}</p></div>
          <div><p class="text-xs text-muted-foreground">Status awal</p><div class="mt-1"><StatusBadge status={row.status} /></div></div>
          <div class="sm:col-span-2"><p class="text-xs text-muted-foreground">Judul</p><p class="font-semibold">{row.title}</p></div>
          <div><p class="text-xs text-muted-foreground">Jenis</p><p class="font-medium">{row.service_type === 'ARTICLE' ? 'Artikel' : 'Media'}</p></div>
          <div><p class="text-xs text-muted-foreground">Tanggal usulan</p><p class="font-medium">{row.publish_date ? new Intl.DateTimeFormat('id-ID', { dateStyle: 'long' }).format(new Date(row.publish_date)) : '—'}</p></div>
        </div>
      </div>

      <div class="mt-7 flex flex-col justify-center gap-3 sm:flex-row">
        <Button href="/admin" variant="outline"><LayoutDashboard /> Ringkasan</Button>
        <Button href={`/admin/content-submissions/${row.id}`} class="bg-blue-500">Lihat detail & timeline <ArrowRight /></Button>
      </div>
    </CardContent>
  </Card>
</div>
