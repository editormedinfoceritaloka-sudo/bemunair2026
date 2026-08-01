<script lang="ts">
  import AdminSidebar from '$lib/admin/AdminSidebar.svelte';
  import { page } from '$app/state';
  import { LoaderCircle } from '@lucide/svelte';

  let { data, children } = $props();
</script>

{#if page.url.pathname === '/admin/login'}
  {@render children()}
{:else if data.user}
  <div class="h-svh overflow-hidden bg-background md:flex">
    <AdminSidebar user={data.user} />
    <div class="flex min-h-0 min-w-0 flex-1 flex-col">
      <header class="flex h-14 shrink-0 items-center border-b border-border/80 bg-card px-4 md:px-6">
        <span class="text-sm font-medium text-muted-foreground">Admin Workspace</span>
      </header>
      <main id="admin-main" class="min-h-0 flex-1 overflow-y-auto p-4 md:p-7 xl:p-8">
        {@render children()}
      </main>
    </div>
  </div>
{:else}
  <div class="grid min-h-svh place-items-center bg-background p-6" role="status">
    <div class="flex items-center gap-3 rounded-xl border bg-card px-4 py-3 text-sm text-muted-foreground">
      <LoaderCircle class="size-4 animate-spin text-blue-600" />
      Memuat workspace admin...
    </div>
  </div>
{/if}
