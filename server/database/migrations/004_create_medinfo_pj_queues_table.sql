CREATE TABLE IF NOT EXISTS medinfo_pj_queues (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  user_id BIGINT UNSIGNED NOT NULL,
  position INT NOT NULL,
  is_current BOOLEAN NOT NULL DEFAULT FALSE,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_medinfo_pj_queues_user_id (user_id),
  KEY idx_medinfo_pj_queues_position (position),
  KEY idx_medinfo_pj_queues_is_current (is_current),
  CONSTRAINT fk_medinfo_pj_queues_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
