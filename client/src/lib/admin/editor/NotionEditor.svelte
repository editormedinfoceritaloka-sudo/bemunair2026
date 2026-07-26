<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { Editor } from '@tiptap/core';
  import { StarterKit } from '@tiptap/starter-kit';
  import { Image } from '@tiptap/extension-image';
  import { Placeholder } from '@tiptap/extension-placeholder';
  import { Typography } from '@tiptap/extension-typography';
  import { BubbleMenu } from '@tiptap/extension-bubble-menu';
  import { Button } from '$lib/components/ui/button';
  import { uploadImageFile } from '$lib/admin/upload-image';
  import { toast } from 'svelte-sonner';
  import { Bold, Italic, Underline, Strikethrough, Code, Link, ImagePlus, Undo2, Redo2, List, ListOrdered, Quote, Minus, Pilcrow, Heading2, Heading3, LoaderCircle } from '@lucide/svelte';

  let { value = '', onchange = (_html: string) => {} }: { value?: string; onchange?: (html: string) => void } = $props();
  let editorElement: HTMLElement;
  let bubbleElement: HTMLElement;
  let wrapper: HTMLElement;
  let imageInput: HTMLInputElement;
  let editor = $state<Editor | null>(null);
  let revision = $state(0);
  let slashOpen = $state(false);
  let slashTop = $state(0);
  let slashLeft = $state(0);
  let selected = $state(0);
  let imageUploading = $state(false);

  const slashItems = [
    { label: 'Teks', description: 'Paragraf biasa', run: () => editor?.chain().focus().setParagraph().run() },
    { label: 'Heading 2', description: 'Judul bagian besar', run: () => editor?.chain().focus().toggleHeading({ level: 2 }).run() },
    { label: 'Heading 3', description: 'Judul bagian kecil', run: () => editor?.chain().focus().toggleHeading({ level: 3 }).run() },
    { label: 'Bullet list', description: 'Daftar berpoin', run: () => editor?.chain().focus().toggleBulletList().run() },
    { label: 'Numbered list', description: 'Daftar bernomor', run: () => editor?.chain().focus().toggleOrderedList().run() },
    { label: 'Quote', description: 'Kutipan penting', run: () => editor?.chain().focus().toggleBlockquote().run() },
    { label: 'Code block', description: 'Blok kode', run: () => editor?.chain().focus().toggleCodeBlock().run() },
    { label: 'Divider', description: 'Garis pemisah', run: () => editor?.chain().focus().setHorizontalRule().run() },
    { label: 'Gambar', description: 'Unggah dari perangkat', run: () => openImagePicker() }
  ];

  function link() {
    if (!editor) return;
    const current = editor.getAttributes('link').href || '';
    const href = window.prompt('URL tautan (https://...)', current);
    if (href === null) return;
    if (!href) {
      editor.chain().focus().extendMarkRange('link').unsetLink().run();
      return;
    }
    editor.chain().focus().extendMarkRange('link').setLink({ href }).run();
  }

  function openImagePicker() {
    if (!imageUploading) imageInput.click();
  }

  async function uploadEditorImage(event: Event) {
    const input = event.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file || !editor) return;
    imageUploading = true;
    try {
      const uploaded = await uploadImageFile(file, 'article');
      editor.chain().focus().setImage({
        src: uploaded.url,
        alt: file.name,
        title: file.name
      }).run();
      toast.success('Gambar berhasil ditambahkan');
    } catch (uploadError) {
      toast.error(uploadError instanceof Error ? uploadError.message : 'Gagal mengunggah gambar');
    } finally {
      imageUploading = false;
      input.value = '';
    }
  }

  function runSlash(index: number) {
    if (!editor) return;
    const pos = editor.state.selection.from;
    editor.chain().focus().deleteRange({ from: Math.max(0, pos - 1), to: pos }).run();
    slashItems[index].run();
    slashOpen = false;
    selected = 0;
  }

  function checkSlash() {
    if (!editor || !wrapper) return;
    const text = editor.state.selection.$from.parent.textContent;
    slashOpen = text === '/';
    if (slashOpen) {
      const coords = editor.view.coordsAtPos(editor.state.selection.from);
      const rect = wrapper.getBoundingClientRect();
      slashTop = coords.bottom - rect.top + 8;
      slashLeft = Math.max(12, Math.min(coords.left - rect.left, rect.width - 300));
      selected = 0;
    }
  }

  onMount(() => {
    editor = new Editor({
      element: editorElement,
      extensions: [
        StarterKit.configure({ heading: { levels: [2, 3] }, link: { openOnClick: false, autolink: true, defaultProtocol: 'https', HTMLAttributes: { class: 'editor-link' } }, underline: {} }),
        Image.configure({ allowBase64: false, HTMLAttributes: { class: 'editor-image' } }),
        Placeholder.configure({ placeholder: ({ node }) => node.type.name === 'heading' ? 'Judul bagian' : 'Tulis cerita, atau ketik / untuk perintah…' }),
        Typography,
        BubbleMenu.configure({ element: bubbleElement, options: { placement: 'top', offset: 8 } })
      ],
      content: value || '<p></p>',
      editorProps: {
        attributes: { class: 'notion-prose' },
        handleKeyDown: (_view, event) => {
          if (!slashOpen) return false;
          if (event.key === 'Escape') { slashOpen = false; return true; }
          if (event.key === 'ArrowDown') { selected = (selected + 1) % slashItems.length; return true; }
          if (event.key === 'ArrowUp') { selected = (selected - 1 + slashItems.length) % slashItems.length; return true; }
          if (event.key === 'Enter') { runSlash(selected); return true; }
          return false;
        }
      },
      onUpdate: ({ editor }) => { revision++; onchange(editor.getHTML()); checkSlash(); },
      onSelectionUpdate: () => { revision++; checkSlash(); }
    });
  });

  onDestroy(() => editor?.destroy());
