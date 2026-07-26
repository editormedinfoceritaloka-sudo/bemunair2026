<script lang="ts">
  import type { MediaSubmissionSetting, User } from '$lib/types';
  import { uploadSubmissionFile, type UploadedSubmissionFile } from '$lib/admin/upload-submission-file';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { ArrowLeft, Check, FileUp, Image, LoaderCircle, Trash2 } from '@lucide/svelte';

  let { service, user, setting, form }: {
    service: 'CONTENT' | 'ARTICLE';
    user: User;
    setting: MediaSubmissionSetting;
    form?: { error?: string };
  } = $props();

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
  let submitting = $state(false);

  const allAccepted = $derived(agreements.every(Boolean));
  const minimumDays = setting.minimum_lead_days ?? setting.MinimumLeadDays ?? (service === 'CONTENT' ? 7 : 3);
  const minimumDate = $derived.by(() => {
    const date = new Date();
    date.setDate(date.getDate() + minimumDays);
    return date.toISOString().slice(0, 10);
  });
  const startTime = setting.publish_time_start ?? setting.PublishTimeStart ?? '08:00';
  const endTime = setting.publish_time_end ?? setting.PublishTimeEnd ?? '17:00';

  async function upload(file: File | undefined, purpose: 'submission_media' | 'submission_brief') {
    if (!file) return;
    uploadError = '';
    if (purpose === 'submission_media') mediaUploading = true; else briefUploading = true;
    try {
      const result = await uploadSubmissionFile(
        file,
        purpose,
        (value) => purpose === 'submission_media' ? mediaProgress = value : briefProgress = value
      );
      if (purpose === 'submission_media') media = result; else brief = result;
    } catch (error) {
      uploadError = error instanceof Error ? error.message : 'Upload gagal.';
    } finally {
      if (purpose === 'submission_media') mediaUploading = false; else briefUploading = false;
    }
  }
</script>

