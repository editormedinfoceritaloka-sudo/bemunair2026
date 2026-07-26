<script lang="ts">
  import { onMount } from 'svelte';
  import type { MediaSubmissionSetting, User } from '$lib/types';
  import { uploadSubmissionFile, type UploadedSubmissionFile } from '$lib/admin/upload-submission-file';
  import MediaSubmissionSopDialog from '$lib/admin/MediaSubmissionSopDialog.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import * as Dialog from '$lib/components/ui/dialog';
  import { ArrowLeft, Check, ExternalLink, FileText, FileUp, Image, Info, LoaderCircle, Save, Trash2 } from '@lucide/svelte';

  let { service, user, setting, form }: {
    service: 'CONTENT' | 'ARTICLE';
    user: User;
    setting: MediaSubmissionSetting;
    form?: { error?: string } | null;
  } = $props();

  let formElement = $state<HTMLFormElement>();
  let sopAccepted = $state(false);
  let reviewOpen = $state(false);
  let confirmed = $state(false);
  let submitting = $state(false);
  let draftRestored = $state(false);
  let media = $state<UploadedSubmissionFile | null>(null);
  let brief = $state<UploadedSubmissionFile | null>(null);
  let mediaProgress = $state(0);
  let briefProgress = $state(0);
  let mediaUploading = $state(false);
  let briefUploading = $state(false);
  let uploadError = $state('');
  let contentFormat = $state('FEED_INSTAGRAM');
  let noSong = $state(true);
  let agreements = $state([false, false, false, false]);
  let captionLength = $state(0);
  let reviewData = $state<Record<string, string>>({});

  const allAccepted = $derived(agreements.every(Boolean));
  const minimumDays = $derived(setting.minimum_lead_days ?? setting.MinimumLeadDays ?? (service === 'CONTENT' ? 7 : 3));
  const minimumDate = $derived.by(() => {
    const date = new Date();
    date.setDate(date.getDate() + minimumDays);
    return date.toISOString().slice(0, 10);
  });
  const startTime = $derived(setting.publish_time_start ?? setting.PublishTimeStart ?? '08:00');
  const endTime = $derived(setting.publish_time_end ?? setting.PublishTimeEnd ?? '17:00');
  const interval = $derived(setting.slot_interval_minutes ?? setting.SlotIntervalMinutes ?? 30);
  const draftKey = $derived(`bemunair:media-draft:${user.id}:${service}`);
  const settingValue = (snake: keyof MediaSubmissionSetting, pascal: keyof MediaSubmissionSetting) =>
    String(setting[snake] ?? setting[pascal] ?? '');
  const references = $derived(service === 'CONTENT'
    ? [
        ['LINK SOP', settingValue('sop_url', 'SOPURL')],
        ['TEMPLATE KEMENTERIAN', settingValue('ministry_template_url', 'MinistryTemplateURL')],
        ['TEMPLATE BRIEF KONTEN', settingValue('brief_template_url', 'BriefTemplateURL')]
      ]
    : [
        ['LINK SOP', settingValue('sop_url', 'SOPURL')],
        ['TEMPLATE ARTIKEL', settingValue('ministry_template_url', 'MinistryTemplateURL')],
        ['TEMPLATE CAPTION', settingValue('caption_template_url', 'CaptionTemplateURL')]
      ]);

  onMount(() => {
    const raw = localStorage.getItem(draftKey);
    if (!raw || !formElement) return;
    try {
      const draft = JSON.parse(raw) as { fields?: Record<string, string>; contentFormat?: string; noSong?: boolean; agreements?: boolean[]; media?: UploadedSubmissionFile; brief?: UploadedSubmissionFile };
      for (const [name, value] of Object.entries(draft.fields || {})) {
        const field = formElement.elements.namedItem(name);
        if (field instanceof HTMLInputElement || field instanceof HTMLTextAreaElement || field instanceof HTMLSelectElement) field.value = value;
      }
      contentFormat = draft.contentFormat || contentFormat;
      noSong = draft.noSong ?? noSong;
      agreements = Array.isArray(draft.agreements) ? draft.agreements : agreements;
      media = draft.media || null;
      brief = draft.brief || null;
      captionLength = (draft.fields?.caption || '').length;
      draftRestored = true;
    } catch {
      localStorage.removeItem(draftKey);
    }
    const warn = (event: BeforeUnloadEvent) => {
      if (!submitting && localStorage.getItem(draftKey)) event.preventDefault();
    };
    window.addEventListener('beforeunload', warn);
    return () => window.removeEventListener('beforeunload', warn);
  });

  function saveDraft() {
    if (!formElement || submitting) return;
    const fields: Record<string, string> = {};
    const data = new FormData(formElement);
    for (const [key, value] of data.entries()) if (typeof value === 'string') fields[key] = value;
    localStorage.setItem(draftKey, JSON.stringify({ fields, contentFormat, noSong, agreements, media, brief }));
  }

  async function upload(file: File | undefined, purpose: 'submission_media' | 'submission_brief') {
    if (!file) return;
    uploadError = '';
    if (purpose === 'submission_media') mediaUploading = true; else briefUploading = true;
    try {
      const result = await uploadSubmissionFile(file, purpose, (value) =>
        purpose === 'submission_media' ? mediaProgress = value : briefProgress = value);
      if (purpose === 'submission_media') media = result; else brief = result;
      saveDraft();
    } catch (error) {
      uploadError = error instanceof Error ? error.message : 'Upload gagal.';
    } finally {
      if (purpose === 'submission_media') mediaUploading = false; else briefUploading = false;
    }
  }

  function prepareReview(event: SubmitEvent) {
    if (confirmed) {
      submitting = true;
      localStorage.removeItem(draftKey);
      return;
    }
    event.preventDefault();
    if (!formElement?.reportValidity()) return;
    const data = new FormData(formElement);
    reviewData = Object.fromEntries([...data.entries()].filter(([, value]) => typeof value === 'string')) as Record<string, string>;
    reviewOpen = true;
  }

  function submitConfirmed() {
    confirmed = true;
    reviewOpen = false;
    requestAnimationFrame(() => formElement?.requestSubmit());
  }
