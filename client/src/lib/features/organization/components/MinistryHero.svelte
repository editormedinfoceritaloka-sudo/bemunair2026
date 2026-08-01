<script lang="ts">
  import type { OrganizationUnit } from '$lib/types';
  let { unit }: { unit: OrganizationUnit } = $props();
</script>

<section class="rounded-[28px] bg-gradient-to-br from-blue-900 via-blue-800 to-blue-500 p-6 text-white shadow-2xl sm:p-10">
  <div class="grid gap-8 md:grid-cols-[180px_1fr] md:items-center">
    <div class="grid aspect-square place-items-center overflow-hidden rounded-3xl border border-white/30 bg-white/90 p-5">
      {#if unit.logo?.url}<img src={unit.logo.url} alt={unit.logo.alt_text} class="size-full object-contain" />{:else}<span class="text-7xl font-black text-blue-700">{unit.name.slice(0, 1)}</span>{/if}
    </div>
    <div>
      <p class="text-xs font-bold uppercase tracking-[0.25em] text-orange-300">{unit.unit_type === 'KEMENKOAN' ? 'Kemenkoan' : 'Kementerian'}</p>
      <h1 class="mt-3 text-4xl font-black tracking-tight sm:text-6xl">{unit.name}</h1>
      {#if unit.description}<p class="mt-5 max-w-2xl text-base leading-8 text-blue-50">{unit.description}</p>{/if}
    </div>
  </div>
</section>
