<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  type SlotId =
    | 'a'
    | 'b'
    | 'c'
    | 'd'
    | 'e'
    | 'f';

  let {
    images,
    title = 'Dokumentasi'
  }: {
    images: string[];
    title?: string;
  } = $props();

  const slots: SlotId[] = [
    'a',
    'b',
    'c',
    'd',
    'e',
    'f'
  ];

  let frameElement: HTMLDivElement;

  let failedSlots = $state<
    Record<SlotId, boolean>
  >({
    a: false,
    b: false,
    c: false,
    d: false,
    e: false,
    f: false
  });

  const displayedImages = $derived(
    images.slice(0, 6)
  );

  $effect(() => {
    displayedImages.join('|');

    failedSlots = {
      a: false,
      b: false,
      c: false,
      d: false,
      e: false,
      f: false
    };
  });

  function handleImageError(
    slot: SlotId
  ): void {
    failedSlots[slot] = true;
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

    <div class="documentation-grid relative z-10">
      {#each slots as slot, index (slot)}
        {@const imageSrc = displayedImages[index]}

        <figure
          class={`documentation-slot slot-${slot}`}
        >
          {#if imageSrc && !failedSlots[slot]}
            <img
              src={imageSrc}
              alt={`Dokumentasi program kerja ${index + 1}`}
              loading="lazy"
              decoding="async"
              onerror={() => {
                handleImageError(slot);
              }}
              class="
                h-full w-full
                object-cover
                object-center
                transition-transform
                duration-500
                hover:scale-[1.04]
              "
            />
          {:else}
            <div
              class="documentation-placeholder"
              role="img"
              aria-label="Dokumentasi belum tersedia"
            ></div>
          {/if}
        </figure>
      {/each}
    </div>
  </div>
</section>

<style>
  .documentation-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    grid-auto-rows: 125px;
    gap: 0.5rem;
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

  .documentation-placeholder {
    width: 100%;
    height: 100%;
    background-color: #f8fafc;
    background-image:
      linear-gradient(
        45deg,
        #e5e7eb 25%,
        transparent 25%
      ),
      linear-gradient(
        -45deg,
        #e5e7eb 25%,
        transparent 25%
      ),
      linear-gradient(
        45deg,
        transparent 75%,
        #e5e7eb 75%
      ),
      linear-gradient(
        -45deg,
        transparent 75%,
        #e5e7eb 75%
      );

    background-size: 24px 24px;
    background-position:
      0 0,
      0 12px,
      12px -12px,
      -12px 0;
  }

  @media (min-width: 640px) {
    .documentation-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
      grid-auto-rows: 160px;
      gap: 0.65rem;
    }

    .documentation-slot {
      border-radius: 14px;
    }
  }

  @media (min-width: 1024px) {
    .documentation-grid {
      height: clamp(330px, 40vw, 450px);
      grid-template-columns: repeat(12, minmax(0, 1fr));
      grid-template-rows: repeat(7, minmax(0, 1fr));
      grid-auto-rows: auto;
      gap: 0.65rem;
    }

    .documentation-slot {
      border-radius: 16px;
    }

    .slot-a {
      grid-column: 1 / 6;
      grid-row: 1 / 4;
    }

    .slot-b {
      grid-column: 6 / 10;
      grid-row: 1 / 4;
    }

    .slot-c {
      grid-column: 10 / 13;
      grid-row: 1 / 5;
    }

    .slot-d {
      grid-column: 1 / 5;
      grid-row: 4 / 8;
    }

    .slot-e {
      grid-column: 5 / 10;
      grid-row: 4 / 8;
    }

    .slot-f {
      grid-column: 10 / 13;
      grid-row: 5 / 8;
    }
  }
</style>