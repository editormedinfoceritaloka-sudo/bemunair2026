UPDATE content_submissions AS submission
JOIN users AS submitter ON submitter.id = submission.submitter_id
LEFT JOIN ministries AS ministry ON ministry.code = UPPER(submission.ministry)
SET submission.submitter_name = submitter.name,
    submission.submitter_phone = submitter.phone,
    submission.ministry_id = ministry.id,
    submission.service_type = IF(submission.submission_type = 'ARTIKEL', 'ARTICLE', 'CONTENT'),
    submission.content_format = CASE
      WHEN submission.submission_type = 'ARTIKEL' THEN NULL
      WHEN submission.submission_type = 'INSTASTORY' THEN 'INSTASTORY'
      ELSE 'FEEDS_REELS_LEGACY'
    END,
    submission.submitted_at = COALESCE(submission.submitted_at, submission.created_at),
    submission.request_code = COALESCE(
      submission.request_code,
      CONCAT('MED-', YEAR(submission.created_at), '-', LPAD(submission.id, 6, '0'))
    )
WHERE submission.id BETWEEN 900001 AND 900999;

UPDATE letter_submissions AS submission
JOIN users AS submitter ON submitter.id = submission.submitter_id
LEFT JOIN ministries AS ministry ON ministry.code = UPPER(submission.ministry)
SET submission.submitter_name = submitter.name,
    submission.submitter_phone = submitter.phone,
    submission.ministry_id = ministry.id,
    submission.submitted_at = COALESCE(submission.submitted_at, submission.created_at),
    submission.request_code = COALESCE(
      submission.request_code,
      CONCAT('SUR-', YEAR(submission.created_at), '-', LPAD(submission.id, 6, '0'))
    )
WHERE submission.id BETWEEN 900001 AND 900999;

DELETE FROM content_submission_status_histories
WHERE submission_id BETWEEN 900001 AND 900999 AND note = 'Development seed state';

INSERT INTO content_submission_status_histories (
  submission_id, actor_id, from_status, to_status, note, created_at
)
SELECT id, submitter_id, NULL, status, 'Development seed state', created_at
FROM content_submissions
WHERE id BETWEEN 900001 AND 900999;

DELETE FROM letter_submission_status_histories
WHERE submission_id BETWEEN 900001 AND 900999 AND note = 'Development seed state';

INSERT INTO letter_submission_status_histories (
  submission_id, actor_id, from_status, to_status, note, created_at
)
SELECT id, submitter_id, NULL, status, 'Development seed state', created_at
FROM letter_submissions
WHERE id BETWEEN 900001 AND 900999;
