<script lang="ts">
  import AdminSidebar from '$lib/admin/AdminSidebar.svelte';
  import * as Sidebar from '$lib/components/ui/sidebar';
  import * as Tooltip from '$lib/components/ui/tooltip';
  let { data, children } = $props();
</script>
{#if data.user}
  <Tooltip.Provider>
    <Sidebar.Provider>
      <AdminSidebar user={data.user} />
      <Sidebar.Inset class="admin-workspace min-w-0 bg-background">
        <header class="sticky top-0 z-20 flex h-14 items-center gap-3 border-b border-border/80 bg-card/90 px-4 backdrop-blur md:px-6">
          <Sidebar.Trigger />
          <div class="h-5 w-px bg-border"></div>
          <span class="text-sm font-medium text-muted-foreground">Admin Workspace</span>
        </header>
        <main class="w-full p-4 md:p-7 xl:p-8">{@render children()}</main>
      </Sidebar.Inset>
    </Sidebar.Provider>
  </Tooltip.Provider>
{:else}
  {@render children()}
{/if}
