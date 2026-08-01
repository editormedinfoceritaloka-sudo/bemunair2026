<script lang="ts">
  import { resolve } from '$app/paths';
  import PageContainer from '$lib/components/layout/PageContainer.svelte';
  import PublicFooter from '$lib/components/layout/PublicFooter.svelte';
  import LeadershipCard from '$lib/features/organization/components/LeadershipCard.svelte';
  import MinistryHero from '$lib/features/organization/components/MinistryHero.svelte';
  import WorkProgramCard from '$lib/features/work-program/components/WorkProgramCard.svelte';
  import type { OrganizationUnit, WorkProgram } from '$lib/types';

  let { data }: { data: { unit: OrganizationUnit; programs: WorkProgram[] } } = $props();
  const leaders = $derived((data.unit.members || []).filter((member) => member.position_type === 'MINISTER' || member.position_type === 'DIRECTOR_GENERAL'));
</script>

<svelte:head><title>{data.unit.name} · BEM UNAIR</title><meta name="description" content={data.unit.description || data.unit.name} /></svelte:head>

<main class="bg-[#f5f7fa] pt-28"><PageContainer><MinistryHero unit={data.unit} /><section class="py-20"><p class="text-xs font-bold uppercase tracking-[0.24em] text-orange-600">Pimpinan kementerian</p><h2 class="mt-3 text-4xl font-black text-blue-900">Menteri dan Dirjen</h2><div class="mt-8 grid gap-6 sm:grid-cols-2">{#each leaders as leader (leader.id)}<LeadershipCard member={leader} />{:else}<p class="rounded-2xl border border-dashed border-blue-200 p-8 text-slate-600">Profil pimpinan belum tersedia.</p>{/each}</div></section><section class="pb-20"><div class="flex flex-wrap items-end justify-between gap-4"><div><p class="text-xs font-bold uppercase tracking-[0.24em] text-orange-600">Karya kementerian</p><h2 class="mt-3 text-4xl font-black text-blue-900">Program kerja</h2></div><a href={resolve(`/kementerian/${data.unit.slug}/program-kerja`)} class="rounded-xl bg-blue-800 px-5 py-3 text-sm font-bold text-white">Lihat semua</a></div><div class="mt-8 grid gap-6 md:grid-cols-2 lg:grid-cols-4">{#each data.programs.slice(0, 8) as program (program.id)}<WorkProgramCard program={program} unitSlug={data.unit.slug} />{/each}</div></section></PageContainer><PublicFooter cabinet={null} /></main>
