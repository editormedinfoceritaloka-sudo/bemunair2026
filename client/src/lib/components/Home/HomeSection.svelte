<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { playShineSound } from '$lib/sound';

  let heroSection: HTMLElement;
  let leftSide: HTMLDivElement;
  let rightSide: HTMLDivElement;
  let bottomSide: HTMLDivElement;
  let centerText: HTMLDivElement;
  let taglineElement: HTMLParagraphElement;
  let introOverlay: HTMLDivElement;
  let cloudLeft: HTMLDivElement;
  let cloudRight: HTMLDivElement;
  let introSky: HTMLDivElement;
  let introLogo: HTMLDivElement;
  let introLogoGlow: HTMLDivElement;
  let introVideo: HTMLVideoElement;

  const logoHold = 3.5;

  let useAlphaVideo = $state(true);

  onMount(() => {
    useAlphaVideo = !/^((?!chrome|android|crios|fxios).)*safari/i.test(
      navigator.userAgent
    );

    const context = gsap.context(() => {
      const reduceMotion = window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

      gsap.set(leftSide, {
        xPercent: -110
      });

      gsap.set(rightSide, {
        xPercent: 110
      });

      gsap.set(bottomSide, {
        yPercent: 100,
        opacity: 0
      });

      gsap.set(centerText, {
        opacity: 0
      });

      gsap.set(taglineElement, {
        autoAlpha: 0,
        y: 16,
        filter: 'blur(6px)'
      });

      gsap.set(introLogo, {
        autoAlpha: 0,
        scale: 0.82,
        filter: 'blur(8px)'
      });

      gsap.set(introLogoGlow, {
        autoAlpha: 0,
        scale: 0.6
      });

      const timeline = gsap.timeline({
        defaults: {
          ease: 'power3.out'
        }
      });

      if (reduceMotion) {
        gsap.set(introOverlay, {
          autoAlpha: 0
        });

        timeline.addLabel('reveal', 0);
      } else {
        timeline
          .to(
            cloudLeft,
            {
              xPercent: -105,
              duration: 2.6,
              ease: 'sine.inOut'
            },
            0.4
          )
          .to(
            cloudLeft,
            {
              opacity: 0,
              duration: 3.2,
              ease: 'power1.in'
            },
            0.4
          )
          .to(
            cloudRight,
            {
              xPercent: 105,
              duration: 2.6,
              ease: 'sine.inOut'
            },
            0.4
          )
          .to(
            cloudRight,
            {
              opacity: 0,
              duration: 3.2,
              ease: 'power1.in'
            },
            0.4
          )
          .to(
            introLogo,
            {
              autoAlpha: 1,
              scale: 1,
              filter: 'blur(0px)',
              duration: 1.6,
              ease: 'power2.out'
            },
            1
          )
          .call(
            () => {
              introVideo.currentTime = 0;
              introVideo.playbackRate = 1.25;
              void introVideo.play().catch(() => {});
            },
            [],
            1
          )
          .to(
            introLogoGlow,
            {
              autoAlpha: 1,
              scale: 1.15,
              duration: 1.2,
              ease: 'sine.out'
            },
            1.6 + logoHold
          )
          .to(
            introLogo,
            {
              y: -6,
              duration: 0.9,
              ease: 'sine.inOut',
              yoyo: true,
              repeat: 1
            },
            2.2 + logoHold
          )
          .call(
            () => {
              const rect = heroSection.getBoundingClientRect();
              const heroVisible =
                rect.bottom > window.innerHeight * 0.4 &&
                rect.top < window.innerHeight * 0.6;

              if (heroVisible) {
                playShineSound();
              }
            },
            [],
            3.3 + logoHold
          )
          .to(
            introLogoGlow,
            {
              autoAlpha: 0,
              scale: 1.4,
              duration: 1,
              ease: 'sine.inOut'
            },
            3.5 + logoHold
          )
          .to(
            introLogo,
            {
              autoAlpha: 0,
              scale: 1.06,
              filter: 'blur(6px)',
              duration: 1,
              ease: 'sine.inOut'
            },
            3.6 + logoHold
          )
          .to(
            introSky,
            {
              autoAlpha: 0,
              duration: 1.6,
              ease: 'sine.inOut'
            },
            3.4 + logoHold
          )
          .set(
            introOverlay,
            {
              autoAlpha: 0
            },
            5 + logoHold
          )
          .addLabel('reveal', 3.9 + logoHold);
      }

      timeline
        .to(
          leftSide,
          {
            xPercent: 0,
            duration: 1.2
          },
          'reveal'
        )
        .to(
          rightSide,
          {
            xPercent: 0,
            duration: 1.2
          },
          'reveal'
        )
        .call(
          () => {
            window.dispatchEvent(
              new CustomEvent('hero-bottom-enter')
            );
          },
          [],
          'reveal+=0.85'
        )
        .to(
          bottomSide,
          {
            yPercent: 0,
            opacity: 1,
            duration: 0.9
          },
          'reveal+=0.85'
        )
        .to(
          centerText,
          {
            opacity: 1,
            duration: 0.8,
            ease: 'power2.out'
          },
          'reveal+=1.55'
        )
        .to(
          taglineElement,
          {
            autoAlpha: 1,
            y: 0,
            filter: 'blur(0px)',
            duration: 0.9,
            ease: 'power2.out',
            clearProps: 'transform,filter'
          },
          'reveal+=1.9'
        );
    }, heroSection);

    return () => {
      context.revert();
    };
  });
</script>

<section
  bind:this={heroSection}
  id="home"
  class="
    hero relative isolate flex items-center justify-center
    overflow-hidden bg-linear-to-b from-blue-200 via-blue-50 to-white
    px-4 sm:px-6
  "
