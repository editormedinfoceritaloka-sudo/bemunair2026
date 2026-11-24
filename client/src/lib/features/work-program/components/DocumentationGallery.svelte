<script lang="ts">
  import type { Documentation } from '$lib/types';
  import type { Snippet } from 'svelte';

  let { items = [], actions }: { items?: Documentation[]; actions?: Snippet<[Documentation]> } = $props();
</script>

{#if items.length}
  <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
    {#each items as item (item.id)}
      <div class="min-w-0 overflow-hidden rounded-2xl border bg-card">
        <figure>
          {#if item.media?.url}
            <div class="relative aspect-[4/3] bg-blue-50">
              <img loading="lazy" src={item.media.url} alt={item.media.alt_text} class="absolute inset-0 size-full object-contain" />
            </div>
          {:else}
            <p class="p-3 text-sm text-muted-foreground">Foto tidak tersedia.</p>
          {/if}
          {#if item.caption}<figcaption class="p-3 text-sm text-blue-900">{item.caption}</figcaption>{/if}
        </figure>
        {#if actions}<div class="p-3">{@render actions(item)}</div>{/if}
      </div>
    {/each}
  </div>
{:else}
  <div class="rounded-[28px] bg-gradient-to-b from-blue-700 via-blue-800 to-blue-900 p-5 shadow-xl sm:p-8"><div class="grid min-h-44 place-items-center rounded-2xl bg-white text-center"><div><p class="text-3xl font-black text-blue-800 sm:text-5xl">Coming Soon</p><p class="mt-2 text-sm text-slate-500">Dokumentasi kegiatan belum tersedia.</p></div></div></div>
{/if}
