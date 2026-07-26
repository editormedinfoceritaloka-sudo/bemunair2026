<script lang="ts">
  import { Inbox } from '@lucide/svelte';
  let { title='Belum ada data', description='Data akan tampil di sini setelah dibuat.', children }: { title?: string; description?: string; children?: import('svelte').Snippet } = $props();
</script>
<div class="grid min-h-52 place-items-center rounded-xl border border-dashed border-border bg-card p-8 text-center">
  <div><div class="mx-auto mb-3 grid size-11 place-items-center rounded-full bg-muted"><Inbox class="size-5 text-muted-foreground" /></div><h3 class="font-semibold text-foreground">{title}</h3><p class="mt-1 max-w-md text-sm text-muted-foreground">{description}</p>{#if children}<div class="mt-4 flex justify-center">{@render children()}</div>{/if}</div>
</div>
