<script lang="ts">
  import { enhance } from '$app/forms';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import { uploadImageFile, type UploadedImage } from '$lib/admin/upload-image';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import type { LetterTemplate } from '$lib/types';
  import { Download, FileText, LoaderCircle, Trash2, UploadCloud } from '@lucide/svelte';
  import { useFormSubmit } from '$lib/hooks/use-form-submit.svelte';

  let { data, form } = $props<{ data: { templates: LetterTemplate[] }; form?: { error?: string } }>();
  let templateInput: HTMLInputElement;
  let uploaded = $state<UploadedImage | null>(null);
  let uploading = $state(false);
  let uploadError = $state('');
  let altText = $state('');
  const done = () => useFormSubmit('Template surat diperbarui');

  async function uploadTemplate(event: Event) {
    const input = event.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;
    uploading = true;
    uploadError = '';
    try {
      uploaded = await uploadImageFile(file, 'letter_template');
      altText = file.name;
    } catch (error) {
      uploaded = null;
      uploadError = error instanceof Error ? error.message : 'Gagal mengunggah template PDF.';
    } finally {
      uploading = false;
      input.value = '';
    }
  }

  function clearUpload() {
    uploaded = null;
    altText = '';
    uploadError = '';
  }
</script>

<PageHeader title="Template Surat" description="Admin Medinfo mengunggah template surat dalam format PDF untuk digunakan oleh admin kementerian." />
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
<div class="grid gap-6 lg:grid-cols-[360px_1fr]">
  <section class="rounded-xl border bg-card p-6">
    <div class="flex items-center gap-3"><div class="grid size-10 place-items-center rounded-xl bg-blue-50 text-blue-700"><UploadCloud class="size-5" /></div><div><h2 class="font-bold text-blue-900">Upload template PDF</h2><p class="text-xs text-muted-foreground">Maksimal 10 MB</p></div></div>
    <form method="POST" action="?/create" use:enhance={done()} class="mt-6 space-y-4">
      <input type="hidden" name="file_id" value={uploaded?.file_id || ''} />
      <input type="hidden" name="file_path" value={uploaded?.file_path || ''} />
      <input type="hidden" name="url" value={uploaded?.url || ''} />
      <input type="hidden" name="file_name" value={uploaded?.name || ''} />
      <input type="hidden" name="mime_type" value={uploaded?.file_type || ''} />
      <input type="hidden" name="size_bytes" value={uploaded?.size || 0} />
      <div class="space-y-2"><Label for="name">Nama template</Label><Input id="name" name="name" placeholder="Undangan kegiatan resmi" required /></div>
      <div class="space-y-2"><Label for="type">Jenis surat</Label><Input id="type" name="type" placeholder="UNDANGAN" required /></div>
      <div class="space-y-2"><Label for="subject">Perihal default <span class="text-black-200">(opsional)</span></Label><Input id="subject" name="subject" /></div>
      <div class="space-y-2"><Label for="template-upload">File PDF</Label><input bind:this={templateInput} id="template-upload" type="file" accept="application/pdf" class="sr-only" onchange={uploadTemplate} /><div class="rounded-xl border border-dashed border-blue-200 bg-blue-50 p-4">{#if uploaded}<div class="flex items-center gap-3"><FileText class="size-8 text-red-600" /><div class="min-w-0 flex-1"><p class="truncate text-sm font-semibold text-blue-900">{uploaded.name}</p><p class="text-xs text-black-300">Upload selesai</p></div><Button type="button" variant="ghost" size="icon" onclick={clearUpload} aria-label="Hapus file"><Trash2 class="text-red-600" /></Button></div>{:else}<button type="button" class="flex w-full flex-col items-center justify-center gap-1 py-3 text-blue-700" onclick={() => templateInput?.click()} disabled={uploading}>{#if uploading}<LoaderCircle class="size-6 animate-spin" /><span class="text-sm font-semibold">Mengunggah PDF...</span>{:else}<UploadCloud class="size-6" /><span class="text-sm font-semibold">Pilih file PDF</span><span class="text-xs text-black-300">PDF - maksimum 10 MB</span>{/if}</button>{/if}</div>{#if uploadError}<p class="text-sm text-red-600">{uploadError}</p>{/if}</div>
      <div class="space-y-2"><Label for="alt_text">Alt text</Label><Input id="alt_text" name="alt_text" bind:value={altText} required disabled={!uploaded} /></div>
      <Button type="submit" class="w-full bg-blue-500" disabled={!uploaded || uploading}>Simpan template</Button>
    </form>
  </section>
  <section class="space-y-4">
    {#if data.templates.length}{#each data.templates as template (template.id)}<article class="flex flex-col gap-4 rounded-xl border bg-card p-5 sm:flex-row sm:items-center sm:justify-between"><div class="flex min-w-0 items-center gap-3"><div class="grid size-11 shrink-0 place-items-center rounded-xl bg-red-50 text-red-600"><FileText class="size-5" /></div><div class="min-w-0"><h2 class="truncate font-semibold text-blue-900">{template.name}</h2><p class="mt-1 text-xs text-muted-foreground">{template.type} - {template.is_active ? 'Aktif' : 'Nonaktif'}</p>{#if !template.file && !template.download_url}<p class="mt-1 text-xs text-amber-700">Belum memiliki file PDF.</p>{/if}</div></div><div class="flex shrink-0 flex-wrap gap-2">{#if template.download_url || template.file?.url}<a href={template.download_url || template.file?.url} target="_blank" rel="external noreferrer" class="inline-flex h-9 items-center gap-2 rounded-lg border px-3 text-sm font-semibold text-blue-700 hover:bg-blue-50"><Download class="size-4" />Unduh PDF</a>{/if}<form method="POST" action="?/delete" use:enhance={done()} onsubmit={(event) => { if (!confirm('Hapus template surat ini?')) event.preventDefault(); }}><input type="hidden" name="id" value={template.id} /><Button type="submit" variant="ghost" size="icon" class="text-red-600" aria-label="Hapus template"><Trash2 class="size-4" /></Button></form></div></article>{/each}{:else}<EmptyState title="Belum ada template surat" description="Upload PDF pertama untuk dibagikan kepada admin kementerian." />{/if}
  </section>
</div>
