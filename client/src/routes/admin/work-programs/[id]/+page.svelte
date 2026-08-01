<script lang="ts">
  import { uploadImageFile, type UploadedImage } from "$lib/admin/upload-image";
  import { enhance } from '$app/forms';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Label } from '$lib/components/ui/label';
  import WorkProgramTimeline from '$lib/features/work-program/components/WorkProgramTimeline.svelte';
  import DocumentationGallery from '$lib/features/work-program/components/DocumentationGallery.svelte';
  import type { WorkProgram } from '$lib/types';
  import { LoaderCircle, Trash2, UploadCloud } from "@lucide/svelte";
  import { useFormSubmit } from '$lib/hooks/use-form-submit.svelte';

  let { data, form } = $props<{ data: { program: WorkProgram }; form?: { error?: string } }>();
  const done = () => useFormSubmit('Data program diperbarui');
  let documentationInput: HTMLInputElement;
  let uploaded = $state<UploadedImage | null>(null);
  let uploading = $state(false);
  let uploadError = $state("");
  let altText = $state("");

  async function uploadDocumentation(event: Event) {
    const input = event.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;
    uploading = true;
    uploadError = "";
    try {
      uploaded = await uploadImageFile(file, "documentation");
      altText = file.name.replace(/\.[^/.]+$/, "").replace(/[-_]+/g, " ");
    } catch (error) {
      uploaded = null;
      uploadError = error instanceof Error ? error.message : "Gagal mengunggah dokumentasi";
    } finally {
      uploading = false;
      input.value = "";
    }
  }

  function clearUpload() {
    uploaded = null;
    altText = "";
    uploadError = "";
  }
</script>

<PageHeader title={data.program.name} description="Kelola timeline dan dokumentasi program kerja."><Button href="/admin/work-programs" variant="outline">Kembali</Button></PageHeader>
{#if form?.error}<p class="mb-5 rounded-xl bg-red-50 p-4 text-sm text-red-700">{form.error}</p>{/if}
<div class="grid gap-6 lg:grid-cols-2"><section class="rounded-2xl border bg-card p-6"><h2 class="text-xl font-black text-blue-900">Timeline</h2><div class="mt-6"><WorkProgramTimeline milestones={data.program.milestones} /></div><form method="POST" action="?/milestone" use:enhance={done()} class="mt-8 space-y-4 border-t pt-6"><div class="space-y-2"><Label for="title">Judul milestone</Label><Input id="title" name="title" required /></div><div class="space-y-2"><Label for="description">Deskripsi</Label><Textarea id="description" name="description" rows={3} /></div><div class="space-y-2"><Label for="status">Status</Label><Input id="status" name="status" value="PLANNED" /></div><Button type="submit">Tambah milestone</Button></form></section><section class="rounded-2xl border bg-card p-6"><h2 class="text-xl font-black text-blue-900">Dokumentasi</h2><div class="mt-6"><DocumentationGallery items={data.program.documentations} /></div><form method="POST" action="?/documentation" use:enhance={done()} class="mt-8 space-y-4 border-t pt-6"><input type="hidden" name="file_id" value={uploaded?.file_id || ""} required /><input type="hidden" name="file_path" value={uploaded?.file_path || ""} /><input type="hidden" name="url" value={uploaded?.url || ""} required /><input type="hidden" name="thumbnail_url" value={uploaded?.thumbnail_url || ""} /><input type="hidden" name="name" value={uploaded?.name || ""} required /><input type="hidden" name="mime_type" value={uploaded?.file_type || ""} required /><input type="hidden" name="size_bytes" value={uploaded?.size || 0} required /><input type="hidden" name="width" value={uploaded?.width || ""} /><input type="hidden" name="height" value={uploaded?.height || ""} /><div class="space-y-2"><Label for="documentation-upload">File dokumentasi</Label><input bind:this={documentationInput} id="documentation-upload" type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={uploadDocumentation} /><div class="flex items-center gap-3 rounded-xl border border-dashed border-blue-200 bg-blue-50 p-4">{#if uploaded}<img src={uploaded.url} alt="Preview dokumentasi" class="size-20 rounded-lg object-cover" /><div class="min-w-0 flex-1"><p class="truncate text-sm font-semibold text-blue-900">{uploaded.name}</p><p class="text-xs text-black-300">Upload selesai</p></div><Button type="button" variant="ghost" size="icon" onclick={clearUpload} aria-label="Hapus file"><Trash2 class="text-red-600" /></Button>{:else}<button type="button" class="flex w-full flex-col items-center justify-center gap-1 py-3 text-blue-700" onclick={() => documentationInput?.click()} disabled={uploading}>{#if uploading}<LoaderCircle class="size-6 animate-spin" /><span class="text-sm font-semibold">Mengunggah…</span>{:else}<UploadCloud class="size-6" /><span class="text-sm font-semibold">Pilih gambar dokumentasi</span><span class="text-xs text-black-300">JPEG, PNG, WebP, GIF, atau AVIF · maksimum 10 MB</span>{/if}</button>{/if}</div>{#if uploadError}<p class="text-sm text-red-600">{uploadError}</p>{/if}</div><div class="space-y-2"><Label for="alt_text">Alt text</Label><Input id="alt_text" name="alt_text" bind:value={altText} required disabled={!uploaded} /></div><div class="space-y-2"><Label for="doc_title">Judul</Label><Input id="doc_title" name="title" /></div><div class="space-y-2"><Label for="caption">Caption</Label><Textarea id="caption" name="caption" rows={3} /></div><label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_cover" /> Jadikan cover</label><Button type="submit" disabled={!uploaded || uploading}>Tambah dokumentasi</Button></form></section></div>