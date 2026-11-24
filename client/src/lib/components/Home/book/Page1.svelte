<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';
  import { playShineSound, playTypeSound } from '$lib/sound';

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
  let typedMessage = $state('');
  let typedName = $state('');
  let typedRole = $state('');
  let showCling = $state(false);

  let reduceMotion = false;
  let stars: HTMLElement[] = [];

  let timeline: gsap.core.Timeline | undefined;
  let typingTweens: gsap.core.Tween[] = [];
  let context: gsap.Context | undefined;
  let scrollTrigger: ScrollTrigger | undefined;

  function typeText(
    value: string,
    onUpdate: (text: string) => void,
    onComplete: () => void
  ): gsap.core.Tween {
    const proxy = { length: 0 };
    const duration = Math.max(0.8, Math.min(7, value.length / 90));

    return gsap.to(proxy, {
      length: value.length,
      duration,
      delay: 0.7,
      ease: 'none',
      onUpdate: () => {
        onUpdate(value.slice(0, Math.round(proxy.length)));

        if (active) {
          playTypeSound();
        }
      },
      onComplete
    });
  }

  function startTyping(): void {
    typingTweens.forEach((tween) => tween.kill());
    typingTweens = [];

    if (reduceMotion) {
      typedMessage = speech.message;
      typedName = speech.name;
      typedRole = speech.role;
      return;
    }

    typedMessage = '';
    typedName = '';
    typedRole = '';
    showCling = false;

    typingTweens.push(
      typeText(
        speech.message,
        (text) => (typedMessage = text),
        () => {
          typedMessage = speech.message;
          typingTweens.push(
            typeText(
              speech.name,
              (text) => (typedName = text),
              () => {
                typedName = speech.name;
                typingTweens.push(
                  typeText(
                    speech.role,
                    (text) => (typedRole = text),
                    () => {
                      typedRole = speech.role;
                      typingTweens.push(
                        gsap.delayedCall(0.45, () => {
                          showCling = true;

                          if (active) {
                            playShineSound();
                          }
                        })
                      );
                    }
                  )
                );
              }
            )
          );
        }
      )
    );
  }

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
      typedMessage = speech.message;
      typedName = speech.name;
      typedRole = speech.role;

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

    startTyping();
  }

  $effect(() => {
    void [active, isInView, mounted];

    playEntrance();
  });

  $effect(() => {
    typingTweens.forEach((tween) => {
      if (active) {
        tween.resume();
      } else {
        tween.pause();
      }
    });
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
          rotation: -15,
          transformOrigin: 'center'
        });

        gsap.set(speechCard, {
          autoAlpha: 1,
          clearProps: 'transform'
        });

        gsap.set(portrait, {
          autoAlpha: 0,
          xPercent: 18,
          scale: 0.97,
          transformOrigin: 'center bottom'
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
      typingTweens.forEach((tween) => tween.kill());
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

    <div
      class="
        absolute -bottom-2 -left-2
        rotate-180
      "
    >
      <img
        src="/landing/speech/star-8.webp"
        alt=""
        draggable="false"
        class="
          speech-star
          block
          w-14 object-contain opacity-65
          sm:w-20
          md:w-24
          lg:w-28
        "
      />
    </div>

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        twinkle
        speech-star
        absolute bottom-[2%]
        left-[25%]
        w-10 -rotate-12
        object-contain opacity-55
        sm:w-12
        md:w-16
      "
    />

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        twinkle twinkle-delay
        speech-star
        absolute top-[7%]
        left-[55%]
        w-8 rotate-12
        object-contain opacity-55
        sm:w-10
        md:w-14
      "
    />
  </div>

  <div class="relative z-10 h-full w-full">
    <div
      bind:this={speechCard}
      class="
        absolute top-[6%]
        left-[3%] z-20
        w-[94%]
        translate-y-0
        will-change-transform
        sm:top-1/2
        sm:left-[4%]
        sm:w-[76%]
        sm:-translate-y-1/2
        md:left-[5%]
        md:w-[70%]
        lg:left-[6%]
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
            -rotate-3
            rounded-xl
            border-2
            border-sky-400
            bg-blue-400/65
            shadow-xl
            [clip-path:polygon(0_0,52%_0,59%_15%,100%_8%,100%_100%,0_100%)]
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
            [clip-path:polygon(0_0,50%_0,57%_13%,100%_13%,100%_100%,0_100%)]
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
                text-[9.5px]
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
              {typedMessage}
            </p>

            <div
              class="
                mt-3
                text-right
                text-[#0a2445]
                sm:mt-4
                md:mt-5
              "
            >
              <p
                class="
                  block text-[9px]
                  leading-tight
                  font-bold
                  sm:text-[10px]
                  md:text-sm
                  lg:text-base
                "
              >
                <span class="group relative isolate inline-grid">
                  <span class="invisible col-start-1 row-start-1">
                    {speech.name}
                  </span>
                  <span
                    aria-hidden="true"
                    class:cling={showCling}
                    class="
                      pointer-events-none absolute inset-x-0
                      top-1/2 -z-10 h-[1.5em]
                      -translate-y-1/2 scale-x-0
                      rounded-full bg-linear-to-r
                      from-transparent via-white/80
                      to-transparent opacity-0
                      shadow-[0_0_14px_rgba(255,255,255,0.75)]
                      transition-[transform,opacity]
                      duration-200 ease-out
                      will-change-[transform,opacity]
                      group-hover:scale-x-100
                      group-hover:opacity-100
                      motion-reduce:transition-none
                    "
                  ></span>
                  <span class="relative z-10 col-start-1 row-start-1">
                    {typedName}
                  </span>
                </span>
              </p>

              <p
                class="
                  mt-0.5 block text-[7px]
                  leading-tight
                  font-medium
                  sm:text-[9px]
                  md:text-xs
                  lg:text-sm
                "
              >
                <span class="group relative isolate inline-grid">
                  <span class="invisible col-start-1 row-start-1">
                    {speech.role}
                  </span>
                  <span
                    aria-hidden="true"
                    class:cling={showCling}
                    class="
                      pointer-events-none absolute inset-x-0
                      top-1/2 -z-10 h-[1.5em]
                      -translate-y-1/2 scale-x-0
                      rounded-full bg-linear-to-r
                      from-transparent via-white/80
                      to-transparent opacity-0
                      shadow-[0_0_14px_rgba(255,255,255,0.75)]
                      transition-[transform,opacity]
                      duration-200 ease-out
                      will-change-[transform,opacity]
                      group-hover:scale-x-100
                      group-hover:opacity-100
                      motion-reduce:transition-none
                    "
                  ></span>
                  <span class="relative z-10 col-start-1 row-start-1">
                    {typedRole}
                  </span>
                </span>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <img
      bind:this={portrait}
      src="https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/PRESIDENT_SENJA"
      alt="Presiden BEM UNAIR 2026"
      draggable="false"
      class="
        pointer-events-none
        absolute -right-[1%]
        bottom-0 z-30
        h-[68%]
        w-auto
        object-contain
        object-bottom
        will-change-transform
        drop-shadow-[0_10px_16px_rgba(0,0,0,0.32)]
        sm:h-[65%]
        md:-right-[2%]
        md:h-[75%]
        lg:h-[85%]
      "
    />

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