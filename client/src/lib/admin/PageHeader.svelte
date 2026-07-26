<script lang="ts">
  let { title, description, children }: { title: string; description?: string; children?: import('svelte').Snippet } = $props();
</script>
<div class="mb-6 flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
  <div><h1 class="text-left text-2xl font-bold tracking-tight text-foreground">{title}</h1>{#if description}<p class="mt-1 text-sm text-muted-foreground">{description}</p>{/if}</div>
  {#if children}<div class="flex flex-wrap gap-2">{@render children()}</div>{/if}
</div>
