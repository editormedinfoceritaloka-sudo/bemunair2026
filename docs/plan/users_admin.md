# Users Admin

- Route: `/admin/users`; covers users list/detail/create/update/delete.
- DataTable fields: name, email, role, ministry, phone and row actions; client-side search and role/ministry filters.
- Create/edit Dialog: name, email, role, conditional ministry, phone; password minimum 8 on create only.
- Delete uses AlertDialog; current signed-in user cannot be deleted from UI. API conflicts show inline errors.
- Responsive: desktop table, mobile stacked user cards.
