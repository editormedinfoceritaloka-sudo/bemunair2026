<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import type { Cabinet } from '$lib/types';

  type CabinetLogo = {
    id: string;
    slug: string;
    name: string;
    src: string;
    alt: string;
    parentOrder: number;
    displayOrder: number;
  };

  let {
    kementrian
  }: {
    kementrian: Cabinet;
  } = $props();

  const cabinetLogos: CabinetLogo[] = $derived.by(() => {
    const logos: CabinetLogo[] = [];

    for (const parent of kementrian.kemenkoan ?? []) {
      for (const child of parent.children ?? []) {
        if (
          child.is_active === false ||
          child.is_published === false
        ) {
          continue;
        }

        const src =
          child.logo?.url ??
          child.logo?.thumbnail_url;

        if (!src || !child.slug) {
          continue;
        }

        logos.push({
          id: String(child.id),
          slug: child.slug,
          name: child.short_name ?? child.name,
          src,
          alt:
            child.logo?.alt_text ??
            `Logo ${child.name}`,
          parentOrder: parent.display_order ?? 0,
          displayOrder: child.display_order ?? 0
        });
      }
    }

    return logos.sort((first, second) => {
      if (first.parentOrder !== second.parentOrder) {
        return first.parentOrder - second.parentOrder;
      }

      return first.displayOrder - second.displayOrder;
    });
  });

  let carousel!: HTMLDivElement;
  let track!: HTMLDivElement;
  let animation: gsap.core.Tween | undefined;

  function pauseCarousel(): void {
    animation?.pause();
  }

  function resumeCarousel(): void {
    animation?.play();
  }

  onMount(() => {
    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (
      reduceMotion ||
      cabinetLogos.length === 0
    ) {
      return;
    }

    const context = gsap.context(() => {
      animation = gsap.to(track, {
        xPercent: -50,
        duration: 28,
        ease: 'none',
        repeat: -1
      });
    }, carousel);

    return () => {
      animation?.kill();
      context.revert();
    };
  });
</script>

<div class="-rotate-2.5 overflow-hidden py-8">
  <div
    bind:this={carousel}
    class="w-full overflow-hidden"
    role="region"
    aria-label="Logo kementerian BEM UNAIR"
    onmouseenter={pauseCarousel}
    onmouseleave={resumeCarousel}
    onfocusin={pauseCarousel}
    onfocusout={resumeCarousel}
  >
    {#if cabinetLogos.length > 0}
      <div
        bind:this={track}
        class="flex w-max will-change-transform"
      >
        {#each [0, 1] as copy}
          <div
            aria-hidden={copy === 1}
            class="
              flex shrink-0 items-center
              gap-5 pr-5
              sm:gap-8 sm:pr-8
              md:gap-10 md:pr-10
            "
          >
            {#each cabinetLogos as logo (`${copy}-${logo.id}`)}
              <a
                href={`/kementrian/${logo.slug}`}
                aria-label={`Lihat ${logo.name}`}
                tabindex={copy === 0 ? 0 : -1}
                class="
                  flex size-24 shrink-0
                  items-center justify-center
                  transition duration-300
                  hover:-translate-y-2
                  hover:scale-105
                  focus-visible:outline-2
                  focus-visible:outline-offset-4
                  focus-visible:outline-blue-700
                  sm:size-32 sm:p-4
                  md:size-40 md:p-5
                "
              >
                <img
                  src={logo.src}
                  alt={copy === 0 ? logo.alt : ''}
                  title={logo.name}
                  draggable="false"
                  loading="lazy"
                  class="
                    h-full w-full
                    object-contain
                  "
                />
              </a>
            {/each}
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>