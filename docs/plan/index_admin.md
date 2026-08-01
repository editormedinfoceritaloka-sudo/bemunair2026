# Admin BEM UNAIR — Architecture & API Coverage

## Goal
Portal operasional role `ADMIN` untuk mengelola users, content submission, letter submission, antrean PJ, artikel, dan health service melalui SvelteKit server-side BFF.

## Routes
`/admin/login`, `/admin`, `/admin/users`, `/admin/content-submissions`, `/admin/content-submissions/new`, `/admin/content-submissions/[id]`, `/admin/letter-submissions`, `/admin/letter-submissions/new`, `/admin/letter-submissions/[id]`, `/admin/medinfo-queue`, `/admin/articles`, `/admin/articles/new`, `/admin/articles/[id]/edit`, `/admin/articles/[id]/preview`, `/admin/system`.

## API coverage
- Auth: `POST /api/v1/auth/login`, `GET /api/v1/auth/me`; public register removed.
- Users: GET list/detail, POST, PUT, DELETE `/api/v1/users`.
- Content: POST multipart, GET list/detail, PUT status, DELETE `/api/v1/content-submissions`.
- Letters: POST, GET list/detail, PUT status, DELETE `/api/v1/letter-submissions`.
- Queue: GET, POST, PUT reorder, DELETE `/api/v1/medinfo-pj/queue`.
- Templates: Admin Medinfo mengelola PDF template surat; Admin Kementerian hanya mengunduh template saat mengajukan surat.
- Articles: public GET list/slug; admin GET list/detail, POST, PUT, publish PUT, DELETE.
- System: backend `/ping` and WA Engine `/health` server-side only.

## Foundation
JWT stays in HttpOnly SameSite=Lax cookie. Admin pages never expose it to browser JavaScript. Server actions forward Bearer auth, normalize API envelopes, clear expired sessions, and redirect to login. UI is BEM editorial: light canvas, navy primary, restrained gold accent, Geist, collapsible shadcn Sidebar, responsive Sheet behavior.

## Shared states
Every page defines skeleton/loading, empty, retry/error, success toast, destructive AlertDialog, disabled pending actions, keyboard focus restoration, and mobile card fallback for wide tables.
