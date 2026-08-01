<script lang="ts">
  import { resolve } from '$app/paths';
  import { paginationPages } from '$lib/hooks/use-pagination.svelte';
  let { page, totalPages, basePath }: { page: number; totalPages: number; basePath: `/kementerian/${string}/program-kerja` } = $props();
</script>

{#if totalPages > 1}
  <nav aria-label="Pagination program kerja" class="mt-10 flex flex-wrap items-center justify-center gap-2">
    {#each paginationPages(page, totalPages) as item (item)}
      <a href={resolve((`${basePath}?page=${item}`) as `/kementerian/${string}/program-kerja`)} aria-current={item === page ? 'page' : undefined} class={`grid size-10 place-items-center rounded-xl text-sm font-bold transition ${item === page ? 'bg-blue-800 text-white' : 'border border-blue-100 bg-white text-blue-800 hover:bg-blue-50'}`}>{item}</a>
    {/each}
  </nav>
{/if}
