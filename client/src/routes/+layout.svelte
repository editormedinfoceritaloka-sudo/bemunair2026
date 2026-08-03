<script lang="ts">
  import './layout.css';

  import { page } from '$app/state';
  import { Toaster } from '$lib/components/ui/sonner';

  import Header from './Header.svelte';
  import Footer from './Footer.svelte';

  let { children } = $props();

  const isAdminRoute = $derived(
    page.url.pathname === '/admin' ||
      page.url.pathname.startsWith('/admin/')
  );

  $effect(() => {
    if (typeof window === 'undefined') {
      return;
    }

    const currentRoute = `${page.url.pathname}${page.url.search}${page.url.hash}`;

    if (currentRoute.startsWith('/admin')) {
      return;
    }

    const storedCurrentRoute = sessionStorage.getItem(
      'current-public-route'
    );

    if (
      storedCurrentRoute &&
      storedCurrentRoute !== currentRoute &&
      !storedCurrentRoute.startsWith('/admin')
    ) {
      sessionStorage.setItem(
        'last-public-route',
        storedCurrentRoute
      );
    }

    sessionStorage.setItem(
      'current-public-route',
      currentRoute
    );
  });
</script>

<div class="app">
  {#if !isAdminRoute}
    <Header />
  {/if}

  <main class="app-content">
    {@render children()}
  </main>

  {#if !isAdminRoute}
    <Footer />
  {/if}
</div>

<Toaster richColors position="top-right" />