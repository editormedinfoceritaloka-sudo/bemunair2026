<script lang="ts">
  import { resolve } from '$app/paths';
  import type { OrganizationUnit } from '$lib/types';
  let { unit }: { unit: OrganizationUnit } = $props();
</script>

<a href={resolve(unit.unit_type === 'KEMENKOAN' ? `/kemenkoan/${unit.slug}` : `/kementerian/${unit.slug}`)} class="group block rounded-2xl border border-blue-200 bg-gradient-to-br from-blue-50 to-blue-300 p-5 shadow-[0_18px_50px_rgba(16,56,100,0.18)] transition hover:-translate-y-1 hover:shadow-[0_24px_60px_rgba(16,56,100,0.25)]">
  <div class="flex items-center justify-between gap-4">
    <div class="grid size-14 shrink-0 place-items-center overflow-hidden rounded-xl border border-white/80 bg-white p-2">
      {#if unit.logo?.url}<img src={unit.logo.url} alt={unit.logo.alt_text} class="size-full object-contain" />{:else}<span class="text-xl font-black text-blue-700">{unit.name.slice(0, 1)}</span>{/if}
    </div>
    <span class="rounded-full bg-white/70 px-3 py-1 text-[10px] font-bold uppercase tracking-wider text-blue-800">{unit.unit_type === 'KEMENKOAN' ? 'Kemenkoan' : 'Kementerian'}</span>
  </div>
  <h3 class="mt-8 text-xl font-black text-blue-900 transition group-hover:text-orange-700">{unit.name}</h3>
  {#if unit.description}<p class="mt-2 line-clamp-3 text-sm leading-6 text-blue-900/75">{unit.description}</p>{/if}
</a>
