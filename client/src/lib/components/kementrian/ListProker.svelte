<script lang="ts">
  import { onMount, tick } from 'svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/state';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import ProkerCard from './ProkerCard.svelte';

  type ProgramKerja = {
    id: string;
    title: string;
    description: string;
    slug: string;
  };

  let {
    programs,
    pageSize = 8,
    queryKey = 'prokerPage'
  }: {
    programs: ProgramKerja[];
    pageSize?: number;
    queryKey?: string;
  } = $props();

  let rootElement!: HTMLDivElement;
  let gridElement!: HTMLDivElement;

  let mounted = $state(false);
  let hasEntered = $state(false);
  let reduceMotion = false;

  let cardsTween: gsap.core.Tween | undefined;
  let entranceTrigger: ScrollTrigger | undefined;

  const normalizedPageSize = $derived(
    Math.max(1, pageSize)
  );

  const totalPages = $derived.by(() => {
    return Math.max(
      1,
      Math.ceil(programs.length / normalizedPageSize)
    );
  });

  const currentPage = $derived.by(() => {
    const queryValue =
      page.url.searchParams.get(queryKey);

    const parsedPage = Number.parseInt(
      queryValue ?? '1',
      10
    );

    if (!Number.isFinite(parsedPage)) {
      return 1;
    }

    return Math.min(
      Math.max(parsedPage, 1),
      totalPages
    );
  });

  const displayedPrograms = $derived.by(() => {
    const startIndex =
      (currentPage - 1) * normalizedPageSize;

    return programs.slice(
      startIndex,
      startIndex + normalizedPageSize
    );
  });

  const canGoPrevious = $derived(
    currentPage > 1
  );

  const canGoNext = $derived(
    currentPage < totalPages
  );

  async function changePage(
    targetPage: number
  ): Promise<void> {
    const nextPage = Math.min(
      Math.max(targetPage, 1),
      totalPages
    );

    if (nextPage === currentPage) {
      return;
    }

    const url = new URL(page.url);

    if (nextPage === 1) {
      url.searchParams.delete(queryKey);
    } else {
      url.searchParams.set(
        queryKey,
        String(nextPage)
      );
    }

    await goto(url, {
      keepFocus: true,
      noScroll: true,
      replaceState: false
    });
  }

  function animateCards(): void {
    if (!gridElement || reduceMotion) {
      return;
    }

    const cards = gsap.utils.toArray<HTMLElement>(
      '[data-program-card]',
      gridElement
    );

    cardsTween?.kill();
    gsap.killTweensOf(cards);

    cardsTween = gsap.fromTo(
      cards,
      {
        autoAlpha: 0,
        y: 45,
        scale: 0.95,
        rotateY: 8,
        transformPerspective: 1000,
        transformOrigin: 'center bottom'
      },
      {
        autoAlpha: 1,
        y: 0,
        scale: 1,
        rotateY: 0,
        duration: 0.65,
        stagger: 0.07,
        ease: 'power3.out',
        clearProps:
          'opacity,visibility,transform,transformPerspective,transformOrigin'
      }
    );
  }

  $effect(() => {
    const activePage = currentPage;

    const activeIds = displayedPrograms
      .map((program) => program.id)
      .join('|');

    if (
      !mounted ||
      !hasEntered ||
      reduceMotion
    ) {
      return;
    }

    void tick().then(() => {
      activePage;
      activeIds;
      animateCards();
    });
  });

  onMount(() => {
    mounted = true;

    gsap.registerPlugin(ScrollTrigger);

    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      hasEntered = true;
      return;
    }

    const cards = gsap.utils.toArray<HTMLElement>(
      '[data-program-card]',
      gridElement
    );

    gsap.set(cards, {
      autoAlpha: 0,
      y: 45,
      scale: 0.95
    });

    entranceTrigger = ScrollTrigger.create({
      trigger: rootElement,
      start: 'top 84%',
      once: true,
      onEnter: () => {
        hasEntered = true;
      }
    });

    return () => {
      cardsTween?.kill();
      entranceTrigger?.kill();
      gsap.killTweensOf(cards);
    };
  });
</script>

<div
  bind:this={rootElement}
  class="w-full"
>
  <div
    bind:this={gridElement}
    aria-live="polite"
    class="
      mx-auto grid w-full max-w-6xl
      grid-cols-1 gap-6 px-4
      sm:grid-cols-2
      lg:grid-cols-4
    "
  >
    {#each displayedPrograms as program (program.id)}
      <ProkerCard
        title={program.title}
        description={program.description}
        slug={program.slug}
      />
    {/each}
  </div>

  {#if totalPages > 1}
    <nav
      aria-label="Navigasi halaman program kerja"
      class="
        mt-10 flex items-center
        justify-center gap-4
      "
    >
      <button
        type="button"
        aria-label="Halaman program kerja sebelumnya"
        disabled={!canGoPrevious}
        onclick={() => {
          void changePage(currentPage - 1);
        }}
        class="
          group flex size-12
          items-center justify-center
          rounded-full
          border-2 border-white/80
          bg-blue-800
          text-white
          shadow-[0_7px_0_rgba(30,64,175,0.32)]
          transition duration-200
          hover:-translate-y-1
          hover:bg-blue-900
          hover:shadow-[0_10px_0_rgba(30,64,175,0.28)]
          active:translate-y-1
          active:shadow-none
          disabled:pointer-events-none
          disabled:translate-y-0
          disabled:opacity-35
          disabled:shadow-none
        "
      >
        <svg
          viewBox="0 0 24 24"
          fill="none"
          aria-hidden="true"
          class="
            size-6 transition-transform duration-200
            group-hover:-translate-x-0.5
          "
        >
          <path
            d="m15 18-6-6 6-6"
            stroke="currentColor"
            stroke-width="2.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>

      <div
        class="
          flex min-w-28 items-center
          justify-center gap-2
          rounded-full
          border border-white/70
          bg-white/30
          px-5 py-3
          text-sm font-black
          text-blue-950
          shadow-sm backdrop-blur-xl
        "
      >
        <span>{currentPage}</span>

        <span class="text-blue-900/45">
          /
        </span>

        <span>{totalPages}</span>
      </div>

      <button
        type="button"
        aria-label="Halaman program kerja berikutnya"
        disabled={!canGoNext}
        onclick={() => {
          void changePage(currentPage + 1);
        }}
        class="
          group flex size-12
          items-center justify-center
          rounded-full
          border-2 border-white/80
          bg-blue-800
          text-white
          shadow-[0_7px_0_rgba(30,64,175,0.32)]
          transition duration-200
          hover:-translate-y-1
          hover:bg-blue-900
          hover:shadow-[0_10px_0_rgba(30,64,175,0.28)]
          active:translate-y-1
          active:shadow-none
          disabled:pointer-events-none
          disabled:translate-y-0
          disabled:opacity-35
          disabled:shadow-none
        "
      >
        <svg
          viewBox="0 0 24 24"
          fill="none"
          aria-hidden="true"
          class="
            size-6 transition-transform duration-200
            group-hover:translate-x-0.5
          "
        >
          <path
            d="m9 6 6 6-6 6"
            stroke="currentColor"
            stroke-width="2.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>
    </nav>
  {/if}
</div>