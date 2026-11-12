
<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import { ArrowRightIcon } from '@lucide/svelte';

  import CardNews from './news/CardNews.svelte';
  import CoverNews from './news/CoverNews.svelte';

  import type { ArticleListItem } from '$lib/types';

  type NewsItem = {
    id: string;
    slug: string;
    title: string;
    excerpt: string;
    coverImage: string;
    category: string;
    publishedAt: string;
  };

  let {
    articles = [],
    logo = '/logo/logo-kabinet.webp',
    organization = 'Kabinet Cerita Loka'
  }: {
    articles?: ArticleListItem[];
    logo?: string;
    organization?: string;
  } = $props();

  // ========================================
  // DATE FORMATTING
  // ========================================

  const dateFormatter = new Intl.DateTimeFormat(
    'id-ID',
    {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
      timeZone: 'Asia/Jakarta'
    }
  );

  const formatCache = new Map<string, string>();

  function formatPublishedAt(
    value: string | null
  ): string {
    if (!value) return '';

    const cached = formatCache.get(value);

    if (cached !== undefined) {
      return cached;
    }

    const date = new Date(value);

    if (Number.isNaN(date.getTime())) {
      formatCache.set(value, value);
      return value;
    }

    const formatted = dateFormatter.format(date);

    formatCache.set(value, formatted);

    return formatted;
  }

  // ========================================
  // DATA TRANSFORMATION
  // ========================================

  const news = $derived(
    articles.map(
      (article): NewsItem => ({
        id: String(article.id),
        slug: article.slug,
        title: article.title,
        excerpt: article.excerpt ?? '',
        coverImage:
          article.cover_image ??
          '/landing/news/test.webp',
        category: 'Berita',
        publishedAt: formatPublishedAt(
          article.published_at
        )
      })
    )
  );

  const featuredNews = $derived(news[0]);

  const latestNews = $derived(
    news.slice(1, 4)
  );

  // ========================================
  // DOM REFERENCES
  // ========================================

  let sectionElement: HTMLElement;
  let glowElement: HTMLDivElement;
  let headingElement: HTMLHeadingElement;

  let featuredElement: HTMLDivElement;
  let latestSectionElement: HTMLElement;
  let latestHeaderElement: HTMLDivElement;
  let cardsElement: HTMLDivElement;

  // ========================================
  // GSAP ANIMATION
  // ========================================

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) return;

    const isMobile = window.matchMedia(
      '(max-width: 639px)'
    ).matches;

    const offset = isMobile ? 24 : 40;

    const context = gsap.context(() => {
      // Initial states

      gsap.set(glowElement, {
        opacity: 0,
        scale: 0.85
      });

      gsap.set(headingElement, {
        autoAlpha: 0,
        y: offset
      });

      if (featuredElement) {
        gsap.set(featuredElement, {
          autoAlpha: 0,
          y: offset + 10,
          scale: 0.985
        });
      }

      if (latestHeaderElement) {
        gsap.set(latestHeaderElement, {
          autoAlpha: 0,
          y: offset
        });
      }

      const cards = cardsElement
        ? gsap.utils.toArray<HTMLElement>(
            '[data-news-card]',
            cardsElement
          )
        : [];

      if (cards.length > 0) {
        gsap.set(cards, {
          autoAlpha: 0,
          y: offset + 5
        });
      }

      // Background glow

      gsap.to(glowElement, {
        opacity: 1,
        scale: 1,
        duration: 1.2,
        ease: 'power3.out',
        scrollTrigger: {
          trigger: sectionElement,
          start: 'top 85%',
          once: true
        }
      });

      // Main heading

      gsap.to(headingElement, {
        autoAlpha: 1,
        y: 0,
        duration: 0.85,
        ease: 'power3.out',
        scrollTrigger: {
          trigger: sectionElement,
          start: 'top 82%',
          once: true
        }
      });

      // Featured news

      if (featuredElement) {
        gsap.to(featuredElement, {
          autoAlpha: 1,
          y: 0,
          scale: 1,
          duration: 0.9,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: featuredElement,
            start: 'top 90%',
            once: true
          }
        });
      }

      // Latest news heading

      if (
        latestSectionElement &&
        latestHeaderElement
      ) {
        gsap.to(latestHeaderElement, {
          autoAlpha: 1,
          y: 0,
          duration: 0.75,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: latestSectionElement,
            start: 'top 90%',
            once: true
          }
        });
      }

      // Latest news cards

      if (cardsElement && cards.length > 0) {
        gsap.to(cards, {
          autoAlpha: 1,
          y: 0,
          duration: 0.8,
          stagger: isMobile ? 0.08 : 0.12,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: cardsElement,
            start: 'top 92%',
            once: true
          }
        });
      }
    }, sectionElement);

    return () => {
      context.revert();
    };
  });
</script>

<section
  id="news"
  aria-labelledby="news-heading"
  bind:this={sectionElement}
  class="
    relative isolate
    w-full
    overflow-hidden
    bg-linear-to-b
    from-white
    via-blue-50/60
    to-blue-50

    pt-16 pb-14

    sm:pt-20 sm:pb-16

    md:pt-24 md:pb-18

    lg:pt-28 lg:pb-20

    xl:pt-32 xl:pb-24
  "
