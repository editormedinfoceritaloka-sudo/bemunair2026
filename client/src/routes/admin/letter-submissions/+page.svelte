<script lang="ts">
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import type { LetterTemplate } from '$lib/types';
  import { Download, FileText } from '@lucide/svelte';

  let { data } = $props<{ data: { templates: LetterTemplate[] } }>();
</script>

<PageHeader title="Template Surat" description="Unduh template surat PDF yang dibutuhkan untuk keperluan kementerian." />

{#if data.templates.length}
  <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
    {#each data.templates as template (template.id)}
      <article class="flex h-full flex-col rounded-xl border bg-card p-5 shadow-sm">
        <div class="flex items-start gap-3">
          <div class="grid size-11 shrink-0 place-items-center rounded-xl bg-red-50 text-red-600">
            <FileText class="size-5" />
          </div>
          <div class="min-w-0">
            <h2 class="font-semibold text-blue-900">{template.name}</h2>
            <p class="mt-1 text-xs font-semibold uppercase tracking-wide text-muted-foreground">{template.type}</p>
          </div>
        </div>
        {#if template.subject}
          <p class="mt-4 text-sm leading-6 text-muted-foreground">{template.subject}</p>
        {:else}
          <p class="mt-4 text-sm leading-6 text-muted-foreground">Template surat resmi BEM UNAIR.</p>
        {/if}
        <div class="mt-auto pt-5">
          {#if template.download_url || template.file?.url}
            <a href={template.download_url || template.file?.url} target="_blank" rel="external noreferrer" class="inline-flex h-10 w-full items-center justify-center gap-2 rounded-lg bg-blue-500 px-4 text-sm font-semibold text-white! transition-colors hover:bg-blue-600 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500">
              <Download class="size-4" />
              Download template PDF
            </a>
          {:else}
            <p class="rounded-lg bg-amber-50 px-3 py-2 text-center text-sm text-amber-700">File PDF belum tersedia.</p>
          {/if}
        </div>
      </article>
    {/each}
  </div>
{:else}
  <EmptyState title="Belum ada template surat" description="Template surat PDF akan tampil di sini setelah diunggah Admin Medinfo." />
{/if}
