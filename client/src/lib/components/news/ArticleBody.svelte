<script lang="ts">
  import DOMPurify from 'isomorphic-dompurify';

  let { body }: { body: string } = $props();

  const safe = $derived(
    DOMPurify.sanitize(body, {
      ALLOWED_TAGS: [
        'p', 'br', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6',
        'ul', 'ol', 'li', 'strong', 'em', 'u', 'a', 'img',
        'blockquote', 'code', 'pre', 'span', 'div',
        'figure', 'figcaption', 'hr'
      ],
      ALLOWED_ATTR: ['href', 'src', 'alt', 'title', 'class']
    })
  );
</script>

<div class="article-body">
  {@html safe}
</div>

<style>
  .article-body {
    position: relative;
    width: 100%;
    max-width: 46rem;
    margin: 0 auto;
    color: #0a2445;
    font-size: 0.925rem;
    line-height: 1.8;
    overflow-wrap: anywhere;
  }

  .article-body :global(p) {
    margin: 1rem 0;
    color: rgb(10 36 69 / 0.8);
    font-weight: 400;
  }

  .article-body :global(h1),
  .article-body :global(h2) {
    margin-top: 2.25rem;
    margin-bottom: 0.8rem;
    color: #0a2445;
    font-size: clamp(1.35rem, 2.8vw, 1.85rem);
    line-height: 1.25;
    font-weight: 800;
    letter-spacing: -0.025em;
  }

  .article-body :global(h3) {
    margin-top: 1.85rem;
    margin-bottom: 0.65rem;
    color: #174487;
    font-size: clamp(1.15rem, 2vw, 1.4rem);
    line-height: 1.3;
    font-weight: 750;
    letter-spacing: -0.02em;
  }

  .article-body :global(h4),
  .article-body :global(h5),
  .article-body :global(h6) {
    margin-top: 1.5rem;
    margin-bottom: 0.6rem;
    color: #0a2445;
    font-size: 1.05rem;
    line-height: 1.4;
    font-weight: 700;
  }

  .article-body :global(strong) {
    color: #0a2445;
    font-weight: 750;
  }

  .article-body :global(em) {
    color: rgb(10 36 69 / 0.75);
  }

  .article-body :global(a) {
    color: #e0621a;
    font-weight: 650;
    text-decoration: underline;
    text-decoration-color: rgb(224 98 26 / 0.35);
    text-underline-offset: 0.2em;
  }

  .article-body :global(a:hover) {
    color: #a33d0a;
  }

  .article-body :global(blockquote) {
    margin: 1.75rem 0;
    border-left: 3px solid #f47b20;
    border-radius: 0 0.6rem 0.6rem 0;
    background: rgb(254 240 227 / 0.65);
    padding: 0.9rem 1.15rem;
    color: #7d2d07;
    font-size: 0.95rem;
    line-height: 1.7;
    font-weight: 600;
  }

  .article-body :global(blockquote p) {
    margin: 0;
    color: inherit;
  }

  .article-body :global(ul),
  .article-body :global(ol) {
    margin: 1rem 0;
    padding-left: 1.5rem;
    color: rgb(10 36 69 / 0.8);
  }

  .article-body :global(ul) {
    list-style: disc;
  }

  .article-body :global(ol) {
    list-style: decimal;
  }

  .article-body :global(li) {
    margin: 0.4rem 0;
    padding-left: 0.2rem;
  }

  .article-body :global(li::marker) {
    color: #e0621a;
    font-weight: 700;
  }

  .article-body :global(img) {
    display: block;
    width: auto;
    max-width: 100%;
    height: auto;
    max-height: 28rem;
    margin: 1.8rem auto;
    border: 3px solid white;
    border-radius: 0.85rem;
    object-fit: contain;
    box-shadow: 0 12px 32px rgb(5 34 70 / 0.12);
  }

  .article-body :global(figure) {
    margin: 1.7rem auto;
  }

  .article-body :global(figcaption) {
    margin-top: 0.35rem;
    color: rgb(10 36 69 / 0.6);
    font-size: 0.8rem;
    line-height: 1.5;
    text-align: center;
  }

  .article-body :global(hr) {
    width: 100%;
    height: 1px;
    margin: 2rem 0;
    border: 0;
    background: linear-gradient(
      to right,
      transparent,
      rgb(10 36 69 / 0.2),
      transparent
    );
  }

  .article-body :global(code) {
    border-radius: 0.3rem;
    background: rgb(10 36 69 / 0.07);
    padding: 0.1rem 0.3rem;
    color: #174487;
    font-size: 0.88em;
  }

  .article-body :global(pre) {
    margin: 1.35rem 0;
    overflow-x: auto;
    border-radius: 0.75rem;
    background: #0a2445;
    padding: 1rem 1.15rem;
    color: white;
    font-size: 0.85rem;
    line-height: 1.65;
  }

  .article-body :global(pre code) {
    background: transparent;
    padding: 0;
    color: inherit;
  }

  @media (min-width: 640px) {
    .article-body {
      font-size: 1rem;
    }
  }
</style>