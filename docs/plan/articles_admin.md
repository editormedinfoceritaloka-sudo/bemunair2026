# Articles Admin

- Route: `/admin/articles`; admin list uses server pagination (`page`, `per_page`, max 50).
- Columns/cards: cover, title, slug, author, DRAFT/PUBLISHED badge, published/updated time and actions.
- Actions: create, edit, protected preview, public preview when published, publish/unpublish and delete.
- Publishing flushes editor save first. Empty state links to create article.
