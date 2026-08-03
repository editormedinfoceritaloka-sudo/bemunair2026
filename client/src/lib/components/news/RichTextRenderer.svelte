<script lang="ts">
  import { onMount } from 'svelte';

  import {
    Editor,
    type JSONContent
  } from '@tiptap/core';

  import { StarterKit } from '@tiptap/starter-kit';
  import Image from '@tiptap/extension-image';
  import Typography from '@tiptap/extension-typography';

  let {
    content
  }: {
    content: JSONContent;
  } = $props();

  let editorElement!: HTMLDivElement;

  let editor = $state.raw<Editor | null>(
    null
  );

  onMount(() => {
    editor = new Editor({
      element: editorElement,

      extensions: [
        StarterKit.configure({
          heading: {
            levels: [2, 3, 4]
          },

          link: {
            openOnClick: true,
            autolink: true,
            linkOnPaste: true,
            HTMLAttributes: {
              target: '_blank',
              rel: 'noopener noreferrer'
            }
          }
        }),

        Image.configure({
          inline: false,
          allowBase64: false
        }),

        Typography
      ],

      content,

      editable: false,

      editorProps: {
        attributes: {
          class: 'blog-prose',
          role: 'document',
          spellcheck: 'false',
          tabindex: '-1'
        }
      }
    });

    return () => {
      editor?.destroy();
      editor = null;
    };
  });

  $effect(() => {
    const currentEditor = editor;
    const nextContent = content;

    if (!currentEditor) {
      return;
    }

    const currentJSON = JSON.stringify(
      currentEditor.getJSON()
    );

    const nextJSON = JSON.stringify(
      nextContent
    );

    if (currentJSON === nextJSON) {
      return;
    }

    currentEditor.commands.setContent(
      nextContent,
      {
        emitUpdate: false
      }
    );
  });
</script>

<div
  class="
    relative overflow-hidden
    rounded-[30px]
    border border-blue-900/10
    bg-white
    px-6 py-9
    shadow-[0_20px_50px_rgba(5,34,70,0.1)]
    sm:px-10 sm:py-12
    lg:px-14 lg:py-14
  "
>
  <div
    aria-hidden="true"
    class="
      pointer-events-none absolute
      -right-24 -top-24
      size-72 rounded-full
      bg-blue-200/30
      blur-[90px]
    "
  ></div>

  <div
    aria-hidden="true"
    class="
      pointer-events-none absolute
      -bottom-32 -left-20
      size-72 rounded-full
      bg-orange-100/35
      blur-[100px]
    "
  ></div>

  <div class="relative z-10">
    <div bind:this={editorElement}></div>
  </div>
</div>

