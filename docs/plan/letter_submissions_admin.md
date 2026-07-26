# Letter Submissions Admin

- Routes: list, create and detail under `/admin/letter-submissions`.
- List filters status, ministry, letter type and deadline; detail shows submitter/PJ, body, notes and timeline.
- Create fields: ministry, letterType, subject, body and RFC3339 deadline. Template selection pre-fills type, subject and body.
- Status transition and destructive behavior mirror content submissions.