>
  <div aria-hidden="true" class="pointer-events-none absolute inset-0">
    <div
      bind:this={leftSide}
      class="absolute top-0 left-0 h-[78%] w-1/3 will-change-transform"
    >
      <img
        src="/landing/home/left.webp"
        alt="Ilustrasi landing bagian kiri"
        class="h-full w-full -rotate-4 object-contain object-[left_top] md:rotate-0"
        draggable="false"
      />
    </div>

    <div
      bind:this={rightSide}
      class="absolute top-0 right-0 h-[78%] w-1/3 will-change-transform"
    >
      <img
        src="/landing/home/right.webp"
        alt="Ilustrasi landing bagian kanan"
        class="h-full w-full object-contain object-[right_top]"
        draggable="false"
      />
    </div>

    <div
      bind:this={bottomSide}
      class="hero-wave absolute inset-x-0 bottom-0 w-full will-change-transform"
    >
      <img
        src="/landing/home/bottom.webp"
        alt="Ilustrasi landing bagian bawah"
        class="block h-auto w-full"
        draggable="false"
      />
    </div>
  </div>

  <div
    bind:this={centerText}
    class="
      relative z-10 flex w-full max-w-5xl min-w-0 flex-col
      items-center gap-3 text-center sm:gap-5 md:gap-6
    "
  >
    <h1
      class="
        hero-title bg-linear-to-b
        from-blue-900
        via-blue-700
        to-blue-400
        bg-clip-text
        leading-none
        font-black
        tracking-[-0.06em]
        text-transparent
        drop-shadow-[0_8px_6px_rgba(15,23,42,0.28)]
      "
    >
      BEM UNAIR
    </h1>

    <p
      class="
        hero-subtitle text-center
        leading-none
        font-black
        tracking-[-0.04em]
        text-orange-600
        drop-shadow-[0_4px_3px_rgba(120,53,15,0.35)]
      "
    >
      Kabinet Cerita Loka
    </p>

    <p
      bind:this={taglineElement}
      class="
        hero-tagline max-w-xl text-center leading-relaxed
      "
    >
      BEM UNAIR sebagai Rumah dan Kompas Gerakan yang Berlandaskan Spirit Panca Loka guna Kebermanfaatan Bersama bagi Airlangga dan Indonesia.
    </p>
  </div>

  <div
    bind:this={introOverlay}
    aria-hidden="true"
    class="pointer-events-none absolute inset-0 z-30 overflow-hidden"
  >
    <div
      bind:this={introSky}
      class="absolute inset-0 bg-linear-to-b from-[#a9cbf0] via-[#c6ddf6] to-[#e4eefa]"
    ></div>

    <div
      bind:this={cloudLeft}
      class="absolute -inset-y-[10%] -left-[10%] w-[72%] will-change-[transform,opacity]"
    >
      <img
        src="/landing/home/cloud-left.webp"
        alt=""
        class="h-full w-full object-cover object-right"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-1.webp"
        alt=""
        class="absolute left-[50%] top-[4%] w-[80%] max-w-none"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-2.webp"
        alt=""
        class="absolute left-[58%] top-[46%] w-[75%] max-w-none"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-3.webp"
        alt=""
        class="absolute left-[70%] top-[24%] w-[60%] max-w-none"
        draggable="false"
      />
    </div>

    <div
      bind:this={cloudRight}
      class="absolute -inset-y-[10%] -right-[10%] w-[72%] will-change-[transform,opacity]"
    >
      <img
        src="/landing/home/cloud-right.webp"
        alt=""
        class="h-full w-full object-cover object-left"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-3.webp"
        alt=""
        class="absolute right-[50%] top-[22%] w-[80%] max-w-none -scale-x-100"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-1.webp"
        alt=""
        class="absolute right-[58%] top-[60%] w-[75%] max-w-none -scale-x-100"
        draggable="false"
      />
      <img
        src="/landing/home/cloud-puff-2.webp"
        alt=""
        class="absolute right-[70%] top-[2%] w-[60%] max-w-none -scale-x-100"
        draggable="false"
      />
    </div>

    <div class="absolute inset-0 flex items-center justify-center">
      <div
        bind:this={introLogoGlow}
        class="absolute size-80 rounded-full bg-[radial-gradient(circle,rgba(255,255,255,0.95)_0%,rgba(191,219,254,0.55)_40%,rgba(255,255,255,0)_70%)] blur-xl md:size-[28rem]"
      ></div>

      <div
        bind:this={introLogo}
        class={`relative will-change-[transform,opacity,filter] ${useAlphaVideo ? '' : 'mix-blend-darken'}`}
      >
        <video
          bind:this={introVideo}
          src={useAlphaVideo
            ? '/landing/home/logo-motion.webm'
            : '/landing/home/logo-motion.mp4'}
          muted
          playsinline
          preload="auto"
          class="size-44 object-contain md:size-72"
        ></video>
      </div>
    </div>
  </div>
</section>

<style>
  .hero {
    --hero-wave-space: clamp(6rem, 18vw, 12rem);
    min-height: 45svh;
    padding-block: clamp(6rem, 12vh, 9rem) var(--hero-wave-space);
  }

  .hero-wave img {
    max-height: var(--hero-wave-space);
  }

  .hero-title {
    font-size: clamp(2.5rem, 9.5vw, 9rem);
  }

  .hero-subtitle {
    font-size: clamp(1.25rem, 3.5vw, 3.5rem);
  }

  .hero-tagline {
    font-size: clamp(0.95rem, 1.65vw, 1.35rem);
  }

  @media (min-width: 40rem) {
    .hero {
      min-height: 80svh;
    }
  }

  @media (min-width: 48rem) {
    .hero {
      min-height: 100svh;
    }
  }
</style>
