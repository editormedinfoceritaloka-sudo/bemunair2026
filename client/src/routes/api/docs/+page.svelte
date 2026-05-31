<script lang="ts">
  import { BookOpen, Braces, Check, ChevronRight, Copy, Menu, Search, Server, X } from '@lucide/svelte';
  import { onMount, tick } from 'svelte';

  type DocItem = { slug: string; title: string; order: number; group?: string };
  type Heading = { id: string; text: string; level: number };

  const fallbackDocs: DocItem[] = [
    { slug: 'overview', title: 'Overview', order: 1, group: 'Start' },
    { slug: 'auth', title: 'Authentication', order: 2, group: 'Core API' },
    { slug: 'users', title: 'Users', order: 3, group: 'Core API' },
    { slug: 'content-submissions', title: 'Content Submissions', order: 4, group: 'Submissions' },
    { slug: 'letter-submissions', title: 'Letter Submissions', order: 5, group: 'Submissions' },
    { slug: 'medinfo-pj-queue', title: 'Medinfo PJ Queue', order: 6, group: 'Admin' },
    { slug: 'letter-templates', title: 'Letter Templates', order: 7, group: 'Admin' },
    { slug: 'wa-engine', title: 'WA Engine', order: 8, group: 'Integration' }
  ];

  let docs: DocItem[] = fallbackDocs;
  let active = 'overview';
  let markdown = '';
  let query = '';
  let loading = true;
  let error = '';
  let drawerOpen = false;
  let copied = '';
  let headings: Heading[] = [];

  const apiBase = import.meta.env.VITE_API_BASE_URL ?? '/api';

  $: activeDoc = docs.find((item) => item.slug === active);
  $: filtered = docs.filter((item) => item.title.toLowerCase().includes(query.toLowerCase()));
  $: grouped = groupDocs(filtered);
  $: endpointCount = (markdown.match(/^##\s+(GET|POST|PUT|DELETE|PATCH)\s+/gm) ?? []).length;
  $: methodCounts = {
    GET: countMethod('GET'),
    POST: countMethod('POST'),
    PUT: countMethod('PUT'),
    DELETE: countMethod('DELETE')
  };
  $: html = renderMarkdown(markdown);

  async function loadIndex() {
    loading = true;
    try {
      await loadDoc(active);
    } catch {
      error = 'Dokumentasi tidak dapat dimuat.';
    } finally {
      loading = false;
    }
  }

  async function loadDoc(slug: string) {
    active = slug;
    drawerOpen = false;
    error = '';
    markdown = '';
    const res = await fetch(`${apiBase.replace(/\/$/, '')}/docs/${slug}`);
    if (!res.ok) {
      error = 'Modul dokumentasi tidak ditemukan.';
      return;
    }
    markdown = await res.text();
    headings = extractHeadings(markdown);
    await tick();
    document.querySelector('.content-scroll')?.scrollTo({ top: 0 });
  }

  function groupDocs(items: DocItem[]) {
    return items.reduce<Record<string, DocItem[]>>((acc, item) => {
      const group = item.group ?? 'API';
      acc[group] = [...(acc[group] ?? []), item];
      return acc;
    }, {});
  }

  function countMethod(method: string) {
    return (markdown.match(new RegExp(`^##\\s+${method}\\s+`, 'gm')) ?? []).length;
  }

  function extractHeadings(md: string): Heading[] {
    return [...md.matchAll(/^(#{1,3})\s+(.+)$/gm)].map((match) => {
      const text = match[2].replace(/`/g, '').trim();
      return { id: slugify(text), text, level: match[1].length };
    });
  }

  function renderMarkdown(md: string) {
    const codeBlocks: string[] = [];
    const withoutCode = md.replace(/```(\w+)?\n([\s\S]*?)```/g, (_, lang, code) => {
      const id = codeBlocks.length;
      codeBlocks.push(`<div class="code-wrap"><button class="copy-code" data-copy="${encodeURIComponent(code)}"><span>Copy</span></button><pre><code class="language-${lang ?? ''}">${escapeHtml(code)}</code></pre></div>`);
      return `@@CODE_${id}@@`;
    });

    let out = escapeHtml(withoutCode)
      .replace(/^---$/gm, '<hr>')
      .replace(/^### (.*)$/gm, (_, text) => `<h3 id="${slugify(text)}"><a href="#${slugify(text)}">${decorateHeading(text)}</a></h3>`)
      .replace(/^## (GET|POST|PUT|DELETE|PATCH) (.*)$/gm, (_, method, path) => `<h2 id="${slugify(`${method} ${path}`)}" class="endpoint-heading"><span class="method method-${method}">${method}</span><code>${path}</code></h2>`)
      .replace(/^## (.*)$/gm, (_, text) => `<h2 id="${slugify(text)}"><a href="#${slugify(text)}">${decorateHeading(text)}</a></h2>`)
      .replace(/^# (.*)$/gm, (_, text) => `<h1 id="${slugify(text)}"><a href="#${slugify(text)}">${decorateHeading(text)}</a></h1>`)
      .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
      .replace(/`([^`]+)`/g, '<code>$1</code>')
      .replace(/\n\|(.+)\|\n\|[-:\s|]+\|\n((?:\|.*\|\n?)+)/g, tableToHtml)
      .replace(/^- (.*)$/gm, '<li>$1</li>')
      .replace(/(<li>.*<\/li>)/gs, '<ul>$1</ul>')
      .replace(/\n\n/g, '</p><p>')
      .replace(/^/, '<p>')
      .replace(/$/, '</p>')
      .replace(/<p>\s*(<h[1-3][\s\S]*?<\/h[1-3]>)\s*<\/p>/g, '$1')
      .replace(/<p>\s*(<div class="code-wrap"[\s\S]*?<\/div>)\s*<\/p>/g, '$1')
      .replace(/<p>\s*(<table[\s\S]*?<\/table>)\s*<\/p>/g, '$1')
      .replace(/<p>\s*(<hr>)\s*<\/p>/g, '$1')
      .replace(/<p>\s*(<ul>[\s\S]*?<\/ul>)\s*<\/p>/g, '$1');

    codeBlocks.forEach((block, id) => {
      out = out.replace(`@@CODE_${id}@@`, block);
    });
    return out;
  }

  function decorateHeading(text: string) {
    return text.replace(/^(Auth|Content-Type|Query|Path|Request|Response|Error|Curl):/i, '<span class="label">$1</span>:');
  }

  function tableToHtml(match: string, header: string, rows: string) {
    const th = header.split('|').filter(Boolean).map((cell) => `<th>${cell.trim()}</th>`).join('');
    const tr = rows.trim().split('\n').map((row) => `<tr>${row.split('|').filter(Boolean).map((cell) => `<td>${cell.trim()}</td>`).join('')}</tr>`).join('');
    return `<table><thead><tr>${th}</tr></thead><tbody>${tr}</tbody></table>`;
  }

  function slugify(text: string) {
    return text.toLowerCase().replace(/`/g, '').replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)/g, '');
  }

  function escapeHtml(value: string) {
    return value.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
  }

  function handleClick(event: MouseEvent) {
    const target = event.target as HTMLElement;
    const button = target.closest<HTMLButtonElement>('button[data-copy]');
    if (button?.dataset.copy) {
      const value = decodeURIComponent(button.dataset.copy);
      navigator.clipboard.writeText(value);
      copied = value.slice(0, 18);
      setTimeout(() => (copied = ''), 1400);
    }
  }

  loadIndex();
  onMount(() => {
    document.addEventListener('click', handleClick);
    return () => document.removeEventListener('click', handleClick);
  });
</script>

<svelte:head><title>API Docs - BEM UNAIR</title></svelte:head>

<main class="docs-shell">
  <aside class:open={drawerOpen}>
    <div class="brand">
      <div>
        <strong>BEM UNAIR</strong>
        <span>API Reference</span>
      </div>
      <button aria-label="Close navigation" on:click={() => (drawerOpen = false)}><X size={18} /></button>
    </div>

    <label class="search">
      <Search size={16} />
      <input bind:value={query} placeholder="Cari modul" />
    </label>

    <nav>
      {#each Object.entries(grouped) as [group, items]}
        <p>{group}</p>
        {#each items as item}
          <button class:active={item.slug === active} on:click={() => loadDoc(item.slug)}>
            <BookOpen size={15} />
            <span>{item.title}</span>
            <ChevronRight size={14} />
          </button>
        {/each}
      {/each}
    </nav>
  </aside>

  <section class="content-shell">
    <header>
      <button class="menu" aria-label="Open navigation" on:click={() => (drawerOpen = true)}><Menu size={20} /></button>
      <div>
        <span>Digital Submission System</span>
        <strong>{activeDoc?.title ?? 'API Docs'}</strong>
      </div>
      <div class="status"><Server size={16} /> <span>Base URL</span> <code>/api</code></div>
    </header>

    <div class="content-scroll">
      {#if loading}
        <div class="loading-grid"><div></div><div></div><div></div></div>
      {:else if error}
        <div class="empty">{error}</div>
      {:else}
        <div class="summary-row">
          <div><Braces size={18} /><span>{endpointCount}</span><small>Endpoints</small></div>
          <div><span>{methodCounts.GET}</span><small>GET</small></div>
          <div><span>{methodCounts.POST}</span><small>POST</small></div>
          <div><span>{methodCounts.PUT}</span><small>PUT</small></div>
          <div><span>{methodCounts.DELETE}</span><small>DELETE</small></div>
        </div>

        <div class="layout">
          <article>{@html html}</article>
          <aside class="toc">
            <strong>On this page</strong>
            {#each headings.slice(0, 18) as heading}
              <a class:sub={heading.level > 2} href={`#${heading.id}`}>{heading.text}</a>
            {/each}
          </aside>
        </div>
      {/if}
    </div>
  </section>

  {#if copied}
    <div class="toast"><Check size={16} /> Copied</div>
  {/if}
</main>

<style>
  :global(body) { margin: 0; font-family: Inter, ui-sans-serif, system-ui, sans-serif; background: #f5f7fa; color: #17202a; }
  .docs-shell { min-height: 100vh; display: grid; grid-template-columns: 300px minmax(0, 1fr); }
  aside { background: #fff; border-right: 1px solid #dbe1ea; padding: 18px; box-sizing: border-box; }
  .docs-shell > aside { position: sticky; top: 0; height: 100vh; overflow: auto; }
  .brand { display: flex; justify-content: space-between; gap: 12px; align-items: flex-start; margin-bottom: 18px; }
  .brand strong { display: block; font-size: 15px; }
  .brand span { color: #667789; font-size: 13px; }
  button { font: inherit; }
  .brand button, .menu { display: none; border: 1px solid #dbe1ea; background: #fff; border-radius: 7px; width: 36px; height: 36px; align-items: center; justify-content: center; }
  .search { display: flex; align-items: center; gap: 9px; border: 1px solid #dbe1ea; border-radius: 8px; padding: 10px 11px; margin-bottom: 18px; color: #667789; }
  .search input { border: 0; outline: 0; width: 100%; font: inherit; min-width: 0; }
  nav p { margin: 18px 6px 7px; color: #8592a3; font-size: 11px; text-transform: uppercase; font-weight: 700; letter-spacing: .08em; }
  nav button { width: 100%; display: grid; grid-template-columns: 18px 1fr 14px; align-items: center; gap: 9px; border: 0; background: transparent; color: #4b5b6d; border-radius: 8px; padding: 10px; text-align: left; }
  nav button.active, nav button:hover { background: #edf4f8; color: #102337; }
  .content-shell { min-width: 0; }
  header { height: 68px; display: flex; align-items: center; justify-content: space-between; gap: 16px; border-bottom: 1px solid #dbe1ea; background: rgba(255,255,255,.9); backdrop-filter: blur(10px); padding: 0 28px; position: sticky; top: 0; z-index: 2; }
  header div:first-of-type span { display: block; color: #6c7a89; font-size: 12px; }
  header div:first-of-type strong { font-size: 16px; }
  .status { display: flex; gap: 8px; align-items: center; color: #526273; font-size: 13px; }
  .status code { background: #e7edf3; border-radius: 6px; padding: 4px 7px; color: #1d3448; }
  .content-scroll { height: calc(100vh - 69px); overflow: auto; }
  .summary-row { max-width: 1160px; margin: 24px auto 0; padding: 0 28px; display: grid; grid-template-columns: 1.6fr repeat(4, 1fr); gap: 12px; }
  .summary-row div { background: #fff; border: 1px solid #dbe1ea; border-radius: 8px; min-height: 70px; padding: 14px; display: flex; align-items: center; gap: 9px; }
  .summary-row span { font-size: 24px; font-weight: 750; color: #14304a; }
  .summary-row small { color: #607286; font-size: 12px; }
  .layout { max-width: 1160px; margin: 0 auto; padding: 20px 28px 80px; display: grid; grid-template-columns: minmax(0, 1fr) 230px; gap: 30px; align-items: start; }
  article { min-width: 0; background: #fff; border: 1px solid #dbe1ea; border-radius: 8px; padding: 34px 38px; line-height: 1.68; }
  .toc { position: sticky; top: 92px; border-radius: 8px; border: 1px solid #dbe1ea; padding: 14px; max-height: calc(100vh - 120px); overflow: auto; }
  .toc strong { display: block; font-size: 13px; margin-bottom: 8px; }
  .toc a { display: block; color: #56677a; text-decoration: none; font-size: 13px; padding: 5px 0; }
  .toc a.sub { padding-left: 12px; font-size: 12px; }
  .loading-grid { padding: 32px; display: grid; gap: 14px; }
  .loading-grid div, .empty { background: #e7edf3; border-radius: 8px; height: 100px; }
  .empty { margin: 32px; height: auto; padding: 28px; background: #fff; border: 1px solid #dbe1ea; }
  .toast { position: fixed; right: 22px; bottom: 22px; background: #102337; color: #fff; border-radius: 8px; padding: 10px 13px; display: flex; gap: 8px; align-items: center; box-shadow: 0 14px 30px rgba(15, 35, 55, .22); }
  :global(article h1) { font-size: 34px; line-height: 1.18; margin: 0 0 18px; color: #102337; }
  :global(article h2) { font-size: 22px; margin: 36px 0 14px; padding-top: 24px; border-top: 1px solid #dbe1ea; color: #122b42; }
  :global(article h3) { font-size: 16px; margin: 24px 0 10px; color: #20384f; }
  :global(article h1 a), :global(article h2 a), :global(article h3 a) { color: inherit; text-decoration: none; }
  :global(.endpoint-heading) { display: flex; gap: 10px; align-items: center; flex-wrap: wrap; }
  :global(.method) { font-size: 12px; min-width: 58px; text-align: center; color: #fff; border-radius: 999px; padding: 4px 9px; font-weight: 800; }
  :global(.method-GET) { background: #0f766e; }
  :global(.method-POST) { background: #1d4ed8; }
  :global(.method-PUT), :global(.method-PATCH) { background: #b45309; }
  :global(.method-DELETE) { background: #b91c1c; }
  :global(article p) { color: #45576a; margin: 12px 0; }
  :global(article table) { border-collapse: collapse; width: 100%; margin: 16px 0 22px; font-size: 13px; overflow: hidden; }
  :global(article th), :global(article td) { border: 1px solid #dbe1ea; padding: 10px 12px; vertical-align: top; }
  :global(article th) { background: #eef3f7; color: #24394d; font-weight: 750; }
  :global(article td) { color: #47596b; }
  :global(article code) { background: #eef3f7; color: #14304a; border-radius: 5px; padding: 2px 5px; font-size: .92em; }
  :global(.code-wrap) { position: relative; margin: 16px 0 24px; }
  :global(.copy-code) { position: absolute; top: 9px; right: 9px; border: 1px solid #42556a; background: #1d2937; color: #fff; border-radius: 6px; padding: 5px 9px; font-size: 12px; }
  :global(pre) { background: #111827; color: #e5e7eb; border-radius: 8px; padding: 18px; overflow: auto; }
  :global(pre code) { background: transparent; color: inherit; padding: 0; }
  :global(hr) { border: 0; border-top: 1px solid #dbe1ea; margin: 28px 0; }
  :global(ul) { padding-left: 20px; color: #45576a; }
  @media (max-width: 980px) {
    .docs-shell { display: block; }
    .docs-shell > aside { position: fixed; z-index: 20; width: 84vw; max-width: 330px; transform: translateX(-100%); transition: transform .2s ease; }
    .docs-shell > aside.open { transform: translateX(0); }
    .brand button, .menu { display: inline-flex; }
    .layout { display: block; padding: 16px; }
    .summary-row { grid-template-columns: repeat(2, minmax(0, 1fr)); padding: 0 16px; }
    .toc { display: none; }
    article { padding: 24px 18px; }
    header { padding: 0 16px; }
    .status { display: none; }
  }
</style>
