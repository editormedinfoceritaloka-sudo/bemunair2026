CREATE TABLE IF NOT EXISTS articles (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  slug VARCHAR(255) NOT NULL,
  title VARCHAR(255) NOT NULL,
  excerpt VARCHAR(500) NULL,
  body LONGTEXT NOT NULL,
  cover_image VARCHAR(500) NULL,
  author_id BIGINT UNSIGNED NOT NULL,
  status ENUM('DRAFT','PUBLISHED') NOT NULL DEFAULT 'DRAFT',
  published_at DATETIME(3) NULL,
  created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_articles_slug (slug),
  KEY idx_articles_author_id (author_id),
  KEY idx_articles_status (status),
  KEY idx_articles_published_at (published_at),
  CONSTRAINT fk_articles_author FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
