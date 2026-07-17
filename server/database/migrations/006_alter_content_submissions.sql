ALTER TABLE content_submissions
  DROP COLUMN platform;

ALTER TABLE content_submissions
  ADD COLUMN title VARCHAR(255) NOT NULL DEFAULT '' AFTER submission_type,
  ADD COLUMN add_song VARCHAR(255) NULL AFTER title,
  ADD COLUMN additional_notes TEXT NULL AFTER caption,
  ADD COLUMN publish_date DATE NULL AFTER additional_notes,
  ADD COLUMN publish_time VARCHAR(5) NULL AFTER publish_date,
  ADD COLUMN design_drive_link VARCHAR(500) NULL AFTER publish_time,
  ADD COLUMN canva_link VARCHAR(500) NULL AFTER design_drive_link,
  ADD COLUMN article_drive_link VARCHAR(500) NULL AFTER canva_link;

ALTER TABLE content_submissions
  MODIFY COLUMN deadline DATETIME NULL;

ALTER TABLE content_submissions
  CHANGE COLUMN brief_file brief_link VARCHAR(500) NULL;
