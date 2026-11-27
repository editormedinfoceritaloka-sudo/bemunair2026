-- Remove the former appointee's seeded photo if an earlier seed assigned it.
UPDATE organization_members AS profile
JOIN ministries AS ministry ON ministry.id = profile.ministry_id
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
JOIN media_assets AS photo ON photo.id = profile.photo_media_id
SET profile.photo_media_id = NULL
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND profile.name = 'Nurnajmi Dwihesti'
  AND photo.imagekit_file_id = 'local-member-seniiora-puput';
