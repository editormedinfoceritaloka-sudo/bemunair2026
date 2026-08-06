<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';
  import type { PageData } from './$types';
  import MenkoCard from '$lib/components/menko/MenkoCard.svelte';
  import KementrianCard from '$lib/components/menko/KementrianCard.svelte';

  let {
    data
  }: {
    data: PageData;
  } = $props();

  const unit = $derived(data.unit);

  const leader = $derived(
    unit.members?.find((member) => member.is_leader) ??
      unit.members?.[0]
  );

  const headingPrefix = $derived(
    unit.unit_type === 'MENKO'
      ? 'Menteri Koordinator'
      : unit.short_name ?? unit.name
  );

  const headingSuffix = $derived(
    unit.unit_type === 'MENKO'
      ? unit.name
          .replace(/^Kementerian Koordinator\s*/i, '')
          .trim()
      : ''
  );

  const leaderName = $derived(
    leader?.name ?? unit.short_name ?? unit.name
  );

  const leaderPosition = $derived(
    leader?.position ?? unit.name
  );

  const leaderImage = $derived(
    leader?.photo?.url ??
      leader?.photo?.thumbnail_url ??
      unit.logo?.url ??
      unit.logo?.thumbnail_url ??
      '/menko/test.png'
  );

  const kementrian = $derived.by(() => {
    return [...(unit.children ?? [])]
      .filter(
        (child) =>
          child.is_active !== false &&
          child.is_published !== false
      )
      .sort(
        (first, second) =>
          (first.display_order ?? 0) -
          (second.display_order ?? 0)
      )
      .map((child) => ({
        id: child.id,
        title: child.name,
        slug: child.slug,
        logo:
          child.logo?.url ??
          child.logo?.thumbnail_url ??
          '/logo/logo-kabinet.png'
      }));
  });

  let sectionElement!: HTMLElement;
  let glowElement!: HTMLDivElement;
  let decorationElement!: HTMLDivElement;

  let headerElement!: HTMLElement;
  let headingElement!: HTMLHeadingElement;
  let descriptionElement!: HTMLParagraphElement;
  let menkoElement!: HTMLDivElement;

  let kementrianSectionElement!: HTMLDivElement;
  let kementrianHeadingElement!: HTMLHeadingElement;
  let kementrianGridElement!: HTMLDivElement;

  let leftStarOneElement!: HTMLImageElement;
  let leftStarTwoElement!: HTMLImageElement;
  let rightStarElement!: HTMLImageElement;

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const previousScrollRestoration =
      window.history.scrollRestoration;

    const previousScrollBehavior =
      document.documentElement.style.scrollBehavior;

    let context: gsap.Context | undefined;
    let secondFrame = 0;

    window.history.scrollRestoration = 'manual';
    document.documentElement.style.scrollBehavior = 'auto';

    function scrollToTop(): void {
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'auto'
      });

      document.documentElement.scrollTop = 0;
      document.body.scrollTop = 0;
    }

    function setupAnimations(): void {
      const reduceMotion = window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

      context = gsap.context(() => {
        const kementrianCards = Array.from(
          kementrianGridElement.children
        ) as HTMLElement[];

        const starElements = [
          leftStarOneElement,
          leftStarTwoElement,
          rightStarElement
        ];

        if (reduceMotion) {
          gsap.set(
            [
              glowElement,
              decorationElement,
              headerElement,
              headingElement,
              descriptionElement,
              menkoElement,
              kementrianHeadingElement,
              ...kementrianCards,
              ...starElements
            ],
            {
              clearProps: 'all'
            }
          );

          return;
        }

        gsap.set(glowElement, {
          opacity: 0,
          scale: 0.65
        });

        gsap.set(decorationElement, {
          opacity: 0,
          x: 100,
          rotate: 25,
          scale: 0.8
        });

        gsap.set(headingElement, {
          y: 75,
          opacity: 0,
          scale: 0.9,
          filter: 'blur(12px)'
        });

        gsap.set(descriptionElement, {
          y: 40,
          opacity: 0,
          filter: 'blur(7px)'
        });

        gsap.set(starElements, {
          opacity: 0,
          scale: 0.45,
          rotate: -18,
          transformOrigin: 'center'
        });

        const headerTimeline = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          },
          scrollTrigger: {
            trigger: headerElement,
            start: 'top 88%',
            once: true
          }
        });

        headerTimeline
          .to(
            glowElement,
            {
              opacity: 1,
              scale: 1,
              duration: 1.5
            },
            0
          )
          .to(
            decorationElement,
            {
              opacity: 1,
              x: 0,
              rotate: 45,
              scale: 1,
              duration: 1.3
            },
            0
          )
          .to(
            starElements,
            {
              opacity: 1,
              scale: 1,
              rotate: 0,
              duration: 0.8,
              stagger: 0.12,
              ease: 'back.out(1.5)'
            },
            0.1
          )
          .to(
            headingElement,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              filter: 'blur(0px)',
              duration: 1
            },
            0.15
          )
          .to(
            descriptionElement,
            {
              y: 0,
              opacity: 1,
              filter: 'blur(0px)',
              duration: 0.85
            },
            0.45
          )
          .set(
            [
              headingElement,
              descriptionElement,
              ...starElements
            ],
            {
              clearProps:
                'opacity,transform,filter,transformOrigin'
            }
          );

        gsap.set(menkoElement, {
          y: 80,
          opacity: 0,
          scale: 0.9,
          rotateY: 14,
          transformPerspective: 1100,
          transformOrigin: 'center bottom'
        });

        gsap
          .timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: menkoElement,
              start: 'top 82%',
              once: true
            }
          })
          .to(menkoElement, {
            y: 0,
            opacity: 1,
            scale: 1,
            rotateY: 0,
            duration: 1,
            ease: 'back.out(1.25)'
          })
          .set(menkoElement, {
            clearProps:
              'opacity,transform,transformPerspective,transformOrigin'
          });

        gsap.set(kementrianHeadingElement, {
          y: 55,
          opacity: 0,
          scale: 0.94,
          filter: 'blur(8px)'
        });

        gsap.set(kementrianCards, {
          y: 80,
          opacity: 0,
          scale: 0.9,
          rotateY: 12,
          transformPerspective: 1100,
          transformOrigin: 'center bottom'
        });

        gsap
          .timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: kementrianSectionElement,
              start: 'top 80%',
              once: true
            }
          })
          .to(kementrianHeadingElement, {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 0.85
          })
          .to(
            kementrianCards,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              rotateY: 0,
              duration: 0.9,
              stagger: 0.13,
              ease: 'back.out(1.25)'
            },
            0.25
          )
          .set(
            [
              kementrianHeadingElement,
              ...kementrianCards
            ],
            {
              clearProps:
                'opacity,transform,filter,transformPerspective,transformOrigin'
            }
          );
      }, sectionElement);

      ScrollTrigger.refresh();

      document.documentElement.style.scrollBehavior =
        previousScrollBehavior;
    }

    scrollToTop();

    const firstFrame = requestAnimationFrame(() => {
      scrollToTop();

      secondFrame = requestAnimationFrame(() => {
        scrollToTop();
        setupAnimations();
      });
    });

    return () => {
      cancelAnimationFrame(firstFrame);
      cancelAnimationFrame(secondFrame);

      context?.revert();

      window.history.scrollRestoration =
        previousScrollRestoration;

      document.documentElement.style.scrollBehavior =
        previousScrollBehavior;
    };
  });
