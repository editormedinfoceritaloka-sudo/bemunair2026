<script lang="ts">
  import PageContainer from '$lib/components/layout/PageContainer.svelte';
  import PublicFooter from '$lib/components/layout/PublicFooter.svelte';
  import SectionHeading from '$lib/components/layout/SectionHeading.svelte';
  import OrganizationCarousel from '$lib/features/organization/components/OrganizationCarousel.svelte';
  import type { Cabinet } from '$lib/types';

  let { data }: { data: { cabinet: Cabinet } } = $props();
</script>

<svelte:head>
  <title>{data.cabinet.meta_title || data.cabinet.name}</title>
  <meta name="description" content={data.cabinet.meta_description || data.cabinet.tagline || 'Profil kabinet BEM UNAIR'} />
</svelte:head>

<main class="overflow-hidden bg-[#f5f7fa] pt-24">
  <section class="relative bg-gradient-to-br from-blue-50 via-white to-blue-200 py-20 sm:py-28">
    <PageContainer>
      <div class="grid items-center gap-12 lg:grid-cols-[1.1fr_0.9fr]">
        <div>
          <p class="text-xs font-bold uppercase tracking-[0.28em] text-orange-600">BEM UNAIR 2026</p>
          <h1 class="mt-5 max-w-3xl text-5xl font-black tracking-tight text-blue-900 sm:text-7xl">{data.cabinet.name}</h1>
          <p class="mt-6 max-w-2xl text-lg leading-8 text-slate-600">{data.cabinet.tagline || data.cabinet.description || 'Mengenal struktur, cerita, dan program kerja kabinet.'}</p>
          <a href="#kemenkoan" class="mt-8 inline-flex rounded-xl bg-orange-500 px-6 py-3 font-bold text-white shadow-lg transition hover:bg-orange-600">Jelajahi kabinet</a>
        </div>
        <div class="relative mx-auto grid aspect-square w-full max-w-md place-items-center rounded-[32px] bg-gradient-to-br from-blue-800 to-blue-500 p-10 shadow-2xl">
          {#if data.cabinet.logo?.url}<img src={data.cabinet.logo.url} alt={data.cabinet.logo.alt_text} class="size-full object-contain" />{:else}<span class="text-center text-8xl font-black text-white">CL</span>{/if}
        </div>
      </div>
    </PageContainer>
  </section>

  <section id="kemenkoan" class="py-20 sm:py-28">
    <PageContainer>
      <SectionHeading eyebrow="Struktur kabinet" title="Kemenkoan" description="Kenali kemenkoan dan kementerian yang menggerakkan kabinet." />
      {#if data.cabinet.kemenkoan?.length}<OrganizationCarousel units={data.cabinet.kemenkoan} />{:else}<p class="mt-10 rounded-2xl border border-dashed border-blue-200 p-8 text-slate-600">Data kemenkoan belum tersedia.</p>{/if}
    </PageContainer>
  </section>

  <PublicFooter cabinet={data.cabinet} />
</main>
