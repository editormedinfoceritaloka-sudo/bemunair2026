UPDATE work_programs SET lifecycle_status = 'PLANNED' WHERE lifecycle_status = 'POSTPONED';
UPDATE work_programs SET lifecycle_status = 'COMPLETED' WHERE lifecycle_status = 'ARCHIVED';
UPDATE work_programs SET lifecycle_status = 'DRAFT' WHERE lifecycle_status = 'CANCELLED';
UPDATE work_programs SET lifecycle_status = 'DRAFT' WHERE is_published = 0 AND lifecycle_status <> 'DRAFT';
UPDATE work_programs SET is_published = 1 WHERE lifecycle_status <> 'DRAFT';
UPDATE work_programs SET is_published = 0 WHERE lifecycle_status = 'DRAFT';
UPDATE work_programs SET published_at = COALESCE(published_at, updated_at, created_at) WHERE is_published = 1 AND published_at IS NULL;
UPDATE work_programs SET published_at = NULL WHERE is_published = 0;
ALTER TABLE work_programs ADD CONSTRAINT chk_work_programs_status CHECK (lifecycle_status IN ('DRAFT', 'PLANNED', 'ONGOING', 'COMPLETED'));
