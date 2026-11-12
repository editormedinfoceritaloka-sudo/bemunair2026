<script lang="ts">
  import { onMount, tick } from 'svelte';
  import { gsap } from 'gsap';
  import Card from './Card.svelte';

  type ApiImage = {
    id?: number;
    file_id?: string;
    url?: string | null;
    thumbnail_url?: string | null;
    name?: string;
    alt_text?: string;
  };

  type KemenkoMember = {
    id: number;
    name: string;
    position: string;
    position_type?: string;
    photo?: ApiImage | null;
    display_order?: number;
    is_leader?: boolean;
  };

  type Kemenkoan = {
    id: number;
    code: string;
    name: string;
    unit_type: string;
    slug: string;
    short_name?: string | null;
    logo?: ApiImage | null;
    display_order?: number;
    is_active?: boolean;
    is_published?: boolean;
    members?: KemenkoMember[];
  };

  type CabinetItem = {
    id: string;
    title: string;
    subtitle?: string;
    image: string;
    backgroundImage: string;
  };

  let {
    autoplayDelay = 5000,
    kemenkoan = []
  }: {
    autoplayDelay?: number;
    kemenkoan?: Kemenkoan[];
  } = $props();

  const cabinetItems = $derived.by<CabinetItem[]>(() =>
    [...kemenkoan]
      .filter(
        (item) =>
          item.unit_type === 'MENKO' &&
          item.is_active !== false &&
          item.is_published !== false
      )
      .sort(
        (first, second) =>
          (first.display_order ?? 0) -
          (second.display_order ?? 0)
      )
      .map((item) => {
        const members = [...(item.members ?? [])].sort(
          (first, second) =>
            (first.display_order ?? 0) -
            (second.display_order ?? 0)
        );

        const leader =
          members.find((member) => member.is_leader === true) ??
          members[0];

        return {
          id: item.slug,
          title:
            leader?.position ??
            item.short_name ??
            item.name,
          subtitle: leader?.name ?? item.name,
          image:
            leader?.photo?.url ??
            leader?.photo?.thumbnail_url ??
            item.logo?.url ??
            item.logo?.thumbnail_url ??
            '/landing/speech/p.webp',
          backgroundImage: '/landing/cabinet/bg-card.webp'
        };
      })
  );

  const loopItems = $derived([
    ...cabinetItems,
    ...cabinetItems,
    ...cabinetItems
  ]);

  let sliderRoot = $state<HTMLDivElement | undefined>(undefined);
  let viewport = $state<HTMLDivElement | undefined>(undefined);
  let track = $state<HTMLDivElement | undefined>(undefined);
  let currentIndex = $state(0);
  let isAnimating = $state(false);
  let mounted = $state(false);

  const cards = new Map<number, HTMLDivElement>();

  let autoplayTimer: ReturnType<typeof setInterval> | undefined;
  let resizeObserver: ResizeObserver | undefined;
  let reduceMotion = false;
  let interactionPaused = false;
  let initialSignature = '';
  let initializeFrame = 0;
  let centerFrame = 0;

  function registerCard(node: HTMLDivElement, index: number) {
    const handleClick = (event: MouseEvent): void => {
      selectCard(event, index);
    };

    cards.set(index, node);
    node.addEventListener('click', handleClick);

    return {
      destroy() {
        node.removeEventListener('click', handleClick);

        if (cards.get(index) === node) {
          cards.delete(index);
        }
      }
    };
  }

  function getCardPosition(index: number): number | null {
    const card = cards.get(index);
    const sliderViewport = viewport;

    if (!card || !sliderViewport) return null;

    return (
      sliderViewport.clientWidth / 2 -
      (card.offsetLeft + card.offsetWidth / 2)
    );
  }

  function centerCard(
    index: number,
    animate = true,
    onComplete?: () => void
  ): void {
    const sliderTrack = track;
    const targetX = getCardPosition(index);

    if (targetX === null || !sliderTrack) {
      isAnimating = false;
      return;
    }

    gsap.killTweensOf(sliderTrack);

    if (!animate || reduceMotion) {
      gsap.set(sliderTrack, { x: targetX });
      isAnimating = false;
      onComplete?.();
      return;
    }

    isAnimating = true;

    gsap.to(sliderTrack, {
      x: targetX,
      duration: 0.75,
      ease: 'power3.inOut',
      overwrite: true,
      onComplete: () => {
        isAnimating = false;
        onComplete?.();
      }
    });
  }

  function normalizePosition(): void {
    const total = cabinetItems.length;

    if (total === 0) return;

    if (currentIndex >= total * 2) {
      currentIndex -= total;
      centerCard(currentIndex, false);
    } else if (currentIndex < total) {
      currentIndex += total;
      centerCard(currentIndex, false);
    }
  }

  function moveToIndex(index: number): void {
    if (isAnimating || cabinetItems.length <= 1) return;

    currentIndex = index;
    centerCard(currentIndex, true, normalizePosition);
  }

  function next(): void {
    moveToIndex(currentIndex + 1);
  }

  function previous(): void {
    moveToIndex(currentIndex - 1);
  }

  function handleNext(): void {
    next();
    restartAutoplay();
  }

  function handlePrevious(): void {
    previous();
    restartAutoplay();
  }

  function selectCard(event: MouseEvent, index: number): void {
    if (index === currentIndex) return;

    event.preventDefault();
    moveToIndex(index);
    restartAutoplay();
  }

  function stopAutoplay(): void {
    if (autoplayTimer === undefined) return;

    clearInterval(autoplayTimer);
    autoplayTimer = undefined;
  }

  function startAutoplay(): void {
    stopAutoplay();

    if (
      !mounted ||
      interactionPaused ||
      reduceMotion ||
      cabinetItems.length <= 1 ||
      document.hidden
    ) {
      return;
    }

    autoplayTimer = setInterval(
      next,
      Math.max(1000, autoplayDelay)
    );
  }

  function restartAutoplay(): void {
    if (!interactionPaused) {
      startAutoplay();
    }
  }

  function pauseSlider(): void {
    interactionPaused = true;
    stopAutoplay();
  }

  function resumeSlider(): void {
    if (
      sliderRoot?.matches(':hover') ||
      sliderRoot?.contains(document.activeElement)
    ) {
      return;
    }

    interactionPaused = false;
    startAutoplay();
  }

  function handleFocusOut(event: FocusEvent): void {
    if (
      event.relatedTarget instanceof Node &&
      sliderRoot?.contains(event.relatedTarget)
    ) {
      return;
    }

    resumeSlider();
  }

  function handleKeydown(event: KeyboardEvent): void {
    if (event.key === 'ArrowLeft') {
      event.preventDefault();
      handlePrevious();
    } else if (event.key === 'ArrowRight') {
      event.preventDefault();
      handleNext();
    }
  }

  function handleVisibilityChange(): void {
    if (document.hidden) {
      stopAutoplay();
    } else {
      startAutoplay();
    }
  }

  $effect(() => {
    const items = cabinetItems;

    if (!mounted) return;

    const signature = items.map((item) => item.id).join('|');

    if (signature === initialSignature) return;

    initialSignature = signature;

    stopAutoplay();
    gsap.killTweensOf(track);

    isAnimating = false;
    currentIndex = items.length;

    void tick().then(() => {
      if (!mounted) return;

      cancelAnimationFrame(initializeFrame);

      initializeFrame = requestAnimationFrame(() => {
        if (!mounted) return;

        if (cabinetItems.length > 0) {
          centerCard(currentIndex, false);
        }

        startAutoplay();
      });
    });
  });

  onMount(() => {
    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    const root = sliderRoot;
    const sliderViewport = viewport;

    root?.addEventListener('mouseenter', pauseSlider);
    root?.addEventListener('mouseleave', resumeSlider);
    root?.addEventListener('focusin', pauseSlider);
    root?.addEventListener('focusout', handleFocusOut);
    root?.addEventListener('keydown', handleKeydown);

    if (sliderViewport) {
      gsap.fromTo(
        sliderViewport,
        {
          autoAlpha: 0,
          y: reduceMotion ? 0 : 24
        },
        {
          autoAlpha: 1,
          y: 0,
          duration: reduceMotion ? 0.01 : 0.85,
          ease: 'power3.out'
        }
      );

      resizeObserver = new ResizeObserver(() => {
        if (cabinetItems.length === 0) return;

        cancelAnimationFrame(centerFrame);

        centerFrame = requestAnimationFrame(() => {
          if (mounted) {
            centerCard(currentIndex, false);
          }
        });
      });

      resizeObserver.observe(sliderViewport);
    }

    document.addEventListener(
      'visibilitychange',
      handleVisibilityChange
    );

    mounted = true;

    return () => {
      mounted = false;

      cancelAnimationFrame(initializeFrame);
      cancelAnimationFrame(centerFrame);

      stopAutoplay();
      resizeObserver?.disconnect();

      root?.removeEventListener('mouseenter', pauseSlider);
      root?.removeEventListener('mouseleave', resumeSlider);
      root?.removeEventListener('focusin', pauseSlider);
      root?.removeEventListener('focusout', handleFocusOut);
      root?.removeEventListener('keydown', handleKeydown);

      document.removeEventListener(
        'visibilitychange',
        handleVisibilityChange
      );

      if (track) {
        gsap.killTweensOf(track);
      }

      if (sliderViewport) {
        gsap.killTweensOf(sliderViewport);
      }

      cards.clear();
    };
  });
