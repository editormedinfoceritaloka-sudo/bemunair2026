<script lang="ts">
  import { onMount, tick } from 'svelte';
  import { gsap } from 'gsap';

  import Card from './Card.svelte';

  type CabinetItem = {
    id: string;
    title: string;
    subtitle?: string;
    image: string;
    backgroundImage: string;
  };

  let {
    autoplayDelay = 5000
  }: {
    autoplayDelay?: number;
  } = $props();

  const cabinetItems: CabinetItem[] = [
    {
      id: 'presiden-bem',
      title: 'Presiden BEM',
      subtitle: 'Kabinet Cerita Loka',
      image: '/landing/speech/p.png',
      backgroundImage: '/landing/cabinet/bg-card.png'
    },
    {
      id: 'wakil-presiden-bem',
      title: 'Wakil Presiden BEM',
      subtitle: 'Kabinet Cerita Loka',
      image: '/landing/speech/w.png',
      backgroundImage: '/landing/cabinet/bg-card.png'
    },
    {
      id: 'kemenkoan-ppk',
      title: 'Kemenkoan PPK',
      subtitle: 'Kabinet Cerita Loka',
      image: '/landing/speech/p.png',
      backgroundImage: '/landing/cabinet/bg-card.png'
    },
    {
      id: 'bendahara-kabinet',
      title: 'Bendahara Kabinet',
      subtitle: 'Kabinet Cerita Loka',
      image: '/landing/speech/w.png',
      backgroundImage: '/landing/cabinet/bg-card.png'
    },
    {
      id: 'sekretaris-kabinet',
      title: 'Sekretaris Kabinet',
      subtitle: 'Kabinet Cerita Loka',
      image: '/landing/speech/p.png',
      backgroundImage: '/landing/cabinet/bg-card.png'
    }
  ];

  const loopItems = [
    ...cabinetItems,
    ...cabinetItems,
    ...cabinetItems
  ];

  let viewport!: HTMLDivElement;
  let track!: HTMLDivElement;

  let currentIndex = $state(cabinetItems.length);
  let isAnimating = $state(false);
  let interactionPaused = $state(false);

  let cards: HTMLDivElement[] = [];
  let autoplayTimer: ReturnType<typeof setInterval> | undefined;
  let resizeObserver: ResizeObserver | undefined;
  let reduceMotion = false;

  const activeIndex = $derived(
    ((currentIndex % cabinetItems.length) + cabinetItems.length) %
      cabinetItems.length
  );

  function registerCard(
    node: HTMLDivElement,
    index: number
  ) {
    cards[index] = node;

    return {
      destroy() {
        cards[index] = undefined as unknown as HTMLDivElement;
      }
    };
  }

  function getCardPosition(index: number): number | null {
    const card = cards[index];

    if (!card || !viewport) {
      return null;
    }

    return (
      viewport.clientWidth / 2 -
      (card.offsetLeft + card.offsetWidth / 2)
    );
  }

  function centerCard(
    index: number,
    animate = true,
    onComplete?: () => void
  ): void {
    const targetX = getCardPosition(index);

    if (targetX === null || !track) {
      return;
    }

    gsap.killTweensOf(track);

    if (!animate || reduceMotion) {
      gsap.set(track, {
        x: targetX
      });

      onComplete?.();
      return;
    }

    isAnimating = true;

    gsap.to(track, {
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

    if (currentIndex >= total * 2) {
      currentIndex -= total;
      centerCard(currentIndex, false);
      return;
    }

    if (currentIndex < total) {
      currentIndex += total;
      centerCard(currentIndex, false);
    }
  }

  function next(): void {
    if (isAnimating || cabinetItems.length <= 1) {
      return;
    }

    currentIndex += 1;

    centerCard(
      currentIndex,
      true,
      normalizePosition
    );
  }

  function selectCard(
    event: MouseEvent,
    index: number
  ): void {
    if (index === currentIndex) {
      return;
    }

    event.preventDefault();

    if (isAnimating) {
      return;
    }

    currentIndex = index;

    centerCard(
      currentIndex,
      true,
      normalizePosition
    );

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
      interactionPaused ||
      reduceMotion ||
      cabinetItems.length <= 1 ||
      document.hidden
    ) {
      return;
    }

    autoplayTimer = setInterval(() => {
      next();
    }, autoplayDelay);
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
    interactionPaused = false;
    startAutoplay();
  }

  onMount(() => {
    let disposed = false;

    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    const handleVisibilityChange = (): void => {
      if (document.hidden) {
        stopAutoplay();
        return;
      }

      startAutoplay();
    };

    const initialize = async (): Promise<void> => {
      await tick();

      if (disposed) {
        return;
      }

      requestAnimationFrame(() => {
        if (disposed) {
          return;
        }

        centerCard(currentIndex, false);

        gsap.fromTo(
          viewport,
          {
            autoAlpha: 0,
            y: 45,
            scale: 0.98
          },
          {
            autoAlpha: 1,
            y: 0,
            scale: 1,
            duration: reduceMotion ? 0.01 : 1,
            ease: 'power3.out'
          }
        );

        startAutoplay();
      });

      resizeObserver = new ResizeObserver(() => {
        centerCard(currentIndex, false);
      });

      resizeObserver.observe(viewport);

      document.addEventListener(
        'visibilitychange',
        handleVisibilityChange
      );
    };

    void initialize();

    return () => {
      disposed = true;

      stopAutoplay();
      resizeObserver?.disconnect();

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
  role="region"
  aria-label="Slider struktur Kabinet Cerita Loka"
  class="relative w-full overflow-hidden py-3 sm:py-5 md:py-6"
  onmouseenter={pauseSlider}
  onmouseleave={resumeSlider}
  onfocusin={pauseSlider}
  onfocusout={resumeSlider}
>
  <div
    bind:this={viewport}
    class="relative w-full overflow-hidden opacity-0"
  >
    <div
      bind:this={track}
      class="
        flex w-max items-center
        gap-0 py-16
        will-change-transform
      "
    >
      {#each loopItems as item, index (`${index}-${item.id}`)}
        {@const distance = Math.abs(index - currentIndex)}

        <div
          use:registerCard={index}
          role="presentation"
          onclick={(event) => selectCard(event, index)}
          class={`
            relative shrink-0
            h-80 w-64
            -mr-24ef
            transition-[transform,opacity,filter] duration-500

            sm:h-105 sm:w-[70vw] sm:max-w-[420px] sm:-mr-36

            md:h-120 md:w-[52vw] md:max-w-[440px] md:-mr-48

            lg:h-130 lg:w-[34vw] lg:max-w-[460px] lg:-mr-52

            xl:h-140 xl:w-[32vw] xl:max-w-[480px] xl:-mr-56

            2xl:h-150 2xl:w-[30vw] 2xl:max-w-[500px] 2xl:-mr-60

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
  </div>
</div>