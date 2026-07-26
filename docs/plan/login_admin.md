# Login Admin

- Route: `/admin/login`; public, redirect authenticated ADMIN to `/admin`.
- Fields: email and password; submit through SvelteKit action to `/api/v1/auth/login`, require `user.role === ADMIN`, then validate `/auth/me`.
- UI: centered BEM-branded shadcn Card, Input, Label, Alert and Button; password visibility toggle and pending state.
- Errors: invalid credentials, non-admin forbidden, server unavailable. Never reveal whether an email exists.
- Acceptance: JWT only in HttpOnly cookie, Enter submits, focus moves to error, logout deletes cookie.
