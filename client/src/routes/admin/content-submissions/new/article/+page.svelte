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

<PageHeader title="Ketentuan Pengajuan Artikel" description="Pastikan naskah dan dokumentasi telah mengikuti standar Medinfo." />
<div class="mb-6 flex items-center gap-2 text-xs font-medium text-muted-foreground">
  <span>1. Pilih layanan</span><span>→</span><span class="rounded-full bg-orange-500 px-3 py-1 text-white">2. Baca ketentuan</span><span>→</span><span>3. Isi formulir</span>
</div>

<div class="max-w-4xl space-y-5">
  <Alert class="border-orange-200 bg-orange-50 text-orange-900">
    <Info class="size-4" />
    <AlertDescription>Artikel akan melalui proses penyuntingan dan dipublikasikan paling cepat H+3 setelah diterima.</AlertDescription>
  </Alert>
  <Card class="border-orange-100 bg-orange-50/40">
    <CardContent class="space-y-6 p-6 md:p-8">
      <section>
        <h2 class="font-serif text-xl font-semibold text-blue-900">Struktur naskah</h2>
        <div class="mt-3 rounded-xl bg-white p-5 font-mono text-sm leading-7 shadow-sm">
          <p>#CeritaHariIni</p><p class="mt-2 font-bold">Headline</p>
          <p class="mt-2">Kepala berita — inti kegiatan.</p>
          <p>Tubuh berita — rangkaian, peserta, tujuan, dan hasil.</p>
          <p>Ekor berita — kesimpulan atau tindak lanjut.</p>
        </div>
      </section>
      <section>
        <h2 class="font-serif text-xl font-semibold text-blue-900">Ketentuan penulisan</h2>
        <ul class="mt-3 list-disc space-y-2 pl-5 text-sm leading-6 text-black-400">
          {#each terms as term}<li>{term}</li>{/each}
          <li>Dokumentasi maksimal empat foto berkualitas tinggi dengan akses tautan terbuka.</li>
          <li>Naskah Google Docs harus dapat diakses untuk proses penyuntingan.</li>
        </ul>
      </section>
      {#if sopUrl}<a href={sopUrl} target="_blank" rel="noreferrer" class="inline-flex items-center gap-2 text-sm font-semibold text-blue-700">Buka SOP lengkap <ExternalLink class="size-4" /></a>{/if}
      <label class="flex cursor-pointer items-start gap-3 rounded-xl border border-orange-200 bg-white p-4">
        <Checkbox bind:checked={accepted} />
        <span class="text-sm">Saya telah membaca ketentuan penulisan dan menyiapkan naskah serta dokumentasi.</span>
      </label>
    </CardContent>
  </Card>
  <div class="flex justify-between">
    <Button href="/admin/content-submissions/new/select" variant="outline"><ArrowLeft /> Kembali</Button>
    <Button href="/admin/content-submissions/new/article/form" disabled={!accepted} class="bg-orange-500 hover:bg-orange-600">Lanjut isi formulir <ArrowRight /></Button>
  </div>
</div>
