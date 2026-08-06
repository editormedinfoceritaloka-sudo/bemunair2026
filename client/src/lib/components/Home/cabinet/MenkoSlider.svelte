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

  const cabinetItems =
    $derived.by<CabinetItem[]>(() => {
      return [...kemenkoan]
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
          const members = [
            ...(item.members ?? [])
          ].sort(
            (first, second) =>
              (first.display_order ?? 0) -
              (second.display_order ?? 0)
          );

          const leader =
            members.find(
              (member) =>
                member.is_leader === true
            ) ?? members[0];

          const image =
            leader?.photo?.url ??
            leader?.photo?.thumbnail_url ??
            item.logo?.url ??
            item.logo?.thumbnail_url ??
            '/landing/speech/p.png';

          return {
            id: item.slug,
            title:
              leader?.position ??
              item.short_name ??
              item.name,
            subtitle:
              leader?.name ??
              item.name,
            image,
            backgroundImage:
              '/landing/cabinet/bg-card.png'
          };
        });
    });

  const loopItems =
    $derived.by<CabinetItem[]>(() => [
      ...cabinetItems,
      ...cabinetItems,
      ...cabinetItems
    ]);

  let sliderRoot = $state<
    HTMLDivElement | undefined
  >(undefined);

  let viewport = $state<
    HTMLDivElement | undefined
  >(undefined);

  let track = $state<
    HTMLDivElement | undefined
  >(undefined);

  let currentIndex = $state(0);
  let initializedItemCount = $state(0);
  let isAnimating = $state(false);
  let interactionPaused = $state(false);

  let cards: Array<
    HTMLDivElement | undefined
  > = [];

  let autoplayTimer:
    | ReturnType<typeof setInterval>
    | undefined;

  let resizeObserver:
    | ResizeObserver
    | undefined;

  let reduceMotion = false;
  let mounted = false;

  function registerCard(
    node: HTMLDivElement,
    index: number
  ) {
    const handleClick = (
      event: MouseEvent
    ): void => {
      selectCard(event, index);
    };

    cards[index] = node;

    node.addEventListener(
      'click',
      handleClick
    );

    return {
      destroy() {
        node.removeEventListener(
          'click',
          handleClick
        );

        if (cards[index] === node) {
          cards[index] = undefined;
        }
      }
    };
  }

  function getCardPosition(
    index: number
  ): number | null {
    const card = cards[index];
    const sliderViewport = viewport;

    if (!card || !sliderViewport) {
      return null;
    }

    return (
      sliderViewport.clientWidth / 2 -
      (
        card.offsetLeft +
        card.offsetWidth / 2
      )
    );
  }

  function centerCard(
    index: number,
    animate = true,
    onComplete?: () => void
  ): void {
    const sliderTrack = track;
    const targetX =
      getCardPosition(index);

    if (
      targetX === null ||
      !sliderTrack
    ) {
      isAnimating = false;
      return;
    }

    gsap.killTweensOf(sliderTrack);

    if (!animate || reduceMotion) {
      gsap.set(sliderTrack, {
        x: targetX
      });

      isAnimating = false;
      onComplete?.();
      return;
    }

    isAnimating = true;

    gsap.to(sliderTrack, {
      x: targetX,
      duration: 0.9,
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

    if (total === 0) {
      return;
    }

    if (currentIndex >= total * 2) {
      currentIndex -= total;

      centerCard(
        currentIndex,
        false
      );

      return;
    }

    if (currentIndex < total) {
      currentIndex += total;

      centerCard(
        currentIndex,
        false
      );
    }
  }

  function moveToIndex(
    index: number
  ): void {
    if (
      isAnimating ||
      cabinetItems.length <= 1
    ) {
      return;
    }

    currentIndex = index;

    centerCard(
      currentIndex,
      true,
      normalizePosition
    );
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

  function selectCard(
    event: MouseEvent,
    index: number
  ): void {
    if (index === currentIndex) {
      return;
    }

    event.preventDefault();

    moveToIndex(index);
    restartAutoplay();
  }

  function stopAutoplay(): void {
    if (!autoplayTimer) {
      return;
    }

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
      autoplayDelay
    );
  }

  function restartAutoplay(): void {
    if (interactionPaused) {
      return;
    }

    startAutoplay();
  }

  function pauseSlider(): void {
    interactionPaused = true;
    stopAutoplay();
  }

  function resumeSlider(): void {
    interactionPaused = false;
    startAutoplay();
  }

  function handleKeydown(
    event: KeyboardEvent
  ): void {
    if (event.key === 'ArrowLeft') {
      event.preventDefault();
      handlePrevious();
      return;
    }

    if (event.key === 'ArrowRight') {
      event.preventDefault();
      handleNext();
    }
  }

  $effect(() => {
    const total = cabinetItems.length;

    if (
      !mounted ||
      total === initializedItemCount
    ) {
      return;
    }

    initializedItemCount = total;
    stopAutoplay();
    cards = [];

    if (total === 0) {
      currentIndex = 0;
      return;
    }

    currentIndex = total;

    void tick().then(() => {
      requestAnimationFrame(() => {
        centerCard(
          currentIndex,
          false
        );

        startAutoplay();
      });
    });
  });

  onMount(() => {
    let disposed = false;
    let initializeFrame = 0;
    let centerFrame = 0;

    mounted = true;

    const root = sliderRoot;

    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    const handleMouseEnter =
      (): void => {
        pauseSlider();
      };

    const handleMouseLeave =
      (): void => {
        resumeSlider();
      };

    const handleFocusIn =
      (): void => {
        pauseSlider();
      };

    const handleFocusOut = (
      event: FocusEvent
    ): void => {
      const nextTarget =
        event.relatedTarget;

      if (
        nextTarget instanceof Node &&
        root?.contains(nextTarget)
      ) {
        return;
      }

      resumeSlider();
    };

    const handleVisibilityChange =
      (): void => {
        if (document.hidden) {
          stopAutoplay();
          return;
        }

        startAutoplay();
      };

    root?.addEventListener(
      'mouseenter',
      handleMouseEnter
    );

    root?.addEventListener(
      'mouseleave',
      handleMouseLeave
    );

    root?.addEventListener(
      'focusin',
      handleFocusIn
    );

    root?.addEventListener(
      'focusout',
      handleFocusOut
    );

    root?.addEventListener(
      'keydown',
      handleKeydown
    );

    const initialize =
      async (): Promise<void> => {
        initializedItemCount =
          cabinetItems.length;

        currentIndex =
          cabinetItems.length;

        await tick();

        if (disposed) {
          return;
        }

        initializeFrame =
          requestAnimationFrame(() => {
            const sliderViewport =
              viewport;

            if (
              disposed ||
              !sliderViewport
            ) {
              return;
            }

            if (
              cabinetItems.length > 0
            ) {
              centerCard(
                currentIndex,
                false
              );
            }

            gsap.fromTo(
              sliderViewport,
              {
                autoAlpha: 0,
                y: 45,
                scale: 0.98
              },
              {
                autoAlpha: 1,
                y: 0,
                scale: 1,
                duration: reduceMotion
                  ? 0.01
                  : 1,
                ease: 'power3.out'
              }
            );

            startAutoplay();
          });

        const sliderViewport =
          viewport;

        if (!sliderViewport) {
          return;
        }

        resizeObserver =
          new ResizeObserver(() => {
            if (
              cabinetItems.length === 0
            ) {
              return;
            }

            cancelAnimationFrame(
              centerFrame
            );

            centerFrame =
              requestAnimationFrame(() => {
                centerCard(
                  currentIndex,
                  false
                );
              });
          });

        resizeObserver.observe(
          sliderViewport
        );

        document.addEventListener(
          'visibilitychange',
          handleVisibilityChange
        );
      };

    void initialize();

    return () => {
      disposed = true;
      mounted = false;

      cancelAnimationFrame(
        initializeFrame
      );

      cancelAnimationFrame(
        centerFrame
      );

      stopAutoplay();
      resizeObserver?.disconnect();

      root?.removeEventListener(
        'mouseenter',
        handleMouseEnter
      );

      root?.removeEventListener(
        'mouseleave',
        handleMouseLeave
      );

      root?.removeEventListener(
        'focusin',
        handleFocusIn
      );

      root?.removeEventListener(
        'focusout',
        handleFocusOut
      );

      root?.removeEventListener(
        'keydown',
        handleKeydown
      );

      document.removeEventListener(
        'visibilitychange',
        handleVisibilityChange
      );

      if (track) {
        gsap.killTweensOf(track);
      }

      if (viewport) {
        gsap.killTweensOf(viewport);
      }
    };
  });
</script>

<div
  bind:this={sliderRoot}
  role="region"
  aria-label="Slider struktur Kabinet Cerita Loka"
  class="
    relative w-full overflow-hidden
    py-3
    sm:py-5
    md:py-6
  "
>
  <div
    bind:this={viewport}
    class="
      relative w-full overflow-hidden
      opacity-0
    "
  >
    {#if cabinetItems.length > 0}
      <div
        bind:this={track}
        class="
          flex w-max items-center
          gap-0 py-16
          will-change-transform
        "
      >
        {#each loopItems as item, index (`${index}-${item.id}`)}
          {@const distance = Math.abs(
            index - currentIndex
          )}

          <div
            use:registerCard={index}
            role="presentation"
            class={`
              relative h-80 w-64 shrink-0
              -mr-24 cursor-pointer
              transition-[transform,opacity,filter]
              duration-500

              sm:h-105 sm:w-[70vw]
              sm:max-w-105
              sm:-mr-36

              md:h-120 md:w-[52vw]
              md:max-w-110
              md:-mr-48

              lg:h-130 lg:w-[34vw]
              lg:max-w-115
              lg:-mr-52

              xl:h-140 xl:w-[32vw]
              xl:max-w-120
              xl:-mr-56

              2xl:h-150 2xl:w-[30vw]
              2xl:max-w-125
              2xl:-mr-60

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
          flex min-h-80 w-full
          items-center justify-center
          px-6 text-center
          text-sm font-semibold
          text-blue-900/70
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
        absolute left-3 top-1/2
        z-50 flex size-11
        -translate-y-1/2
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
        focus-visible:outline-offset-4
        focus-visible:outline-white
        active:scale-95
        disabled:pointer-events-none
        disabled:opacity-40
        sm:left-5 sm:size-12
        md:left-8 md:size-14
      "
    >
      <svg
        viewBox="0 0 24 24"
        aria-hidden="true"
        class="
          size-6 fill-none
          stroke-current
          stroke-2
          sm:size-7
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
        absolute right-3 top-1/2
        z-50 flex size-11
        -translate-y-1/2
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
        focus-visible:outline-offset-4
        focus-visible:outline-white
        active:scale-95
        disabled:pointer-events-none
        disabled:opacity-40
        sm:right-5 sm:size-12
        md:right-8 md:size-14
      "
    >
      <svg
        viewBox="0 0 24 24"
        aria-hidden="true"
        class="
          size-6 fill-none
          stroke-current
          stroke-2
          sm:size-7
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