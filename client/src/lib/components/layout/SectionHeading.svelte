<script lang="ts">
  let { eyebrow, title, description = '' } = $props<{ eyebrow?: string; title: string; description?: string }>();
</script>

<div class="max-w-3xl">
  {#if eyebrow}<p class="text-xs font-bold uppercase tracking-[0.24em] text-orange-600">{eyebrow}</p>{/if}
  <h1 class="mt-2 text-4xl font-black tracking-tight text-blue-900 sm:text-6xl">{title}</h1>
  {#if description}<p class="mt-4 max-w-2xl text-base leading-8 text-slate-600">{description}</p>{/if}
</div>
