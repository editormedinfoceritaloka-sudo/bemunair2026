<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import { playBookSound, playPageSound } from '$lib/sound';

  let {
    title,
    coverImage = '',
    photos = [],
    logo = '/logo/logo-kabinet.webp'
  }: {
    title: string;
    coverImage?: string;
    photos?: string[];
    logo?: string;
  } = $props();

  const photosPerSpread = 10;
  const maxSpreads = 6;

  const leftSlots = [
    'col-span-2 col-start-1 row-span-2 row-start-1',
    'col-start-1 row-start-3',
    'col-start-2 row-start-3',
    'col-start-1 row-start-4',
    'col-start-2 row-start-4'
  ];

  const rightSlots = [
    'col-start-1 row-start-1',
    'col-start-2 row-start-1',
    'col-span-2 col-start-1 row-span-2 row-start-2',
    'col-start-1 row-start-4',
    'col-start-2 row-start-4'
  ];

  function pickPhotos(list: string[]): string[] {
    const limit = photosPerSpread * maxSpreads;

    if (list.length <= limit) {
      return list;
    }

    return Array.from(
      { length: limit },
      (_, index) => list[Math.floor((index * list.length) / limit)]
    );
  }

  const spreads = $derived.by(() => {
    const picked = pickPhotos(photos);
    const count = Math.max(1, Math.ceil(picked.length / photosPerSpread));

    return Array.from({ length: count }, (_, index) => {
      const chunk = picked.slice(
        index * photosPerSpread,
        (index + 1) * photosPerSpread
      );

      return {
        left: leftSlots.map((_, slot) => chunk[slot] ?? null),
        right: rightSlots.map((_, slot) => chunk[slot + leftSlots.length] ?? null)
      };
    });
  });

  const showPlaceholder = $derived(photos.length === 0);

  let rootElement: HTMLDivElement;
  let paperElement: HTMLDivElement;
  let shadowElement: HTMLDivElement;

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    const leaves = gsap.utils.toArray<HTMLElement>(
      '[data-paper-leaf]',
      rootElement
    );

    const lastIndex = leaves.length - 1;

    const context = gsap.context(() => {
      gsap.set(leaves, {
        transformPerspective: 2200,
        rotationY: 0
      });

      if (reduceMotion) {
        gsap.set(leaves[0], {
          rotationY: -180,
          zIndex: 100
        });
        gsap.set(paperElement, {
          xPercent: 0,
          rotation: -2
        });
        gsap.set(shadowElement, {
          left: '0%',
          width: '100%'
        });
        return;
      }

      gsap.set(paperElement, {
        xPercent: -25,
        y: 40,
        rotation: -6,
        scale: 0.85,
        autoAlpha: 0
      });

      gsap.set(shadowElement, {
        left: '50%',
        width: '50%'
      });

      const timeline = gsap.timeline({
        paused: true
      });

      timeline
        .to(paperElement, {
          autoAlpha: 1,
          y: 0,
          scale: 1,
          rotation: -3,
          duration: 0.9,
          ease: 'back.out(1.3)'
        })
        .to(paperElement, {
          keyframes: [
            { rotation: -1.8, duration: 0.09 },
            { rotation: -3.6, duration: 0.09 },
            { rotation: -2.4, duration: 0.09 },
            { rotation: -3, duration: 0.09 }
          ],
          ease: 'none'
        }, 0.75);

      let time = 0.9;

      leaves.forEach((leaf, index) => {
        const isFirst = index === 0;
        const isLast = index === lastIndex;
        const duration = isFirst || isLast ? 1.4 : 1.2;

        timeline.to(leaf, {
          rotationY: -180,
          duration,
          ease: 'power2.inOut'
        }, time);

        timeline.set(leaf, {
          zIndex: 100 + index
        }, time + duration / 2);

        timeline.call(() => {
          if (isFirst || isLast) {
            playBookSound();
          } else {
            playPageSound();
          }
        }, [], time);

        if (isFirst) {
          timeline
            .to(paperElement, {
              xPercent: 0,
              rotation: -2,
              duration,
              ease: 'power2.inOut'
            }, time)
            .to(shadowElement, {
              left: '0%',
              width: '100%',
              duration,
              ease: 'power2.inOut'
            }, time);
        }

        if (isLast) {
          timeline
            .to(paperElement, {
              xPercent: 25,
              rotation: -3,
              duration,
              ease: 'power2.inOut'
            }, time)
            .to(shadowElement, {
              left: '0%',
              width: '50%',
              duration,
              ease: 'power2.inOut'
            }, time);
        }

        time += duration + (isLast ? 0 : 1.8);
      });

      ScrollTrigger.create({
        trigger: rootElement,
        start: 'top 72%',
        once: true,
        onEnter: () => {
          timeline.play();
        }
      });
    }, rootElement);

    return () => {
      context.revert();
    };
  });
