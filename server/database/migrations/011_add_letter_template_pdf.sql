ALTER TABLE letter_templates
  ADD COLUMN media_asset_id BIGINT UNSIGNED NULL AFTER body,
  ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE AFTER media_asset_id,
  ADD COLUMN display_order INT UNSIGNED NOT NULL DEFAULT 0 AFTER is_active,
  ADD KEY idx_letter_templates_media_asset (media_asset_id),
  ADD KEY idx_letter_templates_active_order (is_active, display_order),
  ADD CONSTRAINT fk_letter_templates_media_asset FOREIGN KEY (media_asset_id) REFERENCES media_assets(id) ON DELETE SET NULL;
