<script lang="ts">
  import { resolve } from '$app/paths';

  import CardNews from '$lib/components/Home/news/CardNews.svelte';

  import type { ArticleListItem } from '$lib/types';

  let {
    articles = [],
    search = '',
    logo = '/logo/logo-kabinet.webp',
    organization = 'Kabinet Cerita Loka'
  }: {
    articles?: ArticleListItem[];
    search?: string;
    logo?: string;
    organization?: string;
  } = $props();

  function mapArticle(article: ArticleListItem) {
    return {
      id: String(article.id),
      slug: article.slug,
      title: article.title,
      excerpt: article.excerpt ?? '',
      coverImage: article.cover_image ?? '/landing/news/test.webp',
      publishedAt: article.published_at ?? '',
      authorName: article.author?.name ?? organization
    };
  }

  const items = $derived(articles.map(mapArticle));
</script>

{#if articles.length > 0}
  <div
    class="
      mt-14
      grid
      grid-cols-1
      gap-x-7
      gap-y-14
      md:grid-cols-2
      lg:grid-cols-3
    "
  >
    {#each items as item (item.id)}
      <CardNews {item} {logo} {organization} />
    {/each}
  </div>
{:else}
  <div
    class="
      mx-auto mt-16
      flex max-w-2xl
      flex-col
      items-center
      text-center
    "
  >
    <h2
      class="
        text-3xl
        font-black
        tracking-[-0.045em]
        text-orange-600
        sm:text-4xl
      "
    >
      Tidak Ada Berita
    </h2>

    <p
      class="
        mt-4 max-w-md
        text-sm
        leading-7
        font-medium
        text-blue-950/55
        sm:text-base
      "
    >
      {#if search}
        Tidak ada berita yang cocok dengan pencarian "{search}".
      {:else}
        Belum ada berita yang dipublikasikan.
      {/if}
    </p>

    {#if search}
      <a
        href={resolve('/news')}
        aria-label="Lihat semua berita"
        class="
          mt-7
          inline-flex
          items-center
          justify-center
          rounded-full
          bg-orange-500
          px-7 py-3
          text-sm
          font-black
          text-white
          shadow-[0_6px_0_rgba(163,61,10,0.3)]
          transition
          hover:-translate-y-0.5
          hover:bg-orange-600
          active:translate-y-1
          active:shadow-none
        "
      >
        Lihat Semua Berita
      </a>
    {/if}
  </div>
{/if}