>

  <div
    bind:this={glowElement}
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      top-12 left-1/2
      -translate-x-1/2

      size-[20rem]
      rounded-full
      bg-orange-200/20
      blur-[90px]

      sm:top-16
      sm:size-[32rem]
      sm:blur-[110px]

      lg:size-[46rem]
      lg:blur-[130px]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      -left-24
      top-1/2

      size-56
      rounded-full
      bg-blue-200/20
      blur-[85px]

      sm:-left-32
      sm:size-72
      sm:blur-[100px]

      lg:size-80
      lg:blur-[110px]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      -right-24
      bottom-16

      size-60
      rounded-full
      bg-orange-200/20
      blur-[90px]

      sm:-right-32
      sm:size-80
      sm:blur-[110px]

      lg:size-96
      lg:blur-[120px]
    "
  ></div>


  <div
    class="
      relative z-10
      mx-auto
      w-full
      max-w-7xl
      min-w-0

      px-4

      sm:px-6

      md:px-8

      lg:px-10

      xl:px-12
    "
  >
    <header
      class="
        mx-auto
        flex
        w-full
        max-w-5xl
        min-w-0
        flex-col
        items-center
        justify-center
        text-center
      "
    >
      <h2
        id="news-heading"
        bind:this={headingElement}
        class="
          max-w-full
          py-2

          text-[clamp(2.8rem,9vw,7rem)]
          leading-[0.82]
          font-black
          uppercase
          tracking-[-0.07em]

          text-orange-500

          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]

          drop-shadow-[0_9px_0_rgba(163,61,10,0.45)]

          sm:text-[clamp(3.5rem,9vw,7rem)]
          sm:[-webkit-text-stroke:3px_white] 
        "
      >
        News
      </h2>
    </header>


    {#if featuredNews}
      <div
        bind:this={featuredElement}
        class="
          mt-9
          w-full
          min-w-0

          sm:mt-12

          md:mt-14

          lg:mt-16
        "
      >
        <div
          class="
            w-full
            min-w-0
            max-w-full
          "
        >
          <CoverNews
            item={featuredNews}
            {logo}
            {organization}
          />
        </div>
      </div>
    {/if}


    {#if latestNews.length > 0}
      <section
        aria-labelledby="latest-news-heading"
        bind:this={latestSectionElement}
        class="
          mt-12
          w-full
          min-w-0

          sm:mt-16

          md:mt-18

          lg:mt-20

          xl:mt-24
        "
      >

        <div
          bind:this={latestHeaderElement}
          class="
            mb-6
            flex
            w-full
            min-w-0
            items-end
            justify-between
            gap-3

            border-b
            border-blue-900/15
            pb-4

            sm:mb-8
            sm:gap-5

            lg:mb-10
          "
        >
          <div
            class="
              min-w-0
              flex-1
            "
          >
            <h3
              id="latest-news-heading"
              class="
                text-xl
                leading-tight
                font-black
                tracking-[-0.04em]
                text-blue-950

                sm:text-2xl

                lg:text-3xl
                lg:leading-none
              "
            >
              Berita Terbaru
            </h3>
          </div>

          <!-- View all news link -->

          <a
            href="/news"
            aria-label="Lihat semua berita"
            class="
              group
              inline-flex
              shrink-0
              items-center
              justify-center
              gap-1.5

              whitespace-nowrap

              text-[11px]
              leading-none
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
            <span class="sm:hidden">
              Lihat Semua
            </span>

            <span class="hidden sm:inline">
              Lihat Semua Berita
            </span>

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
          bind:this={cardsElement}
          class="
            grid
            w-full
            min-w-0

            grid-cols-1

            gap-x-5
            gap-y-7

            sm:grid-cols-2
            sm:gap-x-5
            sm:gap-y-9

            md:gap-x-6
            md:gap-y-10

            lg:grid-cols-3
            lg:gap-x-7
            lg:gap-y-10

            xl:gap-y-12
          "
        >
          {#each latestNews as item (item.id)}
            <div
              data-news-card
              class="
                h-full
                w-full
                min-w-0
                max-w-full

                [&>*]:h-full
                [&>*]:max-w-full
              "
            >
              <CardNews
                {item}
                {logo}
                {organization}
              />
            </div>
          {/each}
        </div>
      </section>
    {/if}


    {#if news.length === 0}
      <div
        class="
          mx-auto
          mt-10

          flex
          w-full
          max-w-3xl
          flex-col
          items-center
          justify-center

          rounded-[24px]
          border
          border-orange-900/10
          bg-white/55

          px-5
          py-12

          text-center
          backdrop-blur-sm

          sm:mt-14
          sm:rounded-[28px]
          sm:px-10
          sm:py-16
        "
      >
        <h3
          class="
            text-xl
            font-black
            tracking-[-0.04em]
            text-orange-600

            sm:text-2xl

            lg:text-3xl
          "
        >
          Belum Ada Berita
        </h3>

        <p
          class="
            mt-3
            max-w-md

            text-sm
            leading-6
            font-medium
            text-slate-600

            sm:text-base
            sm:leading-7
          "
        >
          Berita dan cerita terbaru Kabinet Cerita Loka
          akan hadir di sini.
        </p>
      </div>
    {/if}
  </div>
</section>
