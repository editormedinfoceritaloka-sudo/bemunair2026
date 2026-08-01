<script lang="ts">
  import { resolve } from '$app/paths';
  import PageContainer from '$lib/components/layout/PageContainer.svelte';
  import ProgramPagination from '$lib/features/work-program/components/ProgramPagination.svelte';
  import WorkProgramCard from '$lib/features/work-program/components/WorkProgramCard.svelte';
  import type { Meta, OrganizationUnit, WorkProgram } from '$lib/types';

  let { data }: { data: { unit: OrganizationUnit; programs: WorkProgram[]; meta: Meta } } = $props();
</script>

<svelte:head><title>Program Kerja {data.unit.name} · BEM UNAIR</title></svelte:head>

<main class="min-h-screen bg-[#f5f7fa] pb-20 pt-32"><PageContainer><a href={resolve(`/kementerian/${data.unit.slug}`)} class="text-sm font-bold text-blue-700">← Kembali ke kementerian</a><h1 class="mt-6 text-5xl font-black text-blue-900">Program Kerja</h1><p class="mt-3 text-slate-600">{data.unit.name}</p><div class="mt-10 grid gap-6 md:grid-cols-2 lg:grid-cols-4">{#each data.programs as program (program.id)}<WorkProgramCard program={program} unitSlug={data.unit.slug} />{:else}<p class="rounded-2xl border border-dashed border-blue-200 p-8 text-slate-600">Belum ada program kerja.</p>{/each}</div><ProgramPagination page={data.meta?.page || 1} totalPages={data.meta?.total_pages || 0} basePath={`/kementerian/${data.unit.slug}/program-kerja`} /></PageContainer></main>
