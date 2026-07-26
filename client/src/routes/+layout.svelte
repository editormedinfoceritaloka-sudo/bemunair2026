<script lang="ts">
	import './layout.css';
	import { Toaster } from '$lib/components/ui/sonner';

	let { children } = $props();
</script>

<div class="app">
	<main>{@render children()}</main>
</div>
<Toaster richColors position="top-right" />