</script>

{#snippet photoGrid(slots: string[], items: (string | null)[])}
  <div
    class="
      mt-[3%] grid min-h-0 flex-1
      grid-cols-2 grid-rows-4
      gap-[1.8cqw]
    "
  >
    {#each slots as slot, index (index)}
      <div class="relative overflow-hidden rounded-[2px] {slot}">
        {#if items[index]}
          <img
            src={items[index]}
            alt=""
            draggable="false"
            loading="lazy"
            class="absolute inset-0 h-full w-full object-cover"
          />
        {:else if showPlaceholder}
          <div
            class="
              absolute inset-0
              bg-[#e6ddca]
              bg-[repeating-linear-gradient(45deg,rgba(15,23,42,0.06)_0_6px,transparent_6px_12px)]
            "
          ></div>
        {/if}
      </div>
    {/each}
  </div>
{/snippet}

{#snippet pageHeader(start: string, end: string)}
  <div
    class="
      flex items-center justify-between
      border-b-2 border-slate-900
      pb-[1.5%]
      font-serif text-[1.6cqw] font-bold uppercase
      tracking-[0.08em] text-slate-900
    "
  >
    <span>{start}</span>
    <span>{end}</span>
  </div>
{/snippet}

{#snippet leftPage(items: (string | null)[], page: number)}
  <div
    class="
      absolute inset-0
      overflow-hidden rounded-l-[3px]
      bg-[#f7f2e6]
      [backface-visibility:hidden]
      [transform:rotateY(180deg)]
    "
  >
    <div
      class="
        flex h-full flex-col
        px-[4%] pt-[4%] pb-[5%]
      "
    >
      {@render pageHeader('Loka News', `Dokumentasi · ${page}`)}
      {@render photoGrid(leftSlots, items)}
    </div>

    <div
      class="
        absolute inset-y-0 right-0 w-[8%]
        bg-linear-to-l from-black/15 to-transparent
      "
    ></div>
  </div>
{/snippet}

{#snippet rightPage(items: (string | null)[], page: number)}
  <div
    class="
      absolute inset-0
      overflow-hidden rounded-r-[3px]
      bg-[#f7f2e6]
      [backface-visibility:hidden]
    "
  >
    <div
      class="
        flex h-full flex-col
        px-[4%] pt-[4%] pb-[5%]
      "
    >
      {@render pageHeader('Kabinet Cerita Loka', `2026 · ${page}`)}
      {@render photoGrid(rightSlots, items)}
    </div>

    <div
      class="
        absolute inset-y-0 left-0 w-[8%]
        bg-linear-to-r from-black/15 to-transparent
      "
    ></div>
  </div>
{/snippet}

{#snippet frontCover()}
  <div
    class="
      absolute inset-0
      overflow-hidden rounded-r-[3px]
      bg-[#f7f2e6]
      [backface-visibility:hidden]
    "
  >
    <div
      class="
        flex h-full flex-col
        px-[5%] pt-[5%] pb-[5%]
        text-slate-900
      "
    >
      <div
        class="
          border-y-2 border-slate-900
          py-[2%]
          text-center
        "
      >
        <p
          class="
            font-serif text-[5.4cqw] leading-none
            font-black uppercase tracking-[-0.02em]
          "
        >
          Loka News
        </p>
      </div>

      <div
        class="
          mt-[1.5%] flex justify-between
          border-b border-slate-900/60
          pb-[1.5%]
          font-serif text-[1.3cqw] uppercase tracking-[0.1em]
        "
      >
        <span>BEM UNAIR</span>
        <span>Edisi 2026</span>
      </div>

      <p
        class="
          mt-[4%] line-clamp-2
          font-serif text-[2.7cqw] leading-[1.1]
          font-bold
        "
      >
        {title}
      </p>

      <div
        class="
          relative mt-[4%] min-h-0 flex-1
          overflow-hidden rounded-[2px]
          bg-[#e6ddca]
          bg-[repeating-linear-gradient(45deg,rgba(15,23,42,0.06)_0_6px,transparent_6px_12px)]
        "
      >
        {#if coverImage}
          <img
            src={coverImage}
            alt=""
            draggable="false"
            loading="lazy"
            onerror={(event) => {
              (event.currentTarget as HTMLImageElement).style.display = 'none';
            }}
            class="absolute inset-0 h-full w-full object-cover"
          />
        {/if}
      </div>

      <div class="mt-[4%] grid grid-cols-2 gap-[4%]">
        {#each [0, 1] as column (column)}
          <div class="space-y-[0.8cqw]">
            <div class="h-[0.5cqw] rounded-full bg-slate-900/25"></div>
            <div class="h-[0.5cqw] rounded-full bg-slate-900/25"></div>
            <div class="h-[0.5cqw] w-4/5 rounded-full bg-slate-900/25"></div>
          </div>
        {/each}
      </div>
    </div>

    <div
      class="
        absolute inset-y-0 left-0 w-[8%]
        bg-linear-to-r from-black/15 to-transparent
      "
    ></div>
  </div>
{/snippet}

{#snippet backCover()}
  <div
    class="
      absolute inset-0
      overflow-hidden rounded-l-[3px]
      bg-[#f7f2e6]
      [backface-visibility:hidden]
      [transform:rotateY(180deg)]
    "
  >
    <div
      class="
        flex h-full flex-col items-center
        px-[6%] pt-[5%] pb-[6%]
        text-center text-slate-900
      "
    >
      <div
        class="
          w-full
          border-y-2 border-slate-900
          py-[2%]
          font-serif text-[1.6cqw] font-bold uppercase
          tracking-[0.12em]
        "
      >
        Loka News · Edisi 2026
      </div>

      <div class="flex flex-1 flex-col items-center justify-center">
        <img
          src={logo}
          alt=""
          draggable="false"
          loading="lazy"
          class="w-[34%] object-contain"
        />

        <p
          class="
            mt-[8%]
            font-serif text-[3.6cqw] leading-[1.05]
            font-black
          "
        >
          Merajut Cerita,
          <br />
          Tumbuh Bersama
        </p>

        <p
          class="
            mt-[5%]
            font-serif text-[1.5cqw] italic
            text-slate-700
          "
        >
          Sampai jumpa di cerita berikutnya.
        </p>
      </div>

      <div
        class="
          w-full
          border-t border-slate-900/60
          pt-[2%]
          font-serif text-[1.3cqw] uppercase tracking-[0.1em]
        "
      >
        Kabinet Cerita Loka · BEM UNAIR
      </div>
    </div>

    <div
      class="
        absolute inset-y-0 right-0 w-[8%]
        bg-linear-to-l from-black/15 to-transparent
      "
    ></div>
  </div>
{/snippet}

<div
  bind:this={rootElement}
  aria-hidden="true"
  class="
    pointer-events-none absolute inset-0
    flex items-center justify-center
    [container-type:size]
  "
>
  <div
    bind:this={paperElement}
    class="
      relative aspect-[3/2]
      w-[min(100cqw,145cqh)]
    "
  >
    <div
      bind:this={shadowElement}
      class="
        absolute -bottom-[7%] left-0 h-[12%] w-full
        rounded-[50%]
        bg-slate-900/45
        blur-[18px]
      "
    ></div>

    {#each spreads as spread, index (index)}
      <div
        data-paper-leaf
        class="
          absolute inset-y-0 right-0 w-1/2
          origin-left
          [transform-style:preserve-3d]
          shadow-[0_2px_6px_rgba(5,34,70,0.12)]
        "
        style="z-index: {50 - index};"
      >
        {#if index === 0}
          {@render frontCover()}
        {:else}
          {@render rightPage(spreads[index - 1].right, index * 2)}
        {/if}

        {@render leftPage(spread.left, index * 2 + 1)}
      </div>
    {/each}

    <div
      data-paper-leaf
      class="
        absolute inset-y-0 right-0 w-1/2
        origin-left
        [transform-style:preserve-3d]
        shadow-[0_2px_6px_rgba(5,34,70,0.12)]
      "
      style="z-index: {50 - spreads.length};"
    >
      {@render rightPage(spreads[spreads.length - 1].right, spreads.length * 2)}
      {@render backCover()}
    </div>
  </div>
</div>
