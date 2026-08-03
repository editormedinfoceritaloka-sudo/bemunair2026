<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import ChevronLeft from '@lucide/svelte/icons/chevron-left';
  import ChevronRight from '@lucide/svelte/icons/chevron-right';

  import Cover from './book/Cover.svelte';
  import Page1 from './book/Page1.svelte';
  import Page2 from './book/Page2.svelte';

  const totalPages = 3;

  let speechSection: HTMLElement;
  let bookFrame: HTMLDivElement;
  let bookControls: HTMLDivElement;

  let coverPage: HTMLDivElement;
  let pageOne: HTMLDivElement;
  let pageTwo: HTMLDivElement;

  let pages: HTMLDivElement[] = [];
  let currentPage = 0;
  let isAnimating = false;
  let reduceMotion = false;

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    pages = [coverPage, pageOne, pageTwo];

    reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    pages.forEach((page, index) => {
      gsap.set(page, {
        rotationY: 0,
        transformOrigin: 'left center',
        transformStyle: 'preserve-3d',
        backfaceVisibility: 'hidden',
        zIndex: totalPages - index
      });
    });

    const context = gsap.context(() => {
      if (reduceMotion) {
        gsap.set([bookFrame, bookControls], {
          autoAlpha: 1,
          clearProps: 'transform'
        });

        return;
      }

      gsap.set(bookFrame, {
        autoAlpha: 0,
        y: 110,
        scale: 0.94,
        rotationX: 8,
        transformOrigin: 'center bottom'
      });

      gsap.set(bookControls, {
        autoAlpha: 0,
        y: 28
      });

      gsap
        .timeline({
          scrollTrigger: {
            trigger: speechSection,
            start: 'top 78%',
            once: true,
            invalidateOnRefresh: true
          }
        })
        .to(bookFrame, {
          autoAlpha: 1,
          y: 0,
          scale: 1,
          rotationX: 0,
          duration: 1.15,
          ease: 'power3.out',
          clearProps: 'transform,opacity,visibility'
        })
        .to(
          bookControls,
          {
            autoAlpha: 1,
            y: 0,
            duration: 0.65,
            ease: 'power2.out',
            clearProps: 'transform,opacity,visibility'
          },
          '-=0.4'
        );
    }, speechSection);

    const handleKeydown = (event: KeyboardEvent) => {
      if (event.key === 'ArrowRight') {
        nextPage();
      }

      if (event.key === 'ArrowLeft') {
        previousPage();
      }
    };

    window.addEventListener('keydown', handleKeydown);

    return () => {
      window.removeEventListener('keydown', handleKeydown);

      pages.forEach((page) => {
        gsap.killTweensOf(page);
      });

      context.revert();
    };
  });

  function nextPage(): void {
    if (
      isAnimating ||
      currentPage >= totalPages - 1
    ) {
      return;
    }

    const page = pages[currentPage];

    if (!page) {
      return;
    }

    isAnimating = true;

    gsap.to(page, {
      rotationY: -180,
      duration: reduceMotion ? 0.01 : 1.15,
      ease: 'power2.inOut',
      boxShadow: '-20px 12px 32px rgba(15, 23, 42, 0.25)',
      onComplete: () => {
        currentPage += 1;
        isAnimating = false;
      }
    });
  }

  function previousPage(): void {
    if (
      isAnimating ||
      currentPage <= 0
    ) {
      return;
    }

    const previousIndex = currentPage - 1;
    const page = pages[previousIndex];

    if (!page) {
      return;
    }

    isAnimating = true;

    gsap.set(page, {
      zIndex: totalPages - previousIndex
    });

    gsap.to(page, {
      rotationY: 0,
      duration: reduceMotion ? 0.01 : 1.15,
      ease: 'power2.inOut',
      boxShadow: '0 12px 30px rgba(15, 23, 42, 0.15)',
      onComplete: () => {
        currentPage = previousIndex;
        isAnimating = false;
      }
    });
  }
</script>

<section
  bind:this={speechSection}
  id="speech"
  class="
    w-full overflow-hidden
    bg-linear-to-b from-blue-700 via-blue-500 to-blue-100
    px-3 py-24
    sm:px-6 sm:py-26
    md:px-10 md:py-28
    lg:px-16 lg:py-30
  "
>
  <div
    class="
      mx-auto flex w-full max-w-7xl
      flex-col items-center gap-5
    "
  >
    <div
      bind:this={bookFrame}
      class="
        relative h-[400px] w-full
        rounded-xl bg-white/20
        p-3 shadow-2xl backdrop-blur-xl
        [perspective:2200px]
        will-change-transform
        sm:h-[70svh] sm:min-h-[500px] sm:p-5
        md:h-[85svh] md:min-h-[620px] md:p-8
        lg:h-[70vh] lg:max-h-[700px]
      "
    >
      <div
        class="
          relative h-full w-full
          [transform-style:preserve-3d]
        "
      >
        <div
          bind:this={coverPage}
          class="
            absolute inset-0 overflow-hidden rounded-xl
            [backface-visibility:hidden]
            [transform-style:preserve-3d]
            will-change-transform
          "
        >
          <Cover />
        </div>

        <div
          bind:this={pageOne}
          class="
            absolute inset-0 overflow-hidden rounded-xl
            [backface-visibility:hidden]
            [transform-style:preserve-3d]
            will-change-transform
          "
        >
          <Page1 />
        </div>

        <div
          bind:this={pageTwo}
          class="
            absolute inset-0 overflow-hidden rounded-xl
            [backface-visibility:hidden]
            [transform-style:preserve-3d]
            will-change-transform
          "
        >
          <Page2 />
        </div>
      </div>
    </div>

    <div
      bind:this={bookControls}
      class="flex items-center justify-center gap-4"
    >
      <button
        type="button"
        aria-label="Halaman sebelumnya"
        disabled={currentPage === 0 || isAnimating}
        onclick={previousPage}
        class="
          grid size-11 place-items-center rounded-full
          border border-white/60 bg-white/30
          text-blue-900 shadow-lg backdrop-blur-xl
          transition
          hover:bg-white/60
          disabled:pointer-events-none disabled:opacity-35
        "
      >
        <ChevronLeft
          size={22}
          strokeWidth={2.5}
        />
      </button>

      <p
        aria-live="polite"
        class="
          min-w-24 rounded-full
          border border-white/60 bg-white/30
          px-4 py-2 text-center text-sm
          font-bold text-blue-900
          shadow-lg backdrop-blur-xl
        "
      >
        {currentPage + 1} / {totalPages}
      </p>

      <button
        type="button"
        aria-label="Halaman berikutnya"
        disabled={currentPage === totalPages - 1 || isAnimating}
        onclick={nextPage}
        class="
          grid size-11 place-items-center rounded-full
          border border-white/60 bg-white/30
          text-blue-900 shadow-lg backdrop-blur-xl
          transition
          hover:bg-white/60
          disabled:pointer-events-none disabled:opacity-35
        "
      >
        <ChevronRight
          size={22}
          strokeWidth={2.5}
        />
      </button>
    </div>
  </div>
</section>