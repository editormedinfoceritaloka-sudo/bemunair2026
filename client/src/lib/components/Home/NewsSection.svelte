<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import CardNews from './news/CardNews.svelte';
  import CoverNews from './news/CoverNews.svelte';
  import { ArrowRightIcon } from '@lucide/svelte';

  type NewsItem = {
    id: string;
    slug: string;
    title: string;
    excerpt: string;
    coverImage: string;
    category: string;
    publishedAt: string;
  };

  const exampleNews: NewsItem[] = [
    {
      id: 'news-1',
      slug: 'festival-bunga-cerita-loka',
      title: 'Festival Bunga Cerita Loka: Merawat Budaya dan Lingkungan',
      excerpt:
        'Kabinet Cerita Loka menghadirkan festival budaya dan lingkungan sebagai ruang kolaborasi mahasiswa dalam merawat kreativitas, keberagaman, serta kepedulian terhadap alam.',
      coverImage: '/landing/news/test.png',
      category: 'Sosial dan Politik',
      publishedAt: '2026-06-24'
    },
    {
      id: 'news-2',
      slug: 'ruang-ekspresi-mahasiswa',
      title: 'Ruang Ekspresi Mahasiswa Hadir dengan Konsep yang Lebih Segar',
      excerpt:
        'Program ini menjadi wadah bagi mahasiswa Universitas Airlangga untuk menyampaikan gagasan, karya, dan aspirasi melalui aktivitas yang kreatif dan inklusif.',
      coverImage: '/landing/news/test.png',
      category: 'Seni dan Budaya',
      publishedAt: '2026-06-22'
    },
    {
      id: 'news-3',
      slug: 'kolaborasi-kabinet-cerita-loka',
      title: 'Kolaborasi Kabinet Cerita Loka untuk Pengembangan Mahasiswa',
      excerpt:
        'Kabinet Cerita Loka memperkuat kolaborasi lintas kementerian untuk menghadirkan program kerja yang memberikan dampak nyata bagi seluruh mahasiswa.',
      coverImage: '/landing/news/test.png',
      category: 'Pengembangan Mahasiswa',
      publishedAt: '2026-06-20'
    },
    {
      id: 'news-4',
      slug: 'peluncuran-program-unggulan',
      title: 'Kabinet Cerita Loka Meluncurkan Rangkaian Program Unggulan',
      excerpt:
        'Rangkaian program unggulan dirancang untuk memperluas ruang partisipasi mahasiswa serta membangun lingkungan kampus yang aktif, suportif, dan berkelanjutan.',
      coverImage: '/landing/news/test.png',
      category: 'Kabinet Cerita Loka',
      publishedAt: '2026-06-18'
    }
  ];

  let {
    news = exampleNews,
    logo = '/logo/logo-kabinet.png',
    organization = 'Kabinet Cerita Loka'
  }: {
    news?: NewsItem[];
    logo?: string;
    organization?: string;
  } = $props();

  const featuredNews = $derived(news[0]);
  const latestNews = $derived(news.slice(1, 4));

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
        scale: 0.7
      });

      gsap.set(headingElement, {
        y: 70,
        opacity: 0,
        scale: 0.9,
        filter: 'blur(10px)'
      });

      gsap.set(descriptionElement, {
        y: 35,
        opacity: 0,
        filter: 'blur(6px)'
      });

      if (featuredElement) {
        gsap.set(featuredElement, {
          y: 80,
          opacity: 0,
          scale: 0.97,
          transformOrigin: 'center top'
        });
      }

      const introTimeline = gsap.timeline({
        defaults: {
          ease: 'power3.out'
        },
        scrollTrigger: {
          trigger: sectionElement,
          start: 'top 72%',
          once: true
        }
      });

      introTimeline
        .to(
          glowElement,
          {
            opacity: 1,
            scale: 1,
            duration: 1.3
          },
          0
        )
        .to(
          headingElement,
          {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 0.9
          },
          0.1
        )
        .to(
          descriptionElement,
          {
            y: 0,
            opacity: 1,
            filter: 'blur(0px)',
            duration: 0.75
          },
          0.35
        );

      if (featuredElement) {
        introTimeline.to(
          featuredElement,
          {
            y: 0,
            opacity: 1,
            scale: 1,
            duration: 1
          },
          0.55
        );
      }

      introTimeline.set(
        [
          glowElement,
          headingElement,
          descriptionElement,
          featuredElement
        ].filter(Boolean),
        {
          clearProps: 'opacity,transform,filter'
        }
      );

      if (
        latestSectionElement &&
        latestHeaderElement &&
        cardsElement
      ) {
        const cards = gsap.utils.toArray<HTMLElement>(
          '[data-news-card]',
          cardsElement
        );

        gsap.set(latestHeaderElement, {
          y: 45,
          opacity: 0
        });

        gsap.set(cards, {
          y: 70,
          opacity: 0,
          scale: 0.96,
          transformOrigin: 'center top'
        });

        const latestTimeline = gsap.timeline({
          defaults: {
            ease: 'power3.out'
          },
          scrollTrigger: {
            trigger: latestSectionElement,
            start: 'top 78%',
            once: true
          }
        });

        latestTimeline
          .to(latestHeaderElement, {
            y: 0,
            opacity: 1,
            duration: 0.75
          })
          .to(
            cards,
            {
              y: 0,
              opacity: 1,
              scale: 1,
              duration: 0.85,
              stagger: 0.16
            },
            0.25
          )
          .set([latestHeaderElement, ...cards], {
            clearProps: 'opacity,transform'
          });
      }
    }, sectionElement);

    return () => {
      context.revert();
    };
  });
