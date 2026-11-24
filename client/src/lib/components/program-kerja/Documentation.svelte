<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  let {
    images,
    title = 'Dokumentasi'
  }: {
    images: string[];
    title?: string;
  } = $props();

  let frameElement: HTMLDivElement;
  let activeImage = $state<string | null>(null);

  $effect(() => {
    if (typeof document === 'undefined') {
      return;
    }

    document.body.style.overflow = activeImage ? 'hidden' : '';

    return () => {
      document.body.style.overflow = '';
    };
  });

  let failedImages = $state<
    Record<string, boolean>
  >({});

  const visibleImages = $derived(
    images.filter(
      (src) => !failedImages[src]
    )
  );

  $effect(() => {
    images.join('|');

    failedImages = {};
  });

  function handleImageError(
    src: string
  ): void {
    failedImages[src] = true;
  }


  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      return;
    }

    const context = gsap.context(() => {
      const waves = gsap.utils.toArray<SVGGElement>(
        '[data-gallery-wave]',
        frameElement
      );

      const orbs = gsap.utils.toArray<HTMLElement>(
        '[data-gallery-orb]',
        frameElement
      );

      const animations: gsap.core.Tween[] = [];

      waves.forEach((wave, index) => {
        const tween = gsap.to(wave, {
          x: index % 2 === 0 ? 32 : -32,
          y: index % 2 === 0 ? -8 : 8,
          duration: 5.5 + index * 1.2,
          ease: 'sine.inOut',
          repeat: -1,
          yoyo: true,
          paused: true,
          delay: index * 0.25
        });

        animations.push(tween);
      });

      orbs.forEach((orb, index) => {
        const tween = gsap.to(orb, {
          y: index % 2 === 0 ? -18 : 18,
          x: index % 2 === 0 ? 12 : -12,
          duration: 6 + index,
          ease: 'sine.inOut',
          repeat: -1,
          yoyo: true,
          paused: true
        });

        animations.push(tween);
      });

      ScrollTrigger.create({
        trigger: frameElement,
        start: 'top bottom',
        end: 'bottom top',
        onEnter: () => {
          animations.forEach((animation) => {
            animation.play();
          });
        },
        onEnterBack: () => {
          animations.forEach((animation) => {
            animation.play();
          });
        },
        onLeave: () => {
          animations.forEach((animation) => {
            animation.pause();
          });
        },
        onLeaveBack: () => {
          animations.forEach((animation) => {
            animation.pause();
          });
        }
      });
    }, frameElement);

    return () => {
      context.revert();
    };
  });
</script>

