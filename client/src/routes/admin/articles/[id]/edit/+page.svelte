<script lang="ts">import ArticleComposer from '$lib/admin/editor/ArticleComposer.svelte';let {data}=$props();</script>
<svelte:head><title>{data.article.title} · Editor</title></svelte:head>
<ArticleComposer article={data.article}/>
