INSERT INTO media_assets (
    imagekit_file_id,
    url,
    name,
    alt_text,
    mime_type,
    size_bytes,
    purpose,
    status
) VALUES (
    'local-member-seniiora-puput',
    'https://res.cloudinary.com/rvn0bdia/image/upload/v1795411756/PUPUT_DIRJEN_SENIORA.png',
    'PUPUT_DIRJEN_SENIORA.png',
    'Muhammad Fadhil Saputra',
    'image/png',
    0,
    'ORGANIZATION_MEMBER_PHOTO',
    'ACTIVE'
)
ON DUPLICATE KEY UPDATE
    url = VALUES(url),
    name = VALUES(name),
    alt_text = VALUES(alt_text),
    purpose = VALUES(purpose),
    status = VALUES(status),
    updated_at = NOW();

UPDATE organization_members AS profile
JOIN ministries AS ministry ON ministry.id = profile.ministry_id
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
SET profile.is_active = FALSE
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND profile.name = 'Nurnajmi Dwihesti';

UPDATE organization_members AS profile
JOIN ministries AS ministry ON ministry.id = profile.ministry_id
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
SET profile.name = 'Muhammad Fadhil Saputra',
    profile.position = 'Direktur Jenderal Seni dan Olahraga',
    profile.position_type = 'DIRJEN',
    profile.is_active = TRUE,
    profile.display_order = 2,
    profile.photo_media_id = (
        SELECT id FROM media_assets
        WHERE imagekit_file_id = 'local-member-seniiora-puput'
        LIMIT 1
    )
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND profile.name IN ('Puput', 'M. Fadhil');

INSERT INTO organization_members (
    ministry_id,
    name,
    position,
    position_type,
    photo_media_id,
    display_order,
    is_leader,
    is_active
)
SELECT
    ministry.id,
    'Muhammad Fadhil Saputra',
    'Direktur Jenderal Seni dan Olahraga',
    'DIRJEN',
    (
        SELECT id FROM media_assets
        WHERE imagekit_file_id = 'local-member-seniiora-puput'
        LIMIT 1
    ),
    2,
    FALSE,
    TRUE
FROM ministries AS ministry
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND NOT EXISTS (
      SELECT 1
      FROM organization_members AS existing
      WHERE existing.ministry_id = ministry.id
        AND existing.name = 'Muhammad Fadhil Saputra'
        AND existing.is_active = TRUE
  );