<section class="w-full">
  <h2
    class="
      text-center
      text-3xl
      leading-none
      font-black
      tracking-[-0.055em]
      text-orange-600
      [-webkit-text-stroke:1.5px_white]
      [paint-order:stroke_fill]
      drop-shadow-[0_4px_2px_rgba(15,23,42,0.3)]
      sm:text-4xl
      lg:text-5xl
    "
  >
    {title}
  </h2>

  <div
    bind:this={frameElement}
    class="
      relative
      mt-7
      w-full
      overflow-hidden
      rounded-[18px]
      bg-blue-950
      p-2.5
      shadow-[0_14px_28px_rgba(5,34,70,0.22)]
      sm:mt-8
      sm:rounded-[20px]
      sm:p-4
    "
  >
    <div
      aria-hidden="true"
      class="
        pointer-events-none
        absolute inset-0
        z-0
        overflow-hidden
      "
    >
      <div
        data-gallery-orb
        class="
          absolute
          -top-24 -left-20
          size-60
          rounded-full
          bg-blue-500/20
          blur-[65px]
        "
      ></div>

      <div
        data-gallery-orb
        class="
          absolute
          -right-20 -bottom-24
          size-72
          rounded-full
          bg-sky-400/20
          blur-[80px]
        "
      ></div>

      <svg
        viewBox="0 0 1200 500"
        preserveAspectRatio="none"
        class="
          absolute
          inset-0
          h-full w-full
        "
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g data-gallery-wave>
          <path
            d="M-100 140C160 20 340 250 620 120C850 15 1050 180 1300 80V500H-100V140Z"
            fill="#3B82F6"
            fill-opacity="0.18"
          />
        </g>

        <g data-gallery-wave>
          <path
            d="M-100 250C140 120 380 350 640 225C900 100 1110 320 1300 180V500H-100V250Z"
            fill="#38BDF8"
            fill-opacity="0.12"
          />
        </g>

        <g data-gallery-wave>
          <path
            d="M-100 350C170 225 390 430 670 325C940 210 1100 405 1300 290V500H-100V350Z"
            fill="#BFDBFE"
            fill-opacity="0.12"
          />
        </g>
      </svg>
    </div>

    <svg
      viewBox="0 0 1200 500"
      preserveAspectRatio="none"
      class="
        pointer-events-none
        absolute inset-0
        z-20
        h-full w-full
      "
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      aria-hidden="true"
    >
      <g data-gallery-wave>
        <path
          d="M-100 105C180 25 400 195 650 105C900 20 1080 175 1300 95"
          stroke="#DBEAFE"
          stroke-opacity="0.15"
          stroke-width="2"
        />
      </g>

      <g data-gallery-wave>
        <path
          d="M-100 275C160 185 410 355 680 255C920 170 1120 345 1300 250"
          stroke="#BAE6FD"
          stroke-opacity="0.14"
          stroke-width="2"
        />
      </g>

      <g data-gallery-wave>
        <path
          d="M-100 415C160 335 420 490 670 400C930 310 1110 465 1300 375"
          stroke="#FFFFFF"
          stroke-opacity="0.12"
          stroke-width="1.5"
        />
      </g>
    </svg>

    <div
      class="documentation-grid relative z-10"
      data-count={visibleImages.length}
    >
      {#each visibleImages as imageSrc, index (`${imageSrc}:${index}`)}
        <figure class="documentation-slot">
          <button
            type="button"
            class="block h-full w-full cursor-zoom-in"
            onclick={() => activeImage = imageSrc}
            aria-label={`Perbesar dokumentasi ${index + 1}`}
          >
            <img
              src={imageSrc}
              alt={`Dokumentasi program kerja ${index + 1}`}
              loading="lazy"
              decoding="async"
              onerror={() => {
                handleImageError(imageSrc);
              }}
              class="
                h-full w-full
                object-contain
                object-center
              "
            />
          </button>
        </figure>
      {/each}
    </div>
  </div>
</section>

{#if activeImage}
  <div
    class="fixed inset-0 z-[100] grid place-items-center bg-blue-950/90 p-4 backdrop-blur-sm"
    role="dialog"
    aria-modal="true"
    aria-label="Pratinjau dokumentasi"
    tabindex="-1"
    onclick={(event) => {
      if (event.target === event.currentTarget) activeImage = null;
    }}
    onkeydown={(event) => {
      if (event.key === 'Escape') activeImage = null;
    }}
  >
    <button
      type="button"
      class="absolute right-4 top-4 grid size-10 place-items-center rounded-full bg-white/90 text-xl font-black text-blue-900 shadow-lg"
      onclick={() => activeImage = null}
      aria-label="Tutup pratinjau"
    >
      ×
    </button>

    <img
      src={activeImage}
      alt="Pratinjau dokumentasi program kerja"
      class="max-h-[90vh] max-w-full rounded-xl object-contain shadow-2xl"
    />
  </div>
{/if}

<svelte:window
  onkeydown={(event) => {
    if (event.key === 'Escape') activeImage = null;
  }}
/>

<style>
  .documentation-grid {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    grid-auto-rows: 180px;
    gap: 0.65rem;
  }

  .documentation-slot {
    position: relative;
    min-width: 0;
    min-height: 0;
    overflow: hidden;
    border-radius: 12px;
    background-color: white;
    box-shadow: 0 4px 12px rgba(5, 34, 70, 0.08);
  }

  .documentation-grid[data-count='1'] {
    grid-template-columns: 1fr;
    grid-auto-rows: 220px;
  }

  @media (min-width: 640px) {
    .documentation-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
      grid-auto-rows: 160px;
    }

    .documentation-grid[data-count='1'] {
      grid-auto-rows: 300px;
    }

    .documentation-slot {
      border-radius: 14px;
    }
  }

  @media (min-width: 1024px) {
    .documentation-grid:not([data-count='6']) {
      height: auto;
      grid-template-columns: repeat(3, minmax(0, 1fr));
      grid-auto-rows: 220px;
    }

    .documentation-grid[data-count='1'] {
      grid-template-columns: 1fr;
      grid-auto-rows: 380px;
    }

    .documentation-grid[data-count='2'],
    .documentation-grid[data-count='4'] {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }

    .documentation-grid[data-count='6'] {
      height: clamp(330px, 40vw, 450px);
      grid-template-columns: repeat(12, minmax(0, 1fr));
      grid-template-rows: repeat(7, minmax(0, 1fr));
      grid-auto-rows: auto;
    }

    .documentation-slot {
      border-radius: 16px;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(1) {
      grid-column: 1 / 6;
      grid-row: 1 / 4;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(2) {
      grid-column: 6 / 10;
      grid-row: 1 / 4;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(3) {
      grid-column: 10 / 13;
      grid-row: 1 / 5;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(4) {
      grid-column: 1 / 5;
      grid-row: 4 / 8;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(5) {
      grid-column: 5 / 10;
      grid-row: 4 / 8;
    }

    .documentation-grid[data-count='6']
      .documentation-slot:nth-child(6) {
      grid-column: 10 / 13;
      grid-row: 5 / 8;
    }
  }
</style>