</script>

<section
  bind:this={sectionElement}
  id="news"
  aria-labelledby="news-heading"
  class="
    relative min-h-screen overflow-hidden
    bg-blue-50 px-5 pb-24 pt-28
    sm:px-8 sm:pt-32
    lg:px-12 lg:pb-32 lg:pt-36
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none absolute
      left-1/2 top-24
      h-[420px] w-[720px]
      -translate-x-1/2 rounded-full
      bg-blue-200/20 blur-[120px]
      will-change-transform
    "
  ></div>

  <div class="relative mx-auto w-full max-w-6xl">
    <header class="mx-auto max-w-4xl text-center">
      <h2
        bind:this={headingElement}
        id="news-heading"
        class="
          text-6xl leading-none font-black
          tracking-[-0.065em] text-orange-600
          drop-shadow-[0_4px_0_rgba(255,255,255,0.95)]
          will-change-transform
          sm:text-7xl
          lg:text-8xl
        "
      >
        News
      </h2>

      <p
        bind:this={descriptionElement}
        class="
          mx-auto mt-4 max-w-3xl
          text-base leading-7 font-semibold
          text-slate-900
          will-change-transform
          sm:text-lg
          lg:text-xl
        "
      >
        Explore the latest programs, research developments, and cultural
        festivals organized by Cabinet Cerita Loka
      </p>
    </header>

    {#if featuredNews}
      <div
        bind:this={featuredElement}
        class="mt-12 will-change-transform sm:mt-16"
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
        class="mt-16 sm:mt-20"
      >
        <div
          bind:this={latestHeaderElement}
          class="
            mb-8 flex items-end justify-between gap-5
            border-b border-blue-900/15 pb-4
            will-change-transform
          "
        >
          <div>
          </div>

          <a
          href="/news"
          class="
            group inline-flex shrink-0
            items-center gap-2
            text-sm font-bold
            text-blue-700
            transition-colors
            hover:text-blue-950
          "
        >
          <span>Lihat Semua Berita</span>

          <ArrowRightIcon
            size={18}
            strokeWidth={2.5}
            class="
              -rotate-45
              transition-transform duration-300
              group-hover:-translate-y-1
            "
          />
        </a>
        </div>

        <div
          bind:this={cardsElement}
          class="
            grid grid-cols-1
            gap-x-7 gap-y-12
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
  </div>
</section>