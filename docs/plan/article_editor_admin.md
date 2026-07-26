# Notion-style Article Editor Admin

- Routes: `/admin/articles/new`, `/admin/articles/[id]/edit`, protected preview.
- Canvas: borderless large title, ImageKit cover upload with preview/replace/remove actions, excerpt counter, clean centered document body, sticky status/save toolbar.
- Tiptap core: paragraph, H2/H3, bold, italic, underline, strike, code, bullet/ordered list, quote, code block, divider, link, uploaded image, undo/redo and markdown shortcuts.
- Toolbar image action and `/gambar` flow open a local file picker, upload through `POST /api/uploads/imagekit`, then insert the returned ImageKit URL into the document automatically.
- Accepted images: JPEG, PNG, WebP, GIF and AVIF up to 10 MB by default. SVG is intentionally excluded.
- Image upload provides loading, success and failure feedback. Article submission is blocked while the cover is still uploading.
- Bubble menu uses shadcn Button and Tiptap BubbleMenu. `/` opens the block insertion menu.
- First draft is explicit POST; after ID exists, PUT autosaves with 1.5s debounce and Saving/Saved/Error states.
- Preview saves then reloads server-sanitized HTML. Publish flushes pending save before status PUT.
- Validation: title 1–255, excerpt ≤500, cover URL result ≤500, and non-empty body.
- Private ImageKit credentials are server-only. The browser only receives the normalized upload result.
- Out of scope: drag-and-drop block sorting, tables, collaboration and AI writing.
