<script lang="ts">
  import type { Documentation } from '$lib/types';
  let { items = [] }: { items?: Documentation[] } = $props();
</script>

{#if items.length}
  <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
    {#each items as item (item.id)}
      {#if item.media?.url}<figure class="overflow-hidden rounded-2xl bg-blue-100"><img loading="lazy" src={item.media.url} alt={item.media.alt_text} class="aspect-[4/3] size-full object-cover" />{#if item.caption}<figcaption class="p-3 text-sm text-blue-900">{item.caption}</figcaption>{/if}</figure>{/if}
    {/each}
  </div>
{:else}
  <div class="rounded-[28px] bg-gradient-to-b from-blue-700 via-blue-800 to-blue-900 p-5 shadow-xl sm:p-8"><div class="grid min-h-44 place-items-center rounded-2xl bg-white text-center"><div><p class="text-3xl font-black text-blue-800 sm:text-5xl">Coming Soon</p><p class="mt-2 text-sm text-slate-500">Dokumentasi kegiatan belum tersedia.</p></div></div></div>
{/if}
