<script lang="ts">
  import HomeSection from '$lib/components/Home/HomeSection.svelte';
  import SpeechSection from '$lib/components/Home/SpeechSection.svelte';
  import CabinetSection from '$lib/components/Home/CabinetSection.svelte';
  import CalendarSection from '$lib/components/Home/CalendarSection.svelte';
  import NewsSection from '$lib/components/Home/NewsSection.svelte';

  import type {PageData} from './$types';

  export let data: PageData;

</script>

<main>
  <HomeSection />
  <SpeechSection />
  <CabinetSection cabinet={data.cabinet} />
  <CalendarSection />
  <NewsSection />
</main>