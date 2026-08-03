<script lang="ts">
  import { page } from '$app/state';

  import TimelinePelaksanaan from '$lib/components/program-kerja/TimelinePelaksanaan.svelte';
  import Documentation from '$lib/components/program-kerja/Documentation.svelte';

  const slug = $derived(
    page.params.slug
  );

  let title = $state(
    'Penyusunan Standart Operational Procedure'
  );

  let description = $state(
    'Setiap kementerian atau divisi pasti punya cara kerjanya masing-masing. Tanpa adanya standardisasi, ego sektoral mudah muncul, dan koordinasi antar-divisi sering kali terhambat karena perbedaan persepsi alur kerja. Maka dari itu, SOP merupakan salah satu langkah untuk menjaga keberlangsungan roda organisasi BEM UNAIR 2026 selama satu periode kedepan.'
  );

  let startDate = $state(
    '2026-01-15'
  );

  let endDate = $state<
    string | null
  >('2026-01-18');

  const testImages = [
    '/program-kerja/test-1.png',
    '/program-kerja/test-2.png',
    '/program-kerja/test-3.jpeg'
  ];

  const titleLines = $derived.by(() => {
    const words = title
      .trim()
      .split(/\s+/)
      .filter(Boolean);

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

    const splitIndex = Math.floor(
      words.length / 2
    );

    return {
      first: words
        .slice(0, splitIndex)
        .join(' '),

      second: words
        .slice(splitIndex)
        .join(' ')
    };
  });
</script>

<svelte:head>
  <title>{title}</title>

  <meta
    name="description"
    content={description}
  />

  <meta
    name="program-kerja-slug"
    content={slug}
  />
</svelte:head>

<section
  class="
    min-h-screen bg-blue-50
    pb-28 pt-14
    sm:pt-20
    lg:pb-36
    lg:pt-28
  "
>
  <div
    class="
      mx-auto flex w-full max-w-7xl
      flex-col items-center
      px-5 text-center
      sm:px-8
      lg:px-12
    "
  >
    <h1
      class="
        text-[clamp(1.9rem,5.3vw,4.7rem)]
        leading-[0.92] font-black
        tracking-[-0.065em]
      "
    >
      <span
        class="
          block text-blue-300
          [-webkit-text-stroke:2px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_10px_0_rgba(30,64,175,0.55)]
          sm:[-webkit-text-stroke:3px_white]
        "
      >
        {titleLines.first}
      </span>

      {#if titleLines.second}
        <span
          class="
            relative left-[0.55em]
            block text-blue-600
            [-webkit-text-stroke:2px_white]
            [paint-order:stroke_fill]
            drop-shadow-[0_10px_0_rgba(5,34,70,0.78)]
            sm:[-webkit-text-stroke:3px_white]
          "
        >
          {titleLines.second}
        </span>
      {/if}
    </h1>

    <p
      class="
        mt-14 max-w-4xl
        text-lg font-semibold
        text-blue-950/70
      "
    >
      {description}
    </p>

    <div class="mt-16 w-full">
      <TimelinePelaksanaan
        {startDate}
        {endDate}
      />
    </div>

    <div class="mt-20 w-full">
      <Documentation
        images={testImages}
      />
    </div>
  </div>
</section>