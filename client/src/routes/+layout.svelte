<script lang="ts">
	import './layout.css';
	import { Toaster } from '$lib/components/ui/sonner';
	import { page } from '$app/state';
  import Header from './Header.svelte';

	let { children, data } = $props();
</script>

<div class="app">
	{#if !page.url.pathname.startsWith("/admin")}
		<Header cabinet={data.cabinet} />
	{/if}
	<div class="app-content">{@render children()}</div>
</div>
<Toaster richColors position="top-right" />