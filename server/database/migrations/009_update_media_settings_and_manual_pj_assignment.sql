UPDATE media_submission_settings
SET
  sop_url = 'https://docs.google.com/document/d/1RgIdOLXwUJz6pRrktK2fLQ6srvCZn0y8/edit',
  ministry_template_url = 'https://linktr.ee/templateeditingkementerian',
  brief_template_url = 'https://docs.google.com/document/d/1bBIs7dcqL1OMSJWAoXZduy2c6r0MpvLIC84T5t3OZJ8/edit?tab=t.0',
  caption_template_url = NULL,
  minimum_lead_days = 7,
  publish_time_start = '08:00',
  publish_time_end = '17:00',
  slot_interval_minutes = 30,
  terms_json = JSON_ARRAY(
    'Pengajuan yang tidak lengkap dapat menghambat proses peninjauan dan publikasi konten.',
    'Pastikan seluruh informasi, caption, dan file yang dikirimkan benar serta dapat dipertanggungjawabkan.',
    'Link Canva wajib menggunakan akses Anyone with the link can edit.',
    'Jadwal publikasi disesuaikan dengan antrean, urgensi, dan kebijakan Medinfo.',
    'Kebutuhan mendesak harus dicantumkan pada informasi tambahan dan dikomunikasikan kepada PIC.'
  )
WHERE service_type = 'CONTENT';

UPDATE media_submission_settings
SET
  sop_url = 'https://docs.google.com/document/d/1RgIdOLXwUJz6pRrktK2fLQ6srvCZn0y8/edit',
  ministry_template_url = 'https://linktr.ee/templateeditingkementerian',
  brief_template_url = NULL,
  caption_template_url = 'https://docs.google.com/document/d/1--ki9SXtg1tTQQFFNQElnQ4brK1c8mMA6Ww540XCKzk/edit?tab=t.0',
  minimum_lead_days = 3,
  publish_time_start = '08:00',
  publish_time_end = '17:00',
  slot_interval_minutes = 30,
  terms_json = JSON_ARRAY(
    'Artikel wajib diawali tagline #CeritaHariIni dan mengikuti struktur JELITA.',
    'Struktur artikel terdiri dari headline, kepala berita, tubuh berita, dan ekor berita.',
    'Panjang artikel berkisar antara 300 sampai 500 kata.',
    'Artikel melalui proses penyuntingan oleh Kementerian Media dan Informasi.',
    'Artikel dipublikasikan paling cepat H+3 setelah naskah diterima.'
  )
WHERE service_type = 'ARTICLE';

CREATE TABLE content_submission_assignment_histories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  submission_id BIGINT UNSIGNED NOT NULL,
  actor_id BIGINT UNSIGNED NOT NULL,
  from_pj_id BIGINT UNSIGNED NULL,
  to_pj_id BIGINT UNSIGNED NOT NULL,
  note VARCHAR(255) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_content_assignment_submission (submission_id, created_at),
  CONSTRAINT fk_content_assignment_submission FOREIGN KEY (submission_id) REFERENCES content_submissions(id) ON DELETE CASCADE,
  CONSTRAINT fk_content_assignment_actor FOREIGN KEY (actor_id) REFERENCES users(id) ON DELETE RESTRICT,
  CONSTRAINT fk_content_assignment_from_pj FOREIGN KEY (from_pj_id) REFERENCES users(id) ON DELETE SET NULL,
  CONSTRAINT fk_content_assignment_to_pj FOREIGN KEY (to_pj_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE letter_submission_assignment_histories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  submission_id BIGINT UNSIGNED NOT NULL,
  actor_id BIGINT UNSIGNED NOT NULL,
  from_pj_id BIGINT UNSIGNED NULL,
  to_pj_id BIGINT UNSIGNED NOT NULL,
  note VARCHAR(255) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  KEY idx_letter_assignment_submission (submission_id, created_at),
  CONSTRAINT fk_letter_assignment_submission FOREIGN KEY (submission_id) REFERENCES letter_submissions(id) ON DELETE CASCADE,
  CONSTRAINT fk_letter_assignment_actor FOREIGN KEY (actor_id) REFERENCES users(id) ON DELETE RESTRICT,
  CONSTRAINT fk_letter_assignment_from_pj FOREIGN KEY (from_pj_id) REFERENCES users(id) ON DELETE SET NULL,
  CONSTRAINT fk_letter_assignment_to_pj FOREIGN KEY (to_pj_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
