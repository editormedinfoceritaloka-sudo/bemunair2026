<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  let {
    active = true
  }: {
    active?: boolean;
  } = $props();

  const speech = {
    name: 'M. Rizqi Senja Virawan',
    role: 'President of BEM UNAIR 2026',
    message:
      'Halo, Airlangga dan Indonesia! Badan Eksekutif Mahasiswa Universitas Airlangga sebagai organisasi otonom yang menjadi bagian dari masyarakat bertumbuh dan berdampak dengan prinsip kebermanfaatan. Dari Timur Jawa Dwipa, BEM Universitas bertekad untuk mewujudkan organisasi kemahasiswaan yang inklusif dan berkemajuan untuk Airlangga dan Indonesia. Mengusung semangat Panca Loka, kita akan merajut ragam aspirasi menjadi sebuah cerita dan perjalanan yang utuh, memberikan ruang bagi setiap mahasiswa untuk turut andil dalam menciptakan perubahan. Karena untuk membentuk cerita dan perjalanan, kita dapat memaknai bahwa setiap langkah kecil yang dilakukan bersama-sama, akan membawa makna yang jauh lebih besar.'
  };

  let pageRoot: HTMLDivElement;
  let speechCard: HTMLDivElement;
  let speechCopy: HTMLDivElement;
  let portrait: HTMLImageElement;
  let swipeLabel: HTMLParagraphElement;

  let mounted = $state(false);
  let isInView = $state(false);
  let hasAnimated = $state(false);
  let reduceMotion = false;

  let timeline: gsap.core.Timeline | undefined;
  let context: gsap.Context | undefined;
  let scrollTrigger: ScrollTrigger | undefined;
  let stars: HTMLElement[] = [];

  function playEntrance(): void {
    if (
      !mounted ||
      !active ||
      !isInView ||
      hasAnimated
    ) {
      return;
    }

    hasAnimated = true;

    if (reduceMotion) {
      gsap.set(
        [
          speechCard,
          portrait,
          swipeLabel,
          ...stars,
          ...Array.from(speechCopy.children)
        ],
        {
          clearProps: 'all',
          autoAlpha: 1
        }
      );

      return;
    }

    timeline = gsap.timeline({
      defaults: {
        ease: 'power3.out'
      }
    });

    timeline
      .to(
        stars,
        {
          autoAlpha: 1,
          scale: 1,
          rotation: 0,
          duration: 0.7,
          stagger: 0.08
        },
        0
      )
      .to(
        speechCard,
        {
          autoAlpha: 1,
          xPercent: 0,
          scale: 1,
          duration: 1
        },
        0.15
      )
      .to(
        portrait,
        {
          autoAlpha: 1,
          xPercent: 0,
          scale: 1,
          duration: 1.05
        },
        0.25
      )
      .to(
        Array.from(speechCopy.children),
        {
          autoAlpha: 1,
          y: 0,
          duration: 0.65,
          stagger: 0.12,
          ease: 'power2.out'
        },
        0.65
      )
      .to(
        swipeLabel,
        {
          autoAlpha: 1,
          duration: 0.5
        },
        0.9
      );
  }

  $effect(() => {
    active;
    isInView;
    mounted;

    playEntrance();
  });

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    context = gsap.context(() => {
      stars = gsap.utils.toArray<HTMLElement>(
        '.speech-star',
        pageRoot
      );

      if (!reduceMotion) {
        gsap.set(stars, {
          autoAlpha: 0,
          scale: 0.45,
          rotation: -18,
          transformOrigin: 'center'
        });

        gsap.set(speechCard, {
          autoAlpha: 0,
          xPercent: -18,
          scale: 0.97
        });

        gsap.set(portrait, {
          autoAlpha: 0,
          xPercent: 22,
          scale: 0.96,
          transformOrigin: 'center bottom'
        });

        gsap.set(Array.from(speechCopy.children), {
          autoAlpha: 0,
          y: 18
        });

        gsap.set(swipeLabel, {
          autoAlpha: 0
        });
      }

      scrollTrigger = ScrollTrigger.create({
        trigger: pageRoot,
        start: 'top 85%',
        end: 'bottom 15%',
        invalidateOnRefresh: true,

        onEnter: () => {
          isInView = true;
        },

        onEnterBack: () => {
          isInView = true;
        },

        onLeave: () => {
          isInView = false;
        },

        onLeaveBack: () => {
          isInView = false;
        }
      });
    }, pageRoot);

    mounted = true;

    return () => {
      timeline?.kill();
      scrollTrigger?.kill();
      context?.revert();
    };
  });
