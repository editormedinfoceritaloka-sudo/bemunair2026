<script lang="ts">
  import { resolve } from '$app/paths';
  import type { WorkProgram } from '$lib/types';
  let { program, unitSlug }: { program: WorkProgram; unitSlug: string } = $props();
</script>

<a href={resolve(`/kementerian/${unitSlug}/program-kerja/${program.slug}`)} class="group overflow-hidden rounded-2xl border border-blue-100 bg-white shadow-[0_12px_35px_rgba(16,56,100,0.08)] transition hover:-translate-y-1 hover:shadow-[0_18px_45px_rgba(16,56,100,0.16)]">
  <div class="aspect-[16/10] bg-gradient-to-br from-blue-50 to-blue-300">
    {#if program.cover?.url}<img src={program.cover.url} alt={program.cover.alt_text} class="size-full object-cover transition duration-500 group-hover:scale-105" />{:else}<div class="grid size-full place-items-center px-8 text-center text-2xl font-black text-blue-700">{program.name}</div>{/if}
  </div>
  <div class="p-5">
    <div class="flex items-center justify-between gap-3 text-xs font-bold uppercase tracking-wider text-orange-600"><span>{program.execution_month || 'Program kerja'}</span><span>{program.status}</span></div>
    <h3 class="mt-3 text-xl font-black text-blue-900">{program.name}</h3>
    {#if program.short_description}<p class="mt-2 line-clamp-3 text-sm leading-6 text-slate-600">{program.short_description}</p>{/if}
  </div>
</a>
