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
    name: 'Shintya Iftitah Dzikrullah',
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
  let typedMessage = $state('');
  let typedClosing = $state('');
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
      typedClosing = speech.closing;
      typedName = speech.name;
      typedRole = speech.role;
      return;
    }

    typedMessage = '';
    typedClosing = '';
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
              speech.closing,
              (text) => (typedClosing = text),
              () => {
                typedClosing = speech.closing;
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
      typedClosing = speech.closing;
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
        0.1
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
          autoAlpha: 1,
          clearProps: 'transform'
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
        w-16 object-contain opacity-75
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
        w-20 rotate-180
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
        twinkle
        speech-star
        absolute top-[7%]
        left-[25%]
        w-14 -rotate-12
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
        twinkle twinkle-delay
        speech-star
        absolute         right-[15%]
        bottom-[3%]
        w-14 rotate-12
        object-contain opacity-55
        sm:w-12
        md:w-16
      "
    />

    <div
      class="
        absolute left-[2%] bottom-[8%]
        hidden aspect-square w-[34%]
        rounded-full
        bg-[radial-gradient(circle,rgba(147,197,253,0.35)_0%,rgba(147,197,253,0)_70%)]
        sm:block
      "
    ></div>

    <div
      class="
        absolute bottom-[7%] right-[34%]
        hidden h-16 w-36
        bg-[radial-gradient(rgba(255,255,255,0.4)_1px,transparent_1.6px)]
        bg-size-[12px_12px]
        opacity-60
        [mask-image:radial-gradient(ellipse,black_40%,transparent_75%)]
        sm:block
        md:h-20 md:w-44
      "
    ></div>

    <span
      class="
        absolute top-[1%] right-[9%]
        hidden
        font-serif text-[110px]
        leading-none
        text-white/10
        sm:block
        md:text-[150px]
      "
    >
      ”
    </span>

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        twinkle twinkle-delay
        speech-star
        absolute top-[14%] right-[30%]
        hidden w-7 rotate-[18deg]
        object-contain opacity-45
        sm:block
        md:w-9
      "
    />

    <img
      src="/landing/speech/star-4.webp"
      alt=""
      draggable="false"
      class="
        twinkle
        speech-star
        absolute bottom-[16%] left-[42%]
        hidden w-5 -rotate-[10deg]
        object-contain opacity-40
        sm:block
        md:w-7
      "
    />
  </div>

  <div class="relative z-10 h-full w-full">
    <img
      bind:this={portrait}
      src="https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/VICE_PRESIDENT_SHINTYA"
      alt="Wakil Presiden BEM UNAIR 2026"
      draggable="false"
      class="
        pointer-events-none
        absolute -bottom-[1%]
        -left-[8%] z-30
        h-[68%]
        w-auto
        object-contain
        object-bottom
        will-change-transform
        drop-shadow-[0_10px_16px_rgba(0,0,0,0.32)]
        sm:-left-[1%]
        sm:h-[65%]
        md:-left-[2%]
        md:h-[78%]
        lg:h-[85%]
      "
    />

    <div
      aria-hidden="true"
      class="
        pointer-events-none
        absolute bottom-[6%] right-[4%] z-20
        flex w-[44%] flex-col items-end
        text-right
        sm:hidden
      "
    >
      <span
        class="
          -mb-5
          font-serif text-[72px]
          leading-none
          text-white/25
        "
      >
        ”
      </span>

      <p
        class="
          mt-2
          text-[8px]
          font-bold uppercase
          tracking-[0.25em]
          text-sky-200/90
        "
      >
        Vice President
      </p>

      <p
        class="
          mt-1
          whitespace-nowrap
          text-[13px]
          leading-tight
          font-black
          text-white
        "
      >
        Merajut Cerita,
        <br />
        Tumbuh Bersama
      </p>

      <span class="mt-2.5 h-px w-14 bg-white/40"></span>

      <p
        class="
          mt-2
          text-[8px]
          font-semibold uppercase
          tracking-[0.2em]
          text-white/70
        "
      >
        Kabinet Cerita Loka
      </p>
    </div>

    <div
      bind:this={speechCard}
      class="
        absolute top-[6%]
        right-[3%] z-20
        w-[94%]
        translate-y-0
        will-change-transform
        sm:top-1/2
        sm:right-[4%]
        sm:w-[76%]
        sm:-translate-y-1/2
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
            max-h-[62%]
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
            <div class="speech-message">
              <span
                class="speech-wrap-shape speech-wrap-left"
                aria-hidden="true"
              ></span>

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

              <p
                class="
                  mt-3
                  text-[9.5px]
                  leading-[1.5]
                  text-[#0a2445]
                  sm:text-[8px]
                  md:mt-4
                  md:text-[10px]
                  lg:text-xs
                "
              >
                {typedClosing}
              </p>
            </div>

            <div
              class="
                mt-2
                text-right
                text-[#0a2445]
                sm:mt-3 sm:text-left
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

<style>
  @media (max-width: 639px) {
    .speech-message::after {
      content: '';
      display: block;
      clear: both;
    }

    /*
      Kontur kepala + bahu (mobile), versi mirror: foto di kiri.
      Float selebar teks, tingginya juga jadi tinggi minimum paragraf.
      - sumbu X pakai %, relatif ke lebar kolom teks
      - sumbu Y pakai px, relatif ke atas paragraf
      Atas (y 0-26px)   : takik/notch kiri card
      Tengah            : kosong -> teks melebar penuh
      y ~128px & ke bawah : hijab -> leher -> bahu
    */
    .speech-wrap-shape {
      display: block;
      width: 100%;
      height: 200px;
      margin: 0;
    }

    .speech-wrap-left {
      float: left;
      shape-outside: polygon(
        0 0,
        47.7% 0,
        44.5% 26px,
        0.1% 26px,
        0.1% 134px,
        23.7% 134px,
        29.4% 144px,
        33.2% 154px,
        35.7% 184px,
        35.1% 200px,
        0 200px
      );
    }
  }

  @media (min-width: 640px) {
    .speech-wrap-shape {
      display: none;
    }
  }
</style>