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

  let currentPage = $state(0);
  let isAnimating = $state(false);

  let reduceMotion = false;

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
      duration: reduceMotion ? 0.01 : 0.9,
      ease: 'power2.inOut',
      boxShadow:
        '-14px 8px 24px rgba(15, 23, 42, 0.25)',
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
      duration: reduceMotion ? 0.01 : 0.9,
      ease: 'power2.inOut',
      boxShadow:
        '0 8px 20px rgba(15, 23, 42, 0.15)',
      onComplete: () => {
        currentPage = previousIndex;
        isAnimating = false;
      }
    });
  }

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    pages = [
      coverPage,
      pageOne,
      pageTwo
    ];

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
        gsap.set(
          [bookFrame, bookControls],
          {
            autoAlpha: 1,
            clearProps: 'transform'
          }
        );

        return;
      }

      gsap.set(bookFrame, {
        autoAlpha: 0,
        y: 65,
        scale: 0.96,
        rotationX: 5,
        transformOrigin: 'center bottom'
      });

      gsap.set(bookControls, {
        autoAlpha: 0,
        y: 18
      });

      gsap
        .timeline({
          scrollTrigger: {
            trigger: speechSection,
            start: 'top 80%',
            once: true,
            invalidateOnRefresh: true
          }
        })
        .to(bookFrame, {
          autoAlpha: 1,
          y: 0,
          scale: 1,
          rotationX: 0,
          duration: 0.95,
          ease: 'power3.out',
          clearProps:
            'transform,opacity,visibility'
        })
        .to(
          bookControls,
          {
            autoAlpha: 1,
            y: 0,
            duration: 0.55,
            ease: 'power2.out',
            clearProps:
              'transform,opacity,visibility'
          },
          '-=0.3'
        );
    }, speechSection);

    const handleKeydown = (
      event: KeyboardEvent
    ): void => {
      if (
        event.altKey ||
        event.ctrlKey ||
        event.metaKey
      ) {
        return;
      }

      const target = event.target;

      if (
        target instanceof HTMLElement &&
        (
          target.matches(
            'input, textarea, select'
          ) ||
          target.isContentEditable
        )
      ) {
        return;
      }

      const rect =
        speechSection.getBoundingClientRect();

      if (
        rect.bottom <= 0 ||
        rect.top >= window.innerHeight
      ) {
        return;
      }

      if (event.key === 'ArrowRight') {
        event.preventDefault();
        nextPage();
      }

      if (event.key === 'ArrowLeft') {
        event.preventDefault();
        previousPage();
      }
    };

    window.addEventListener(
      'keydown',
      handleKeydown
    );

    return () => {
      window.removeEventListener(
        'keydown',
        handleKeydown
      );

      pages.forEach((page) => {
        gsap.killTweensOf(page);
      });

      context.revert();
    };
  });
</script>

<section
  bind:this={speechSection}
  id="speech"
  class="
    w-full overflow-hidden
    bg-linear-to-b
    from-blue-700
    via-blue-500
    to-blue-100
    px-3 pt-14 pb-4
    sm:px-5 sm:pt-16 sm:pb-6
    md:px-8 md:pt-20 md:pb-8
    lg:px-10 lg:pt-24 lg:pb-10 
  "
>
  <div
    class="
      mx-auto flex w-full
      max-w-4xl
      flex-col items-center
      gap-3
      sm:gap-4
    "
  >
    <div
      bind:this={bookFrame}
      class="
        relative
        h-[330px]
        w-full
        rounded-xl
        bg-white/20
        p-2
        shadow-2xl
        backdrop-blur-xl
        [perspective:1800px]
        will-change-transform
        sm:h-[410px]
        sm:p-3
        md:h-[490px]
        md:p-4
        lg:h-[540px]
        lg:p-5
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
            absolute inset-0
            overflow-hidden
            rounded-xl
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
            absolute inset-0
            overflow-hidden
            rounded-xl
            [backface-visibility:hidden]
            [transform-style:preserve-3d]
            will-change-transform
          "
        >
          <Page1
            active={currentPage === 1}
          />
        </div>

        <div
          bind:this={pageTwo}
          class="
            absolute inset-0
            overflow-hidden
            rounded-xl
            [backface-visibility:hidden]
            [transform-style:preserve-3d]
            will-change-transform
          "
        >
          <Page2
            active={currentPage === 2}
          />
        </div>
      </div>
    </div>

    <div
      bind:this={bookControls}
      class="
        flex items-center
        justify-center gap-3
      "
    >
      <button
        type="button"
        aria-label="Halaman sebelumnya"
        disabled={currentPage === 0 || isAnimating}
        onclick={previousPage}
        class="
          grid size-9
          place-items-center
          rounded-full
          border border-white/60
          bg-white/30
          text-blue-900
          shadow-lg
          backdrop-blur-xl
          transition
          hover:bg-white/60
          disabled:pointer-events-none
          disabled:opacity-35
          sm:size-10
        "
      >
        <ChevronLeft
          size={18}
          strokeWidth={2.5}
        />
      </button>

      <p
        aria-live="polite"
        class="
          min-w-18
          rounded-full
          border border-white/60
          bg-white/30
          px-3 py-1.5
          text-center
          text-xs font-bold
          text-blue-900
          shadow-lg
          backdrop-blur-xl
          sm:min-w-20
          sm:text-sm
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
          grid size-9
          place-items-center
          rounded-full
          border border-white/60
          bg-white/30
          text-blue-900
          shadow-lg
          backdrop-blur-xl
          transition
          hover:bg-white/60
          disabled:pointer-events-none
          disabled:opacity-35
          sm:size-10
        "
      >
        <ChevronRight
          size={18}
          strokeWidth={2.5}
        />
      </button>
    </div>
  </div>
</section>