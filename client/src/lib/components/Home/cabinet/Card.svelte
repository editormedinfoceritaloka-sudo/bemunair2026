<script lang="ts">
  import { resolve } from '$app/paths';
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { playFlipSound } from '$lib/sound';

  let {
    id,
    title,
    subtitle,
    image,
    pathType = 'menko',
    active = false,
    backgroundImage = '/landing/cabinet/bg-card.webp'
  }: {
    id: string;
    title: string;
    subtitle?: string;
    image: string | null;
    pathType?: 'menko' | 'kementerian';
    active?: boolean;
    backgroundImage?: string;
  } = $props();

  const splitFrontTitle = $derived(
    id === 'sekretaris-kabinet' || id === 'bendahara-kabinet'
  );

  let cardElement!: HTMLDivElement;
  let frontFace!: HTMLDivElement;
  let backFace!: HTMLDivElement;
  let frontTitle!: HTMLDivElement;

  let flipAnimation: gsap.core.Timeline | undefined;
  let reduceMotion = false;
  let isFlipped = false;

  function showDetail(): void {
    if (!active || !cardElement || isFlipped) return;

    isFlipped = true;
    playFlipSound();
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
    if (!cardElement || !isFlipped) return;

    isFlipped = false;
    playFlipSound();
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

  function isTouchOnly(): boolean {
    return (
      typeof window !== 'undefined' &&
      window.matchMedia('(hover: none)').matches
    );
  }

  function handleClick(event: MouseEvent): void {
    if (!active) {
      event.preventDefault();
      return;
    }

    if (isTouchOnly() && !isFlipped) {
      event.preventDefault();
      showDetail();
    }
  }

  function handleMouseEnter(): void {
    if (isTouchOnly()) return;
    showDetail();
  }

  function handleMouseLeave(): void {
    if (isTouchOnly()) return;
    hideDetail();
  }

  function handleFocusIn(): void {
    if (isTouchOnly()) return;
    showDetail();
  }

  function handleFocusOut(): void {
    if (isTouchOnly()) return;
    hideDetail();
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
  href={pathType === 'kementerian'
    ? resolve('/kementrian/[slug]', { slug: id })
    : resolve('/menko/[slug]', { slug: id })}
  aria-label={`Lihat detail ${title}`}
  aria-disabled={!active}
  tabindex={active ? 0 : -1}
  onclick={handleClick}
  onmouseenter={handleMouseEnter}
  onmouseleave={handleMouseLeave}
  onfocusin={handleFocusIn}
  onfocusout={handleFocusOut}
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
            {#if splitFrontTitle}
              {title.split(' ')[0]}<br />{title
                .split(' ')
                .slice(1)
                .join(' ')}
            {:else}
              {title}
            {/if}
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
          aria-hidden="true"
          class={`
            pointer-events-none absolute inset-x-0 z-[25]
            ${id === 'menko-komunikasi-dan-informasi'
              ? 'top-[15%]'
              : id === 'menko-ppk'
                ? 'top-[14%]'
                  : id === 'menko-kemahasiswaan' ||
                      id === 'menko-pergerakan' ||
                      id === 'menko-kemasyarakatan'
                    ? 'top-[13%]'
                    : id === 'pendayagunaan-aparatur-kabinet'
                      ? 'top-[19%]'
                      : id === 'sekretaris-kabinet' ||
                          id === 'bendahara-kabinet'
                        ? 'top-[15%]'
                        : 'top-[19%]'}
            h-[26%] rounded-full
            bg-gradient-to-r from-transparent via-white/90
            to-transparent blur-md
            shadow-[0_0_24px_rgba(255,255,255,0.8)]
            transition-opacity duration-500
            ${active ? 'opacity-100' : 'opacity-0'}
          `}
        ></div>

        <div
          class={`
            absolute inset-x-0 z-30
            ${id === 'menko-komunikasi-dan-informasi'
              ? 'top-[15%]'
              : id === 'menko-ppk'
                ? 'top-[14%]'
                : id === 'pendayagunaan-aparatur-kabinet'
                  ? 'top-[19%]'
                  : id === 'sekretaris-kabinet' ||
                      id === 'bendahara-kabinet'
                    ? 'top-[15%]'
                    : 'top-[19%]'}
            flex flex-col items-center px-3 text-center
          `}
        >
          <h3
            class={`
              max-w-[90%] break-words uppercase
              ${title.length > 28 ? 'mt-0' : 'mt-6'}
              text-sm font-extrabold leading-[1.05]
              tracking-[-0.025em] text-blue-800
              drop-shadow-[0_1px_1px_rgba(255,255,255,0.95)]
              sm:text-base md:text-lg xl:text-xl
            `}
          >
            {title}
          </h3>

          {#if subtitle}
            <p
              class="
                mt-1 max-w-[90%] break-words text-[8px]
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
          class={`
            pointer-events-none absolute inset-x-0
            bottom-0 z-20 flex w-full
            items-stretch justify-center overflow-hidden
            h-[66%]
          `}
        >
          {#if image}
            <img
              src={image}
              alt={title}
              draggable="false"
              class="
                block h-full w-full
                object-contain object-bottom
              "
            />
          {:else}
            <span aria-hidden="true" class="absolute inset-0 grid place-items-center text-6xl font-bold text-blue-700">{(subtitle ?? title).slice(0, 1)}</span>
          {/if}
        </div>
      </div>
    </div>
  </div>
</a>