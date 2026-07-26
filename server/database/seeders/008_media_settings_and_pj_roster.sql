INSERT INTO medinfo_pj_queues (user_id, position, is_current)
SELECT
  users.id,
  ROW_NUMBER() OVER (ORDER BY users.id),
  FALSE
FROM users
JOIN ministries ON ministries.id = users.ministry_id
WHERE users.role = 'ADMIN_MEDINFO' AND ministries.code = 'MEDINFO'
ON DUPLICATE KEY UPDATE is_current = FALSE;

-- Data development mengikuti aturan satu PJ hanya menangani satu task aktif.
DELETE FROM content_submission_assignment_histories WHERE submission_id BETWEEN 900001 AND 900999;
DELETE FROM letter_submission_assignment_histories WHERE submission_id BETWEEN 900001 AND 900999;
UPDATE content_submissions SET assigned_pj_id = NULL WHERE id BETWEEN 900001 AND 900999;
UPDATE letter_submissions SET assigned_pj_id = NULL WHERE id BETWEEN 900001 AND 900999;
UPDATE content_submissions SET assigned_pj_id = (SELECT id FROM users WHERE email = "mentri.medinfo@bem.unair.ac.id") WHERE id = 900001;
UPDATE letter_submissions SET assigned_pj_id = (SELECT id FROM users WHERE email = "editor.medinfo@bem.unair.ac.id") WHERE id = 900002;

UPDATE media_submission_settings
SET
  sop_url = 'https://docs.google.com/document/d/1RgIdOLXwUJz6pRrktK2fLQ6srvCZn0y8/edit',
  ministry_template_url = 'https://linktr.ee/templateeditingkementerian',
  brief_template_url = 'https://docs.google.com/document/d/1bBIs7dcqL1OMSJWAoXZduy2c6r0MpvLIC84T5t3OZJ8/edit?tab=t.0',
  caption_template_url = NULL,
  minimum_lead_days = 7
WHERE service_type = 'CONTENT';

UPDATE media_submission_settings
SET
  sop_url = 'https://docs.google.com/document/d/1RgIdOLXwUJz6pRrktK2fLQ6srvCZn0y8/edit',
  ministry_template_url = 'https://linktr.ee/templateeditingkementerian',
  brief_template_url = NULL,
  caption_template_url = 'https://docs.google.com/document/d/1--ki9SXtg1tTQQFFNQElnQ4brK1c8mMA6Ww540XCKzk/edit?tab=t.0',
  minimum_lead_days = 3
WHERE service_type = 'ARTICLE';
