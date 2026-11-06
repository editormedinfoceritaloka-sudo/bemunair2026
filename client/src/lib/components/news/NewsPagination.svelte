<script lang="ts">
  import { ArrowLeftIcon, ArrowRightIcon } from '@lucide/svelte';
  import { resolve } from '$app/paths';

  import type { Meta } from '$lib/types';

  type PaginationItem = number | 'ellipsis';

  let {
    meta,
    search = ''
  }: {
    meta: Meta;
    search?: string;
  } = $props();

  const pages = $derived.by((): PaginationItem[] => {
    const total = meta.total_pages;
    const current = meta.page;

    if (total <= 1) {
      return [];
    }

    const values: number[] = [];

    const add = (page: number): void => {
      if (page >= 1 && page <= total && !values.includes(page)) {
        values.push(page);
      }
    };

    add(1);
    add(total);
    add(current - 1);
    add(current);
    add(current + 1);

    values.sort((a, b) => a - b);

    const result: PaginationItem[] = [];

    for (let index = 0; index < values.length; index++) {
      const currentPage = values[index];
      const previousPage = values[index - 1];

      if (index > 0 && currentPage - previousPage > 1) {
        result.push('ellipsis');
      }

      result.push(currentPage);
    }

    return result;
  });

  function pageQuery(page: number): string {
    const query: string[] = [];

    if (search.trim()) {
      query.push(`search=${encodeURIComponent(search.trim())}`);
    }

    query.push(`page=${page}`);
    query.push(`per_page=${meta.per_page}`);

    return query.join('&');
  }
</script>

{#if meta.total_pages > 1}
  <nav
    aria-label="Pagination berita"
    class="
      mt-16
      flex flex-wrap
      items-center
      justify-center
      gap-2
      sm:mt-20
    "
  >
    {#if meta.page > 1}
      <a
        href={resolve(`/news?${pageQuery(meta.page - 1)}`)}
        aria-label="Halaman sebelumnya"
        class="
          group
          inline-flex
          h-11
          items-center
          justify-center
          gap-2
          rounded-full
          border
          border-blue-900/10
          bg-white
          px-5
          text-sm
          font-bold
          text-blue-950
          transition
          hover:border-orange-300
          hover:text-orange-600
        "
      >
        <ArrowLeftIcon
          size={17}
          strokeWidth={2.5}
          class="transition-transform group-hover:-translate-x-1"
        />

        <span class="hidden sm:inline">Sebelumnya</span>
      </a>
    {/if}

    {#each pages as item, index (`${item}-${index}`)}
      {#if item === 'ellipsis'}
        <span
          aria-hidden="true"
          class="
            inline-flex
            size-11
            items-center
            justify-center
            text-sm
            font-bold
            text-blue-950/35
          "
        >
          ...
        </span>
      {:else}
        <a
          href={resolve(`/news?${pageQuery(item)}`)}
          aria-label={`Halaman ${item}`}
          aria-current={item === meta.page ? 'page' : undefined}
          class="
            inline-flex
            size-11
            items-center
            justify-center
            rounded-full
            border
            text-sm
            font-black
            transition
            {
              item === meta.page
                ? 'border-orange-500 bg-orange-500 text-white shadow-[0_4px_0_rgba(163,61,10,0.3)]'
                : 'border-blue-900/10 bg-white text-blue-950 hover:border-orange-300 hover:text-orange-600'
            }
          "
        >
          {item}
        </a>
      {/if}
    {/each}

    {#if meta.page < meta.total_pages}
      <a
        href={resolve(`/news?${pageQuery(meta.page + 1)}`)}
        aria-label="Halaman berikutnya"
        class="
          group
          inline-flex
          h-11
          items-center
          justify-center
          gap-2
          rounded-full
          border
          border-blue-900/10
          bg-white
          px-5
          text-sm
          font-bold
          text-blue-950
          transition
          hover:border-orange-300
          hover:text-orange-600
        "
      >
        <span class="hidden sm:inline">Selanjutnya</span>

        <ArrowRightIcon
          size={17}
          strokeWidth={2.5}
          class="transition-transform group-hover:translate-x-1"
        />
      </a>
    {/if}
  </nav>
{/if}
