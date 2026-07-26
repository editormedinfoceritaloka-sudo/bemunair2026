<script lang="ts">
  import { onDestroy, untrack } from 'svelte';
  import { goto } from '$app/navigation';
  import NotionEditor from './NotionEditor.svelte';
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { uploadImageFile } from '$lib/admin/upload-image';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Label } from '$lib/components/ui/label';
  import { Save, Eye, Send, ArrowLeft, Check, LoaderCircle, AlertCircle, ImagePlus, Trash2, UploadCloud } from '@lucide/svelte';
  import { toast } from 'svelte-sonner';
  import type { Article } from '$lib/types';

  let { article = null, error = '' }: { article?: Article | null; error?: string } = $props();
  const initial = untrack(() => article);
  let title = $state(initial?.title || '');
  let excerpt = $state(initial?.excerpt || '');
  let cover = $state(initial?.cover_image || '');
  let body = $state(initial?.body || '<p></p>');
  let saveState = $state<'idle' | 'dirty' | 'saving' | 'saved' | 'error'>('idle');
  let coverUploading = $state(false);
  let coverInput: HTMLInputElement;
  let initialized = $state(false);
  let timer: ReturnType<typeof setTimeout> | undefined;

  const isEdit = Boolean(initial);
  const endpoint = initial ? `/admin/articles/${initial.id}/edit/save` : '';

  function payload() {
    return {
      title: title.trim(),
      excerpt: excerpt.trim() || null,
      cover_image: cover.trim() || null,
      body
    };
  }

  function valid() {
    if (coverUploading) {
      toast.error('Tunggu upload cover selesai');
      return false;
    }
    if (!title.trim()) {
      toast.error('Judul artikel wajib diisi');
      return false;
    }
    if (!body || body === '<p></p>') {
      toast.error('Isi artikel belum ditulis');
      return false;
    }
    if (excerpt.length > 500) {
      toast.error('Ringkasan maksimal 500 karakter');
      return false;
    }
    return true;
  }

  async function uploadCover(event: Event) {
    const input = event.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;
    coverUploading = true;
    try {
      const uploaded = await uploadImageFile(file, 'cover');
      cover = uploaded.url;
      toast.success('Cover berhasil diunggah');
    } catch (uploadError) {
      toast.error(uploadError instanceof Error ? uploadError.message : 'Gagal mengunggah cover');
    } finally {
      coverUploading = false;
      input.value = '';
    }
  }

  async function save(showToast = false) {
    if (!isEdit || !valid()) return false;
    if (timer) clearTimeout(timer);
    saveState = 'saving';
    try {
      const response = await fetch(endpoint, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload())
      });
      if (!response.ok) throw new Error((await response.json()).message || 'Gagal menyimpan');
      saveState = 'saved';
      if (showToast) toast.success('Artikel tersimpan');
      return true;
    } catch (saveError) {
      saveState = 'error';
      toast.error(saveError instanceof Error ? saveError.message : 'Gagal menyimpan');
      return false;
    }
  }

  async function publish() {
    if (!(await save())) return;
    const response = await fetch(endpoint, { method: 'POST' });
    if (response.ok) {
      toast.success(article?.status === 'PUBLISHED' ? 'Artikel dikembalikan ke draft' : 'Artikel dipublikasikan');
      await goto('/admin/articles');
    } else {
      toast.error('Status publikasi gagal diperbarui');
    }
  }

  async function preview() {
    if (await save()) window.open(`/admin/articles/${article?.id}/preview`, '_blank');
  }

  $effect(() => {
    title;
    excerpt;
    cover;
    body;
    if (!isEdit) return;
    if (!initialized) {
      initialized = true;
      return;
    }
    saveState = 'dirty';
    if (timer) clearTimeout(timer);
    timer = setTimeout(() => save(), 1500);
  });

  onDestroy(() => {
    if (timer) clearTimeout(timer);
  });
</script>

