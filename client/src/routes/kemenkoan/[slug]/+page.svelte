<script lang="ts">
  import PageContainer from '$lib/components/layout/PageContainer.svelte';
  import PublicFooter from '$lib/components/layout/PublicFooter.svelte';
  import MinistryHero from '$lib/features/organization/components/MinistryHero.svelte';
  import OrganizationCard from '$lib/features/organization/components/OrganizationCard.svelte';
  import type { OrganizationUnit, Cabinet } from '$lib/types';

  let { data }: { data: { unit: OrganizationUnit; cabinet?: Cabinet | null } } = $props();
</script>

<svelte:head><title>{data.unit.name} · BEM UNAIR</title><meta name="description" content={data.unit.description || data.unit.name} /></svelte:head>

<main class="bg-[#f5f7fa] pt-28"><PageContainer><MinistryHero unit={data.unit} /><section class="py-20"><p class="text-xs font-bold uppercase tracking-[0.24em] text-orange-600">Struktur di bawah koordinasi</p><h2 class="mt-3 text-4xl font-black text-blue-900">Kementerian</h2><div class="mt-8 grid gap-6 md:grid-cols-2 lg:grid-cols-3">{#each data.unit.children || [] as child (child.id)}<OrganizationCard unit={child} />{:else}<p class="rounded-2xl border border-dashed border-blue-200 p-8 text-slate-600">Belum ada kementerian.</p>{/each}</div></section></PageContainer><PublicFooter cabinet={data.cabinet || null} /></main>
