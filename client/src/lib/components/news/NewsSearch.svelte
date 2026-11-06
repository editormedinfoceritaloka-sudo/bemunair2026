<script lang="ts">
  import { SearchIcon, XIcon } from '@lucide/svelte';
  import { resolve } from '$app/paths';

  let {
    search = '',
    perPage = 6
  }: {
    search?: string;
    perPage?: number;
  } = $props();
</script>

<div class="mx-auto mt-12 w-full max-w-2xl sm:mt-14">
  <form
    method="GET"
    action="/news"
    class="flex w-full items-center gap-3"
  >
    <div class="group relative min-w-0 flex-1">
      <SearchIcon
        size={20}
        strokeWidth={2.3}
        class="
          pointer-events-none
          absolute left-5
          top-1/2
          -translate-y-1/2
          text-blue-950/40
          transition-colors
          group-focus-within:text-orange-500
        "
      />

      <label class="sr-only" for="news-search">Cari berita</label>
      <input
        id="news-search"
        type="search"
        name="search"
        value={search}
        autocomplete="off"
        placeholder="Cari berita..."
        class="
          h-14 w-full
          rounded-full
          border
          border-blue-900/10
          bg-white/90
          pl-13 pr-5
          text-sm
          font-semibold
          text-blue-950
          shadow-[0_10px_30px_rgba(10,36,69,0.06)]
          outline-none
          transition
          placeholder:text-blue-950/35
          hover:border-blue-900/20
          focus:border-orange-400
          focus:ring-4
          focus:ring-orange-100
          sm:h-16
          sm:text-base
        "
      />
    </div>

    <input
      type="hidden"
      name="per_page"
      value={perPage}
    />

    <button
      type="submit"
      class="
        inline-flex
        h-14 shrink-0
        items-center
        justify-center
        rounded-full
        bg-orange-500
        px-7
        text-sm
        font-black
        text-white
        shadow-[0_6px_0_rgba(163,61,10,0.35)]
        transition
        hover:-translate-y-0.5
        hover:bg-orange-600
        active:translate-y-1
        active:shadow-none
        sm:h-16
        sm:px-9
      "
    >
      Cari
    </button>
  </form>

  {#if search}
    <div class="mt-4 flex items-center justify-between gap-5 px-2">
      <p class="min-w-0 text-sm font-semibold text-blue-950/55">
        Hasil pencarian untuk
        <strong class="font-black text-blue-950">"{search}"</strong>
      </p>

      <a
        href={resolve(`/news?per_page=${perPage}`)}
        class="
          group inline-flex
          shrink-0
          items-center
          gap-1.5
          text-sm
          font-bold
          text-orange-600
          transition-colors
          hover:text-orange-800
        "
      >
        <XIcon
          size={16}
          strokeWidth={2.5}
          class="transition-transform group-hover:rotate-90"
        />

        <span>Reset</span>
      </a>
    </div>
  {/if}
</div>