</script>

<div
  bind:this={sliderRoot}
  role="region"
  aria-label="Slider struktur Kabinet Cerita Loka"
  class="
    relative w-full overflow-hidden
    py-2 sm:py-3 md:py-4
  "
>
  <div
    bind:this={viewport}
    class="relative w-full overflow-hidden opacity-0"
  >
    {#if cabinetItems.length > 0}
      <div
        bind:this={track}
        class="
          flex w-max items-center gap-0
          py-8 will-change-transform
          sm:py-10 md:py-12
        "
      >
        {#each loopItems as item, index (`${index}-${item.id}`)}
          {@const distance = Math.abs(index - currentIndex)}

          <div
            use:registerCard={index}
            role="presentation"
            class={`
              relative h-[17rem] w-48 shrink-0
              -mr-16 cursor-pointer
              transition-[transform,opacity,filter]
              duration-500

              sm:h-[21rem] sm:w-60 sm:-mr-20
              md:h-[24rem] md:w-68 md:-mr-24
              lg:h-[26rem] lg:w-76 lg:-mr-28
              xl:h-[28rem] xl:w-80 xl:-mr-32

              ${
                distance === 0
                  ? 'z-30 scale-100 opacity-100'
                  : distance === 1
                    ? 'z-20 scale-[0.9] opacity-[0.76]'
                    : distance === 2
                      ? 'z-10 scale-[0.8] opacity-[0.42]'
                      : 'z-0 scale-[0.72] opacity-10'
              }
            `}
          >
            <Card
              id={item.id}
              title={item.title}
              subtitle={item.subtitle}
              image={item.image}
              backgroundImage={item.backgroundImage}
              active={index === currentIndex}
            />
          </div>
        {/each}
      </div>
    {:else}
      <div
        class="
          flex min-h-48 w-full
          items-center justify-center
          px-4 text-center
          text-xs font-semibold text-blue-900/70
        "
      >
        Data kementerian koordinator belum tersedia.
      </div>
    {/if}
  </div>

  {#if cabinetItems.length > 1}
    <button
      type="button"
      aria-label="Tampilkan kementerian koordinator sebelumnya"
      onclick={handlePrevious}
      disabled={isAnimating}
      class="
        absolute left-2 top-1/2 z-50
        flex size-9 -translate-y-1/2
        items-center justify-center
        rounded-full
        border border-white/70
        bg-blue-950/75
        text-white
        shadow-lg
        backdrop-blur-md
        transition
        hover:scale-110
        hover:bg-blue-900
        focus-visible:outline-2
        focus-visible:outline-offset-2
        focus-visible:outline-white
        active:scale-95
        disabled:pointer-events-none
        disabled:opacity-40
        sm:left-4 sm:size-10
        md:left-6 md:size-11
      "
    >
      <svg
        viewBox="0 0 24 24"
        aria-hidden="true"
        class="
          size-5 fill-none
          stroke-current stroke-2
          sm:size-6
        "
      >
        <path
          d="m15 18-6-6 6-6"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </button>

    <button
      type="button"
      aria-label="Tampilkan kementerian koordinator berikutnya"
      onclick={handleNext}
      disabled={isAnimating}
      class="
        absolute right-2 top-1/2 z-50
        flex size-9 -translate-y-1/2
        items-center justify-center
        rounded-full
        border border-white/70
        bg-blue-950/75
        text-white
        shadow-lg
        backdrop-blur-md
        transition
        hover:scale-110
        hover:bg-blue-900
        focus-visible:outline-2
        focus-visible:outline-offset-2
        focus-visible:outline-white
        active:scale-95
        disabled:pointer-events-none
        disabled:opacity-40
        sm:right-4 sm:size-10
        md:right-6 md:size-11
      "
    >
      <svg
        viewBox="0 0 24 24"
        aria-hidden="true"
        class="
          size-5 fill-none
          stroke-current stroke-2
          sm:size-6
        "
      >
        <path
          d="m9 18 6-6-6-6"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </button>
  {/if}
</div>