<script lang="ts">
  import { resolve } from '$app/paths';
  import type { Cabinet } from '$lib/types';
  import { publicNavigation } from '$lib/constants/public-navigation';

  let { cabinet }: { cabinet: Cabinet | null } = $props();
</script>

<footer class="mt-24 bg-gradient-to-b from-blue-800 via-blue-700 to-blue-500 px-5 py-14 text-white sm:px-8 lg:px-10">
  <div class="mx-auto grid max-w-7xl gap-10 md:grid-cols-[1.2fr_0.8fr_1fr]">
    <div>
      <p class="text-xs font-bold uppercase tracking-[0.24em] text-blue-100">BEM UNAIR</p>
      <h2 class="mt-3 text-3xl font-black">{cabinet?.name || 'Kabinet'}</h2>
      <p class="mt-3 max-w-sm text-sm leading-7 text-blue-50">{cabinet?.tagline || 'Ruang bertumbuh, berbagi cerita, dan menghadirkan dampak.'}</p>
    </div>
    <nav aria-label="Navigasi footer" class="grid content-start gap-3 text-sm text-blue-50">
      {#each publicNavigation as item (item.href)}
        <a class="transition hover:text-white" href={resolve(item.href)}>{item.label}</a>
      {/each}
    </nav>
    <div class="text-sm text-blue-50">
      <p class="font-semibold text-white">Kontak</p>
      <p class="mt-3 leading-7">Universitas Airlangga<br />Surabaya, Jawa Timur</p>
    </div>
  </div>
  <div class="mx-auto mt-12 max-w-7xl border-t border-white/20 pt-5 text-xs text-blue-100">© 2026 BEM UNAIR</div>
</footer>
