<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { resolve } from '$app/paths';
  import { page } from '$app/state';
  import { gsap } from 'gsap';

  let sectionElement!: HTMLElement;
  let glowElement!: HTMLDivElement;
  let orbitElement!: HTMLDivElement;
  let badgeElement!: HTMLDivElement;
  let codeElement!: HTMLDivElement;
  let titleElement!: HTMLHeadingElement;
  let descriptionElement!: HTMLParagraphElement;
  let actionsElement!: HTMLDivElement;
  let logoElement!: HTMLDivElement;

  let previousRoute = $state<string | null>(null);

  const homePath = resolve('/');

  const statusCode = $derived(page.status || 404);

  const errorTitle = $derived(
    statusCode === 404
      ? 'Cerita Ini Belum Ditemukan'
      : 'Cerita Sedang Terhenti'
  );

  const errorDescription = $derived(
    statusCode === 404
      ? 'Halaman yang kamu cari mungkin telah dipindahkan, dihapus, atau belum menjadi bagian dari perjalanan Cerita Loka.'
      : 'Terjadi kendala saat membuka halaman ini. Kamu dapat kembali ke halaman sebelumnya atau melanjutkan ke halaman utama.'
  );

  function getPreviousRoute(): string | null {
    const savedRoute = sessionStorage.getItem(
      'last-public-route'
    );

    if (
      !savedRoute ||
      !savedRoute.startsWith('/') ||
      savedRoute.startsWith('/admin') ||
      savedRoute === page.url.pathname
    ) {
      return null;
    }

    return savedRoute;
  }

  async function goHome(): Promise<void> {
    await goto(homePath);
  }

  async function goBack(): Promise<void> {
    if (previousRoute) {
      await goto(previousRoute);
      return;
    }

    if (window.history.length > 1) {
      window.history.back();
      return;
    }

    await goto(homePath);
  }

  onMount(() => {
    previousRoute = getPreviousRoute();

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      return;
    }

    const context = gsap.context(() => {
      gsap.set(
        [
          badgeElement,
          codeElement,
          titleElement,
          descriptionElement,
          actionsElement,
          logoElement
        ],
        {
          opacity: 0
        }
      );

      gsap.set(badgeElement, {
        y: -30,
        scale: 0.9
      });

      gsap.set(codeElement, {
        y: 80,
        scale: 0.72,
        rotateX: 18,
        filter: 'blur(14px)',
        transformPerspective: 1000
      });

      gsap.set(titleElement, {
        y: 45,
        filter: 'blur(8px)'
      });

      gsap.set(descriptionElement, {
        y: 35,
        filter: 'blur(6px)'
      });

      gsap.set(actionsElement, {
        y: 30,
        scale: 0.95
      });

      gsap.set(logoElement, {
        y: 30,
        scale: 0.75,
        rotate: -8
      });

      gsap.set(glowElement, {
        opacity: 0,
        scale: 0.65
      });

      const timeline = gsap.timeline({
        defaults: {
          ease: 'power3.out'
        }
      });

      timeline
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
          badgeElement,
          {
            y: 0,
            scale: 1,
            opacity: 1,
            duration: 0.7
          },
          0.15
        )
        .to(
          codeElement,
          {
            y: 0,
            scale: 1,
            rotateX: 0,
            opacity: 1,
            filter: 'blur(0px)',
            duration: 1.1,
            ease: 'back.out(1.35)'
          },
          0.25
        )
        .to(
          logoElement,
          {
            y: 0,
            scale: 1,
            rotate: 0,
            opacity: 1,
            duration: 0.85,
            ease: 'back.out(1.6)'
          },
          0.65
        )
        .to(
          titleElement,
          {
            y: 0,
            opacity: 1,
            filter: 'blur(0px)',
            duration: 0.8
          },
          0.65
        )
        .to(
          descriptionElement,
          {
            y: 0,
            opacity: 1,
            filter: 'blur(0px)',
            duration: 0.75
          },
          0.85
        )
        .to(
          actionsElement,
          {
            y: 0,
            scale: 1,
            opacity: 1,
            duration: 0.75
          },
          1
        );

      gsap.to(logoElement, {
        y: -10,
        rotate: 2,
        duration: 2.6,
        repeat: -1,
        yoyo: true,
        ease: 'sine.inOut'
      });

      gsap.to(orbitElement, {
        rotate: 360,
        duration: 24,
        repeat: -1,
        ease: 'none'
      });

      gsap.to(codeElement, {
        textShadow:
          '0 18px 0 rgba(5,34,70,0.72), 0 32px 45px rgba(4,29,64,0.38)',
        duration: 1.8,
        repeat: -1,
        yoyo: true,
        ease: 'sine.inOut'
      });
    }, sectionElement);

    return () => {
      context.revert();
    };
  });
</script>

<svelte:head>
  <title>{statusCode} | Cerita Loka</title>
  <meta
    name="description"
    content="Halaman tidak ditemukan di website Kabinet Cerita Loka."
  />
</svelte:head>

