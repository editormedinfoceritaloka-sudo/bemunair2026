<script lang="ts">
  import { resolve } from '$app/paths';
  import type { Cabinet } from '$lib/types';
  import { publicNavigation } from '$lib/constants/public-navigation';

  let { cabinet }: { cabinet: Cabinet | null } = $props();
</script>

<header class="fixed z-50 flex w-full items-center justify-center px-4 py-4">
  <div class="flex w-full max-w-7xl items-center justify-between gap-4 rounded-2xl border border-white/60 bg-white/30 p-2 shadow-sm backdrop-blur-md">
    <a href={resolve('/kabinet')} class="flex min-w-0 items-center gap-3 rounded-xl px-3 py-2">
      <img src={cabinet?.logo?.url || '/logo/logo-bem.png'} alt={cabinet?.logo?.alt_text || 'Logo BEM UNAIR'} class="size-9 shrink-0 rounded-lg object-contain" />
      <span class="min-w-0">
        <span class="block truncate text-xs font-bold uppercase tracking-[0.18em] text-blue-800">BEM UNAIR</span>
        <span class="block truncate text-sm font-semibold text-blue-900">{cabinet?.name || 'Kabinet'}</span>
      </span>
    </a>

    <nav aria-label="Navigasi publik" class="hidden items-center gap-1 rounded-xl bg-white/35 p-1 text-sm font-semibold text-blue-800 md:flex">
      {#each publicNavigation as item (item.href)}
        <a class="rounded-lg px-4 py-2 transition hover:bg-white/70" href={resolve(item.href)}>{item.label}</a>
      {/each}
      {#each cabinet?.kemenkoan || [] as unit (unit.id)}
        <a class="rounded-lg px-4 py-2 transition hover:bg-white/70" href={resolve(`/kemenkoan/${unit.slug}`)}>{unit.short_name || unit.name}</a>
      {/each}
    </nav>

    <a href={resolve('/admin')} class="hidden rounded-xl bg-blue-700 px-4 py-2 text-sm font-bold text-white transition hover:bg-blue-800 sm:inline-flex">Admin</a>
  </div>
</header>
