<script lang="ts">
  import AdminSidebar from '$lib/admin/AdminSidebar.svelte';
  import { page } from '$app/state';
  import { LoaderCircle, Menu, X } from '@lucide/svelte';

  let { data, children } = $props();
  let mobileNavOpen = $state(false);

  $effect(() => {
    void page.url.pathname;
    mobileNavOpen = false;
  });
</script>

<a href="#admin-main" class="sr-only focus:not-sr-only focus:fixed focus:left-4 focus:top-4 focus:z-50 focus:rounded-md focus:bg-blue-600 focus:px-3 focus:py-2 focus:text-white">Skip to content</a>

{#if page.url.pathname === '/admin/login'}
  {@render children()}
{:else if data.user}
  <div class="h-svh overflow-hidden bg-background md:flex">
    <AdminSidebar user={data.user} ministryLogo={data.ministryLogo} />
    {#if mobileNavOpen}
      <div class="fixed inset-0 z-50 md:hidden">
        <button type="button" class="absolute inset-0 bg-black/40" onclick={() => mobileNavOpen = false} aria-label="Tutup navigasi admin"></button>
        <div class="relative h-full w-72">
          <AdminSidebar user={data.user} ministryLogo={data.ministryLogo} mobile />
          <button type="button" class="absolute right-3 top-3 grid size-9 place-items-center rounded-lg border bg-card" onclick={() => mobileNavOpen = false} aria-label="Tutup navigasi admin">
            <X class="size-5" />
          </button>
        </div>
      </div>
    {/if}
    <div class="flex min-h-0 min-w-0 flex-1 flex-col">
      <header class="flex h-14 shrink-0 items-center gap-3 border-b border-border/80 bg-card px-4 md:px-6">
        <button type="button" class="grid size-9 place-items-center rounded-lg border md:hidden" onclick={() => mobileNavOpen = true} aria-label="Buka navigasi admin" aria-expanded={mobileNavOpen}>
          <Menu class="size-5" />
        </button>
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
