UPDATE work_programs wp
JOIN (
  SELECT id, ROW_NUMBER() OVER (PARTITION BY slug ORDER BY id) AS rn
  FROM work_programs
) dup ON dup.id = wp.id AND dup.rn > 1
SET wp.slug = CONCAT(wp.slug, '-', wp.id);

ALTER TABLE work_programs DROP INDEX idx_work_programs_slug;
ALTER TABLE work_programs ADD UNIQUE KEY idx_work_programs_slug (slug);
