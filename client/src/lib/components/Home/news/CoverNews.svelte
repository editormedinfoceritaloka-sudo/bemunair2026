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

<article class="group">
  <a
    href={`/news/${encodeURIComponent(item.slug)}`}
    aria-label={`Baca berita utama ${item.title}`}
    class="
      grid items-stretch gap-6
      md:grid-cols-[1.08fr_0.92fr]
      lg:gap-9
    "
  >
    <div
      class="
        relative min-h-[260px] overflow-hidden
        rounded-[20px] bg-blue-100
        sm:min-h-[340px]
        md:min-h-[390px]
      "
    >
      <img
        src={item.coverImage}
        alt={item.title}
        draggable="false"
        class="
          absolute inset-0 h-full w-full object-cover
          transition duration-700
          group-hover:scale-[1.035]
        "
      />

      <div
        class="
          pointer-events-none absolute inset-0
          bg-gradient-to-t from-black/15 via-transparent to-transparent
        "
      ></div>
    </div>

    <div class="flex flex-col justify-center py-1 md:py-4">
      <h3
        class="
          text-3xl leading-[1.08] font-extrabold
          tracking-[-0.045em] text-slate-950
          transition-colors duration-200
          group-hover:text-blue-700
          sm:text-4xl
          lg:text-[42px]
        "
      >
        {item.title}
      </h3>

      <p
        class="
          mt-5 line-clamp-5
          text-sm leading-7 text-slate-700
          sm:text-base
        "
      >
        {item.excerpt}
      </p>

      <div class="mt-7 flex items-center gap-3">
        <img
          src={logo}
          alt={organization}
          draggable="false"
          class="h-12 w-12 shrink-0 object-contain"
        />

        <div class="min-w-0">
          <p
            class="
              truncate text-sm font-bold
              text-slate-950 sm:text-base
            "
          >
            {item.category}
          </p>

          <p class="mt-1 text-xs font-medium text-slate-600 sm:text-sm">
            {formatDate(item.publishedAt)}
          </p>
        </div>
      </div>
    </div>
  </a>
</article>