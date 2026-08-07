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

<main>
  <HomeSection />
  <SpeechSection />
  <CabinetSection cabinet={data.cabinet} />
  <CalendarSection cabinet={data.cabinet} />
  <NewsSection />
</main>