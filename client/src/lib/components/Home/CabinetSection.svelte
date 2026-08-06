<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import LogoCarrousel from './cabinet/LogoCarrousel.svelte';
  import MenkoSlider from './cabinet/MenkoSlider.svelte';
  import type { Cabinet } from '$lib/types';

  let { cabinet }:{ cabinet: Cabinet } = $props();

  let sectionElement!: HTMLElement;
  let headingElement!: HTMLDivElement;
  let titleElement!: HTMLHeadingElement;
  let subtitleElement!: HTMLParagraphElement;
  let carouselElement!: HTMLDivElement;
  let sliderElement!: HTMLDivElement;

  onMount(() => {
    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      gsap.set(
        [
          headingElement,
          titleElement,
          subtitleElement,
          carouselElement,
          sliderElement
        ],
        {
          clearProps: 'all'
        }
      );

      return;
    }

    const context = gsap.context(() => {
      gsap.set(titleElement, {
        y: 70,
        scale: 0.88,
        opacity: 0
      });

      gsap.set(subtitleElement, {
        y: 35,
        opacity: 0
      });

      gsap.set(carouselElement, {
        y: 50,
        opacity: 0
      });

      gsap.set(sliderElement, {
        y: 70,
        scale: 0.96,
        opacity: 0
      });
    }, sectionElement);

    let animation: gsap.core.Timeline | undefined;

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (!entry.isIntersecting) {
          return;
        }

        animation?.kill();

        animation = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          }
        });

        animation
          .to(titleElement, {
            y: 0,
            scale: 1,
            opacity: 1,
            duration: 0.9
          })
          .to(
            subtitleElement,
            {
              y: 0,
              opacity: 1,
              duration: 0.7
            },
            '-=0.55'
          )
          .to(
            carouselElement,
            {
              y: 0,
              opacity: 1,
              duration: 0.85
            },
            '-=0.35'
          )
          .to(
            sliderElement,
            {
              y: 0,
              scale: 1,
              opacity: 1,
              duration: 0.95
            },
            '-=0.5'
          );

        observer.unobserve(sectionElement);
      },
      {
        threshold: 0.2
      }
    );

    observer.observe(sectionElement);

    return () => {
      observer.disconnect();
      animation?.kill();
      context.revert();
    };
  });
</script>

<section
  bind:this={sectionElement}
  id="cabinet"
  class="
    min-h-screen w-full overflow-hidden
    bg-linear-to-b from-blue-100 via-blue-50 to-white
    py-24
    sm:py-28
    md:py-32
  "
>
  <div class="mx-auto flex w-full max-w-7xl flex-col items-center">
    <div
      bind:this={headingElement}
      class="
        flex flex-col items-center
        text-center font-black uppercase
        leading-none tracking-[-0.05em]
      "
    >
      <h2
        bind:this={titleElement}
        class="
          whitespace-nowrap
          text-[clamp(2.3rem,8vw,7rem)]
          text-orange-600
          [-webkit-text-stroke:1.5px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_5px_3px_rgba(15,23,42,0.3)]
          will-change-[opacity,transform]
          sm:[-webkit-text-stroke:2px_white]
          md:[-webkit-text-stroke:3px_white]
        "
      >
        Kementerian
      </h2>

      <p
        bind:this={subtitleElement}
        class="
          mt-3 whitespace-nowrap
          text-[clamp(1.1rem,3vw,2.5rem)]
          text-blue-700
          [-webkit-text-stroke:1px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_3px_2px_rgba(15,23,42,0.3)]
          will-change-[opacity,transform]
          sm:mt-4
          sm:[-webkit-text-stroke:1.5px_white]
          md:[-webkit-text-stroke:2px_white]
        "
      >
        Cerita Loka
      </p>
    </div>

    <div
      bind:this={carouselElement}
      class="
        mt-6 w-full
        will-change-[opacity,transform]
        sm:mt-8
        md:mt-10
      "
    >
      <LogoCarrousel kementrian={cabinet} />
    </div>

    <div
      bind:this={sliderElement}
      class="
        mt-10 w-full
        will-change-[opacity,transform]
        sm:mt-14
        md:mt-16
      "
    >
      <MenkoSlider autoplayDelay={3000}  kemenkoan={cabinet.kemenkoan} />
    </div>
  </div>
</section>