</script>

<div
  bind:this={pageRoot}
  class="
    relative h-full w-full overflow-hidden rounded-xl
    bg-linear-to-br from-blue-900 via-blue-800 to-blue-500
  "
>
  <div
    aria-hidden="true"
    class="pointer-events-none absolute inset-0 overflow-hidden"
  >
    <img
      src="/landing/speech/star-8.png"
      alt=""
      draggable="false"
      class="
        speech-star absolute -top-2 -right-2
        w-16 object-contain opacity-75
        sm:w-20 md:w-24 lg:w-28
      "
    />

    <img
      src="/landing/speech/star-8.png"
      alt=""
      draggable="false"
      class="
        speech-star absolute -bottom-3 -left-3
        w-20 rotate-180 object-contain opacity-65
        sm:w-24 md:w-28 lg:w-32
      "
    />

    <img
      src="/landing/speech/star-4.png"
      alt=""
      draggable="false"
      class="
        speech-star absolute bottom-[2%] left-[25%]
        w-12 -rotate-12 object-contain opacity-55
        sm:w-16 md:w-20
      "
    />

    <img
      src="/landing/speech/star-4.png"
      alt=""
      draggable="false"
      class="
        speech-star absolute top-[7%] left-[55%]
        w-10 rotate-12 object-contain opacity-55
        sm:w-14 md:w-20
      "
    />
  </div>

  <div class="relative z-10 h-full w-full">
    <div
      bind:this={speechCard}
      class="
        absolute top-1/2 left-[3%] z-20
        w-[82%] -translate-y-1/2
        will-change-transform
        sm:left-[4%] sm:w-[77%]
        md:left-[5%] md:w-[72%]
        lg:left-[6%] lg:w-[68%]
      "
    >
      <div class="relative">
        <div
          aria-hidden="true"
          class="
            absolute -inset-x-1 -top-7 bottom-1
            -rotate-3 rounded-xl
            border-2 border-sky-400
            bg-blue-400/65 shadow-xl
            [clip-path:polygon(0_0,52%_0,59%_15%,100%_8%,100%_100%,0_100%)]
            sm:-top-9
            md:-top-11
          "
        ></div>

        <div
          class="
            relative overflow-hidden
            rounded-xl bg-blue-100/95
            px-5 pt-12 pb-5
            shadow-xl backdrop-blur-md
            [clip-path:polygon(0_0,50%_0,57%_13%,100%_13%,100%_100%,0_100%)]
            sm:px-7 sm:pt-14 sm:pb-6
            md:px-8 md:pt-16 md:pb-7
            lg:px-10 lg:pt-20 lg:pb-8
          "
        >
          <div bind:this={speechCopy}>
            <p
              class="
                text-[8px] leading-[1.65] text-black-900
                sm:text-[10px] sm:leading-[1.7]
                md:text-xs md:leading-[1.75]
                lg:text-sm lg:leading-[1.8]
              "
            >
              {speech.message}
            </p>

            <div class="mt-4 text-right text-black-900 md:mt-6">
              <p
                class="
                  text-[10px] leading-tight font-bold
                  sm:text-xs
                  md:text-base
                  lg:text-lg
                "
              >
                {speech.name}
              </p>

              <p
                class="
                  text-[8px] leading-tight font-medium
                  sm:text-[10px]
                  md:text-sm
                  lg:text-base
                "
              >
                {speech.role}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <img
      bind:this={portrait}
      src="/landing/speech/p.png"
      alt="Presiden BEM UNAIR 2026"
      draggable="false"
      class="
        pointer-events-none absolute -right-[1%] bottom-0 z-30
        h-[65%] w-auto object-contain object-bottom
        will-change-transform
        drop-shadow-[0_14px_20px_rgba(0,0,0,0.32)]
        sm:h-[74%]
        md:-right-[2%] md:h-[84%]
        lg:h-[91%]
      "
    />

    <p
      bind:this={swipeLabel}
      class="
        absolute top-1/2 right-1 z-40
        hidden -translate-y-1/2 rotate-180
        text-[7px] font-medium tracking-wide text-white/85
        [writing-mode:vertical-rl]
        md:block lg:text-[9px]
      "
    >
      swipe to flip page
    </p>
  </div>
</div>