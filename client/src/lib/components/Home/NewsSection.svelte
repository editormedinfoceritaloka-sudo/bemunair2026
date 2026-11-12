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

  const dateFormatter = new Intl.DateTimeFormat('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  });
  const formatCache = new Map<string, string>();

  function formatPublishedAt(value: string | null): string {
    if (!value) {
      return '';
    }

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

  const featuredNews = $derived(
    news[0]
  );

  const latestNews = $derived(
    news.slice(1, 4)
  );

  let sectionElement!: HTMLElement;
  let glowElement!: HTMLDivElement;
  let headingElement!: HTMLHeadingElement;
  let descriptionElement!: HTMLParagraphElement;
  let featuredElement!: HTMLDivElement;
  let latestSectionElement!: HTMLDivElement;
  let latestHeaderElement!: HTMLDivElement;
  let cardsElement!: HTMLDivElement;

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      return;
    }

    const context = gsap.context(() => {
      gsap.set(glowElement, {
        opacity: 0,
        scale: 0.75
      });

      gsap.set(
        [
          headingElement,
          descriptionElement
        ],
        {
          opacity: 0,
          y: 40
        }
      );

      if (featuredElement) {
        gsap.set(featuredElement, {
          opacity: 0,
          y: 50,
          scale: 0.98
        });
      }

      if (latestHeaderElement) {
        gsap.set(latestHeaderElement, {
          opacity: 0,
          y: 35
        });
      }

      if (cardsElement) {
        gsap.set(
          cardsElement.querySelectorAll(
            '[data-news-card]'
          ),
          {
            opacity: 0,
            y: 45
          }
        );
      }

      gsap.to(glowElement, {
        opacity: 1,
        scale: 1,
        duration: 1.4,
        ease: 'power3.out',
        scrollTrigger: {
          trigger: sectionElement,
          start: 'top 80%',
          once: true
        }
      });

      gsap.to(
        [
          headingElement,
          descriptionElement
        ],
        {
          opacity: 1,
          y: 0,
          duration: 0.85,
          stagger: 0.12,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: sectionElement,
            start: 'top 75%',
            once: true
          }
        }
      );

      if (featuredElement) {
        gsap.to(featuredElement, {
          opacity: 1,
          y: 0,
          scale: 1,
          duration: 0.95,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: featuredElement,
            start: 'top 86%',
            once: true
          }
        });
      }

      if (latestHeaderElement) {
        gsap.to(latestHeaderElement, {
          opacity: 1,
          y: 0,
          duration: 0.75,
          ease: 'power3.out',
          scrollTrigger: {
            trigger: latestSectionElement,
            start: 'top 86%',
            once: true
          }
        });
      }

      if (cardsElement) {
        gsap.to(
          cardsElement.querySelectorAll(
            '[data-news-card]'
          ),
          {
            opacity: 1,
            y: 0,
            duration: 0.8,
            stagger: 0.14,
            ease: 'power3.out',
            scrollTrigger: {
              trigger: cardsElement,
              start: 'top 88%',
              once: true
            }
          }
        );
      }
    }, sectionElement);

    return () => {
      context.revert();
    };
  });
</script>

<section
  id="news"
  bind:this={sectionElement}
  class="
    relative min-h-screen
    w-full overflow-hidden
    bg-linear-to-b
    from-white
    via-blue-50/60
    to-blue-100
    py-24
    sm:py-28
    lg:py-32
  "
>
  <div
    bind:this={glowElement}
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      left-1/2 top-16
      size-[34rem]
      -translate-x-1/2
      rounded-full
      bg-orange-200/20
      blur-[130px]
      will-change-transform
      sm:size-[44rem]
      lg:size-[54rem]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      -left-32 top-1/2
      size-80
      rounded-full
      bg-blue-200/20
      blur-[110px]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      -right-36 bottom-20
      size-96
      rounded-full
      bg-orange-200/20
      blur-[120px]
    "
  ></div>

  <div
    class="
      relative z-10
      mx-auto
      w-full max-w-7xl
      px-5
      sm:px-8
      lg:px-12
    "
  >
    <header
      class="
        mx-auto
        flex max-w-5xl
        flex-col
        items-center
        text-center
      "
    >
      <h2
        bind:this={headingElement}
        class="
          text-[clamp(4rem,10vw,8rem)]
          leading-[0.82]
          font-black
          uppercase
          tracking-[-0.07em]
          text-orange-500
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_9px_0_rgba(163,61,10,0.45)]
          will-change-transform
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
          mt-12
          will-change-transform
          sm:mt-16
        "
      >
        <CoverNews
          item={featuredNews}
          {logo}
          {organization}
        />
      </div>
    {/if}

    {#if latestNews.length > 0}
      <div
        bind:this={latestSectionElement}
        class="
          mt-16
          sm:mt-20
          lg:mt-24
        "
      >
        <div
          bind:this={latestHeaderElement}
          class="
            mb-8
            flex items-end
            justify-between
            gap-5
            border-b
            border-blue-900/15
            pb-4
            will-change-transform
            sm:mb-10
          "
        >
          <div>
            <h3
              class="
                text-2xl
                leading-none
                font-black
                tracking-[-0.04em]
                text-blue-950
                sm:text-3xl
              "
            >
              Berita Terbaru
            </h3>
          </div>

          <a
            href="/news"
            aria-label="Lihat semua berita"
            class="
              group
              inline-flex shrink-0
              items-center gap-2
              text-xs
              font-bold
              text-orange-600
              transition-colors
              duration-300
              hover:text-orange-800
              sm:text-sm
            "
          >
            <span>
              Lihat Semua Berita
            </span>

            <ArrowRightIcon
              size={18}
              strokeWidth={2.5}
              class="
                -rotate-45
                transition-transform
                duration-300
                group-hover:-translate-y-1
                group-hover:translate-x-1
              "
            />
          </a>
        </div>

        <div
          bind:this={cardsElement}
          class="
            grid
            grid-cols-1
            gap-x-7
            gap-y-12
            md:grid-cols-2
            lg:grid-cols-3
          "
        >
          {#each latestNews as item (item.id)}
            <div
              data-news-card
              class="will-change-transform"
            >
              <CardNews
                {item}
                {logo}
                {organization}
              />
            </div>
          {/each}
        </div>
      </div>
    {/if}

    {#if news.length === 0}
      <div
        class="
          mx-auto
          mt-14
          flex max-w-8xl
          flex-col
          items-center
          justify-center
          rounded-[28px]
          border
          border-orange-900/10
          bg-white/55
          px-6
          py-16
          text-center
          backdrop-blur-sm
          sm:px-10
        "
      >
        <h3
          class="
            text-2xl
            font-black
            tracking-[-0.04em]
            text-orange-600
            sm:text-3xl
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
          "
        >
          Berita dan cerita terbaru Kabinet Cerita Loka
          akan hadir di sini.
        </p>
      </div>
    {/if}
  </div>
</section>