</script>

<svelte:head>
  <title>{leaderPosition} | BEM UNAIR</title>
  <meta
    name="description"
    content={unit.description}
  />
</svelte:head>

<section
  bind:this={sectionElement}
  class="
    relative min-h-screen overflow-hidden
    bg-gradient-to-b
    from-blue-800
    via-blue-600
    to-blue-50
    px-5 pb-28 pt-14
    sm:px-8 sm:pt-20
    lg:px-12 lg:pb-36 lg:pt-28
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none absolute
      left-1/2 top-0
      h-[520px] w-[900px]
      -translate-x-1/2
      rounded-full
      bg-blue-400/10
      blur-[120px]
    "
  ></div>

  <div
    bind:this={decorationElement}
    class="
      pointer-events-none absolute
      -right-24 top-20
      h-80 w-80 rotate-45
      border border-blue-300/10
    "
  ></div>

  <div
    class="
      relative mx-auto flex w-full max-w-7xl
      flex-col items-center
    "
  >
    <header
      bind:this={headerElement}
      class="flex w-full flex-col items-center text-center"
    >
      <h1
        bind:this={headingElement}
        class="
          max-w-4xl
          text-4xl leading-[0.98] font-black
          tracking-[-0.055em]
          text-blue-700
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_10px_0_rgba(5,34,70,0.75)]
          sm:text-5xl
          md:text-6xl
          lg:text-7xl
        "
      >
        {headingPrefix}

        {#if headingSuffix}
          <span class="block">
            {headingSuffix}
          </span>
        {/if}
      </h1>

      <p
        bind:this={descriptionElement}
        class="
          mt-7 max-w-3xl
          text-sm leading-6 font-semibold
          text-white
          drop-shadow-[0_2px_2px_rgba(0,0,0,0.5)]
          sm:text-base sm:leading-7
        "
      >
        {unit.description}
      </p>
    </header>

    <div
      bind:this={menkoElement}
      class="mt-10 flex justify-center sm:mt-12"
    >
      <MenkoCard
        name={leaderName}
        position={leaderPosition}
        image={leaderImage}
      />
    </div>

    <div
      bind:this={kementrianSectionElement}
      class="mt-14 w-full sm:mt-16"
    >
      <div class="flex w-full flex-col items-center">
        <h2
          bind:this={kementrianHeadingElement}
          class="
            text-center
            text-4xl leading-none font-black
            tracking-[-0.055em]
            text-blue-700
            [-webkit-text-stroke:1.5px_white]
            [paint-order:stroke_fill]
            drop-shadow-[0_8px_0_rgba(5,34,70,0.75)]
            sm:text-5xl
            lg:text-6xl
          "
        >
          Kementerian
        </h2>
      </div>

      <div
        bind:this={kementrianGridElement}
        class="
          mx-auto mt-10 grid
          w-full max-w-5xl
          grid-cols-2
          gap-5
          sm:gap-7
          md:grid-cols-3
          lg:gap-10
        "
      >
        {#each kementrian as item (item.id)}
          <KementrianCard
            title={item.title}
            slug={item.slug}
            logo={item.logo}
          />
        {/each}
      </div>
    </div>
  </div>

  <img
    bind:this={leftStarOneElement}
    src="/menko/b-3-left.png"
    alt=""
    class="
      absolute left-0 top-1/3
      size-28
      md:size-40
    "
  />

  <img
    bind:this={leftStarTwoElement}
    src="/menko/b-3-left.png"
    alt=""
    class="
      absolute left-0 top-2/3
      size-28
      md:size-40
    "
  />

  <img
    bind:this={rightStarElement}
    src="/menko/b-3-right.png"
    alt=""
    class="
      absolute right-0 top-1/5
      size-28
      md:size-40
    "
  />
</section>