<section
  bind:this={sectionElement}
  class="
    relative flex min-h-screen
    items-center justify-center
    overflow-hidden
    bg-gradient-to-b
    from-blue-950
    via-blue-700
    to-blue-50
    px-5 pb-24 pt-32
    sm:px-8 sm:pt-36
    lg:px-12
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none absolute
      left-1/2 top-16
      h-[560px] w-[900px]
      -translate-x-1/2
      rounded-full
      bg-blue-300/20
      blur-[140px]
    "
  ></div>

  <div
    bind:this={orbitElement}
    class="
      pointer-events-none absolute
      left-1/2 top-1/2
      h-[620px] w-[620px]
      -translate-x-1/2
      -translate-y-1/2
      rounded-full
      border border-white/10
    "
  >
    <span
      class="
        absolute left-1/2 top-0
        h-3 w-3 -translate-x-1/2
        rounded-full bg-white/80
        shadow-[0_0_22px_rgba(255,255,255,0.9)]
      "
    ></span>

    <span
      class="
        absolute bottom-[12%] right-[5%]
        h-2 w-2 rounded-full
        bg-blue-100/80
        shadow-[0_0_18px_rgba(219,234,254,0.8)]
      "
    ></span>
  </div>

  <div
    class="
      pointer-events-none absolute
      -right-28 top-32
      h-80 w-80 rotate-45
      border border-blue-200/10
    "
  ></div>

  <div
    class="
      pointer-events-none absolute
      -bottom-36 -left-32
      h-96 w-96 rotate-12
      rounded-[80px]
      border border-blue-900/10
    "
  ></div>

  <div
    class="
      relative z-10 mx-auto flex
      w-full max-w-5xl
      flex-col items-center
      text-center
    "
  >

    <div class="relative mt-6 sm:mt-8">
      <div
        bind:this={codeElement}
        aria-label={`Kode kesalahan ${statusCode}`}
        class="
          select-none
          text-[clamp(8rem,27vw,18rem)]
          leading-[0.72]
          font-black
          tracking-[-0.095em]
          text-blue-700
          [-webkit-text-stroke:3px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_18px_0_rgba(5,34,70,0.72)]
          sm:[-webkit-text-stroke:4px_white]
        "
      >
        {statusCode}
      </div>

      <div
        bind:this={logoElement}
        class="
          absolute left-1/2 top-1/2
          flex size-20
          -translate-x-1/2 -translate-y-1/2
          items-center justify-center
          rounded-[24px]
          border-4 border-white
          bg-gradient-to-b
          from-white
          via-blue-50
          to-blue-200
          shadow-[0_18px_40px_rgba(5,34,70,0.38)]
          sm:size-28
          sm:rounded-[30px]
        "
      >
        <img
          src="/logo/logo-kabinet.png"
          alt="Logo Kabinet Cerita Loka"
          class="h-[74%] w-[74%] object-contain"
        />
      </div>
    </div>

    <h1
      bind:this={titleElement}
      class="
        mt-10 max-w-4xl
        text-4xl leading-[0.96]
        font-black tracking-[-0.055em]
        text-blue-800
        [-webkit-text-stroke:1.5px_white]
        [paint-order:stroke_fill]
        drop-shadow-[0_6px_0_rgba(5,34,70,0.25)]
        sm:text-5xl
        md:text-6xl
      "
    >
      {errorTitle}
    </h1>

    <p
      bind:this={descriptionElement}
      class="
        mt-6 max-w-2xl
        text-sm leading-7 font-semibold
        text-white
        sm:text-base
        md:text-lg md:leading-8
      "
    >
      {errorDescription}
    </p>

    <div
      bind:this={actionsElement}
      class="
        mt-9 flex w-full max-w-xl
        flex-col justify-center gap-4
        sm:flex-row
      "
    >
      <button
        type="button"
        onclick={goHome}
        class="
          group inline-flex min-h-14
          flex-1 items-center justify-center gap-3
          rounded-2xl
          border-2 border-white
          bg-blue-700
          px-6 py-3
          text-sm font-black
          text-white
          shadow-[0_10px_0_rgba(30,64,175,0.32)]
          transition duration-300
          hover:-translate-y-1
          hover:bg-blue-800
          hover:shadow-[0_14px_0_rgba(30,64,175,0.28)]
          active:translate-y-1
          active:shadow-none
          sm:text-base
        "
      >
        <svg
          viewBox="0 0 24 24"
          fill="none"
          aria-hidden="true"
          class="
            size-5 transition-transform duration-300
            group-hover:-translate-x-1
          "
        >
          <path
            d="M3 11.5 12 4l9 7.5"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M5.5 10.5V20h13v-9.5"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>

        Kembali ke Home
      </button>

      <button
        type="button"
        onclick={goBack}
        class="
          group inline-flex min-h-14
          flex-1 items-center justify-center gap-3
          rounded-2xl
          border-2 border-blue-800/20
          bg-white/80
          px-6 py-3
          text-sm font-black
          text-blue-800
          shadow-[0_10px_0_rgba(30,64,175,0.16)]
          backdrop-blur-xl
          transition duration-300
          hover:-translate-y-1
          hover:bg-white
          hover:shadow-[0_14px_0_rgba(30,64,175,0.14)]
          active:translate-y-1
          active:shadow-none
          sm:text-base
        "
      >
        <svg
          viewBox="0 0 24 24"
          fill="none"
          aria-hidden="true"
          class="
            size-5 transition-transform duration-300
            group-hover:-translate-x-1
          "
        >
          <path
            d="m10 6-6 6 6 6"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M4 12h16"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
          />
        </svg>

        Kembali ke Halaman Terakhir
      </button>
    </div>

  </div>
</section>