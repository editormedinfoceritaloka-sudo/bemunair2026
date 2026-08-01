<script lang="ts">
  import { onMount } from 'svelte';
  import { ChevronLeft, ChevronRight } from '@lucide/svelte';
  import { resolve } from '$app/paths';
  import type { OrganizationUnit } from '$lib/types';
  import { useGsapCarousel } from '$lib/hooks/use-gsap-carousel.svelte';
  import { useReducedMotion } from '$lib/hooks/use-reduced-motion.svelte';

  let { units }: { units: OrganizationUnit[] } = $props();
  let root: HTMLDivElement;
  let active = $state(0);
  const motion = useReducedMotion();
  const carousel = useGsapCarousel();

  function move(direction: number) {
    if (!units.length) return;
    active = (active + direction + units.length) % units.length;
    carousel.animate(root, active, units.length, motion.value);
  }

  function handleKeydown(event: KeyboardEvent) {
    if (!root.contains(document.activeElement)) return;
    if (event.key === 'ArrowRight') move(1);
    if (event.key === 'ArrowLeft') move(-1);
  }

  onMount(() => {
    carousel.animate(root, active, units.length, motion.value);
    return () => carousel.destroy();
  });
</script>

<svelte:window onkeydown={handleKeydown} />

<div class="relative overflow-hidden px-2 py-8 sm:px-10" bind:this={root} role="region" aria-label="Kumpulan Kemenkoan">
  <div class="relative mx-auto h-[330px] max-w-5xl">
    {#each units as unit, index (unit.id)}
      <a data-carousel-card href={resolve(`/kemenkoan/${unit.slug}`)} aria-label={`Buka ${unit.name}`} class="absolute left-1/2 top-1/2 block w-[min(78vw,390px)] -translate-x-1/2 -translate-y-1/2 rounded-[22px] border-2 border-blue-600 bg-gradient-to-br from-blue-50 via-blue-200 to-blue-500 p-6 shadow-2xl">
        <div class="grid aspect-[1.15] place-items-center rounded-2xl border border-white/70 bg-white/70 p-8">
          {#if unit.logo?.url}<img src={unit.logo.url} alt={unit.logo.alt_text} class="size-32 object-contain" />{:else}<span class="text-7xl font-black text-blue-700">{unit.name.slice(0, 1)}</span>{/if}
        </div>
        <h3 class="mt-5 text-center text-xl font-black text-blue-900">{unit.name}</h3>
      </a>
      <span class="sr-only">{index === active ? 'Aktif' : ''}</span>
    {/each}
  </div>
  <button type="button" aria-label="Kartu sebelumnya" onclick={() => move(-1)} class="absolute left-1 top-1/2 grid size-11 -translate-y-1/2 place-items-center rounded-full bg-blue-800 text-white shadow-lg transition hover:bg-blue-700 sm:left-4"><ChevronLeft /></button>
  <button type="button" aria-label="Kartu berikutnya" onclick={() => move(1)} class="absolute right-1 top-1/2 grid size-11 -translate-y-1/2 place-items-center rounded-full bg-blue-800 text-white shadow-lg transition hover:bg-blue-700 sm:right-4"><ChevronRight /></button>
</div>
