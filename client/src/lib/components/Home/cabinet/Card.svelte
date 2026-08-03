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
  let flipAnimation: gsap.core.Tween | undefined;
  let reduceMotion = false;

  function showDetail(): void {
    if (!active) {
      return;
    }

    flipAnimation?.kill();

    flipAnimation = gsap.to(cardElement, {
      rotationY: 180,
      duration: reduceMotion ? 0.01 : 0.75,
      ease: 'power3.inOut',
      overwrite: true
    });
  }

  function hideDetail(): void {
    flipAnimation?.kill();

    flipAnimation = gsap.to(cardElement, {
      rotationY: 0,
      duration: reduceMotion ? 0.01 : 0.65,
      ease: 'power3.inOut',
      overwrite: true
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
      transformStyle: 'preserve-3d'
    });

    return () => {
      flipAnimation?.kill();
      gsap.killTweensOf(cardElement);
    };
  });
</script>

<a
  href={`/cabinet/${id}`}
  aria-label={`Lihat detail ${title}`}
  aria-disabled={!active}
  tabindex={active ? 0 : -1}
  onclick={handleClick}
  onmouseenter={showDetail}
  onmouseleave={hideDetail}
  onfocusin={showDetail}
  onfocusout={hideDetail}
  class={`
    block h-full w-full rounded-[28px] outline-none
    transition-[filter] duration-500
    ${
      active
        ? 'cursor-pointer drop-shadow-[0_38px_48px_rgba(15,54,100,0.52)]'
        : 'cursor-default drop-shadow-[0_22px_30px_rgba(15,54,100,0.3)]'
    }
  `}
>
  <div class="h-full w-full [perspective:2200px]">
    <div
      bind:this={cardElement}
      class="
        relative h-full w-full
        [transform-style:preserve-3d]
        will-change-transform
      "
    >
      <div
        class="
          absolute inset-0 overflow-hidden rounded-[28px]
          [backface-visibility:hidden]
          [-webkit-backface-visibility:hidden]
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
            relative z-10 flex h-full w-full
            flex-col items-center justify-center
            px-6 text-center
          "
        >
          <h3
            class="
              max-w-[82%]
              text-xl leading-[0.92] font-black uppercase
              tracking-[-0.05em] text-blue-700
              drop-shadow-[0_2px_1px_rgba(255,255,255,0.95)]
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
        class="
          absolute inset-0 overflow-hidden rounded-[28px]
          [backface-visibility:hidden]
          [-webkit-backface-visibility:hidden]
          [transform:rotateY(180deg)]
          [-webkit-transform:rotateY(180deg)]
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
            px-6 text-center
          "
        >
          <h3
            class="
              max-w-[88%]
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

        <img
          src={image}
          alt=""
          draggable="false"
          class="
            pointer-events-none absolute
            inset-x-0 bottom-0 z-20
            mx-auto h-[98%] w-auto max-w-[115%]
            object-contain object-bottom
            drop-shadow-[0_26px_36px_rgba(15,23,42,0.48)]
            sm:h-[102%]
            md:h-[106%]
            lg:h-[110%]
          "
        />
      </div>
    </div>
  </div>
</a>