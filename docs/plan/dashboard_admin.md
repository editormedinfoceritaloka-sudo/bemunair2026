# Dashboard Admin

- Route: `/admin`; parallel-load users, content, letters, queue, templates, article meta and health.
- Cards: total users, open content, open letters, published/draft articles; queue current PJ and health summary.
- Operational panels: combined upcoming deadlines, status distribution, newest submissions, quick actions.
- Monthly operational calendar:
  - content submissions use `deadline` (or `publish_date` as fallback);
  - letter submissions use `deadline`;
  - published articles use `published_at`;
  - previous/next month navigation and “Hari ini” shortcut;
  - each agenda is color-coded by module and opens a shadcn Dialog with status, ministry/PJ, WIB time, and a link to its existing detail route;
  - mobile cells use compact colored indicators while retaining an accessible click target and full detail dialog.
- Components: shadcn Card, Badge, Progress, Skeleton, Alert, Empty and responsive tables.
- Calendar components: shadcn Card, Button, Dialog, existing StatusBadge, and Lucide module icons.
- Empty datasets render onboarding actions instead of zero-value charts.
