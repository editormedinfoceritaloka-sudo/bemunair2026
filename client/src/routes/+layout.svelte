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

  const title = $derived(
    page.data.seo?.title ??
      'BEM Universitas Airlangga 2026 | Kabinet Cerita Loka'
  );

  const description = $derived(
    page.data.seo?.description ??
      'Website resmi BEM Universitas Airlangga 2026 Kabinet Cerita Loka.'
  );

  const image = $derived(
    page.data.seo?.image
      ? new URL(page.data.seo.image, page.url.origin).href
      : `${page.url.origin}/og-image.png`
  );

  const canonical = $derived(
    `${page.url.origin}${page.url.pathname}`
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

<svelte:head>
  <title>{title}</title>

  <meta
    name="description"
    content={description}
  />

  <meta
    name="robots"
    content={isAdminRoute ? 'noindex, nofollow' : 'index, follow'}
  />

  <meta
    name="theme-color"
    content="#1d4ed8"
  />

  {#if !isAdminRoute}
    <link
      rel="canonical"
      href={canonical}
    />

    <meta
      property="og:title"
      content={title}
    />

    <meta
      property="og:description"
      content={description}
    />

    <meta
      property="og:type"
      content="website"
    />

    <meta
      property="og:site_name"
      content="BEM Universitas Airlangga"
    />

    <meta
      property="og:locale"
      content="id_ID"
    />

    <meta
      property="og:url"
      content={canonical}
    />

    <meta
      property="og:image"
      content={image}
    />

    <meta
      name="twitter:card"
      content="summary_large_image"
    />

    <meta
      name="twitter:title"
      content={title}
    />

    <meta
      name="twitter:description"
      content={description}
    />

    <meta
      name="twitter:image"
      content={image}
    />
  {/if}
</svelte:head>

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