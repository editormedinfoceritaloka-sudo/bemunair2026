<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import {
    ArrowLeftIcon,
    ArrowRightIcon,
    ArrowUpRightIcon,
    PauseIcon,
    PlayIcon
  } from '@lucide/svelte';
  import { resolve } from '$app/paths';

  import type { ArticleListItem } from '$lib/types';

  let {
    articles = [],
    logo = '/logo/logo-kabinet.webp',
    organization = 'Kabinet Cerita Loka'
  }: {
    articles?: ArticleListItem[];
    logo?: string;
    organization?: string;
  } = $props();

  const AUTOPLAY_SECONDS = 6;
  const SWIPE_THRESHOLD = 48;
  const FALLBACK_COVER = '/landing/news/test.webp';

  const dateFormatter = new Intl.DateTimeFormat('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    timeZone: 'Asia/Jakarta'
  });

  function formatDate(value: string | null): string {
    if (!value) {
      return '';
    }

    const date = new Date(value);

    return Number.isNaN(date.getTime()) ? value : dateFormatter.format(date);
  }

  function pad(value: number): string {
    return String(value).padStart(2, '0');
  }

  const slides = $derived(
    articles.slice(0, 3).map((article) => ({
      id: String(article.id),
      slug: article.slug,
      title: article.title,
      excerpt: article.excerpt ?? '',
      cover: article.cover_image ?? FALLBACK_COVER,
      author: article.author?.name ?? organization,
      publishedAt: article.published_at,
      publishedLabel: formatDate(article.published_at)
    }))
  );

  const total = $derived(slides.length);

  let active = $state(0);
  let reducedMotion = $state(false);
  let userPaused = $state(false);
  let hovering = $state(false);
  let focusWithin = $state(false);
  let inView = $state(false);

  const autoplayEnabled = $derived(total > 1 && !reducedMotion);
  const autoplayRunning = $derived(
    autoplayEnabled && !userPaused && !hovering && !focusWithin && inView
  );

  let sectionElement: HTMLElement;
  let headerElement: HTMLDivElement;
  let carouselElement: HTMLDivElement;
  let viewportElement: HTMLDivElement;
  let trackElement: HTMLDivElement;

  let progressTween: gsap.core.Tween | null = null;
  let mounted = false;

  function getSlides(): HTMLElement[] {
    return trackElement
      ? Array.from(trackElement.querySelectorAll<HTMLElement>('[data-slide]'))
      : [];
  }

  function getProgressFills(): HTMLElement[] {
    return carouselElement
      ? Array.from(
          carouselElement.querySelectorAll<HTMLElement>('[data-progress]')
        )
      : [];
  }

  // Lebar slide & padding track memakai satuan cqw (relatif ke viewport
  // carousel) sehingga slide pertama sudah di tengah pada x = 0, juga saat
  // SSR. Geser track sebesar selisih offsetLeft saja.
  function offsetFor(index: number): number {
    const nodes = getSlides();
    const first = nodes[0];
    const target = nodes[index];

    if (!first || !target) {
      return 0;
    }

    return first.offsetLeft - target.offsetLeft;
  }

  function animateTo(index: number, immediate = false) {
    const nodes = getSlides();

    if (!trackElement || nodes.length === 0) {
      return;
    }

    const instant = immediate || reducedMotion;
    const duration = instant ? 0 : 0.85;

    gsap.to(trackElement, {
      x: offsetFor(index),
      duration,
      ease: 'power3.inOut',
      overwrite: 'auto'
    });

    nodes.forEach((node, nodeIndex) => {
      const isActive = nodeIndex === index;

      gsap.to(node, {
        opacity: isActive ? 1 : 0.4,
        scale: isActive ? 1 : 0.94,
        duration,
        ease: 'power3.out',
        overwrite: 'auto'
      });
    });

    if (instant) {
      return;
    }

    const current = nodes[index];
    const image = current.querySelector('[data-slide-image]');
    const reveal = current.querySelectorAll('[data-slide-reveal]');

    if (image) {
      gsap.fromTo(
        image,
        { scale: 1.12 },
        { scale: 1, duration: 1.2, ease: 'power3.out', overwrite: 'auto' }
      );
    }

    if (reveal.length > 0) {
      gsap.fromTo(
        reveal,
        { autoAlpha: 0, y: 20 },
        {
          autoAlpha: 1,
          y: 0,
          duration: 0.6,
          delay: 0.25,
          stagger: 0.07,
          ease: 'power3.out',
          overwrite: 'auto'
        }
      );
    }
  }

  function restartProgress() {
    progressTween?.kill();
    progressTween = null;

    const fills = getProgressFills();

    if (fills.length === 0) {
      return;
    }

    gsap.set(fills, { scaleX: 0 });

    const fill = fills[active];

    if (!fill) {
      return;
    }

    if (!autoplayEnabled) {
      gsap.set(fill, { scaleX: 1 });
      return;
    }

    progressTween = gsap.to(fill, {
      scaleX: 1,
      duration: AUTOPLAY_SECONDS,
      ease: 'none',
      paused: !autoplayRunning,
      onComplete: () => goTo(active + 1)
    });
  }

  function goTo(index: number) {
    if (total < 2) {
      return;
    }

    const next = (index + total) % total;

    if (next === active) {
      return;
    }

    active = next;

    if (!mounted) {
      return;
    }

    animateTo(next);
    restartProgress();
  }

  function handleSlideClick(event: MouseEvent, index: number) {
    // Klik pada slide samping (yang mengintip) memutar carousel,
    // bukan langsung membuka artikel.
    if (index !== active) {
      event.preventDefault();
      goTo(index);
    }
  }

  function toggleAutoplay() {
    userPaused = !userPaused;
  }

  $effect(() => {
    const running = autoplayRunning;

    if (!progressTween) {
      return;
    }

    if (running) {
      progressTween.resume();
    } else {
      progressTween.pause();
    }
  });

  onMount(() => {
    if (total === 0) {
      return;
    }

    gsap.registerPlugin(ScrollTrigger);

    const motionQuery = window.matchMedia('(prefers-reduced-motion: reduce)');

    reducedMotion = motionQuery.matches;

    const context = gsap.context(() => {
      gsap.set(getSlides(), { transformOrigin: '50% 50%' });

      if (!reducedMotion) {
        gsap.set([headerElement, carouselElement], { autoAlpha: 0, y: 36 });

        gsap.to([headerElement, carouselElement], {
          autoAlpha: 1,
          y: 0,
          duration: 0.85,
          stagger: 0.12,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: sectionElement,
            start: 'top 85%',
            once: true
          }
        });
      }

      const visibility = ScrollTrigger.create({
        trigger: sectionElement,
        start: 'top bottom',
        end: 'bottom top',
        onToggle: (self) => {
          inView = self.isActive;
        }
      });

      inView = visibility.isActive;
    }, sectionElement);

    mounted = true;
    animateTo(active, true);
    restartProgress();

    const handleMotionChange = () => {
      reducedMotion = motionQuery.matches;
      animateTo(active, true);
      restartProgress();
    };

    const resizeObserver = new ResizeObserver(() => {
      gsap.killTweensOf(trackElement);
      gsap.set(trackElement, { x: offsetFor(active) });
    });

    resizeObserver.observe(viewportElement);

    const handleKeydown = (event: KeyboardEvent) => {
      if (event.key === 'ArrowRight') {
        event.preventDefault();
        goTo(active + 1);
      }

      if (event.key === 'ArrowLeft') {
        event.preventDefault();
        goTo(active - 1);
      }
    };

    const handlePointerEnter = (event: PointerEvent) => {
      if (event.pointerType === 'mouse') {
        hovering = true;
      }
    };

    // focusout/pointerleave bisa terpicu saat DOM carousel dilepas (mis.
    // klik link slide aktif lalu {#key} me-remount). Tunda mutasi state agar
    // tidak terjadi di tengah teardown Svelte (state_unsafe_mutation).
    const deferIfMounted = (update: () => void) => {
      queueMicrotask(() => {
        if (mounted) {
          update();
        }
      });
    };

    const handlePointerLeave = () => {
      deferIfMounted(() => {
        hovering = false;
      });
    };

    // Jeda hanya untuk fokus keyboard; klik mouse pada tombol tidak
    // menghentikan autoplay secara permanen (hover sudah menjeda).
    const handleFocusIn = (event: FocusEvent) => {
      const target = event.target as HTMLElement | null;

      focusWithin = target?.matches(':focus-visible') ?? true;
    };

    const handleFocusOut = (event: FocusEvent) => {
      if (!carouselElement.contains(event.relatedTarget as Node | null)) {
        deferIfMounted(() => {
          focusWithin = false;
        });
      }
    };

    // Swipe hanya untuk sentuhan/pen; drag mouse dibiarkan agar klik link
    // tetap berfungsi normal.
    let swipePointer: number | null = null;
    let swipeStartX = 0;
    let swipeStartY = 0;

    const handlePointerDown = (event: PointerEvent) => {
      if (event.pointerType === 'mouse') {
        return;
      }

      swipePointer = event.pointerId;
      swipeStartX = event.clientX;
      swipeStartY = event.clientY;
    };

    const handlePointerUp = (event: PointerEvent) => {
      if (swipePointer !== event.pointerId) {
        return;
      }

      swipePointer = null;

      const deltaX = event.clientX - swipeStartX;
      const deltaY = event.clientY - swipeStartY;

      if (
        Math.abs(deltaX) > SWIPE_THRESHOLD &&
        Math.abs(deltaX) > Math.abs(deltaY)
      ) {
        goTo(active + (deltaX < 0 ? 1 : -1));
      }
    };

    const handlePointerCancel = () => {
      swipePointer = null;
    };

    motionQuery.addEventListener('change', handleMotionChange);
    carouselElement.addEventListener('keydown', handleKeydown);
    carouselElement.addEventListener('pointerenter', handlePointerEnter);
    carouselElement.addEventListener('pointerleave', handlePointerLeave);
    carouselElement.addEventListener('focusin', handleFocusIn);
    carouselElement.addEventListener('focusout', handleFocusOut);
    viewportElement.addEventListener('pointerdown', handlePointerDown);
    viewportElement.addEventListener('pointerup', handlePointerUp);
    viewportElement.addEventListener('pointercancel', handlePointerCancel);

    return () => {
      mounted = false;

      motionQuery.removeEventListener('change', handleMotionChange);
      carouselElement.removeEventListener('keydown', handleKeydown);
      carouselElement.removeEventListener('pointerenter', handlePointerEnter);
      carouselElement.removeEventListener('pointerleave', handlePointerLeave);
      carouselElement.removeEventListener('focusin', handleFocusIn);
      carouselElement.removeEventListener('focusout', handleFocusOut);
      viewportElement.removeEventListener('pointerdown', handlePointerDown);
      viewportElement.removeEventListener('pointerup', handlePointerUp);
      viewportElement.removeEventListener('pointercancel', handlePointerCancel);
      resizeObserver.disconnect();

      progressTween?.kill();
      progressTween = null;

      const nodes = getSlides();

      gsap.killTweensOf([
        trackElement,
        ...nodes,
        ...nodes.flatMap((node) =>
          Array.from(
            node.querySelectorAll('[data-slide-image], [data-slide-reveal]')
          )
        ),
        ...getProgressFills()
      ]);

      context.revert();
    };
  });
