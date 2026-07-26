<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
  import { Button } from '$lib/components/ui/button';
  import { FileImage, Newspaper, ArrowRight, Clock3 } from '@lucide/svelte';
</script>

<svelte:head><title>Pilih Pengajuan Media · BEM UNAIR</title></svelte:head>

<PageHeader
  title="Pengajuan Media Medinfo"
  description="Pilih layanan yang sesuai. Setiap pengajuan akan tersimpan atas nama akun yang sedang login."
/>

<div class="mb-6 flex items-center gap-2 text-xs font-medium text-muted-foreground">
  <span class="rounded-full bg-blue-500 px-3 py-1 text-white">1. Pilih layanan</span>
  <span>2. Baca ketentuan</span><span>→</span><span>3. Isi formulir</span>
</div>

<div class="grid max-w-4xl gap-5 md:grid-cols-2">
  <Card class="group border-blue-100 bg-blue-50/40 transition hover:-translate-y-0.5 hover:border-blue-300 hover:shadow-md">
    <CardHeader>
      <div class="mb-3 grid size-12 place-items-center rounded-xl bg-blue-500 text-white"><FileImage /></div>
      <CardTitle>Pengajuan Konten</CardTitle>
    </CardHeader>
    <CardContent class="space-y-5">
      <p class="text-sm leading-6 text-muted-foreground">Feed, reels, Instastory, desain, dan publikasi Instagram BEM UNAIR.</p>
      <div class="flex items-center gap-2 text-xs text-blue-700"><Clock3 class="size-4" />Diajukan paling lambat H-7</div>
      <Button href="/admin/content-submissions/new/content" class="w-full bg-blue-500">Pilih Konten <ArrowRight /></Button>
    </CardContent>
  </Card>

  <Card class="group border-orange-100 bg-orange-50/50 transition hover:-translate-y-0.5 hover:border-orange-300 hover:shadow-md">
    <CardHeader>
      <div class="mb-3 grid size-12 place-items-center rounded-xl bg-orange-500 text-white"><Newspaper /></div>
      <CardTitle>Pengajuan Artikel</CardTitle>
    </CardHeader>
    <CardContent class="space-y-5">
      <p class="text-sm leading-6 text-muted-foreground">Berita, liputan kegiatan, dokumentasi, dan artikel website BEM UNAIR.</p>
      <div class="flex items-center gap-2 text-xs text-orange-700"><Clock3 class="size-4" />Publikasi paling cepat H+3</div>
      <Button href="/admin/content-submissions/new/article" class="w-full bg-orange-500 hover:bg-orange-600">Pilih Artikel <ArrowRight /></Button>
    </CardContent>
  </Card>
</div>