{#if form?.error}<Alert variant="destructive" class="mb-5"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
{#if uploadError}<Alert variant="destructive" class="mb-5"><AlertDescription>{uploadError}</AlertDescription></Alert>{/if}

<form
  method="POST"
  enctype="multipart/form-data"
  class="max-w-5xl space-y-6"
  onsubmit={() => submitting = true}
>
  <input type="hidden" name="service_type" value={service} />
  <input type="hidden" name="submission_type" value={service === 'ARTICLE' ? 'ARTIKEL' : contentFormat} />
  <input type="hidden" name="content_format" value={service === 'CONTENT' ? contentFormat : ''} />
  <input type="hidden" name="ministry" value={user.ministry || ''} />
  <input type="hidden" name="design_drive_link" value={media?.url || ''} />
  <input type="hidden" name="brief_link" value={service === 'CONTENT' ? brief?.url || '' : '-'} />

  <Card class={service === 'CONTENT' ? 'border-blue-100 bg-blue-50/30' : 'border-orange-100 bg-orange-50/30'}>
    <CardHeader><CardTitle>1. Data pengaju</CardTitle></CardHeader>
    <CardContent class="grid gap-4 sm:grid-cols-2">
      <div class="space-y-2"><Label for="submitter_name">Nama lengkap</Label><Input id="submitter_name" value={user.name} disabled /><p class="text-xs text-muted-foreground">Mengikuti akun yang sedang login.</p></div>
      <div class="space-y-2"><Label for="ministry_view">Kementerian</Label><Input id="ministry_view" value={user.ministry || 'Belum ditentukan'} disabled /></div>
      <div class="space-y-2 sm:col-span-2"><Label for="submitter_phone">Nomor WhatsApp</Label><Input id="submitter_phone" name="submitter_phone" type="tel" value={user.phone || ''} placeholder="081234567890" pattern="[+]?[0-9]{10,15}" required /></div>
    </CardContent>
  </Card>

  <Card>
    <CardHeader><CardTitle>2. Detail publikasi</CardTitle></CardHeader>
    <CardContent class="grid gap-4 sm:grid-cols-2">
      {#if service === 'CONTENT'}
        <div class="space-y-2 sm:col-span-2">
          <Label for="content_format">Jenis pengajuan</Label>
          <div class="grid gap-3 sm:grid-cols-3">
            {#each [['FEED_INSTAGRAM', 'Feed Instagram'], ['REELS_INSTAGRAM', 'Reels Instagram'], ['INSTASTORY', 'Instastory']] as option}
              <label class="flex cursor-pointer items-center gap-3 rounded-xl border bg-card p-4 has-[:checked]:border-blue-400 has-[:checked]:bg-blue-50">
                <input type="radio" bind:group={contentFormat} value={option[0]} /><span class="text-sm font-semibold">{option[1]}</span>
              </label>
            {/each}
          </div>
        </div>
      {/if}
      <div class="space-y-2"><Label for="publish_date">Tanggal publikasi</Label><Input id="publish_date" name="publish_date" type="date" min={minimumDate} required /><p class="text-xs text-muted-foreground">Tanggal minimal {minimumDays} hari dari hari ini.</p></div>
      <div class="space-y-2"><Label for="publish_time">Waktu publikasi</Label><Input id="publish_time" name="publish_time" type="time" min={startTime} max={endTime} required /><p class="text-xs text-muted-foreground">Usulan waktu {startTime}–{endTime}; dapat disesuaikan Medinfo.</p></div>
      <div class="space-y-2 sm:col-span-2"><Label for="title">{service === 'ARTICLE' ? 'Judul artikel atau nama kegiatan' : 'Judul konten atau nama kegiatan'}</Label><Input id="title" name="title" maxlength={180} required /></div>
    </CardContent>
  </Card>

  <Card>
    <CardHeader><CardTitle>3. Materi publikasi</CardTitle></CardHeader>
    <CardContent class="grid gap-5">
      {#if service === 'CONTENT'}
        {#if contentFormat === 'REELS_INSTAGRAM'}
          <div class="rounded-xl border p-4">
            <label class="flex items-center gap-2 text-sm"><input type="checkbox" bind:checked={noSong} /> Tidak menggunakan lagu</label>
            {#if !noSong}<div class="mt-4 grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label for="song_title">Judul lagu</Label><Input id="song_title" name="song_title" required /></div><div class="space-y-2"><Label for="song_artist">Penyanyi</Label><Input id="song_artist" name="song_artist" required /></div><div class="space-y-2"><Label for="add_song">Bagian lagu</Label><Input id="add_song" name="add_song" placeholder="00:15–00:45" required /></div></div>{/if}
          </div>
        {/if}
        <div class="space-y-2"><Label for="caption">Caption Instagram</Label><Textarea id="caption" name="caption" rows={7} maxlength={2200} required /><p class="text-xs text-muted-foreground">Paragraf, emoji, mention, dan hashtag akan dipertahankan.</p></div>
        <div class="space-y-2"><Label for="canva_link">Link Canva <span class="text-muted-foreground">(opsional untuk video final)</span></Label><Input id="canva_link" name="canva_link" type="url" placeholder="https://www.canva.com/..." /></div>

        <div class="grid gap-4 sm:grid-cols-2">
          <div class="rounded-xl border border-dashed bg-muted/20 p-5">
            <div class="mb-3 flex items-center gap-2 font-semibold"><Image class="size-5 text-blue-600" />File final media</div>
            {#if media}
              <div class="rounded-lg bg-card p-3"><p class="truncate text-sm font-medium">{media.name}</p><p class="text-xs text-muted-foreground">{Math.ceil(media.size / 1024 / 1024)} MB · upload selesai</p><Button type="button" variant="ghost" size="sm" class="mt-2 text-red-600" onclick={() => media = null}><Trash2 /> Hapus</Button></div>
            {:else}
              <Input type="file" accept="image/jpeg,image/png,image/webp,video/mp4,video/quicktime" disabled={mediaUploading} onchange={(event) => upload(event.currentTarget.files?.[0], 'submission_media')} required />
              <p class="mt-2 text-xs text-muted-foreground">JPG/PNG/WebP/MP4/MOV, maksimal 100 MB.</p>
              {#if mediaUploading}<p class="mt-2 text-xs text-blue-700">Mengunggah {mediaProgress}%</p>{/if}
            {/if}
          </div>
          <div class="rounded-xl border border-dashed bg-muted/20 p-5">
            <div class="mb-3 flex items-center gap-2 font-semibold"><FileUp class="size-5 text-orange-600" />Brief konten</div>
            {#if brief}
              <div class="rounded-lg bg-card p-3"><p class="truncate text-sm font-medium">{brief.name}</p><p class="text-xs text-muted-foreground">Upload selesai</p><Button type="button" variant="ghost" size="sm" class="mt-2 text-red-600" onclick={() => brief = null}><Trash2 /> Hapus</Button></div>
            {:else}
              <Input type="file" accept=".pdf,.doc,.docx" disabled={briefUploading} onchange={(event) => upload(event.currentTarget.files?.[0], 'submission_brief')} required />
              <p class="mt-2 text-xs text-muted-foreground">PDF/DOC/DOCX, maksimal 15 MB.</p>
              {#if briefUploading}<p class="mt-2 text-xs text-blue-700">Mengunggah {briefProgress}%</p>{/if}
            {/if}
          </div>
        </div>
      {:else}
        <input type="hidden" name="caption" value="Pengajuan artikel" />
        <div class="space-y-2"><Label for="documentation_drive_link">Link Drive dokumentasi acara</Label><Input id="documentation_drive_link" name="documentation_drive_link" type="url" placeholder="https://drive.google.com/..." required /><p class="text-xs text-muted-foreground">Maksimal empat foto dan pastikan tautan dapat dibuka.</p></div>
        <div class="space-y-2"><Label for="article_drive_link">Link Google Docs isi artikel</Label><Input id="article_drive_link" name="article_drive_link" type="url" placeholder="https://docs.google.com/..." required /><p class="text-xs text-muted-foreground">Naskah 300–500 kata dan dapat diakses tim Medinfo.</p></div>
        <div class="space-y-2"><Label for="required_information">Informasi yang wajib dicantumkan</Label><Textarea id="required_information" name="required_information" rows={4} placeholder="Narasumber, jabatan, kutipan, sponsor, data peserta..." /></div>
        <div class="space-y-2"><Label for="article_caption">Caption Instagram</Label><Textarea id="article_caption" name="caption" rows={6} maxlength={2200} required /></div>
      {/if}
      <div class="space-y-2"><Label for="additional_notes">Keterangan tambahan <span class="text-muted-foreground">(opsional)</span></Label><Textarea id="additional_notes" name="additional_notes" rows={4} /></div>
    </CardContent>
  </Card>

  <Card>
    <CardHeader><CardTitle>4. Pernyataan pengaju</CardTitle></CardHeader>
    <CardContent class="space-y-3">
      {#each [
        'Saya telah membaca seluruh SOP Kementerian Media dan Informasi BEM UNAIR 2026.',
        'Saya memastikan seluruh informasi benar dan dapat dipertanggungjawabkan.',
        'Saya memahami jadwal publikasi menyesuaikan kebijakan dan antrean Medinfo.',
        service === 'CONTENT' ? 'Saya telah mengisi brief konten dengan lengkap dan jelas.' : 'Saya memastikan naskah artikel lengkap, jelas, dan sesuai ketentuan penulisan.'
      ] as text, index}
        <label class="flex cursor-pointer items-start gap-3 rounded-lg border p-3 text-sm"><input type="checkbox" bind:checked={agreements[index]} required /><span>{text}</span></label>
      {/each}
    </CardContent>
  </Card>

  <div class="flex flex-col-reverse justify-between gap-3 sm:flex-row">
    <Button href={`/admin/content-submissions/new/${service.toLowerCase()}`} variant="outline"><ArrowLeft /> Kembali</Button>
    <Button type="submit" class={service === 'CONTENT' ? 'bg-blue-500' : 'bg-orange-500 hover:bg-orange-600'} disabled={!allAccepted || mediaUploading || briefUploading || (service === 'CONTENT' && (!media || !brief)) || submitting}>
      {#if submitting}<LoaderCircle class="animate-spin" /> Mengirim...{:else}<Check /> Periksa dan kirim pengajuan{/if}
    </Button>
  </div>
</form>
