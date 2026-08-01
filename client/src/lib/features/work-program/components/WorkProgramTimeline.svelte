<script lang="ts">
  import type { Milestone } from '$lib/types';
  let { milestones = [] }: { milestones?: Milestone[] } = $props();
</script>

{#if milestones.length}
  <ol class="relative space-y-6 border-l-2 border-blue-200 pl-7">
    {#each milestones as milestone (milestone.id)}
      <li class="relative rounded-2xl border border-blue-100 bg-blue-50/70 p-5">
        <span class="absolute -left-[2.15rem] top-6 size-4 rounded-full border-4 border-white bg-orange-500"></span>
        <div class="flex flex-wrap items-center justify-between gap-3"><h3 class="font-black text-blue-900">{milestone.title}</h3><span class="text-xs font-bold uppercase tracking-wider text-blue-600">{milestone.status}</span></div>
        {#if milestone.description}<p class="mt-2 text-sm leading-6 text-slate-600">{milestone.description}</p>{/if}
      </li>
    {/each}
  </ol>
{:else}
  <p class="rounded-2xl border border-dashed border-blue-200 bg-blue-50 p-6 text-sm text-slate-600">Timeline program kerja belum tersedia.</p>
{/if}