</script>

<MediaSubmissionSopDialog {service} {setting} onaccepted={() => sopAccepted = true} />

<div class="mx-auto max-w-5xl" class:pointer-events-none={!sopAccepted} aria-hidden={!sopAccepted}>
  <header class="mb-8 text-center">
    <h1 class="submission-title">
      <span class="text-blue-600">Pengajuan {service === 'CONTENT' ? 'Konten' : 'Artikel'}</span>
      <span class="text-blue-400">Medinfo BEM UNAIR 2026</span>
    </h1>
    <p class="mt-4 text-sm text-muted-foreground">Lengkapi seluruh informasi berikut. Data akan tersimpan sebagai draft pada perangkat ini.</p>
  </header>

  {#if form?.error}<Alert variant="destructive" class="mb-5"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
  {#if uploadError}<Alert variant="destructive" class="mb-5"><AlertDescription>{uploadError}</AlertDescription></Alert>{/if}
  {#if draftRestored}<Alert class="mb-5 border-blue-200 bg-blue-50"><Save class="size-4" /><AlertDescription>Draft terakhir berhasil dipulihkan.</AlertDescription></Alert>{/if}

  <section class="mb-6 grid gap-3 sm:grid-cols-3">
    {#each references as reference}
      {#if reference[1]}
        <a href={reference[1]} target="_blank" rel="noreferrer" class="flex items-center justify-between gap-3 rounded-xl border border-blue-100 bg-white p-4 text-sm font-bold text-blue-800 underline-offset-4 shadow-sm hover:border-blue-300 hover:underline">
          <span class="flex items-center gap-2"><FileText class="size-4" />{reference[0]}</span><ExternalLink class="size-4 shrink-0" />
        </a>
      {/if}
    {/each}
  </section>

  <form bind:this={formElement} method="POST" enctype="multipart/form-data" class="submission-panel space-y-8" oninput={saveDraft} onchange={saveDraft} onsubmit={prepareReview}>
    <input type="hidden" name="service_type" value={service} />
    <input type="hidden" name="submission_type" value={service === 'ARTICLE' ? 'ARTIKEL' : contentFormat} />
    <input type="hidden" name="content_format" value={service === 'CONTENT' ? contentFormat : ''} />
    <input type="hidden" name="ministry" value={user.ministry || ''} />
    {#if service === 'CONTENT' && contentFormat !== 'REELS_INSTAGRAM'}
      <input type="hidden" name="design_drive_link" value={media?.url || ''} />
    {/if}
    <input type="hidden" name="brief_link" value={service === 'CONTENT' ? brief?.url || '' : '-'} />
    <input type="hidden" name="media_file_id" value={media?.file_id || ''} />
    <input type="hidden" name="media_file_name" value={media?.name || ''} />
    <input type="hidden" name="media_file_mime_type" value={media?.mime_type || ''} />
    <input type="hidden" name="media_file_size" value={media?.size || 0} />
    <input type="hidden" name="brief_file_id" value={brief?.file_id || ''} />
    <input type="hidden" name="brief_file_name" value={brief?.name || ''} />
    <input type="hidden" name="brief_file_mime_type" value={brief?.mime_type || ''} />
    <input type="hidden" name="brief_file_size" value={brief?.size || 0} />

    <section class="form-section">
      <h2>1. Informasi Pengaju</h2>
      <div class="grid gap-5 sm:grid-cols-2">
        <div class="field"><Label for="submitter_name">Nama Lengkap Menteri</Label><Input id="submitter_name" value={user.name} disabled /><p>Mengikuti akun yang sedang login.</p></div>
        <div class="field"><Label for="ministry_view">Kementerian</Label><Input id="ministry_view" value={user.ministry || 'Belum ditentukan'} disabled /></div>
        <div class="field sm:col-span-2"><Label for="submitter_phone">Nomor WhatsApp *</Label><Input id="submitter_phone" name="submitter_phone" type="tel" value={user.phone || ''} placeholder="Contoh: 6281234567890" pattern="[+]?[0-9]+" minlength={10} maxlength={15} required /><p>Gunakan nomor aktif; format 08, 62, atau +62 akan dinormalisasi.</p></div>
      </div>
    </section>

    <section class="form-section">
      <h2>2. Detail Publikasi</h2>
      <div class="grid gap-5 sm:grid-cols-2">
        {#if service === 'CONTENT'}
          <div class="field sm:col-span-2">
            <Label>Jenis Pengajuan *</Label>
            <div class="grid gap-3 sm:grid-cols-3">
              {#each [['FEED_INSTAGRAM', 'Feed Instagram'], ['REELS_INSTAGRAM', 'Reels Instagram'], ['INSTASTORY', 'Instastory']] as option}
                <label class="flex cursor-pointer items-center gap-3 rounded-xl border border-blue-100 bg-white p-4 has-[:checked]:border-blue-500 has-[:checked]:bg-blue-50">
                  <input type="radio" bind:group={contentFormat} value={option[0]} onchange={saveDraft} /><span class="text-sm font-semibold">{option[1]}</span>
                </label>
              {/each}
            </div>
          </div>
        {/if}
        <div class="field"><Label for="publish_date">Tanggal Publikasi *</Label><Input id="publish_date" name="publish_date" type="date" min={minimumDate} required /><p>Minimal {minimumDays} hari dari hari ini.</p></div>
        <div class="field"><Label for="publish_time">Waktu Publikasi *</Label><Input id="publish_time" name="publish_time" type="time" min={startTime} max={endTime} step={interval * 60} required /><p>Tersedia {startTime}–{endTime} WIB dalam interval {interval} menit.</p></div>
        <div class="field sm:col-span-2"><Label for="title">{service === 'ARTICLE' ? 'Judul Artikel / Nama Kegiatan' : 'Judul Konten / Nama Kegiatan'} *</Label><Input id="title" name="title" minlength={5} maxlength={150} required /></div>
      </div>
    </section>

    <section class="form-section">
      <h2>3. Materi {service === 'CONTENT' ? 'Konten' : 'Artikel'}</h2>
      <div class="grid gap-5">
        {#if service === 'CONTENT'}
          {#if contentFormat === 'REELS_INSTAGRAM'}
            <Alert class="border-blue-200 bg-blue-50"><Info class="size-4" /><AlertDescription>Unggah video final ke Google Drive. Pastikan akses menggunakan “Anyone with the link can view”, lalu masukkan tautannya di bawah ini.</AlertDescription></Alert>
            <div class="field"><Label for="design_drive_link">Link Google Drive Video Final *</Label><Input id="design_drive_link" name="design_drive_link" type="url" placeholder="https://drive.google.com/..." required /><p>Video tidak diunggah ke ImageKit agar file besar lebih mudah dikelola.</p></div>
            <div class="rounded-xl border border-blue-200 bg-white p-4">
              <label class="flex items-center gap-2 text-sm"><input type="checkbox" bind:checked={noSong} onchange={saveDraft} /> Tidak menggunakan lagu</label>
              {#if !noSong}<div class="mt-4 grid gap-4 sm:grid-cols-3"><div class="field"><Label for="song_title">Judul Lagu *</Label><Input id="song_title" name="song_title" required /></div><div class="field"><Label for="song_artist">Penyanyi *</Label><Input id="song_artist" name="song_artist" required /></div><div class="field"><Label for="add_song">Bagian Lagu *</Label><Input id="add_song" name="add_song" placeholder="00:15–00:45" required /></div></div>{/if}
            </div>
          {:else}
            <div class="field">
              <Label>File Gambar Final *</Label>
              <div class="dropzone">
                {#if media}
                  <div class="flex items-center gap-4"><img src={media.thumbnail_url || media.url} alt="Preview materi" class="size-20 rounded-lg object-cover" /><div class="min-w-0 flex-1"><p class="truncate font-semibold">{media.name}</p><p>{Math.ceil(media.size / 1024 / 1024)} MB · upload selesai</p></div><Button type="button" variant="ghost" size="icon" class="text-red-700" onclick={() => { media = null; saveDraft(); }}><Trash2 /></Button></div>
                {:else}
                  <Image class="mx-auto size-8 text-blue-600" /><p class="mt-2 font-semibold text-black-700">Pilih gambar final berkualitas tinggi</p><p>JPG, PNG, atau WebP · maksimum 20 MB</p>
                  <label class="mt-3 inline-flex cursor-pointer rounded-lg bg-blue-500 px-4 py-2 text-sm font-semibold text-white hover:bg-blue-600"><input class="sr-only" type="file" accept="image/jpeg,image/png,image/webp" disabled={mediaUploading} onchange={(event) => upload(event.currentTarget.files?.[0], 'submission_media')} required />Pilih gambar</label>
                  {#if mediaUploading}<p class="mt-2 text-blue-700">Mengunggah {mediaProgress}%</p>{/if}
                {/if}
              </div>
            </div>
          {/if}
          <div class="field"><Label for="caption">Caption Instagram *</Label><Textarea id="caption" name="caption" rows={8} maxlength={2200} required oninput={(event) => captionLength = event.currentTarget.value.length} /><p class="text-right">{captionLength}/2200 karakter</p></div>
          <div class="field"><Label for="canva_link">Link Canva {contentFormat === 'REELS_INSTAGRAM' ? '(opsional)' : '*'}</Label><Input id="canva_link" name="canva_link" type="url" placeholder="https://www.canva.com/design/..." required={contentFormat !== 'REELS_INSTAGRAM'} /><p>Pastikan akses “Anyone with the link can edit”.</p></div>
          <div class="field"><Label>Upload Brief Konten *</Label><div class="dropzone">{#if brief}<div class="flex items-center gap-3"><FileText class="size-8 text-blue-600" /><div class="min-w-0 flex-1"><p class="truncate font-semibold">{brief.name}</p><p>Upload selesai</p></div><Button type="button" variant="ghost" size="icon" class="text-red-700" onclick={() => { brief = null; saveDraft(); }}><Trash2 /></Button></div>{:else}<FileUp class="mx-auto size-8 text-blue-600" /><p class="mt-2 font-semibold text-black-700">Pilih brief yang telah dilengkapi</p><p>PDF, DOC, atau DOCX · maksimum 15 MB</p><label class="mt-3 inline-flex cursor-pointer rounded-lg bg-blue-500 px-4 py-2 text-sm font-semibold text-white hover:bg-blue-600"><input class="sr-only" type="file" accept=".pdf,.doc,.docx" disabled={briefUploading} onchange={(event) => upload(event.currentTarget.files?.[0], 'submission_brief')} required />Pilih brief</label>{#if briefUploading}<p class="mt-2 text-blue-700">Mengunggah {briefProgress}%</p>{/if}{/if}</div></div>
        {:else}
          <div class="field"><Label for="documentation_drive_link">Link Drive Dokumentasi Acara *</Label><Input id="documentation_drive_link" name="documentation_drive_link" type="url" placeholder="https://drive.google.com/drive/folders/..." required /><p>Maksimal empat foto terbaik. Pastikan akses “Anyone with the link can view”.</p></div>
          <div class="field"><Label for="article_drive_link">Link Google Docs Isi Artikel *</Label><Input id="article_drive_link" name="article_drive_link" type="url" placeholder="https://docs.google.com/document/..." required /><p>Naskah 300–500 kata, memuat #CeritaHariIni dan struktur JELITA, serta dapat diakses tim Medinfo.</p></div>
          <div class="field"><Label for="required_information">Informasi yang Wajib Dicantumkan</Label><Textarea id="required_information" name="required_information" rows={5} placeholder="Narasumber, jabatan, kutipan, sponsor, mitra, atau istilah khusus..." /></div>
          <div class="field"><Label for="article_caption">Caption Instagram *</Label><Textarea id="article_caption" name="caption" rows={8} maxlength={2200} required oninput={(event) => captionLength = event.currentTarget.value.length} /><p class="text-right">{captionLength}/2200 karakter</p></div>
        {/if}
        <div class="field"><Label for="additional_notes">Keterangan Tambahan <span class="font-normal text-muted-foreground">(opsional)</span></Label><Textarea id="additional_notes" name="additional_notes" rows={5} placeholder="Informasi khusus, akun kolaborator, kebutuhan mendesak, atau catatan lain..." /></div>
      </div>
    </section>

    <section class="form-section">
      <h2>4. Pernyataan Pengaju</h2>
      <div class="grid gap-3">
        {#each [
          'Saya telah membaca seluruh SOP Kementerian Media dan Informasi BEM UNAIR 2026.',
          'Saya memastikan seluruh informasi benar dan dapat dipertanggungjawabkan.',
          'Saya memahami jadwal publikasi menyesuaikan kebijakan dan antrean Medinfo.',
          service === 'CONTENT' ? 'Saya telah mengisi brief konten dengan lengkap dan jelas.' : 'Saya memastikan naskah artikel di Google Docs telah diisi dengan lengkap dan jelas.'
        ] as text, index}
          <label class="flex cursor-pointer items-start gap-3 rounded-xl border border-blue-100 bg-white p-4 text-sm leading-6"><input class="mt-1 size-5 accent-blue-500" type="checkbox" bind:checked={agreements[index]} onchange={saveDraft} required /><span>{text}</span></label>
        {/each}
      </div>
    </section>

    <div class="flex flex-col-reverse justify-between gap-3 sm:flex-row">
      <Button href="/admin/content-submissions/new/select" variant="outline"><ArrowLeft /> Kembali</Button>
      <Button type="submit" class="bg-blue-500 hover:bg-blue-600" disabled={!allAccepted || mediaUploading || briefUploading || (service === 'CONTENT' && ((contentFormat !== 'REELS_INSTAGRAM' && !media) || !brief)) || submitting}>
        {#if submitting}<LoaderCircle class="animate-spin" /> Mengirim...{:else}<Check /> Periksa dan Kirim Pengajuan{/if}
      </Button>
    </div>
  </form>
</div>

<Dialog.Root bind:open={reviewOpen}>
  <Dialog.Content class="max-h-[90vh] overflow-y-auto sm:max-w-xl">
    <Dialog.Header><Dialog.Title>Periksa Pengajuan</Dialog.Title><Dialog.Description>Pastikan informasi berikut sudah benar sebelum dikirim.</Dialog.Description></Dialog.Header>
    <dl class="grid gap-3 rounded-xl bg-blue-50 p-4 text-sm">
      <div><dt class="font-semibold text-muted-foreground">Pengaju</dt><dd>{user.name} · {user.ministry || 'Tanpa kementerian'}</dd></div>
      <div><dt class="font-semibold text-muted-foreground">Jenis</dt><dd>{service === 'ARTICLE' ? 'Artikel' : contentFormat.replaceAll('_', ' ')}</dd></div>
      <div><dt class="font-semibold text-muted-foreground">Judul</dt><dd>{reviewData.title || '—'}</dd></div>
      <div><dt class="font-semibold text-muted-foreground">Jadwal usulan</dt><dd>{reviewData.publish_date || '—'} · {reviewData.publish_time || '—'} WIB</dd></div>
      {#if service === 'CONTENT'}<div><dt class="font-semibold text-muted-foreground">Materi</dt><dd>{contentFormat === 'REELS_INSTAGRAM' ? reviewData.design_drive_link : media?.name || '—'}</dd></div><div><dt class="font-semibold text-muted-foreground">Brief</dt><dd>{brief?.name || '—'}</dd></div>{:else}<div><dt class="font-semibold text-muted-foreground">Naskah</dt><dd class="break-all">{reviewData.article_drive_link || '—'}</dd></div>{/if}
    </dl>
    <Dialog.Footer><Button type="button" variant="outline" onclick={() => reviewOpen = false}>Kembali Edit</Button><Button type="button" class="bg-blue-500 hover:bg-blue-600" onclick={submitConfirmed}>Kirim Pengajuan</Button></Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>

<style>
  .submission-title { margin: 0; font-weight: 800; line-height: 1.08; letter-spacing: -0.02em; }
  .submission-title span { display: block; font-size: clamp(2rem, 5vw, 2.75rem); text-shadow: 0 3px 0 var(--white-600), 1px 4px 3px rgb(0 0 0 / 0.14); }
  .submission-panel { border: 1px solid var(--blue-200); border-radius: 1rem; background: linear-gradient(180deg, var(--blue-100), var(--blue-50)); padding: clamp(1rem, 4vw, 2.5rem); box-shadow: 0 18px 40px rgb(10 36 69 / 0.08); }
  .form-section { display: grid; gap: 1.25rem; border-bottom: 1px solid rgb(59 116 179 / 0.18); padding-bottom: 2rem; }
  .form-section:last-of-type { border-bottom: 0; }
  .form-section h2 { font-size: 1.125rem; font-weight: 750; color: var(--blue-900); }
  .field { display: grid; gap: 0.5rem; }
  .field :global(label) { font-weight: 700; color: var(--black-700); }
  .field > p, .dropzone p { font-size: 0.75rem; color: var(--black-300); }
  .field :global(input:not([type='radio']):not([type='checkbox'])), .field :global(textarea) { background: var(--white-50); }
  .dropzone { border: 1px dashed var(--blue-300); border-radius: 0.75rem; background: var(--white-50); padding: 1.25rem; text-align: center; }
</style>
