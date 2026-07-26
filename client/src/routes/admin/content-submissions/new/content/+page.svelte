<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Card, CardContent } from '$lib/components/ui/card';
  import { Button } from '$lib/components/ui/button';
  import { Checkbox } from '$lib/components/ui/checkbox';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { ArrowLeft, ArrowRight, ExternalLink, Info } from '@lucide/svelte';
  let { data } = $props();
  let accepted = $state(false);
  const setting = data.setting;
  const terms = setting.terms || [];
  const sopUrl = setting.sop_url || setting.SOPURL;
</script>

<PageHeader title="Ketentuan Pengajuan Konten" description="Baca SOP dan ketentuan sebelum melanjutkan ke formulir." />
<div class="mb-6 flex items-center gap-2 text-xs font-medium text-muted-foreground">
  <span>1. Pilih layanan</span><span>→</span><span class="rounded-full bg-blue-500 px-3 py-1 text-white">2. Baca ketentuan</span><span>→</span><span>3. Isi formulir</span>
</div>

<div class="max-w-4xl space-y-5">
  <Alert class="border-blue-200 bg-blue-50 text-blue-900">
    <Info class="size-4" />
    <AlertDescription>Pengaju wajib membaca SOP Kementerian Media dan Informasi BEM UNAIR 2026 secara menyeluruh.</AlertDescription>
  </Alert>

  <Card class="border-blue-100 bg-blue-50/40">
    <CardContent class="space-y-6 p-6 md:p-8">
      <section>
        <h2 class="font-serif text-xl font-semibold text-blue-900">Jenis publikasi</h2>
        <div class="mt-4 grid gap-3 sm:grid-cols-3">
          <div class="rounded-xl bg-white p-4 shadow-sm"><h3 class="font-semibold">Feed Instagram</h3><p class="mt-1 text-sm text-muted-foreground">Desain feed dan caption untuk akun BEM UNAIR.</p></div>
          <div class="rounded-xl bg-white p-4 shadow-sm"><h3 class="font-semibold">Reels Instagram</h3><p class="mt-1 text-sm text-muted-foreground">Video final, caption, serta referensi lagu bila diperlukan.</p></div>
          <div class="rounded-xl bg-white p-4 shadow-sm"><h3 class="font-semibold">Instastory</h3><p class="mt-1 text-sm text-muted-foreground">Materi vertikal singkat untuk publikasi story.</p></div>
        </div>
      </section>

      <section>
        <h2 class="font-serif text-xl font-semibold text-blue-900">Ketentuan umum</h2>
        <ul class="mt-3 list-disc space-y-2 pl-5 text-sm leading-6 text-black-400">
          {#each terms as term}<li>{term}</li>{/each}
          <li>Pengajuan dilakukan paling lambat tujuh hari sebelum tanggal publikasi.</li>
          <li>Waktu publikasi merupakan usulan dan dapat berubah mengikuti antrean Medinfo.</li>
        </ul>
      </section>

      {#if sopUrl}
        <a href={sopUrl} target="_blank" rel="noreferrer" class="inline-flex items-center gap-2 text-sm font-semibold text-blue-700">Buka SOP lengkap <ExternalLink class="size-4" /></a>
      {/if}

      <label class="flex cursor-pointer items-start gap-3 rounded-xl border border-blue-200 bg-white p-4">
        <Checkbox bind:checked={accepted} />
        <span class="text-sm">Saya telah membaca dan memahami SOP serta seluruh ketentuan pengajuan konten.</span>
      </label>
    </CardContent>
  </Card>

  <div class="flex justify-between">
    <Button href="/admin/content-submissions/new/select" variant="outline"><ArrowLeft /> Kembali</Button>
    <Button href="/admin/content-submissions/new/content/form" disabled={!accepted} class="bg-blue-500">Lanjut isi formulir <ArrowRight /></Button>
  </div>
</div>
