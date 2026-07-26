# Medinfo PJ Queue Admin

- Route: `/admin/medinfo-queue`; covers list, add, reorder and delete.
- Ordered cards identify current PJ, position, ministry and contact.
- Add Combobox includes only MENTRI users not already queued.
- Reorder sends the complete ordered row IDs; optimistic UI rolls back on failure. First item becomes current.
- Remove uses confirmation and explains assignment impact.