<div class="mx-auto max-w-5xl pb-20">
  <div class="sticky top-0 z-20 -mx-4 mb-6 flex flex-wrap items-center gap-2 border-b bg-background/95 px-4 py-3 backdrop-blur md:-mx-8 md:px-8">
    <Button href="/admin/articles" variant="ghost" size="sm"><ArrowLeft />Artikel</Button>
    <div class="h-5 border-l"></div>
    {#if article}
      <StatusBadge status={article.status} />
      <span class="ml-auto flex items-center gap-1.5 text-xs text-muted-foreground">
        {#if saveState === 'saving'}<LoaderCircle class="size-3.5 animate-spin" />Menyimpan…
        {:else if saveState === 'saved'}<Check class="size-3.5 text-green-700" />Tersimpan
        {:else if saveState === 'error'}<AlertCircle class="size-3.5 text-red-600" />Gagal menyimpan
        {:else if saveState === 'dirty'}Perubahan belum tersimpan{/if}
      </span>
      <Button type="button" variant="outline" size="sm" onclick={() => save(true)}><Save />Simpan</Button>
      <Button type="button" variant="outline" size="sm" onclick={preview}><Eye />Preview</Button>
      <Button type="button" size="sm" class={article.status === 'PUBLISHED' ? 'bg-black-400' : 'bg-orange-500'} onclick={publish}><Send />{article.status === 'PUBLISHED' ? 'Jadikan draft' : 'Publikasikan'}</Button>
    {:else}
      <span class="ml-auto text-xs text-muted-foreground">Draft baru belum disimpan</span>
      <Button type="submit" form="article-create" class="bg-blue-500" disabled={coverUploading}><Save />Simpan draft</Button>
    {/if}
  </div>

  {#if error}<div class="mb-5 rounded-lg border border-red-200 bg-red-50 p-3 text-sm text-red-700">{error}</div>{/if}

  <form id="article-create" method="POST" class="space-y-6" onsubmit={(event) => { if (!valid()) event.preventDefault(); }}>
    <div class="space-y-5 rounded-xl border bg-card p-5 md:p-8">
      <Input name="title" bind:value={title} maxlength={255} placeholder="Judul artikel" class="h-auto border-0 px-0 font-serif text-4xl font-bold shadow-none focus-visible:ring-0 md:text-5xl" required />
      <div class="grid gap-5 md:grid-cols-[1fr_280px]">
        <div class="space-y-2">
          <Label for="excerpt">Ringkasan</Label>
          <Textarea id="excerpt" name="excerpt" bind:value={excerpt} maxlength={500} rows={5} placeholder="Ringkasan singkat untuk kartu artikel..." />
          <p class="text-right text-xs text-black-200">{excerpt.length}/500</p>
        </div>
        <div class="space-y-2">
          <Label for="cover-upload">Cover artikel</Label>
          <input bind:this={coverInput} id="cover-upload" type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={uploadCover} />
          <input type="hidden" name="cover_image" value={cover} />
          {#if cover}
            <div class="group relative overflow-hidden rounded-xl border border-border bg-muted">
              <img src={cover} alt="Preview cover artikel" class="aspect-[16/9] w-full object-cover" />
              <div class="absolute inset-x-0 bottom-0 flex justify-end gap-1 bg-black-900/65 p-2 opacity-0 transition-opacity group-hover:opacity-100 group-focus-within:opacity-100">
                <Button type="button" size="sm" variant="secondary" onclick={() => coverInput.click()} disabled={coverUploading}><ImagePlus />Ganti</Button>
                <Button type="button" size="icon-sm" variant="destructive" onclick={() => cover = ''} aria-label="Hapus cover"><Trash2 /></Button>
              </div>
            </div>
          {:else}
            <button type="button" onclick={() => coverInput.click()} disabled={coverUploading} class="flex aspect-[16/9] w-full flex-col items-center justify-center gap-2 rounded-xl border border-dashed border-blue-200 bg-blue-50 px-4 text-center text-blue-700 transition-colors hover:border-blue-400 hover:bg-blue-100 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring disabled:opacity-60">
              {#if coverUploading}<LoaderCircle class="size-6 animate-spin" /><span class="text-sm font-medium">Mengunggah cover…</span>
              {:else}<UploadCloud class="size-6" /><span class="text-sm font-semibold">Unggah cover</span><span class="text-xs text-black-300">JPEG, PNG, WebP, GIF, atau AVIF · maks. 10 MB</span>{/if}
            </button>
          {/if}
        </div>
      </div>
    </div>

    <NotionEditor value={body} onchange={(html) => body = html} />
    <input type="hidden" name="body" value={body} />
  </form>
</div>
