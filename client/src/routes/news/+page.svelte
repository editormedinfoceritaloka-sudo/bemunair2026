<script lang="ts">
  import NewsSearch from '$lib/components/news/NewsSearch.svelte';
  import NewsList from '$lib/components/news/NewsList.svelte';
  import NewsPagination from '$lib/components/news/NewsPagination.svelte';

  import type { PageData } from './$types';

  let {
    data
  }: {
    data: PageData;
  } = $props();
</script>

<main
  class="
    relative
    min-h-screen
    overflow-hidden
    bg-linear-to-b
    from-white
    via-blue-50/60
    to-blue-100
    pb-28
    pt-28
    sm:pt-32
    lg:pb-36
    lg:pt-40
  "
>
  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      left-1/2
      top-8
      size-[42rem]
      -translate-x-1/2
      rounded-full
      bg-orange-200/20
      blur-[140px]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      -left-40
      top-[45%]
      size-96
      rounded-full
      bg-blue-200/25
      blur-[120px]
    "
  ></div>

  <div
    class="
      relative z-10
      mx-auto
      w-full
      max-w-7xl
      px-5
      sm:px-8
      lg:px-12
    "
  >
    <header class="mx-auto max-w-4xl text-center">
      <h1
        class="
          text-[clamp(4rem,11vw,9rem)]
          leading-[0.82]
          font-black
          uppercase
          tracking-[-0.075em]
          text-orange-500
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_9px_0_rgba(163,61,10,0.45)]
          sm:[-webkit-text-stroke:3px_white]
        "
      >
        News
      </h1>

      <p
        class="
          mx-auto
          mt-7
          max-w-2xl
          text-base
          leading-7
          font-semibold
          text-slate-900
          sm:text-lg
          sm:leading-8
          lg:text-xl
        "
      >
        Jelajahi cerita terbaru, perkembangan program,
        kegiatan mahasiswa, dan perjalanan
        Kabinet Cerita Loka BEM Universitas Airlangga 2026.
      </p>
    </header>

    <NewsSearch
      search={data.filters.search}
      perPage={data.filters.per_page}
    />

    <div
      class="
        mt-16
        flex
        items-end
        justify-between
        gap-5
        border-b
        border-blue-900/15
        pb-4
        sm:mt-20
      "
    >
      <h2
        class="
          text-2xl
          leading-none
          font-black
          tracking-[-0.04em]
          text-blue-950
          sm:text-3xl
        "
      >
        {data.filters.search ? 'Hasil Pencarian' : 'Semua Berita'}
      </h2>

      <p class="shrink-0 text-xs font-semibold text-blue-950/45 sm:text-sm">
        {data.meta.total}
        artikel
      </p>
    </div>

    {#key data.filters.search}
      <NewsList
        articles={data.articles}
        search={data.filters.search}
      />
    {/key}

    <NewsPagination
      meta={data.meta}
      search={data.filters.search}
    />
  </div>
</main>
