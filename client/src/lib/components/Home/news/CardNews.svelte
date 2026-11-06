<script lang="ts">
  import { resolve } from '$app/paths';

  type NewsItem = {
    slug: string;
    title: string;
    excerpt: string;
    coverImage: string;
    publishedAt: string;
    authorName?: string;
  };

  let {
    item,
    logo = '/logo/logo-kabinet.webp',
    organization = 'Kabinet Cerita Loka'
  }: {
    item: NewsItem;
    logo?: string;
    organization?: string;
  } = $props();

  function formatDate(value: string): string {
    if (!value) {
      return '';
    }

    const date = new Date(value);

    if (Number.isNaN(date.getTime())) {
      return value;
    }

    return new Intl.DateTimeFormat('id-ID', {
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    }).format(date);
  }
</script>

<article class="group min-w-0">
  <a
    href={resolve(`/news/${encodeURIComponent(item.slug)}`)}
    aria-label={`Baca berita ${item.title}`}
    class="block"
  >
    <div
      class="
        relative aspect-[1.45/1] w-full
        overflow-hidden rounded-[18px]
        bg-blue-100
      "
    >
      <img
        src={item.coverImage}
        alt={item.title}
        draggable="false"
        loading="lazy"
        decoding="async"
        width="640"
        height="442"
        class="
          h-full w-full object-cover
          transition duration-500
          group-hover:scale-[1.04]
        "
      />

      <div
        class="
          pointer-events-none absolute inset-0
          bg-gradient-to-t from-black/10 via-transparent to-transparent
          opacity-0 transition duration-300
          group-hover:opacity-100
        "
      ></div>
    </div>

    <div class="pt-4">
      <h3
        class="
          line-clamp-2
          text-lg leading-[1.2] font-extrabold
          tracking-[-0.025em] text-slate-950
          transition-colors duration-200
          group-hover:text-blue-700
          sm:text-xl
        "
      >
        {item.title}
      </h3>

      {#if item.excerpt}
        <p class="mt-3 line-clamp-4 text-sm leading-6 text-slate-700">
          {item.excerpt}
        </p>
      {/if}

      <div class="mt-5 flex items-center gap-3">
        <img
          src={logo}
          alt={organization}
          draggable="false"
          class="h-10 w-10 shrink-0 object-contain"
        />

        <div class="min-w-0">
          <p class="truncate text-sm font-bold text-slate-950">
            {item.authorName ?? organization}
          </p>

          {#if item.publishedAt}
            <time
              datetime={item.publishedAt}
              class="mt-0.5 block text-xs font-medium text-slate-600"
            >
              {formatDate(item.publishedAt)}
            </time>
          {/if}
        </div>
      </div>
    </div>
  </a>
</article>
