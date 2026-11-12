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
    name: 'Shintya Iftitah',
    role: 'Vice President of BEM UNAIR 2026',
    message:
      'Halo Ksatria Airlangga! Menjadi mahasiswa adalah sebuah privilege untuk menghadirkan makna yang lebih besar. Badan Eksekutif Mahasiswa hadir untuk mengkolektifkan semangat kebermaknaan itu melalui berbagai gagasan yang diharapkan mampu menjawab kebutuhan mahasiswa dan masyarakat Indonesia. Universitas Airlangga adalah cerita tentang rumah dengan selaksa cinta di dalamnya. Berpartisipasi dalam menuliskan cerita bersamanya adalah pengalaman paling menyenangkan yang tidak boleh sampai dilewatkan.',
    closing:
      'Salam hangat dari kami. BEM Unair Kabinet Cerita Loka'
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
  let stars: HTMLElement[] = [];

  let timeline: gsap.core.Timeline | undefined;
  let context: gsap.Context | undefined;
  let scrollTrigger: ScrollTrigger | undefined;

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

    const copyElements =
      Array.from(speechCopy.children);

    if (reduceMotion) {
      gsap.set(
        [
          speechCard,
          portrait,
          swipeLabel,
          ...stars,
          ...copyElements
        ],
        {
          autoAlpha: 1,
          clearProps: 'all'
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
          duration: 0.6,
          stagger: 0.07
        },
        0
      )
      .to(
        portrait,
        {
          autoAlpha: 1,
          xPercent: 0,
          scale: 1,
          duration: 0.9
        },
        0.1
      )
      .to(
        speechCard,
        {
          autoAlpha: 1,
          xPercent: 0,
          scale: 1,
          duration: 0.85
        },
        0.2
      )
      .to(
        copyElements,
        {
          autoAlpha: 1,
          y: 0,
          duration: 0.5,
          stagger: 0.1,
          ease: 'power2.out'
        },
        0.5
      )
      .to(
        swipeLabel,
        {
          autoAlpha: 1,
          duration: 0.4
        },
        0.75
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
          scale: 0.5,
          rotation: 15,
          transformOrigin: 'center'
        });

        gsap.set(portrait, {
          autoAlpha: 0,
          xPercent: -18,
          scale: 0.97,
          transformOrigin: 'center bottom'
        });

        gsap.set(speechCard, {
          autoAlpha: 0,
          xPercent: 14,
          scale: 0.98
        });

        gsap.set(
          Array.from(speechCopy.children),
          {
            autoAlpha: 0,
            y: 12
          }
        );

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
    relative h-full w-full
    overflow-hidden rounded-xl
    bg-linear-to-br
    from-blue-900
    via-blue-800
    to-blue-500
  "
>
  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute inset-0
      overflow-hidden
    "
  >
    <img
      src="/landing/speech/star-8.webp"
      alt=""
      draggable="false"
      class="
        speech-star
        absolute -top-2 -right-2
        w-12 object-contain opacity-75
        sm:w-16
        md:w-20
        lg:w-24
      "
    />

    <img
      src="/landing/speech/star-8.webp"
      alt=""
      draggable="false"
      class="
        speech-star
        absolute -bottom-2 -left-2
        w-14 rotate-180
        object-contain opacity-65
        sm:w-20
        md:w-24
        lg:w-28
      "
    />

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        speech-star
        absolute top-[7%]
        left-[25%]
        w-8 -rotate-12
        object-contain opacity-55
        sm:w-10
        md:w-14
      "
    />

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        speech-star
        absolute right-[15%]
        bottom-[3%]
        w-9 rotate-12
        object-contain opacity-55
        sm:w-12
        md:w-16
      "
    />
  </div>

  <div class="relative z-10 h-full w-full">
    <img
      bind:this={portrait}
      src="/landing/speech/w.webp"
      alt="Wakil Presiden BEM UNAIR 2026"
      draggable="false"
      class="
        pointer-events-none
        absolute -bottom-[1%]
        -left-[1%] z-30
        h-[38%]
        w-auto
        object-contain
        object-bottom
        will-change-transform
        drop-shadow-[0_10px_16px_rgba(0,0,0,0.32)]
        sm:h-[65%]
        md:-left-[2%]
        md:h-[78%]
        lg:h-[85%]
      "
    />

    <div
      bind:this={speechCard}
      class="
        absolute top-1/2
        right-[3%] z-20
        w-[84%]
        -translate-y-1/2
        will-change-transform
        sm:right-[4%]
        sm:w-[76%]
        md:right-[5%]
        md:w-[70%]
        lg:right-[6%]
        lg:w-[66%]
      "
    >
      <div class="relative">
        <div
          aria-hidden="true"
          class="
            absolute
            -inset-x-1
            -top-5 bottom-1
            rotate-3
            rounded-xl
            border-2
            border-sky-400
            bg-blue-400/65
            shadow-xl
            [clip-path:polygon(0_8%,41%_15%,48%_0,100%_0,100%_100%,0_100%)]
            sm:-top-6
            md:-top-8
          "
        ></div>

        <div
          class="
            relative
            max-h-[275px]
            overflow-y-auto
            rounded-xl
            bg-blue-100/95
            px-3
            pt-8
            pb-3
            shadow-xl
            backdrop-blur-md
            [clip-path:polygon(0_13%,43%_13%,50%_0,100%_0,100%_100%,0_100%)]
            [scrollbar-width:thin]
            sm:max-h-[350px]
            sm:px-4
            sm:pt-10
            sm:pb-4
            md:max-h-[420px]
            md:px-5
            md:pt-12
            md:pb-5
            lg:max-h-[470px]
            lg:px-6
            lg:pt-14
            lg:pb-6
          "
        >
          <div bind:this={speechCopy}>
            <p
              class="
                text-justify
                text-[7px]
                leading-[1.5]
                text-[#0a2445]
                sm:text-[8px]
                sm:leading-[1.6]
                md:text-[10px]
                md:leading-[1.65]
                lg:text-xs
                lg:leading-[1.7]
              "
            >
              {speech.message}
            </p>

            <p
              class="
                mt-3
                text-[7px]
                leading-[1.5]
                text-[#0a2445]
                sm:text-[8px]
                md:mt-4
                md:text-[10px]
                lg:text-xs
              "
            >
              {speech.closing}
            </p>

            <div
              class="
                mt-2
                text-left
                text-[#0a2445]
                sm:mt-3
              "
            >
              <p
                class="
                  text-[9px]
                  leading-tight
                  font-bold
                  sm:text-[10px]
                  md:text-sm
                  lg:text-base
                "
              >
                {speech.name}
              </p>

              <p
                class="
                  mt-0.5
                  text-[7px]
                  leading-tight
                  font-medium
                  sm:text-[9px]
                  md:text-xs
                  lg:text-sm
                "
              >
                {speech.role}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <p
      bind:this={swipeLabel}
      class="
        absolute top-1/2
        right-1 z-40
        hidden
        -translate-y-1/2
        rotate-180
        text-[7px]
        font-medium
        tracking-wide
        text-white/85
        [writing-mode:vertical-rl]
        md:block
        lg:text-[8px]
      "
    >
      swipe to flip page
    </p>
  </div>
</div>