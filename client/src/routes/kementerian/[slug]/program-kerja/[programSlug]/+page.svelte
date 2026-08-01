<script lang="ts">
  import { resolve } from '$app/paths';
  import PageContainer from '$lib/components/layout/PageContainer.svelte';
  import DocumentationGallery from '$lib/features/work-program/components/DocumentationGallery.svelte';
  import WorkProgramTimeline from '$lib/features/work-program/components/WorkProgramTimeline.svelte';
  import type { WorkProgram } from '$lib/types';

  let { data }: { data: { program: WorkProgram } } = $props();
</script>

<svelte:head><title>{data.program.name} · BEM UNAIR</title><meta name="description" content={data.program.short_description || data.program.name} /></svelte:head>

<main class="min-h-screen bg-[#f5f7fa] pb-24 pt-32"><PageContainer><a href={resolve('/kabinet')} class="text-sm font-bold text-blue-700">← Kembali ke kabinet</a><div class="mt-6 grid gap-10 lg:grid-cols-[1.1fr_0.9fr]"><div><p class="text-xs font-bold uppercase tracking-[0.24em] text-orange-600">{data.program.status}</p><h1 class="mt-3 text-5xl font-black tracking-tight text-blue-900">{data.program.name}</h1><p class="mt-5 text-lg leading-8 text-slate-600">{data.program.description || data.program.short_description}</p></div><div class="overflow-hidden rounded-3xl bg-blue-100">{#if data.program.cover?.url}<img src={data.program.cover.url} alt={data.program.cover.alt_text} class="aspect-[4/3] size-full object-cover" />{:else}<div class="grid aspect-[4/3] place-items-center p-8 text-center text-3xl font-black text-blue-700">{data.program.name}</div>{/if}</div></div><section class="mt-16 grid gap-12 lg:grid-cols-2"><div><h2 class="text-3xl font-black text-blue-900">Timeline</h2><div class="mt-6"><WorkProgramTimeline milestones={data.program.milestones} /></div></div><div><h2 class="text-3xl font-black text-blue-900">Dokumentasi</h2><div class="mt-6"><DocumentationGallery items={data.program.documentations} /></div></div></section></PageContainer></main>
