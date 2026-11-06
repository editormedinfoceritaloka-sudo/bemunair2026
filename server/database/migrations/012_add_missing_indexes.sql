-- Indexes for hot lookup paths. Existing migrations 001-011 are not modified.
ALTER TABLE work_programs ADD KEY idx_work_programs_slug (slug);
ALTER TABLE content_submissions ADD KEY idx_content_submissions_status_deadline (status, deadline);
ALTER TABLE letter_submissions ADD KEY idx_letter_submissions_status_deadline (status, deadline);
ALTER TABLE media_assets ADD KEY idx_media_assets_status (status);