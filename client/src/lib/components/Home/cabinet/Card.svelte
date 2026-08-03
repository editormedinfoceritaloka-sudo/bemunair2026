<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';

  let {
    id,
    title,
    subtitle,
    image,
    active = false,
    backgroundImage = '/landing/cabinet/bg-card.png'
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
    if (!active || !cardElement) {
      return;
    }

    flipAnimation?.kill();

    flipAnimation = gsap.timeline({
      defaults: {
        overwrite: true
      }
    });

    flipAnimation
      .to(
        frontTitle,
        {
          autoAlpha: 0,
          scale: 0.92,
          duration: reduceMotion ? 0.01 : 0.16,
          ease: 'power2.out'
        },
        0
      )
      .to(
        cardElement,
        {
          rotationY: 180,
          duration: reduceMotion ? 0.01 : 0.75,
          ease: 'power3.inOut'
        },
        reduceMotion ? 0 : 0.08
      );
  }

  function hideDetail(): void {
    if (!cardElement) {
      return;
    }

    flipAnimation?.kill();

    flipAnimation = gsap.timeline({
      defaults: {
        overwrite: true
      }
    });

    flipAnimation
      .to(
        cardElement,
        {
          rotationY: 0,
          duration: reduceMotion ? 0.01 : 0.65,
          ease: 'power3.inOut'
        },
        0
      )
      .set(frontTitle, {
        autoAlpha: 0,
        scale: 0.92
      })
      .to(frontTitle, {
        autoAlpha: 1,
        scale: 1,
        duration: reduceMotion ? 0.01 : 0.18,
        ease: 'power2.out'
      });
  }

  function handleClick(event: MouseEvent): void {
    if (!active) {
      event.preventDefault();
    }
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
  href={`/kemenkoan/${encodeURIComponent(id)}`}
  aria-label={`Lihat detail ${title}`}
  aria-disabled={!active}
  tabindex={active ? 0 : -1}
  onclick={handleClick}
  onmouseenter={showDetail}
  onmouseleave={hideDetail}
  onfocusin={showDetail}
  onfocusout={hideDetail}
  class={`
    block h-full w-full rounded-[30px] outline-none
    transition-[filter] duration-500
    ${active ? 'cursor-pointer' : 'cursor-default'}
  `}
>
  <div class="relative h-full w-full [perspective:2200px]">
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
          absolute inset-0 overflow-hidden rounded-[30px]
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
            px-6 text-center
            will-change-[opacity,transform]
          "
        >
          <h3
            class="
              max-w-[84%]
              text-xl leading-[0.92] font-black uppercase
              tracking-[-0.05em] text-blue-700
              sm:text-2xl
              md:text-3xl
              xl:text-4xl
            "
          >
            {title}
          </h3>

          {#if subtitle}
            <p
              class="
                mt-3 text-[9px] font-semibold uppercase
                tracking-[0.14em] text-blue-600
                sm:text-[11px]
                md:text-xs
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
          absolute inset-0 overflow-hidden rounded-[30px]
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
            absolute inset-x-0 top-[5%] z-30
            flex flex-col items-center
            px-5 text-center
          "
        >
          <h3
            class="
              mt-16 max-w-[88%]
              text-lg leading-[0.92] font-black uppercase
              tracking-[-0.045em] text-blue-800
              drop-shadow-[0_2px_1px_rgba(255,255,255,0.95)]
              sm:text-xl
              md:text-2xl
              xl:text-3xl
            "
          >
            {title}
          </h3>

          {#if subtitle}
            <p
              class="
                mt-1.5 text-[8px] font-semibold uppercase
                tracking-[0.12em] text-blue-600
                sm:text-[10px]
                md:text-xs
              "
            >
              {subtitle}
            </p>
          {/if}
        </div>

        <div
          class="
            pointer-events-none absolute
            inset-x-0 bottom-0 z-20
            flex h-[70%] w-full
            items-end justify-center
            overflow-hidden
          "
        >
          <img
            src={image}
            alt={title}
            draggable="false"
            class="
              block h-full w-auto
              max-w-full
              object-contain object-bottom
            "
          />
        </div>
      </div>
    </div>
  </div>
</a>