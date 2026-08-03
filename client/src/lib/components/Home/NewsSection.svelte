<script lang="ts">
  import CardNews from './news/CardNews.svelte';
  import CoverNews from './news/CoverNews.svelte';

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
      coverImage: '/landing/news/news-1.jpg',
      category: 'Sosial dan Politik',
      publishedAt: '2026-06-24'
    },
    {
      id: 'news-2',
      slug: 'ruang-ekspresi-mahasiswa',
      title: 'Ruang Ekspresi Mahasiswa Hadir dengan Konsep yang Lebih Segar',
      excerpt:
        'Program ini menjadi wadah bagi mahasiswa Universitas Airlangga untuk menyampaikan gagasan, karya, dan aspirasi melalui aktivitas yang kreatif dan inklusif.',
      coverImage: '/landing/news/news-2.jpg',
      category: 'Seni dan Budaya',
      publishedAt: '2026-06-22'
    },
    {
      id: 'news-3',
      slug: 'kolaborasi-kabinet-cerita-loka',
      title: 'Kolaborasi Kabinet Cerita Loka untuk Pengembangan Mahasiswa',
      excerpt:
        'Kabinet Cerita Loka memperkuat kolaborasi lintas kementerian untuk menghadirkan program kerja yang memberikan dampak nyata bagi seluruh mahasiswa.',
      coverImage: '/landing/news/news-3.jpg',
      category: 'Pengembangan Mahasiswa',
      publishedAt: '2026-06-20'
    },
    {
      id: 'news-4',
      slug: 'peluncuran-program-unggulan',
      title: 'Kabinet Cerita Loka Meluncurkan Rangkaian Program Unggulan',
      excerpt:
        'Rangkaian program unggulan dirancang untuk memperluas ruang partisipasi mahasiswa serta membangun lingkungan kampus yang aktif, suportif, dan berkelanjutan.',
      coverImage: '/landing/news/news-4.jpg',
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
</script>

<section
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
    class="
      pointer-events-none absolute
      left-1/2 top-24
      h-[420px] w-[720px]
      -translate-x-1/2 rounded-full
      bg-blue-200/20 blur-[120px]
    "
  ></div>

  <div class="relative mx-auto w-full max-w-6xl">
    <header class="mx-auto max-w-4xl text-center">
      <h2
        id="news-heading"
        class="
          text-6xl leading-none font-black
          tracking-[-0.065em] text-orange-600
          drop-shadow-[0_4px_0_rgba(255,255,255,0.95)]
          sm:text-7xl
          lg:text-8xl
        "
      >
        News
      </h2>

      <p
        class="
          mx-auto mt-4 max-w-3xl
          text-base leading-7 font-semibold
          text-slate-900
          sm:text-lg
          lg:text-xl
        "
      >
        Explore the latest programs, research developments, and cultural
        festivals organized by Cabinet Cerita Loka
      </p>
    </header>

    {#if featuredNews}
      <div class="mt-12 sm:mt-16">
        <CoverNews
          item={featuredNews}
          {logo}
          {organization}
        />
      </div>
    {/if}

    {#if latestNews.length > 0}
      <div class="mt-16 sm:mt-20">
        <div
          class="
            mb-8 flex items-end justify-between gap-5
            border-b border-blue-900/15 pb-4
          "
        >
          <div>
            <p
              class="
                text-xs font-extrabold uppercase
                tracking-[0.18em] text-blue-700
              "
            >
              Latest Stories
            </p>

            <h3
              class="
                mt-1 text-3xl font-black
                tracking-[-0.04em] text-slate-950
                sm:text-4xl
              "
            >
              All News
            </h3>
          </div>

          <a
            href="/news"
            class="
              shrink-0 text-sm font-bold
              text-blue-700 transition-colors
              hover:text-blue-950
            "
          >
            View all news
          </a>
        </div>

        <div
          class="
            grid grid-cols-1
            gap-x-7 gap-y-12
            md:grid-cols-2
            lg:grid-cols-3
          "
        >
          {#each latestNews as item (item.id)}
            <CardNews
              {item}
              {logo}
              {organization}
            />
          {/each}
        </div>
      </div>
    {/if}
  </div>
</section>