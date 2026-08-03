<script lang="ts">
	import './layout.css';
	import { Toaster } from '$lib/components/ui/sonner';
	import { page } from '$app/state';
  import Header from './Header.svelte';
  import Footer from './Footer.svelte';

	let { children } = $props();
</script>

<div class="app">
	{#if !page.url.pathname.startsWith("/admin")}
		<Header />
	{/if}
	<div class="app-content">{@render children()}</div>
	{#if !page.url.pathname.startsWith("/admin")}
		<Footer />
	{/if}
</div>
<Toaster richColors position="top-right" />