</script>

<div class="editor-shell" bind:this={wrapper}>
  <span class="hidden">{revision}</span>
  <input bind:this={imageInput} type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={uploadEditorImage} />
  <div class="editor-toolbar" aria-label="Toolbar editor">
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleHeading({ level: 2 }).run()} class={editor?.isActive('heading', { level: 2 }) ? 'active' : ''}><Heading2 /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleHeading({ level: 3 }).run()} class={editor?.isActive('heading', { level: 3 }) ? 'active' : ''}><Heading3 /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleBulletList().run()} class={editor?.isActive('bulletList') ? 'active' : ''}><List /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleOrderedList().run()} class={editor?.isActive('orderedList') ? 'active' : ''}><ListOrdered /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleBlockquote().run()} class={editor?.isActive('blockquote') ? 'active' : ''}><Quote /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={openImagePicker} disabled={imageUploading} aria-label="Unggah gambar" title="Unggah gambar">{#if imageUploading}<LoaderCircle class="animate-spin" />{:else}<ImagePlus />{/if}</Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().setHorizontalRule().run()}><Minus /></Button>
    <span class="mx-1 h-5 border-l"></span>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().undo().run()} disabled={!editor?.can().undo()}><Undo2 /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().redo().run()} disabled={!editor?.can().redo()}><Redo2 /></Button>
  </div>
  <div bind:this={bubbleElement} class="bubble-menu">
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleBold().run()} class={editor?.isActive('bold') ? 'active' : ''}><Bold /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleItalic().run()} class={editor?.isActive('italic') ? 'active' : ''}><Italic /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleUnderline().run()} class={editor?.isActive('underline') ? 'active' : ''}><Underline /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleStrike().run()} class={editor?.isActive('strike') ? 'active' : ''}><Strikethrough /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={() => editor?.chain().focus().toggleCode().run()} class={editor?.isActive('code') ? 'active' : ''}><Code /></Button>
    <Button type="button" variant="ghost" size="icon" onclick={link} class={editor?.isActive('link') ? 'active' : ''}><Link /></Button>
  </div>
  <div bind:this={editorElement} class="editor-canvas"></div>
  {#if slashOpen}
    <div class="slash-menu" style:top={`${slashTop}px`} style:left={`${slashLeft}px`}>
      <div class="border-b px-3 py-2 text-[11px] font-semibold uppercase tracking-widest text-black-200">Basic blocks</div>
      {#each slashItems as item, index}
        <button type="button" class:selected={index === selected} onmouseenter={() => selected = index} onclick={() => runSlash(index)}>
          <span class="grid size-9 place-items-center rounded-md border bg-card"><Pilcrow class="size-4" /></span>
          <span><strong>{item.label}</strong><small>{item.description}</small></span>
        </button>
      {/each}
    </div>
  {/if}
</div>

<style>
  .editor-shell{position:relative;border:1px solid var(--border);border-radius:12px;background:var(--card);min-height:520px}.editor-toolbar{position:sticky;top:0;z-index:10;display:flex;align-items:center;gap:1px;overflow-x:auto;border-bottom:1px solid var(--border);background:color-mix(in srgb,var(--card) 96%,transparent);padding:6px}.editor-toolbar :global(button),.bubble-menu :global(button){width:32px;height:32px}:global(.active){background:var(--border)!important;color:var(--blue-700)!important}.bubble-menu{display:flex;gap:1px;border:1px solid var(--border);border-radius:8px;background:var(--card);padding:3px;box-shadow:0 10px 30px rgba(12,12,13,.14);visibility:hidden}.editor-canvas{min-height:470px;padding:48px clamp(24px,8vw,96px) 100px}.slash-menu{position:absolute;z-index:30;width:288px;max-height:360px;overflow:auto;border:1px solid var(--border);border-radius:10px;background:var(--card);box-shadow:0 18px 50px rgba(12,12,13,.18)}.slash-menu button{display:flex;width:100%;align-items:center;gap:10px;padding:7px 10px;text-align:left}.slash-menu button.selected,.slash-menu button:hover{background:var(--white-500)}.slash-menu strong,.slash-menu small{display:block}.slash-menu strong{font-size:13px}.slash-menu small{font-size:11px;color:var(--black-300)}
  :global(.notion-prose){outline:none;min-height:340px;color:var(--black-500);font-size:17px;line-height:1.8}:global(.notion-prose p){margin:.45em 0}:global(.notion-prose h2){font-family:Georgia,serif;font-size:1.75rem;font-weight:700;line-height:1.25;margin:1.5em 0 .45em}:global(.notion-prose h3){font-family:Georgia,serif;font-size:1.3rem;font-weight:700;line-height:1.3;margin:1.25em 0 .35em}:global(.notion-prose ul){list-style:disc;padding-left:1.5rem}:global(.notion-prose ol){list-style:decimal;padding-left:1.5rem}:global(.notion-prose blockquote){border-left:3px solid var(--orange-500);margin:1.2rem 0;padding:.25rem 0 .25rem 1.25rem;color:var(--black-300)}:global(.notion-prose pre){background:var(--black-800);color:var(--white-200);padding:1rem;border-radius:8px;overflow:auto}:global(.notion-prose hr){border:0;border-top:1px solid var(--white-700);margin:2rem 0}:global(.notion-prose img){display:block;max-width:100%;border-radius:8px;margin:1.5rem auto}:global(.notion-prose a){color:var(--blue-600);text-decoration:underline}:global(.notion-prose .is-editor-empty:first-child::before){color:var(--black-200);content:attr(data-placeholder);float:left;height:0;pointer-events:none}
</style>
