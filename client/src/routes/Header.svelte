<script lang="ts">
  import { onMount } from 'svelte';
  import { resolve } from '$app/paths';
  import { page } from '$app/state';
  import { afterNavigate } from '$app/navigation';
  import { gsap } from 'gsap';
  import { publicNavigation } from '$lib/constants/public-navigation';

  let headerElement!: HTMLElement;

  let mounted = $state(false);
  let isScrolled = $state(false);
  let activeHash = $state('#home');

  let headerVisible = false;
  let reducedMotion = false;
  let isInteracting = false;

  let scrollFrame: number | null = null;
  let hideTimer: number | null = null;
  let fallbackTimer: number | null = null;

  const landingPath = resolve('/');
  const HIDE_DELAY = 3000;

  function normalizePath(pathname: string): string {
    return pathname.replace(/\/+$/, '') || '/';
  }

  const isLanding = $derived(
    normalizePath(page.url.pathname) === normalizePath(landingPath)
  );

  function getNavigationHref(hash: string): string {
    return isLanding ? hash : `${landingPath}${hash}`;
  }

  function isActive(hash: string): boolean {
    return isLanding && activeHash === hash;
  }

  function clearHideTimer(): void {
    if (hideTimer !== null) {
      window.clearTimeout(hideTimer);
      hideTimer = null;
    }
  }

  function clearFallbackTimer(): void {
    if (fallbackTimer !== null) {
      window.clearTimeout(fallbackTimer);
      fallbackTimer = null;
    }
  }

  function hideHeader(): void {
    if (!headerElement || !headerVisible || isInteracting) {
      return;
    }

    clearHideTimer();

    headerVisible = false;

    gsap.killTweensOf(headerElement);

    if (reducedMotion) {
      gsap.set(headerElement, {
        yPercent: -130,
        opacity: 0
      });

      return;
    }

    gsap.to(headerElement, {
      yPercent: -130,
      opacity: 0,
      duration: 0.65,
      ease: 'power3.inOut',
      overwrite: true
    });
  }

  function scheduleHide(): void {
    clearHideTimer();

    if (!headerVisible || isInteracting) {
      return;
    }

    hideTimer = window.setTimeout(() => {
      hideTimer = null;
      hideHeader();
    }, HIDE_DELAY);
  }

  function showHeader(): void {
    if (!headerElement) {
      return;
    }

    clearFallbackTimer();

    if (headerVisible) {
      scheduleHide();
      return;
    }

    headerVisible = true;

    gsap.killTweensOf(headerElement);

    if (reducedMotion) {
      gsap.set(headerElement, {
        yPercent: 0,
        opacity: 1
      });

      scheduleHide();
      return;
    }

    gsap.to(headerElement, {
      yPercent: 0,
      opacity: 1,
      duration: 0.9,
      ease: 'power3.out',
      overwrite: true,
      onComplete: () => {
        if (headerVisible) {
          gsap.set(headerElement, {
            clearProps: 'transform,opacity'
          });
        }
      }
    });

    scheduleHide();
  }

  function updateActiveSection(): void {
    if (!isLanding) {
      return;
    }

    const threshold = window.innerHeight * 0.35;

    let currentHash = '#home';
    let closestTop = -Infinity;

    for (const item of publicNavigation) {
      const sectionId = item.hash.replace(/^#/, '');
      const section = document.getElementById(sectionId);

      if (!section) {
        continue;
      }

      const rect = section.getBoundingClientRect();

      if (
        rect.top <= threshold &&
        rect.top > closestTop
      ) {
        closestTop = rect.top;
        currentHash = item.hash;
      }
    }

    if (activeHash !== currentHash) {
      activeHash = currentHash;
    }
  }

  function updateScrollState(): void {
    const nextScrolled = window.scrollY > 40;

    if (isScrolled !== nextScrolled) {
      isScrolled = nextScrolled;
    }

    updateActiveSection();
  }

  function handleScroll(): void {
    showHeader();

    if (scrollFrame !== null) {
      return;
    }

    scrollFrame = window.requestAnimationFrame(() => {
      scrollFrame = null;
      updateScrollState();
    });
  }

  function handleResize(): void {
    if (scrollFrame !== null) {
      return;
    }

    scrollFrame = window.requestAnimationFrame(() => {
      scrollFrame = null;
      updateScrollState();
    });
  }

  function handleHeroBottomEnter(): void {
    showHeader();
  }

  function handleHashChange(): void {
    showHeader();
    handleResize();
  }

  function handleInteractionStart(): void {
    isInteracting = true;
    clearHideTimer();
    showHeader();
  }

  function handleInteractionEnd(): void {
    if (
      headerElement?.matches(':hover') ||
      headerElement?.contains(document.activeElement)
    ) {
      return;
    }

    isInteracting = false;
    scheduleHide();
  }

  function handleDocumentVisibility(): void {
    if (document.hidden) {
      clearHideTimer();
    } else if (headerVisible) {
      scheduleHide();
    }
  }

  function scheduleInitialReveal(): void {
    clearFallbackTimer();

    if (headerVisible) {
      return;
    }

    fallbackTimer = window.setTimeout(() => {
      fallbackTimer = null;
      showHeader();
    }, 2400);
  }

  $effect(() => {
    if (!mounted) {
      return;
    }

    const pathname = page.url.pathname;
    const hash = page.url.hash;

    const onLanding =
      normalizePath(pathname) === normalizePath(landingPath);

    if (!onLanding || hash) {
      showHeader();
    }

    if (
      onLanding &&
      hash &&
      publicNavigation.some((item) => item.hash === hash)
    ) {
      activeHash = hash;
    }
  });

  afterNavigate(({ to }) => {
    if (!mounted || !to) {
      return;
    }

    const destinationPath = normalizePath(to.url.pathname);

    const onLanding =
      destinationPath === normalizePath(landingPath);

    if (!onLanding || to.url.hash) {
      showHeader();
    } else if (!headerVisible) {
      scheduleInitialReveal();
    }

    window.requestAnimationFrame(() => {
      updateScrollState();
    });
  });

  onMount(() => {
    reducedMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    gsap.set(headerElement, {
      yPercent: -130,
      opacity: 0
    });

    window.addEventListener(
      'hero-bottom-enter',
      handleHeroBottomEnter
    );

    window.addEventListener(
      'scroll',
      handleScroll,
      { passive: true }
    );

    window.addEventListener(
      'resize',
      handleResize
    );

    window.addEventListener(
      'hashchange',
      handleHashChange
    );

    document.addEventListener(
      'visibilitychange',
      handleDocumentVisibility
    );

    mounted = true;

    updateScrollState();

    if (
      !isLanding ||
      page.url.hash ||
      window.scrollY > 24
    ) {
      showHeader();
    } else {
      scheduleInitialReveal();
    }

    return () => {
      mounted = false;

      window.removeEventListener(
        'hero-bottom-enter',
        handleHeroBottomEnter
      );

      window.removeEventListener(
        'scroll',
        handleScroll
      );

      window.removeEventListener(
        'resize',
        handleResize
      );

      window.removeEventListener(
        'hashchange',
        handleHashChange
      );

      document.removeEventListener(
        'visibilitychange',
        handleDocumentVisibility
      );

      if (scrollFrame !== null) {
        window.cancelAnimationFrame(scrollFrame);
        scrollFrame = null;
      }

      clearHideTimer();
      clearFallbackTimer();

      gsap.killTweensOf(headerElement);
    };
  });
</script>

<header
  role="banner"
  bind:this={headerElement}
  onmouseenter={handleInteractionStart}
  onmouseleave={handleInteractionEnd}
  onfocusin={handleInteractionStart}
  onfocusout={() => {
    window.requestAnimationFrame(handleInteractionEnd);
  }}
  class="
    fixed inset-x-0 top-0 z-50
    px-1 py-1.5
    sm:px-4 sm:py-2
    md:px-6 md:py-3 backdrop-blur-[10px]
  "
>
  <div
    class={`
      mx-auto flex w-full max-w-7xl
      items-center justify-between
      gap-0.5
      sm:gap-2
      lg:gap-7
      transform-gpu
      transition-[transform,filter]
      duration-500
      ease-out
      motion-reduce:transition-none
      ${
        isScrolled
          ? 'translate-y-1 scale-[0.975] drop-shadow-[0_14px_24px_rgba(15,23,42,0.18)]'
          : 'translate-y-0 scale-100'
      }
    `}
  >
    <a
      href={landingPath}
      aria-label="Kembali ke halaman utama"
      class={`
        flex shrink-0 items-center
        gap-0.5 rounded-md
        border border-white/70
        px-1 py-1
        shadow-sm backdrop-blur-xl
        transition duration-200
        sm:gap-1.5
        sm:rounded-lg
        sm:px-2 sm:py-1.5
        md:rounded-xl
        md:px-3 md:py-2
        ${
          isLanding
            ? 'bg-white/25 hover:bg-white/40'
            : 'bg-blue-950/40 hover:bg-blue-950/55'
        }
      `}
    >
      <div
        class="
          flex shrink-0 items-center
          gap-px
          sm:gap-0.5
          md:gap-1
        "
      >
        <img
          class="size-4 object-contain sm:size-6 md:size-8"
          src="/logo/logo-bem.webp"
          alt="Logo BEM UNAIR"
        />

        <img
          class="size-4 object-contain sm:size-6 md:size-8"
          src="/logo/logo-kabinet.webp"
          alt="Logo Kabinet Cerita Loka"
        />
      </div>

      <div class="min-w-0 leading-none">
        <p
          class="
            whitespace-nowrap
            text-[5px] font-extrabold
            transition-colors duration-200
            sm:text-[7px]
            md:text-[10px]
          "
          style:color={isLanding ? '#000000' : '#ffffff'}
        >
          Kabinet
        </p>

        <p
          class="
            mt-0.5 whitespace-nowrap
            text-[5px] font-medium
            transition-colors duration-200
            sm:text-[7px]
            md:mt-1 md:text-[10px]
          "
          style:color={isLanding
            ? 'rgb(0 0 0 / 0.7)'
            : 'rgb(255 255 255 / 0.9)'}
        >
          Cerita Loka
        </p>
      </div>
    </a>

    <nav
      aria-label="Navigasi publik"
      class={`
        flex w-fit min-w-0 flex-none
        items-center justify-center
        rounded-md
        border border-white/70
        p-px
        shadow-sm backdrop-blur-xl
        transition-colors duration-200
        sm:rounded-lg
        sm:p-0.5
        md:rounded-xl
        md:px-4
        md:py-1
        ${
          isLanding
            ? 'bg-white/25'
            : 'bg-blue-950/40'
        }
      `}
    >
      {#each publicNavigation as item (item.hash)}
        <a
          href={getNavigationHref(item.hash)}
          aria-current={isActive(item.hash) ? 'location' : undefined}
          style:color={isLanding
            ? undefined
            : '#ffffff'}
          class={`
            whitespace-nowrap rounded-sm
            px-2 py-1
            text-[7px] font-semibold
            transition duration-200
            sm:rounded-md
            sm:px-3
            sm:text-[9px]
            md:rounded-lg
            md:px-3
            md:py-2
            md:text-xs
            lg:px-5
            lg:text-sm
            ${
              isLanding
                ? isActive(item.hash)
                  ? 'bg-white/45 text-blue-900 shadow-sm'
                  : 'text-blue-600 hover:bg-white/35 hover:text-blue-800'
                : 'hover:bg-white/20'
            }
          `}
        >
          {item.label}
        </a>
      {/each}
    </nav>

    <div
      class={`
        flex shrink-0 items-center
        gap-0.5 rounded-md
        border border-white/70
        px-1 py-1
        shadow-sm backdrop-blur-xl
        transition duration-200
        sm:gap-1
        sm:rounded-lg
        sm:px-2
        sm:py-1.5
        md:gap-2
        md:rounded-xl
        md:px-4
        md:py-2
        ${
          isLanding
            ? 'bg-white/25 hover:bg-white/40'
            : 'bg-blue-950/40 hover:bg-blue-950/55'
        }
      `}
    >
      <span
        class="
          whitespace-nowrap text-right
          text-[5px] leading-[1.05]
          transition-colors duration-200
          sm:text-[7px]
          md:text-[10px]
        "
        style:color={isLanding
          ? 'rgb(0 0 0 / 0.7)'
          : 'rgb(255 255 255 / 0.9)'}
      >
        <span class="font-semibold">
          Merajut Cerita
        </span>

        <br />

        <span class="font-medium">
          Tumbuh Bersama
        </span>
      </span>

      <span
        aria-hidden="true"
        class="
          text-sm leading-none font-black
          transition-colors duration-200
          sm:text-xl
          md:text-3xl
        "
        style:color={isLanding ? '#1e3a8a' : '#ffffff'}
      >
        #
      </span>
    </div>
  </div>
</header>