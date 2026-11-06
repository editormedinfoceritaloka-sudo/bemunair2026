<script lang="ts">
  import DOMPurify from 'isomorphic-dompurify';

  let {
    body
  }: {
    body: string;
  } = $props();

  const safe = $derived(
    DOMPurify.sanitize(body, {
      ALLOWED_TAGS: ['p', 'br', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'ul', 'ol', 'li', 'strong', 'em', 'u', 'a', 'img', 'blockquote', 'code', 'pre', 'span', 'div', 'figure', 'figcaption'],
      ALLOWED_ATTR: ['href', 'src', 'alt', 'title', 'class'],
      ALLOWED_URI_REGEXP: /^(?:(?:https?|mailto):|[^a-z]|[a-z+.\-]+(?:[^a-z+.\-:]|$))/i
    })
  );
</script>

<div class="article-body">
  {@html safe}
</div>

<style>
  .article-body {
    position: relative;
    max-width: 52rem;
    margin: 0 auto;
    color: #0a2445;
    font-size: 1rem;
    line-height: 1.9;
  }

  .article-body :global(p) {
    margin: 1.35rem 0;
    color: rgb(10 36 69 / 0.78);
    font-weight: 500;
  }

  .article-body :global(h2) {
    margin-top: 3.5rem;
    margin-bottom: 1.25rem;
    color: #0a2445;
    font-size: clamp(1.75rem, 4vw, 2.6rem);
    line-height: 1.05;
    font-weight: 900;
    letter-spacing: -0.045em;
  }

  .article-body :global(h3) {
    margin-top: 2.75rem;
    margin-bottom: 1rem;
    color: #174487;
    font-size: clamp(1.35rem, 3vw, 1.8rem);
    line-height: 1.15;
    font-weight: 900;
    letter-spacing: -0.035em;
  }

  .article-body :global(strong) {
    color: #0a2445;
    font-weight: 850;
  }

  .article-body :global(em) {
    color: rgb(10 36 69 / 0.7);
  }

  .article-body :global(a) {
    color: #e0621a;
    font-weight: 750;
    text-decoration: underline;
    text-decoration-color: rgb(224 98 26 / 0.3);
    text-underline-offset: 0.2em;
  }

  .article-body :global(a:hover) {
    color: #a33d0a;
  }

  .article-body :global(blockquote) {
    margin: 2.5rem 0;
    border-left: 5px solid #f47b20;
    border-radius: 0 1rem 1rem 0;
    background: rgb(254 240 227 / 0.75);
    padding: 1.5rem 1.75rem;
    color: #7d2d07;
    font-size: 1.08rem;
    line-height: 1.75;
    font-weight: 700;
  }

  .article-body :global(blockquote p) {
    margin: 0;
    color: inherit;
  }

  .article-body :global(ul),
  .article-body :global(ol) {
    margin: 1.5rem 0;
    padding-left: 1.75rem;
    color: rgb(10 36 69 / 0.78);
  }

  .article-body :global(ul) {
    list-style: disc;
  }

  .article-body :global(ol) {
    list-style: decimal;
  }

  .article-body :global(li) {
    margin: 0.65rem 0;
    padding-left: 0.35rem;
  }

  .article-body :global(li::marker) {
    color: #e0621a;
    font-weight: 800;
  }

  .article-body :global(img) {
    display: block;
    width: 100%;
    height: auto;
    margin: 2.75rem auto;
    border: 6px solid white;
    border-radius: 1.75rem;
    object-fit: cover;
    box-shadow: 0 24px 60px rgb(5 34 70 / 0.18);
  }

  .article-body :global(hr) {
    width: 100%;
    height: 1px;
    margin: 3.5rem 0;
    border: 0;
    background: linear-gradient(
      to right,
      transparent,
      rgb(10 36 69 / 0.2),
      transparent
    );
  }

  .article-body :global(code) {
    border-radius: 0.4rem;
    background: rgb(10 36 69 / 0.07);
    padding: 0.15rem 0.4rem;
    color: #174487;
    font-size: 0.9em;
  }

  .article-body :global(pre) {
    margin: 2rem 0;
    overflow-x: auto;
    border-radius: 1.25rem;
    background: #0a2445;
    padding: 1.5rem;
    color: white;
  }

  .article-body :global(pre code) {
    background: transparent;
    padding: 0;
    color: inherit;
  }

  @media (min-width: 640px) {
    .article-body {
      font-size: 1.08rem;
    }
  }
</style>
