# Content Submissions Admin

- Routes: list, create and detail under `/admin/content-submissions`.
- List supports search plus status/type/ministry/PJ filters and deadline ordering.
- Create is multipart and conditional: FEEDS_REELS/INSTASTORY require publish date/time, design drive and Canva; ARTIKEL requires article drive; all require title, caption and brief link.
- Detail groups submitter, brief, publication, design links, assignment, notes and timestamps.
- Status actions only permit PENDING→IN_REVIEW and IN_REVIEW→APPROVED/REJECTED; rejection requires notes in UI. Delete requires confirmation.
