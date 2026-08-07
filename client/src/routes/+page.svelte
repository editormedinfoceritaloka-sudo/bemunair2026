<script lang="ts">
  import { onMount } from 'svelte';

  import HomeSection from '$lib/components/Home/HomeSection.svelte';
  import SpeechSection from '$lib/components/Home/SpeechSection.svelte';
  import CabinetSection from '$lib/components/Home/CabinetSection.svelte';
  import CalendarSection from '$lib/components/Home/CalendarSection.svelte';
  import NewsSection from '$lib/components/Home/NewsSection.svelte';

  import type { PageData } from './$types';

  let {
    data
  }: {
    data: PageData;
  } = $props();

  $effect(() => {
    console.log('data', data);
  });

  onMount(() => {
    const calendarPrograms =
      data.cabinet?.kemenkoan
        ?.flatMap((kemenkoan) => [
          ...(kemenkoan.programs ?? []).map((program) => ({
            ...program,
            unit_name: kemenkoan.name,
            unit_slug: kemenkoan.slug
          })),
          ...(kemenkoan.children ?? []).flatMap((ministry) =>
            (ministry.programs ?? []).map((program) => ({
              ...program,
              unit_name: ministry.name,
              unit_slug: ministry.slug
            }))
          )
        ])
        .filter((program) => program.start_date) ?? [];

    console.log('Calendar programs:', calendarPrograms);
  });
</script>

<svelte:head>
  <title>BEM Universitas Airlangga 2026 | Kabinet Cerita Loka</title>

  <meta
    name="description"
    content="Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka. Temukan informasi kabinet, kementerian, program kerja, agenda, dan berita terbaru BEM UNAIR."
  />

  <meta
    name="keywords"
    content="BEM UNAIR, BEM Universitas Airlangga, BEM UNAIR 2026, Kabinet Cerita Loka, Universitas Airlangga, organisasi mahasiswa UNAIR"
  />

  <meta name="author" content="BEM Universitas Airlangga" />
  <meta name="robots" content="index, follow" />

  <meta
    property="og:title"
    content="BEM Universitas Airlangga 2026 | Kabinet Cerita Loka"
  />

  <meta
    property="og:description"
    content="Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka."
  />

  <meta property="og:type" content="website" />
  <meta property="og:locale" content="id_ID" />
  <meta property="og:image" content="/og-image.png" />

  <meta name="twitter:card" content="summary_large_image" />

  <meta
    name="twitter:title"
    content="BEM Universitas Airlangga 2026 | Kabinet Cerita Loka"
  />

  <meta
    name="twitter:description"
    content="Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka."
  />

  <meta name="twitter:image" content="/og-image.png" />
  <meta name="theme-color" content="#1d4ed8" />
</svelte:head>

<main>
  <HomeSection />
  <SpeechSection />
  <CabinetSection cabinet={data.cabinet} />
  <CalendarSection cabinet={data.cabinet} />
  <NewsSection />
</main>