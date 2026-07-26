CREATE TABLE ministries (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  code VARCHAR(50) NOT NULL,
  name VARCHAR(120) NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_ministries_code (code),
  UNIQUE KEY idx_ministries_name (name),
  KEY idx_ministries_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO ministries (code, name)
SELECT DISTINCT UPPER(REPLACE(TRIM(ministry), ' ', '_')), TRIM(ministry)
FROM users
WHERE ministry IS NOT NULL AND TRIM(ministry) <> ''
ON DUPLICATE KEY UPDATE name = VALUES(name);

ALTER TABLE users
  MODIFY COLUMN role ENUM('ADMIN','MENTRI','ADMIN_MEDINFO') NOT NULL,
  ADD COLUMN ministry_id BIGINT UNSIGNED NULL AFTER role;

UPDATE users AS user
LEFT JOIN ministries AS ministry
  ON ministry.code = UPPER(REPLACE(TRIM(user.ministry), ' ', '_'))
SET user.ministry_id = ministry.id;

UPDATE users
SET role = 'ADMIN_MEDINFO'
WHERE role = 'ADMIN' OR UPPER(COALESCE(ministry, '')) = 'MEDINFO';

UPDATE users SET role = 'ADMIN' WHERE role = 'MENTRI';

ALTER TABLE users
  MODIFY COLUMN role ENUM('ADMIN','ADMIN_MEDINFO') NOT NULL,
  ADD KEY idx_users_ministry_id (ministry_id),
  ADD CONSTRAINT fk_users_ministry FOREIGN KEY (ministry_id) REFERENCES ministries(id) ON DELETE SET NULL;

ALTER TABLE content_submissions
  MODIFY COLUMN status ENUM(
    'PENDING','IN_REVIEW','APPROVED','REJECTED',
    'DRAFT','SUBMITTED','PENDING_REVIEW','REVISION_REQUIRED',
    'REVISION_SUBMITTED','SCHEDULED','PUBLISHED'
  ) NOT NULL DEFAULT 'DRAFT',
  ADD COLUMN request_code VARCHAR(30) NULL AFTER id,
  ADD COLUMN service_type ENUM('CONTENT','ARTICLE') NOT NULL DEFAULT 'CONTENT' AFTER ministry,
  ADD COLUMN content_format VARCHAR(50) NULL AFTER service_type,
  ADD COLUMN submitter_name VARCHAR(100) NOT NULL DEFAULT '' AFTER submitter_id,
  ADD COLUMN submitter_phone VARCHAR(30) NULL AFTER submitter_name,
  ADD COLUMN ministry_id BIGINT UNSIGNED NULL AFTER ministry,
  ADD COLUMN song_title VARCHAR(180) NULL AFTER add_song,
  ADD COLUMN song_artist VARCHAR(180) NULL AFTER song_title,
  ADD COLUMN song_start_seconds INT UNSIGNED NULL AFTER song_artist,
  ADD COLUMN song_end_seconds INT UNSIGNED NULL AFTER song_start_seconds,
  ADD COLUMN documentation_drive_link VARCHAR(500) NULL AFTER article_drive_link,
  ADD COLUMN required_information TEXT NULL AFTER documentation_drive_link,
  ADD COLUMN confirmed_publish_at DATETIME(3) NULL AFTER deadline,
  ADD COLUMN submitted_at DATETIME(3) NULL AFTER confirmed_publish_at;

UPDATE content_submissions AS submission
JOIN users AS submitter ON submitter.id = submission.submitter_id
LEFT JOIN ministries AS ministry
  ON ministry.code = UPPER(REPLACE(TRIM(submission.ministry), ' ', '_'))
SET submission.submitter_name = submitter.name,
    submission.submitter_phone = submitter.phone,
    submission.ministry_id = ministry.id,
    submission.service_type = IF(submission.submission_type = 'ARTIKEL', 'ARTICLE', 'CONTENT'),
    submission.content_format = CASE
      WHEN submission.submission_type = 'ARTIKEL' THEN NULL
      WHEN submission.submission_type = 'INSTASTORY' THEN 'INSTASTORY'
      ELSE 'FEEDS_REELS_LEGACY'
    END,
    submission.status = CASE
      WHEN submission.status IN ('PENDING', 'IN_REVIEW') THEN 'PENDING_REVIEW'
      ELSE submission.status
    END,
    submission.submitted_at = submission.created_at,
    submission.request_code = CONCAT('MED-', YEAR(submission.created_at), '-', LPAD(submission.id, 6, '0'));

ALTER TABLE content_submissions
  MODIFY COLUMN status ENUM(
    'DRAFT','SUBMITTED','PENDING_REVIEW','REVISION_REQUIRED',
    'REVISION_SUBMITTED','APPROVED','SCHEDULED','PUBLISHED','REJECTED'
  ) NOT NULL DEFAULT 'DRAFT',
  ADD UNIQUE KEY idx_content_submissions_request_code (request_code),
  ADD KEY idx_content_submissions_service_type (service_type),
  ADD KEY idx_content_submissions_ministry_id (ministry_id),
  ADD CONSTRAINT fk_content_submissions_ministry FOREIGN KEY (ministry_id) REFERENCES ministries(id) ON DELETE SET NULL;

CREATE TABLE content_submission_status_histories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  submission_id BIGINT UNSIGNED NOT NULL,
  actor_id BIGINT UNSIGNED NULL,
  from_status VARCHAR(40) NULL,
  to_status VARCHAR(40) NOT NULL,
  note TEXT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_content_history_submission (submission_id, created_at),
  KEY idx_content_history_actor (actor_id),
  CONSTRAINT fk_content_history_submission FOREIGN KEY (submission_id) REFERENCES content_submissions(id) ON DELETE CASCADE,
  CONSTRAINT fk_content_history_actor FOREIGN KEY (actor_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO content_submission_status_histories (submission_id, actor_id, from_status, to_status, note, created_at)
SELECT id, submitter_id, NULL, status, 'Status awal dimigrasikan', created_at
FROM content_submissions;

CREATE TABLE content_submission_attachments (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  submission_id BIGINT UNSIGNED NULL,
  uploaded_by BIGINT UNSIGNED NOT NULL,
  imagekit_file_id VARCHAR(100) NOT NULL,
  purpose ENUM('FINAL_MEDIA','BRIEF_DOCUMENT') NOT NULL,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(1000) NOT NULL,
  mime_type VARCHAR(120) NOT NULL,
  size_bytes BIGINT UNSIGNED NOT NULL,
  status ENUM('STAGED','ATTACHED') NOT NULL DEFAULT 'STAGED',
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_content_attachments_file_id (imagekit_file_id),
  KEY idx_content_attachments_submission (submission_id),
  KEY idx_content_attachments_uploader (uploaded_by, status),
  CONSTRAINT fk_content_attachments_submission FOREIGN KEY (submission_id) REFERENCES content_submissions(id) ON DELETE CASCADE,
  CONSTRAINT fk_content_attachments_uploader FOREIGN KEY (uploaded_by) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE letter_submissions
  MODIFY COLUMN status ENUM(
    'PENDING','IN_REVIEW','APPROVED','REJECTED',
    'DRAFT','SUBMITTED','PENDING_REVIEW','REVISION_REQUIRED',
    'REVISION_SUBMITTED','COMPLETED'
  ) NOT NULL DEFAULT 'DRAFT',
  ADD COLUMN request_code VARCHAR(30) NULL AFTER id,
  ADD COLUMN submitter_name VARCHAR(100) NOT NULL DEFAULT '' AFTER submitter_id,
  ADD COLUMN submitter_phone VARCHAR(30) NULL AFTER submitter_name,
  ADD COLUMN ministry_id BIGINT UNSIGNED NULL AFTER ministry,
  ADD COLUMN submitted_at DATETIME(3) NULL AFTER deadline;

UPDATE letter_submissions AS submission
JOIN users AS submitter ON submitter.id = submission.submitter_id
LEFT JOIN ministries AS ministry
  ON ministry.code = UPPER(REPLACE(TRIM(submission.ministry), ' ', '_'))
SET submission.submitter_name = submitter.name,
    submission.submitter_phone = submitter.phone,
    submission.ministry_id = ministry.id,
    submission.status = CASE
      WHEN submission.status IN ('PENDING', 'IN_REVIEW') THEN 'PENDING_REVIEW'
      ELSE submission.status
    END,
    submission.submitted_at = submission.created_at,
    submission.request_code = CONCAT('SUR-', YEAR(submission.created_at), '-', LPAD(submission.id, 6, '0'));

ALTER TABLE letter_submissions
  MODIFY COLUMN status ENUM(
    'DRAFT','SUBMITTED','PENDING_REVIEW','REVISION_REQUIRED',
    'REVISION_SUBMITTED','APPROVED','REJECTED','COMPLETED'
  ) NOT NULL DEFAULT 'DRAFT',
  ADD UNIQUE KEY idx_letter_submissions_request_code (request_code),
  ADD KEY idx_letter_submissions_ministry_id (ministry_id),
  ADD CONSTRAINT fk_letter_submissions_ministry FOREIGN KEY (ministry_id) REFERENCES ministries(id) ON DELETE SET NULL;

CREATE TABLE letter_submission_status_histories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  submission_id BIGINT UNSIGNED NOT NULL,
  actor_id BIGINT UNSIGNED NULL,
  from_status VARCHAR(40) NULL,
  to_status VARCHAR(40) NOT NULL,
  note TEXT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_letter_history_submission (submission_id, created_at),
  KEY idx_letter_history_actor (actor_id),
  CONSTRAINT fk_letter_history_submission FOREIGN KEY (submission_id) REFERENCES letter_submissions(id) ON DELETE CASCADE,
  CONSTRAINT fk_letter_history_actor FOREIGN KEY (actor_id) REFERENCES users(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO letter_submission_status_histories (submission_id, actor_id, from_status, to_status, note, created_at)
SELECT id, submitter_id, NULL, status, 'Status awal dimigrasikan', created_at
FROM letter_submissions;

CREATE TABLE media_submission_settings (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  service_type ENUM('CONTENT','ARTICLE') NOT NULL,
  sop_url VARCHAR(1000) NULL,
  ministry_template_url VARCHAR(1000) NULL,
  brief_template_url VARCHAR(1000) NULL,
  caption_template_url VARCHAR(1000) NULL,
  pic_name VARCHAR(120) NULL,
  pic_whatsapp VARCHAR(30) NULL,
  terms_json JSON NOT NULL,
  minimum_lead_days INT UNSIGNED NOT NULL,
  publish_time_start VARCHAR(5) NOT NULL DEFAULT '08:00',
  publish_time_end VARCHAR(5) NOT NULL DEFAULT '17:00',
  slot_interval_minutes INT UNSIGNED NOT NULL DEFAULT 30,
  daily_capacity INT UNSIGNED NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_media_settings_service_type (service_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO media_submission_settings (
  service_type, terms_json, minimum_lead_days
) VALUES
  (
    'CONTENT',
    JSON_ARRAY(
      'Pengajuan tidak lengkap dapat menghambat proses peninjauan dan publikasi.',
      'Caption dan seluruh materi harus benar serta dapat dipertanggungjawabkan.',
      'Link Canva harus memberikan akses pengeditan kepada tim Medinfo.',
      'Jadwal publikasi menyesuaikan antrean, urgensi, dan kebijakan Medinfo.'
    ),
    7
  ),
  (
    'ARTICLE',
    JSON_ARRAY(
      'Artikel diawali tagline #CeritaHariIni dan mengikuti struktur JELITA.',
      'Naskah memiliki headline, kepala berita, tubuh berita, dan ekor berita.',
      'Panjang artikel 300 sampai 500 kata dan akan melalui penyuntingan Medinfo.',
      'Publikasi paling cepat tiga hari setelah artikel diterima.'
    ),
    3
  );

CREATE TABLE publication_blackouts (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  service_type ENUM('CONTENT','ARTICLE') NULL,
  blackout_date DATE NOT NULL,
  reason VARCHAR(255) NULL,
  created_by BIGINT UNSIGNED NOT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_publication_blackout (service_type, blackout_date),
  CONSTRAINT fk_publication_blackout_creator FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE request_sequences (
  prefix VARCHAR(10) NOT NULL,
  year SMALLINT UNSIGNED NOT NULL,
  last_number BIGINT UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (prefix, year)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE articles
  ADD COLUMN source_submission_id BIGINT UNSIGNED NULL AFTER author_id,
  ADD UNIQUE KEY idx_articles_source_submission (source_submission_id),
  ADD CONSTRAINT fk_articles_source_submission FOREIGN KEY (source_submission_id) REFERENCES content_submissions(id) ON DELETE SET NULL;

CREATE TABLE notification_outbox (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  recipient VARCHAR(30) NOT NULL,
  event_type VARCHAR(80) NOT NULL,
  payload JSON NOT NULL,
  status ENUM('PENDING','SENT','FAILED') NOT NULL DEFAULT 'PENDING',
  attempts INT UNSIGNED NOT NULL DEFAULT 0,
  available_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  sent_at DATETIME(3) NULL,
  last_error TEXT NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_notification_outbox_pending (status, available_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