<style>
  :global(.blog-prose) {
    width: 100%;
    color: #334155;
    font-size: 1rem;
    font-weight: 500;
    line-height: 1.9;
    outline: none;
  }

  :global(.blog-prose > *:first-child) {
    margin-top: 0;
  }

  :global(.blog-prose > *:last-child) {
    margin-bottom: 0;
  }

  :global(.blog-prose p) {
    margin: 1.35rem 0;
    color: #334155;
    line-height: 1.9;
  }

  :global(.blog-prose > p:first-child) {
    color: #1e3a8a;
    font-size: 1.12rem;
    font-weight: 650;
    line-height: 1.85;
  }

  :global(.blog-prose h2) {
    position: relative;
    margin-top: 3.5rem;
    margin-bottom: 1.2rem;
    padding-bottom: 1rem;
    color: #1d4ed8;
    font-size: clamp(
      1.75rem,
      4vw,
      2.6rem
    );
    font-weight: 950;
    line-height: 1.02;
    letter-spacing: -0.055em;
  }

  :global(.blog-prose h2::after) {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 4.5rem;
    height: 4px;
    border-radius: 999px;
    background: linear-gradient(
      90deg,
      #f97316,
      #fb923c
    );
    content: '';
  }

  :global(.blog-prose h3) {
    margin-top: 2.8rem;
    margin-bottom: 0.9rem;
    color: #1e40af;
    font-size: clamp(
      1.4rem,
      3vw,
      1.95rem
    );
    font-weight: 900;
    line-height: 1.1;
    letter-spacing: -0.04em;
  }

  :global(.blog-prose h4) {
    margin-top: 2.25rem;
    margin-bottom: 0.75rem;
    color: #1e3a8a;
    font-size: 1.25rem;
    font-weight: 900;
    line-height: 1.2;
  }

  :global(.blog-prose strong) {
    color: #172554;
    font-weight: 900;
  }

  :global(.blog-prose em) {
    color: #1e40af;
    font-style: italic;
  }

  :global(.blog-prose u) {
    text-decoration-color: #60a5fa;
    text-decoration-thickness: 2px;
    text-underline-offset: 3px;
  }

  :global(.blog-prose s) {
    color: #64748b;
    text-decoration-color: #f97316;
    text-decoration-thickness: 2px;
  }

  :global(.blog-prose a) {
    color: #1d4ed8;
    font-weight: 800;
    text-decoration-line: underline;
    text-decoration-color: #93c5fd;
    text-decoration-thickness: 2px;
    text-underline-offset: 4px;
    transition:
      color 180ms ease,
      text-decoration-color 180ms ease;
  }

  :global(.blog-prose a:hover) {
    color: #172554;
    text-decoration-color: #f97316;
  }

  :global(.blog-prose ul),
  :global(.blog-prose ol) {
    margin: 1.75rem 0;
    padding-left: 1.75rem;
  }

  :global(.blog-prose ul) {
    list-style-type: disc;
  }

  :global(.blog-prose ol) {
    list-style-type: decimal;
  }

  :global(.blog-prose li) {
    margin: 0.8rem 0;
    padding-left: 0.45rem;
    color: #334155;
    line-height: 1.75;
  }

  :global(.blog-prose li::marker) {
    color: #2563eb;
    font-weight: 900;
  }

  :global(.blog-prose li p) {
    margin: 0;
  }

  :global(.blog-prose blockquote) {
    position: relative;
    margin: 2.6rem 0;
    overflow: hidden;
    border-left: 7px solid #2563eb;
    border-radius: 0 22px 22px 0;
    background: linear-gradient(
      135deg,
      rgba(219, 234, 254, 0.95),
      rgba(147, 197, 253, 0.35)
    );
    padding: 1.75rem 2rem;
    color: #172554;
    font-size: 1.1rem;
    font-weight: 800;
    line-height: 1.75;
    box-shadow:
      inset 0 0 0 1px
        rgba(37, 99, 235, 0.08),
      0 12px 28px
        rgba(30, 64, 175, 0.08);
  }

  :global(.blog-prose blockquote::before) {
    position: absolute;
    top: -0.8rem;
    right: 1.25rem;
    color: rgba(37, 99, 235, 0.12);
    content: '“';
    font-family: Georgia, serif;
    font-size: 7rem;
    font-weight: 900;
    line-height: 1;
  }

  :global(.blog-prose blockquote p) {
    position: relative;
    z-index: 1;
    margin: 0;
    color: inherit;
  }

  :global(.blog-prose hr) {
    height: 4px;
    margin: 3.5rem 0;
    border: 0;
    border-radius: 999px;
    background: linear-gradient(
      90deg,
      transparent,
      #93c5fd,
      #2563eb,
      #f97316,
      transparent
    );
  }

  :global(.blog-prose img) {
    display: block;
    width: 100%;
    max-height: 650px;
    margin: 2.5rem auto;
    border: 5px solid white;
    border-radius: 22px;
    object-fit: cover;
    object-position: center;
    box-shadow:
      0 20px 45px
      rgba(5, 34, 70, 0.18);
  }

  :global(.blog-prose code) {
    border: 1px solid #bfdbfe;
    border-radius: 7px;
    background: #eff6ff;
    padding: 0.18rem 0.42rem;
    color: #1e3a8a;
    font-family:
      'Fira Mono',
      'JetBrains Mono',
      monospace;
    font-size: 0.88em;
    font-weight: 750;
  }

  :global(.blog-prose pre) {
    margin: 2rem 0;
    overflow-x: auto;
    border: 1px solid
      rgba(147, 197, 253, 0.18);
    border-radius: 20px;
    background: linear-gradient(
      145deg,
      #071b36,
      #0b2850
    );
    padding: 1.5rem;
    color: #dbeafe;
    box-shadow:
      0 20px 36px
      rgba(5, 34, 70, 0.18);
  }

  :global(.blog-prose pre code) {
    border: 0;
    background: transparent;
    padding: 0;
    color: inherit;
    font-size: 0.9rem;
    font-weight: 500;
    line-height: 1.75;
  }

  @media (min-width: 640px) {
    :global(.blog-prose) {
      font-size: 1.075rem;
    }

    :global(.blog-prose > p:first-child) {
      font-size: 1.22rem;
    }
  }
</style>