</script>

<aside
  bind:this={sectionElement}
  aria-labelledby="related-news-heading"
  class="
    relative z-10
    mx-auto
    mt-24
    w-full
    max-w-7xl
    px-5
    sm:px-8
    lg:mt-32
    lg:px-12
  "
>
  <div
    bind:this={headerElement}
    class="
      flex
      items-end
      justify-between
      gap-4
      border-b
      border-blue-900/15
      pb-4
      sm:gap-6
    "
  >
    <div class="min-w-0">
      <p
        class="
          text-[11px]
          font-black
          uppercase
          tracking-[0.24em]
          text-orange-500
          sm:text-xs
        "
      >
        Baca Juga
      </p>

      <h2
        id="related-news-heading"
        class="
          mt-2
          text-[clamp(1.75rem,4.5vw,3rem)]
          leading-none
          font-black
          tracking-[-0.05em]
          text-blue-950
        "
      >
        Berita Lainnya
      </h2>
    </div>

    <a
      href={resolve('/news')}
      class="
        group
        inline-flex
        shrink-0
        items-center
        gap-1.5
        text-xs
        font-bold
        text-orange-600
        transition-colors
        duration-300
        hover:text-orange-800
        focus-visible:rounded-sm
        focus-visible:outline-2
        focus-visible:outline-offset-4
        focus-visible:outline-orange-600
        sm:gap-2
        sm:text-sm
      "
    >
      <span class="sm:hidden">Lihat Semua</span>
      <span class="hidden sm:inline">Lihat Semua Berita</span>

      <ArrowRightIcon
        size={18}
        strokeWidth={2.5}
        class="
          size-4
          -rotate-45
          transition-transform
          duration-300
          group-hover:-translate-y-1
          group-hover:translate-x-1
          motion-reduce:transform-none
          sm:size-[18px]
        "
      />
    </a>
  </div>

  <div
    bind:this={carouselElement}
    role="region"
    aria-roledescription="carousel"
    aria-label="Berita lainnya"
    class="mt-8 sm:mt-10"
  >
    <div
      bind:this={viewportElement}
      class="@container -mx-5 touch-pan-y overflow-hidden py-4 sm:-mx-8 lg:mx-0"
    >
      <div
        bind:this={trackElement}
        aria-live={autoplayRunning ? 'off' : 'polite'}
        class="
          relative
          flex
          gap-4
          px-[6cqw]
          will-change-transform
          sm:gap-6
          lg:px-[9cqw]
        "
      >
        {#each slides as slide, index (slide.id)}
          <div
            data-slide
            role="group"
            aria-roledescription="slide"
            aria-label={`${index + 1} dari ${total}: ${slide.title}`}
            aria-hidden={index !== active}
            class="w-[88cqw] shrink-0 lg:w-[82cqw]"
          >
            <a
              href={resolve('/news/[slug]', { slug: slide.slug })}
              tabindex={index === active ? 0 : -1}
              draggable="false"
              onclick={(event) => handleSlideClick(event, index)}
              class="
                group
                grid
                h-full
                overflow-hidden
                rounded-[24px]
                border-[5px]
                border-white
                bg-white
                shadow-[0_20px_50px_rgba(5,34,70,0.16)]
                focus-visible:outline-none
                focus-visible:ring-4
                focus-visible:ring-orange-300/60
                sm:rounded-[28px]
                lg:grid-cols-[1.15fr_1fr]
              "
            >
              <div
                class="
                  relative
                  aspect-[16/10]
                  overflow-hidden
                  bg-blue-100
                  lg:aspect-auto
                  lg:min-h-[23rem]
                "
              >
                <div
                  class="
                    absolute inset-0
                    transition-transform
                    duration-700
                    ease-out
                    group-hover:scale-[1.04]
                    motion-reduce:transition-none
                  "
                >
                  <img
                    data-slide-image
                    src={slide.cover}
                    alt={slide.title}
                    loading="lazy"
                    decoding="async"
                    draggable="false"
                    width="960"
                    height="600"
                    class="h-full w-full object-cover object-center"
                  />
                </div>

                <div
                  aria-hidden="true"
                  class="
                    pointer-events-none
                    absolute inset-0
                    bg-gradient-to-t
                    from-blue-950/35
                    via-transparent
                    to-transparent
                  "
                ></div>

                <span
                  class="
                    absolute
                    left-4 top-4
                    rounded-full
                    bg-white/90
                    px-3 py-1
                    text-[10px]
                    font-black
                    uppercase
                    tracking-[0.2em]
                    text-blue-900
                    shadow-sm
                    sm:text-[11px]
                  "
                >
                  Berita
                </span>
              </div>

              <div
                class="
                  flex
                  min-w-0
                  flex-col
                  justify-center
                  gap-4
                  p-5
                  sm:p-8
                  lg:p-10
                "
              >
                <p
                  data-slide-reveal
                  class="
                    text-xs
                    font-black
                    tracking-[0.2em]
                    text-orange-500
                  "
                >
                  {pad(index + 1)} / {pad(total)}
                </p>

                <h3
                  data-slide-reveal
                  class="
                    line-clamp-3
                    text-xl
                    leading-[1.12]
                    font-black
                    tracking-[-0.04em]
                    text-blue-950
                    transition-colors
                    duration-200
                    group-hover:text-blue-700
                    sm:text-2xl
                    lg:text-3xl
                  "
                >
                  {slide.title}
                </h3>

                {#if slide.excerpt}
                  <p
                    data-slide-reveal
                    class="
                      line-clamp-3
                      text-sm
                      leading-6
                      font-medium
                      text-blue-950/65
                      sm:text-base
                      sm:leading-7
                    "
                  >
                    {slide.excerpt}
                  </p>
                {/if}

                <div
                  data-slide-reveal
                  class="
                    mt-1
                    flex
                    items-center
                    justify-between
                    gap-4
                    border-t
                    border-blue-900/10
                    pt-4
                    sm:mt-2
                    sm:pt-5
                  "
                >
                  <div class="flex min-w-0 items-center gap-3">
                    <img
                      src={logo}
                      alt=""
                      draggable="false"
                      class="
                        size-10
                        shrink-0
                        rounded-full
                        border-2
                        border-white
                        bg-blue-900
                        object-contain
                        p-1
                        shadow-md
                      "
                    />

                    <div class="min-w-0">
                      <p class="truncate text-sm font-black text-blue-950">
                        {slide.author}
                      </p>

                      {#if slide.publishedAt}
                        <time
                          datetime={slide.publishedAt}
                          class="
                            mt-0.5
                            block
                            text-xs
                            font-semibold
                            text-blue-950/55
                          "
                        >
                          {slide.publishedLabel}
                        </time>
                      {/if}
                    </div>
                  </div>

                  <span
                    aria-hidden="true"
                    class="
                      grid
                      size-11
                      shrink-0
                      place-items-center
                      rounded-full
                      bg-gradient-to-br
                      from-orange-400
                      to-orange-600
                      text-white
                      shadow-[0_4px_0_#c2410c]
                      transition-transform
                      duration-300
                      group-hover:-translate-y-0.5
                    "
                  >
                    <ArrowUpRightIcon size={18} strokeWidth={2.5} />
                  </span>
                </div>
              </div>
            </a>
          </div>
        {/each}
      </div>
    </div>

    {#if total > 1}
      <div
        class="
          mt-6
          flex
          items-center
          justify-between
          gap-4
          sm:mt-8
          lg:px-[9%]
        "
      >
        <div class="flex items-center gap-2">
          {#each slides as slide, index (slide.id)}
            <button
              type="button"
              aria-label={`Tampilkan berita ${index + 1}: ${slide.title}`}
              aria-current={index === active ? 'true' : undefined}
              onclick={() => goTo(index)}
              class={[
                'relative h-2.5 overflow-hidden rounded-full transition-[width,background-color] duration-500 ease-out focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-orange-300/60',
                index === active
                  ? 'w-12 bg-blue-900/15'
                  : 'w-2.5 bg-blue-900/20 hover:bg-blue-900/35'
              ]}
            >
              <span
                data-progress
                aria-hidden="true"
                style="transform: scaleX(0); transform-origin: left center;"
                class="absolute inset-0 rounded-full bg-orange-500"
              ></span>
            </button>
          {/each}
        </div>

        <div class="flex items-center gap-2 sm:gap-3">
          {#if autoplayEnabled}
            <button
              type="button"
              aria-label={userPaused
                ? 'Putar otomatis carousel'
                : 'Jeda putar otomatis carousel'}
              aria-pressed={userPaused}
              onclick={toggleAutoplay}
              class="
                grid
                size-10
                place-items-center
                rounded-full
                border
                border-blue-900/15
                bg-white
                text-blue-900
                shadow-sm
                transition
                hover:bg-blue-50
                focus-visible:outline-none
                focus-visible:ring-4
                focus-visible:ring-blue-300/60
              "
            >
              {#if userPaused}
                <PlayIcon size={16} strokeWidth={2.5} />
              {:else}
                <PauseIcon size={16} strokeWidth={2.5} />
              {/if}
            </button>
          {/if}

          <button
            type="button"
            aria-label="Berita sebelumnya"
            onclick={() => goTo(active - 1)}
            class="
              grid
              size-11
              place-items-center
              rounded-full
              bg-blue-800
              text-white
              shadow-[0_4px_0_#1e3a8a]
              transition
              hover:-translate-y-0.5
              hover:bg-blue-700
              active:translate-y-0.5
              active:shadow-none
              focus-visible:outline-none
              focus-visible:ring-4
              focus-visible:ring-blue-300/60
            "
          >
            <ArrowLeftIcon size={18} strokeWidth={2.5} />
          </button>

          <button
            type="button"
            aria-label="Berita berikutnya"
            onclick={() => goTo(active + 1)}
            class="
              grid
              size-11
              place-items-center
              rounded-full
              bg-blue-800
              text-white
              shadow-[0_4px_0_#1e3a8a]
              transition
              hover:-translate-y-0.5
              hover:bg-blue-700
              active:translate-y-0.5
              active:shadow-none
              focus-visible:outline-none
              focus-visible:ring-4
              focus-visible:ring-blue-300/60
            "
          >
            <ArrowRightIcon size={18} strokeWidth={2.5} />
          </button>
        </div>
      </div>
    {/if}
  </div>
</aside>
