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

  let { kementrian }: { kementrian: Cabinet } = $props();

  const cabinetLogos = $derived.by<CabinetLogo[]>(() => {
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
          child.logo?.url ?? child.logo?.thumbnail_url;

        if (!src || !child.slug) continue;

        logos.push({
          id: String(child.id),
          slug: child.slug,
          name: child.short_name ?? child.name,
          src,
          alt: child.logo?.alt_text ?? `Logo ${child.name}`,
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

  let carousel = $state<HTMLDivElement | undefined>(undefined);
  let track = $state<HTMLDivElement | undefined>(undefined);
  let mounted = $state(false);
  let reduceMotion = false;
  let animation: gsap.core.Tween | undefined;

  function pauseCarousel(): void {
    animation?.pause();
  }

  function resumeCarousel(): void {
    if (!carousel?.contains(document.activeElement)) {
      animation?.play();
    }
  }

  function handleFocusOut(event: FocusEvent): void {
    if (
      event.relatedTarget instanceof Node &&
      carousel?.contains(event.relatedTarget)
    ) {
      return;
    }

    resumeCarousel();
  }

  $effect(() => {
    const items = cabinetLogos;
    const element = track;
    const root = carousel;

    if (
      !mounted ||
      reduceMotion ||
      !root ||
      !element ||
      items.length === 0
    ) {
      return;
    }

    const context = gsap.context(() => {
      animation = gsap.to(element, {
        xPercent: -50,
        duration: 34,
        ease: 'none',
        repeat: -1
      });
    }, root);

    return () => {
      animation?.kill();
      animation = undefined;
      context.revert();
    };
  });

  onMount(() => {
    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    const root = carousel;

    root?.addEventListener('mouseenter', pauseCarousel);
    root?.addEventListener('mouseleave', resumeCarousel);
    root?.addEventListener('focusin', pauseCarousel);
    root?.addEventListener('focusout', handleFocusOut);

    mounted = true;

    return () => {
      mounted = false;

      root?.removeEventListener('mouseenter', pauseCarousel);
      root?.removeEventListener('mouseleave', resumeCarousel);
      root?.removeEventListener('focusin', pauseCarousel);
      root?.removeEventListener('focusout', handleFocusOut);
    };
  });
</script>

<div class="overflow-hidden py-3 sm:py-4">
  <div
    bind:this={carousel}
    class="w-full overflow-hidden"
    role="region"
    aria-label="Logo kementerian BEM UNAIR"
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
              gap-3 pr-3
              sm:gap-5 sm:pr-5
              md:gap-6 md:pr-6
            "
          >
            {#each cabinetLogos as logo (`${copy}-${logo.id}`)}
              <a
                href={`/kementrian/${encodeURIComponent(logo.slug)}`}
                aria-label={`Lihat ${logo.name}`}
                tabindex={copy === 0 ? 0 : -1}
                class="
                  flex size-16 shrink-0 items-center
                  justify-center rounded-xl p-1.5
                  transition-transform duration-300
                  hover:-translate-y-1
                  hover:scale-[1.03]
                  focus-visible:outline-2
                  focus-visible:outline-offset-2
                  focus-visible:outline-blue-700
                  sm:size-20 sm:p-2
                  md:size-24 md:p-3
                "
              >
                <img
                  src={logo.src}
                  alt={copy === 0 ? logo.alt : ''}
                  title={logo.name}
                  draggable="false"
                  loading="lazy"
                  class="h-full w-full object-contain"
                />
              </a>
            {/each}
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>