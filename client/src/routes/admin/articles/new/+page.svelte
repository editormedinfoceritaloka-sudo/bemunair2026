<script lang="ts">import ArticleComposer from '$lib/admin/editor/ArticleComposer.svelte';let {form}=$props();</script>
<svelte:head><title>Artikel baru · Admin BEM UNAIR</title></svelte:head>
<ArticleComposer error={form?.error||''}/>
