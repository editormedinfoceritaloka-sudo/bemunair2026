<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';

  let {
    id,
    title,
    subtitle,
    image,
    active = false,
    backgroundImage = '/landing/cabinet/bg-card.webp'
  }: {
    id: string;
    title: string;
    subtitle?: string;
    image: string;
    active?: boolean;
    backgroundImage?: string;
  } = $props();

  let cardElement!: HTMLDivElement;
  let frontFace!: HTMLDivElement;
  let backFace!: HTMLDivElement;
  let frontTitle!: HTMLDivElement;

  let flipAnimation: gsap.core.Timeline | undefined;
  let reduceMotion = false;

  function showDetail(): void {
    if (!active || !cardElement) return;

    flipAnimation?.kill();

    flipAnimation = gsap.timeline({
      defaults: { overwrite: true }
    });

    flipAnimation
      .to(
        frontTitle,
        {
          autoAlpha: 0,
          scale: 0.96,
          duration: reduceMotion ? 0.01 : 0.14,
          ease: 'power2.out'
        },
        0
      )
      .to(
        cardElement,
        {
          rotationY: 180,
          duration: reduceMotion ? 0.01 : 0.65,
          ease: 'power3.inOut'
        },
        reduceMotion ? 0 : 0.06
      );
  }

  function hideDetail(): void {
    if (!cardElement) return;

    flipAnimation?.kill();

    flipAnimation = gsap.timeline({
      defaults: { overwrite: true }
    });

    flipAnimation
      .to(
        cardElement,
        {
          rotationY: 0,
          duration: reduceMotion ? 0.01 : 0.55,
          ease: 'power3.inOut'
        },
        0
      )
      .set(frontTitle, {
        autoAlpha: 0,
        scale: 0.96
      })
      .to(frontTitle, {
        autoAlpha: 1,
        scale: 1,
        duration: reduceMotion ? 0.01 : 0.17,
        ease: 'power2.out'
      });
  }

  function handleClick(event: MouseEvent): void {
    if (!active) event.preventDefault();
  }

  $effect(() => {
    if (!active && cardElement) {
      hideDetail();
    }
  });

  onMount(() => {
    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    gsap.set(cardElement, {
      rotationY: 0,
      transformOrigin: 'center center',
      transformStyle: 'preserve-3d',
      force3D: true
    });

    gsap.set(frontFace, {
      rotationY: 0,
      backfaceVisibility: 'hidden',
      transformStyle: 'preserve-3d',
      force3D: true
    });

    gsap.set(backFace, {
      rotationY: 180,
      backfaceVisibility: 'hidden',
      transformStyle: 'preserve-3d',
      force3D: true
    });

    gsap.set(frontTitle, {
      autoAlpha: 1,
      scale: 1
    });

    return () => {
      flipAnimation?.kill();

      gsap.killTweensOf([
        cardElement,
        frontFace,
        backFace,
        frontTitle
      ]);
    };
  });
</script>

<a
  href={`/menko/${encodeURIComponent(id)}`}
  aria-label={`Lihat detail ${title}`}
  aria-disabled={!active}
  tabindex={active ? 0 : -1}
  onclick={handleClick}
  onmouseenter={showDetail}
  onmouseleave={hideDetail}
  onfocusin={showDetail}
  onfocusout={hideDetail}
  class={`
    block h-full w-full rounded-[20px] outline-none
    transition-[filter] duration-300
    focus-visible:ring-2 focus-visible:ring-blue-700
    focus-visible:ring-offset-2
    ${active ? 'cursor-pointer' : 'cursor-default'}
  `}
>
  <div class="relative h-full w-full [perspective:1800px]">
    <div
      bind:this={cardElement}
      class="
        relative h-full w-full
        [transform-style:preserve-3d]
        will-change-transform
      "
    >
      <div
        bind:this={frontFace}
        class="
          absolute inset-0 overflow-hidden rounded-[20px]
          [backface-visibility:hidden]
          [-webkit-backface-visibility:hidden]
          [transform:rotateY(0deg)_translateZ(0)]
        "
      >
        <img
          src={backgroundImage}
          alt=""
          draggable="false"
          class="absolute inset-0 h-full w-full object-fill"
        />

        <div
          bind:this={frontTitle}
          class="
            relative z-10 flex h-full w-full
            flex-col items-center justify-center
            px-4 text-center
            will-change-[opacity,transform]
          "
        >
          <h3
            class="
              max-w-[90%] text-base font-extrabold
              uppercase leading-[1.05]
              tracking-[-0.025em] text-blue-700
              sm:text-lg md:text-xl xl:text-2xl
            "
          >
            {title}
          </h3>

          {#if subtitle}
            <p
              class="
                mt-2 max-w-[90%] text-[8px]
                font-semibold uppercase leading-snug
                tracking-[0.08em] text-blue-600
                sm:text-[9px] md:text-[10px]
              "
            >
              {subtitle}
            </p>
          {/if}
        </div>
      </div>

      <div
        bind:this={backFace}
        class="
          absolute inset-0 overflow-hidden rounded-[20px]
          [backface-visibility:hidden]
          [-webkit-backface-visibility:hidden]
          [transform:rotateY(180deg)_translateZ(0)]
        "
      >
        <img
          src={backgroundImage}
          alt=""
          draggable="false"
          class="absolute inset-0 h-full w-full object-fill"
        />

        <div
          class="
            absolute inset-x-0 top-[9%] z-30
            flex flex-col items-center px-3 text-center
          "
        >
          <h3
            class="
              max-w-[90%] text-sm font-extrabold
              uppercase leading-[1.05]
              tracking-[-0.025em] text-blue-800
              drop-shadow-[0_1px_1px_rgba(255,255,255,0.95)]
              sm:text-base md:text-lg xl:text-xl
            "
          >
            {title}
          </h3>

          {#if subtitle}
            <p
              class="
                mt-1 max-w-[90%] text-[8px]
                font-semibold uppercase leading-snug
                tracking-[0.07em] text-blue-600
                sm:text-[9px] md:text-[10px]
              "
            >
              {subtitle}
            </p>
          {/if}
        </div>

        <div
          class="
            pointer-events-none absolute inset-x-0
            bottom-0 z-20 flex h-[64%] w-full
            items-end justify-center overflow-hidden
          "
        >
          <img
            src={image}
            alt={title}
            draggable="false"
            class="
              block h-full w-auto max-w-full
              object-contain object-bottom
            "
          />
        </div>
      </div>
    </div>
  </div>
</a>