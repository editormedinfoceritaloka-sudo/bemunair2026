<script lang="ts">
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
    item,
    logo = '/logo/logo-kabinet.png',
    organization = 'Kabinet Cerita Loka'
  }: {
    item: NewsItem;
    logo?: string;
    organization?: string;
  } = $props();

  function formatDate(value: string): string {
    const date = new Date(value);

    if (Number.isNaN(date.getTime())) {
      return value;
    }

    return new Intl.DateTimeFormat('en-GB', {
      day: '2-digit',
      month: 'long',
      year: 'numeric'
    }).format(date);
  }
</script>

<article class="group min-w-0">
  <a
    href={`/news/${encodeURIComponent(item.slug)}`}
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

      <p
        class="
          mt-3 line-clamp-4
          text-sm leading-6 text-slate-700
        "
      >
        {item.excerpt}
      </p>

      <div class="mt-5 flex items-center gap-3">
        <img
          src={logo}
          alt={organization}
          draggable="false"
          class="h-10 w-10 shrink-0 object-contain"
        />

        <div class="min-w-0">
          <p
            class="
              truncate text-sm font-bold
              text-slate-950
            "
          >
            {item.category}
          </p>

          <p class="mt-0.5 text-xs font-medium text-slate-600">
            {formatDate(item.publishedAt)}
          </p>
        </div>
      </div>
    </div>
  </a>
</article>