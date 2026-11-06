<script lang="ts">
  import { ArrowLeftIcon } from '@lucide/svelte';
  import { resolve } from '$app/paths';

  import ArticleBody from '$lib/components/news/ArticleBody.svelte';

  import type { PageData } from './$types';

  let {
    data
  }: {
    data: PageData;
  } = $props();

  const article = $derived(data.article);
  const readTime = $derived(data.readTime);

  const titleLines = $derived.by(() => {
    const words = article.title.trim().split(/\s+/).filter(Boolean);

    if (words.length === 0) {
      return {
        first: '',
        second: ''
      };
    }

    if (words.length === 1) {
      return {
        first: words[0],
        second: ''
      };
    }

    const splitIndex = Math.ceil(words.length / 2);

    return {
      first: words.slice(0, splitIndex).join(' '),
      second: words.slice(splitIndex).join(' ')
    };
  });

  function formatDate(value: string | null): string {
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

<article
  class="
    relative
    min-h-screen
    overflow-hidden
    bg-blue-50
    pb-28
    pt-28
    sm:pt-32
    lg:pb-36
    lg:pt-40
  "
>
  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute
      left-1/2
      top-10
      size-[42rem]
      -translate-x-1/2
      rounded-full
      bg-blue-300/15
      blur-[140px]
    "
  ></div>

  <header
    class="
      relative z-10
      mx-auto
      flex w-full
      max-w-7xl
      flex-col
      items-center
      px-5
      text-center
      sm:px-8
      lg:px-12
    "
  >
    <h1
      class="
        max-w-6xl
        text-[clamp(2.5rem,6.4vw,6rem)]
        leading-[0.9]
        font-black
        tracking-[-0.065em]
      "
    >
      <span
        class="
          block
          text-blue-300
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_9px_0_rgba(30,64,175,0.5)]
          sm:[-webkit-text-stroke:3px_white]
        "
      >
        {titleLines.first}
      </span>

      {#if titleLines.second}
        <span
          class="
            relative
            left-[0.15em]
            block
            text-blue-700
            [-webkit-text-stroke:2px_white]
            [paint-order:stroke_fill]
            drop-shadow-[0_9px_0_rgba(5,34,70,0.68)]
            sm:[-webkit-text-stroke:3px_white]
          "
        >
          {titleLines.second}
        </span>
      {/if}
    </h1>

    {#if article.excerpt}
      <p
        class="
          mt-9
          max-w-3xl
          text-base
          leading-7
          font-semibold
          text-blue-950/65
          sm:text-lg
          sm:leading-8
        "
      >
        {article.excerpt}
      </p>
    {/if}

    <div
      class="
        mt-8
        flex flex-wrap
        items-center
        justify-center
        gap-x-4
        gap-y-4
      "
    >
      <div class="flex items-center gap-3">
        <img
          src="/logo/logo-kabinet.webp"
          alt="Kabinet Cerita Loka"
          class="
            size-12
            rounded-full
            border-2
            border-white
            bg-blue-900
            object-contain
            p-1
            shadow-md
          "
        />

        <div class="text-left">
          <p class="text-sm font-black text-blue-950">
            {article.author.name}
          </p>

          <p class="text-xs font-semibold text-blue-950/55">
            BEM Universitas Airlangga 2026
          </p>
        </div>
      </div>

      {#if article.published_at}
        <span
          aria-hidden="true"
          class="
            hidden
            h-9 w-px
            bg-blue-900/15
            sm:block
          "
        ></span>

        <div
          class="
            flex
            items-center
            gap-2
            text-xs
            font-bold
            text-blue-950/55
            sm:text-sm
          "
        >
          <time datetime={article.published_at}>
            {formatDate(article.published_at)}
          </time>

          <span aria-hidden="true">•</span>

          <span>{readTime} menit baca</span>
        </div>
      {/if}
    </div>
  </header>

  {#if article.cover_image}
    <div
      class="
        relative z-10
        mx-auto
        mt-14
        w-full
        max-w-7xl
        px-5
        sm:px-8
        lg:px-12
      "
    >
      <figure
        class="
          relative
          overflow-hidden
          rounded-[28px]
          border-[6px]
          border-white
          bg-blue-950
          shadow-[0_24px_60px_rgba(5,34,70,0.24)]
        "
      >
        <div
          aria-hidden="true"
          class="
            pointer-events-none
            absolute inset-0 z-10
            bg-gradient-to-t
            from-blue-950/20
            via-transparent
            to-transparent
          "
        ></div>

        <img
          src={article.cover_image}
          alt={`Cover ${article.title}`}
          loading="eager"
          fetchpriority="high"
          class="
            aspect-[16/8]
            w-full
            object-cover
            object-center
          "
        />
      </figure>
    </div>
  {/if}

  <main
    class="
      relative z-10
      mx-auto
      mt-16
      w-full
      max-w-7xl
      px-5
      sm:px-8
    "
  >
    <div
      aria-hidden="true"
      class="
        pointer-events-none
        absolute
        -left-20
        top-40
        size-72
        rounded-full
        bg-blue-300/15
        blur-[100px]
      "
    ></div>

    <ArticleBody body={article.body} />
  </main>

  <footer
    class="
      relative z-10
      mx-auto
      mt-16
      flex w-full
      max-w-4xl
      flex-col
      gap-6
      px-5
      sm:px-8
    "
  >
    <div
      class="
        h-px w-full
        bg-gradient-to-r
        from-transparent
        via-blue-900/20
        to-transparent
      "
    ></div>

    <div
      class="
        flex
        flex-col
        items-center
        justify-between
        gap-5
        text-center
        sm:flex-row
        sm:text-left
      "
    >
      <div>
        <p class="text-sm font-black text-blue-950">
          Kabinet Cerita Loka
        </p>

        <p class="mt-1 text-xs font-medium text-blue-950/50">
          Bersama merajut cerita untuk Airlangga dan Indonesia.
        </p>
      </div>

      <a
        href={resolve('/news')}
        class="
          group
          inline-flex
          items-center
          justify-center
          gap-2
          rounded-full
          bg-orange-500
          px-6 py-3
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
        <ArrowLeftIcon
          size={17}
          strokeWidth={2.5}
          class="transition-transform group-hover:-translate-x-1"
        />

        <span>Kembali ke News</span>
      </a>
    </div>
  </footer>
</article>
