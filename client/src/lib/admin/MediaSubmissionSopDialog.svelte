<script lang="ts">
  import type { MediaSubmissionSetting } from '$lib/types';
  import * as Dialog from '$lib/components/ui/dialog';
  import { Button } from '$lib/components/ui/button';
  import { Checkbox } from '$lib/components/ui/checkbox';
  import { ExternalLink, TriangleAlert } from '@lucide/svelte';

  let {
    service,
    setting,
    onaccepted
  }: {
    service: 'CONTENT' | 'ARTICLE';
    setting: MediaSubmissionSetting;
    onaccepted: () => void;
  } = $props();

  let open = $state(true);
  let accepted = $state(false);
  const value = (snake: keyof MediaSubmissionSetting, pascal: keyof MediaSubmissionSetting) =>
    String(setting[snake] ?? setting[pascal] ?? '');
  const sopUrl = $derived(value('sop_url', 'SOPURL'));

  $effect(() => {
    if (!accepted && !open) open = true;
  });

  function continueToForm() {
    if (!accepted) return;
    onaccepted();
    open = false;
  }
</script>

<Dialog.Root bind:open>
  <Dialog.Content
    showCloseButton={false}
    class="max-h-[92vh] overflow-y-auto border-blue-100 bg-white-50 p-0 sm:max-w-3xl"
  >
    <div class="space-y-7 px-5 py-7 sm:px-9 sm:py-9">
      <Dialog.Header class="text-center">
        <Dialog.Title class="media-sop-title">
          <span class="media-sop-title-primary">Pengajuan {service === 'CONTENT' ? 'Konten' : 'Artikel'}</span>
          <span class="media-sop-title-secondary">Medinfo BEM UNAIR 2026</span>
        </Dialog.Title>
        <Dialog.Description class="sr-only">
          Ketentuan wajib sebelum mengisi formulir pengajuan {service === 'CONTENT' ? 'konten' : 'artikel'}.
        </Dialog.Description>
      </Dialog.Header>

      <div class="flex items-start justify-center gap-2 text-center font-bold leading-6 text-black-900">
        <TriangleAlert class="mt-0.5 size-5 shrink-0 fill-yellow-500 text-yellow-700" />
        <p>Sebelum mengisi formulir ini, pemohon WAJIB membaca dan memahami SOP Kementerian Media dan Informasi BEM UNAIR 2026 secara keseluruhan.</p>
        <TriangleAlert class="mt-0.5 hidden size-5 shrink-0 fill-yellow-500 text-yellow-700 sm:block" />
      </div>

      {#if service === 'ARTICLE'}
        <section>
          <h2 class="text-base font-bold text-black-900">Ketentuan Penulisan Berita/Liputan/Artikel:</h2>
          <ul class="mt-4 list-disc space-y-3 pl-5 text-sm leading-6 text-black-700">
            <li>Setiap artikel wajib diawali dengan tagline <strong>#CeritaHariIni</strong>. Struktur artikel terdiri dari Headline: JELITA (Jejak Liputan Cerita Loka), kepala berita, tubuh berita, dan ekor berita.</li>
            <li>Panjang artikel berkisar antara 300–500 kata.</li>
            <li>Hasil akhir dipublikasikan pada website dan Instagram BEM UNAIR 2026.</li>
            <li>Setiap artikel melalui proses penyuntingan oleh Kementerian Media dan Informasi sebelum dipublikasikan.</li>
            <li>Artikel dipublikasikan paling cepat H+3 setelah naskah diterima, dengan mempertimbangkan hasil penyuntingan.</li>
          </ul>
        </section>
      {:else}
        <section>
          <h2 class="text-base font-bold text-black-900">Keterangan Jenis Pengajuan:</h2>
          <ul class="mt-4 list-disc space-y-3 pl-5 text-sm leading-6 text-black-700">
            <li><strong>Feed Instagram:</strong> publikasi konten feed melalui akun Instagram BEM UNAIR.</li>
            <li><strong>Reels Instagram:</strong> video final dan caption disiapkan mandiri oleh pengaju, lalu dipublikasikan melalui Instagram BEM UNAIR dan dapat dikolaborasikan dengan akun terkait.</li>
            <li><strong>Instastory:</strong> materi vertikal singkat yang telah disiapkan untuk publikasi Instagram Story BEM UNAIR.</li>
          </ul>
        </section>
        <section>
          <h2 class="text-base font-bold text-black-900">Ketentuan Umum:</h2>
          <ul class="mt-4 list-disc space-y-3 pl-5 text-sm leading-6 text-black-700">
            <li>Pengajuan yang tidak lengkap dapat menghambat proses peninjauan dan publikasi konten.</li>
            <li>Pastikan seluruh informasi, caption, dan file benar serta dapat dipertanggungjawabkan.</li>
            <li>Link Canva wajib menggunakan akses “Anyone with the link can edit”.</li>
            <li>Jadwal publikasi menyesuaikan antrean, urgensi, dan kebijakan Medinfo.</li>
            <li>Kebutuhan mendesak harus dicantumkan pada informasi tambahan dan dikomunikasikan kepada PIC.</li>
            <li>Dengan melanjutkan, pemohon dianggap memahami dan menyetujui SOP Medinfo BEM UNAIR 2026.</li>
          </ul>
        </section>
      {/if}

      {#if sopUrl}
        <a href={sopUrl} target="_blank" rel="noreferrer" class="inline-flex items-center gap-2 font-semibold text-blue-700 underline underline-offset-4 hover:text-blue-900">
          Buka SOP lengkap <ExternalLink class="size-4" />
        </a>
      {/if}

      <label class="flex cursor-pointer items-start gap-3 rounded-xl border border-blue-200 bg-blue-50 p-4">
        <Checkbox bind:checked={accepted} class="mt-0.5 size-5" />
        <span class="text-sm leading-6 text-black-700">Saya telah membaca dan memahami SOP serta seluruh ketentuan pengajuan.</span>
      </label>

      <Dialog.Footer class="gap-3 sm:justify-between">
        <Button href="/admin/content-submissions/new/select" variant="outline">Kembali</Button>
        <Button type="button" disabled={!accepted} onclick={continueToForm} class="bg-blue-500 hover:bg-blue-600">
          Selanjutnya
        </Button>
      </Dialog.Footer>
    </div>
  </Dialog.Content>
</Dialog.Root>

<style>
  .media-sop-title {
    margin: 0;
    text-align: center;
    font-weight: 800;
    line-height: 1.08;
    letter-spacing: -0.02em;
  }
  .media-sop-title span { display: block; }
  .media-sop-title-primary {
    color: var(--blue-600);
    font-size: clamp(1.875rem, 5vw, 2.75rem);
    text-shadow: 0 3px 0 var(--white-600), 1px 4px 3px rgb(0 0 0 / 0.16);
  }
  .media-sop-title-secondary {
    color: var(--blue-400);
    font-size: clamp(1.75rem, 4.8vw, 2.625rem);
    text-shadow: 0 3px 0 var(--white-600), 1px 4px 3px rgb(0 0 0 / 0.14);
  }
</style>
