-- ============================================================
-- DEMO WORK PROGRAM DOCUMENTATIONS - KABINET CERITA LOKA 2026
--
-- Depends on:
--   000_ministries.sql
--   009_demo_program_kerja.sql
--
-- Scope:
--   108 demo work programs
--   3 documentation items per program
--   Expected total: 324 documentation rows
--
-- Strategy:
--   - Resolve work_program_id from ministry code + program slug.
--   - Create deterministic demo MediaAsset rows.
--   - Reuse each ministry logo URL as a safe local placeholder image.
--   - Link the MediaAsset rows into work_program_documentations.
--   - Documentation #1 becomes the program cover only when cover_media_id is NULL.
--   - Idempotent: safe to run repeatedly.
--
-- Target: MySQL 8
-- ============================================================

SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci;

START TRANSACTION;

-- ============================================================
-- 1. EXACT TARGET PROGRAMS FROM 009_demo_program_kerja.sql
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_demo_program_documentation_targets;

CREATE TEMPORARY TABLE seed_demo_program_documentation_targets (
    unit_code VARCHAR(50) NOT NULL,
    slug      VARCHAR(200) NOT NULL,
    PRIMARY KEY (unit_code, slug)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

INSERT INTO seed_demo_program_documentation_targets (
    unit_code,
    slug
) VALUES
    ('ADKESMA', 'adkesma-airlangga-mendengar'),
    ('ADKESMA', 'adkesma-pusat-advokasi-mahasiswa'),
    ('ADKESMA', 'adkesma-kawan-ukt'),
    ('ADKESMA', 'adkesma-beasiswa-update'),
    ('ADKESMA', 'adkesma-posko-akademik'),
    ('ADKESMA', 'adkesma-ruang-aman-kampus'),
    ('ADKESMA', 'adkesma-advokasi-inklusif'),
    ('ADKESMA', 'adkesma-survei-kesejahteraan-mahasiswa'),
    ('PENGPROF', 'pengprof-career-preparation-class'),
    ('PENGPROF', 'pengprof-cv-clinic'),
    ('PENGPROF', 'pengprof-linkedin-lab'),
    ('PENGPROF', 'pengprof-internship-hub'),
    ('PENGPROF', 'pengprof-alumni-insight'),
    ('PENGPROF', 'pengprof-professional-mentoring'),
    ('PENGPROF', 'pengprof-skill-development-workshop'),
    ('PENGPROF', 'pengprof-company-visit'),
    ('PENGPROF', 'pengprof-airlangga-career-expo'),
    ('PSDM', 'psdm-cerita-kader'),
    ('PSDM', 'psdm-leadership-camp'),
    ('PSDM', 'psdm-upgrading-internal'),
    ('PSDM', 'psdm-talent-mapping-mahasiswa'),
    ('LH', 'lh-green-campus-movement'),
    ('LH', 'lh-bank-sampah-mahasiswa'),
    ('LH', 'lh-airlangga-earth-week'),
    ('KEMENKES', 'kemenkes-health-check-day'),
    ('KEMENKES', 'kemenkes-mental-health-talk'),
    ('KEMENKES', 'kemenkes-konseling-sebaya'),
    ('KEMENKES', 'kemenkes-donor-darah-airlangga'),
    ('KEMENKES', 'kemenkes-kampus-tanpa-rokok'),
    ('KEMENKES', 'kemenkes-nutrition-week'),
    ('KEMENKES', 'kemenkes-first-aid-training'),
    ('KEMENKES', 'kemenkes-healthy-movement'),
    ('KEMENKES', 'kemenkes-inclusive-health-campaign'),
    ('KEMENKES', 'kemenkes-health-information-center'),
    ('PENGMAS', 'pengmas-desa-mitra'),
    ('PENGMAS', 'pengmas-cerita-loka-mengabdi'),
    ('PENGMAS', 'pengmas-volunteer-academy'),
    ('PENGMAS', 'pengmas-bakti-pendidikan'),
    ('PENGMAS', 'pengmas-pasar-berdaya'),
    ('PENGMAS', 'pengmas-aksi-tanggap-bencana'),
    ('PENGMAS', 'pengmas-community-impact-report'),
    ('HUBLU', 'hublu-external-partnership-forum'),
    ('HUBLU', 'hublu-campus-visit'),
    ('HUBLU', 'hublu-international-student-connect'),
    ('HUBLU', 'hublu-partnership-database'),
    ('HUBLU', 'hublu-delegation-support'),
    ('HUBLU', 'hublu-stakeholder-gathering'),
    ('MEDINFO', 'medinfo-cerita-hari-ini'),
    ('MEDINFO', 'medinfo-media-request-center'),
    ('MEDINFO', 'medinfo-visual-branding-kit'),
    ('MEDINFO', 'medinfo-bem-podcast'),
    ('MEDINFO', 'medinfo-social-media-insight'),
    ('MEDINFO', 'medinfo-website-kabinet-cerita-loka'),
    ('MEDINFO', 'medinfo-documentation-hub'),
    ('MEDINFO', 'medinfo-media-training'),
    ('MEDINFO', 'medinfo-annual-media-report'),
    ('SINEMA', 'sinema-airlangga-collaboration-forum'),
    ('SINEMA', 'sinema-ormawa-connect'),
    ('SINEMA', 'sinema-festival-cerita-loka'),
    ('SINEMA', 'sinema-student-community-day'),
    ('SINEMA', 'sinema-collaborative-project-incubator'),
    ('KKG', 'kkg-safe-space-campaign'),
    ('KKG', 'kkg-gender-equality-class'),
    ('KKG', 'kkg-support-circle'),
    ('KKG', 'kkg-inclusive-campus-audit'),
    ('SOSPOL', 'sospol-kajian-isu-strategis'),
    ('SOSPOL', 'sospol-airlangga-policy-forum'),
    ('SOSPOL', 'sospol-aksi-mahasiswa'),
    ('SOSPOL', 'sospol-sekolah-politik'),
    ('SOSPOL', 'sospol-ruang-demokrasi'),
    ('SOSPOL', 'sospol-policy-brief-series'),
    ('SOSPOL', 'sospol-pemilu-kampus-watch'),
    ('SOSPOL', 'sospol-jejaring-gerakan-mahasiswa'),
    ('EDIGPRO', 'edigpro-digital-business-bootcamp'),
    ('EDIGPRO', 'edigpro-creative-product-lab'),
    ('EDIGPRO', 'edigpro-umkm-goes-digital'),
    ('EDIGPRO', 'edigpro-startup-ideation-day'),
    ('EDIGPRO', 'edigpro-student-marketplace'),
    ('EDIGPRO', 'edigpro-financial-literacy-class'),
    ('EDIGPRO', 'edigpro-creative-business-expo'),
    ('RISKEL', 'riskel-research-academy'),
    ('RISKEL', 'riskel-scientific-writing-clinic'),
    ('RISKEL', 'riskel-pkm-center'),
    ('RISKEL', 'riskel-data-analysis-workshop'),
    ('RISKEL', 'riskel-research-mentoring'),
    ('RISKEL', 'riskel-journal-club'),
    ('RISKEL', 'riskel-innovation-challenge'),
    ('RISKEL', 'riskel-research-funding-information'),
    ('RISKEL', 'riskel-conference-preparation-class'),
    ('RISKEL', 'riskel-airlangga-research-showcase'),
    ('SENIIORA', 'seniiora-airlangga-art-festival'),
    ('SENIIORA', 'seniiora-cerita-loka-sports-league'),
    ('SENIIORA', 'seniiora-creative-class'),
    ('SENIIORA', 'seniiora-talent-showcase'),
    ('SENIIORA', 'seniiora-community-sports-day'),
    ('SENIIORA', 'seniiora-culture-appreciation-week'),
    ('BENKAB', 'benkab-budget-planning-clinic'),
    ('BENKAB', 'benkab-treasury-dashboard'),
    ('BENKAB', 'benkab-financial-accountability-week'),
    ('PAK', 'pak-organizational-performance-review'),
    ('PAK', 'pak-sop-center'),
    ('PAK', 'pak-internal-audit-kabinet'),
    ('PAK', 'pak-human-resource-data-center'),
    ('PAK', 'pak-cabinet-evaluation-forum'),
    ('SESKAB', 'seskab-administration-center'),
    ('SESKAB', 'seskab-correspondence-clinic'),
    ('SESKAB', 'seskab-archive-digitalization'),
    ('SESKAB', 'seskab-cabinet-meeting-management');

-- ============================================================
-- 2. DOCUMENTATION SLOTS
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_demo_program_documentation_slots;

CREATE TEMPORARY TABLE seed_demo_program_documentation_slots (
    slot_no       TINYINT UNSIGNED NOT NULL,
    title_template VARCHAR(100) NOT NULL,
    caption_template VARCHAR(255) NOT NULL,
    PRIMARY KEY (slot_no)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

INSERT INTO seed_demo_program_documentation_slots (
    slot_no,
    title_template,
    caption_template
) VALUES
    (
        1,
        'Dokumentasi Utama',
        'Dokumentasi utama pelaksanaan program kerja.'
    ),
    (
        2,
        'Dokumentasi Kegiatan',
        'Dokumentasi aktivitas dan partisipasi dalam program kerja.'
    ),
    (
        3,
        'Dokumentasi Penutup',
        'Dokumentasi penutup, hasil, dan evaluasi program kerja.'
    );

-- ============================================================
-- 3. CREATE DEMO MEDIA ASSETS
--
-- imagekit_file_id dibuat deterministik dari slug + slot,
-- sehingga seeder dapat dijalankan ulang tanpa duplikasi.
--
-- URL menggunakan logo kementerian yang sudah tersedia.
-- Kalau relasi logo belum tersedia, fallback ke path lokal
-- /images/Kementerian/<CODE>.png.
-- ============================================================

INSERT INTO media_assets (
    uploaded_by,
    imagekit_file_id,
    file_path,
    url,
    thumbnail_url,
    name,
    alt_text,
    caption,
    mime_type,
    size_bytes,
    width,
    height,
    purpose,
    status,
    created_at,
    updated_at
)
SELECT
    NULL AS uploaded_by,

    CONCAT(
        'demo-work-doc-',
        MD5(CONCAT(program.slug, ':', slot.slot_no))
    ) AS imagekit_file_id,

    COALESCE(
        ministry_logo.file_path,
        CONCAT('/images/Kementerian/', ministry.code, '.png')
    ) AS file_path,

    COALESCE(
        ministry_logo.url,
        CONCAT('/images/Kementerian/', ministry.code, '.png')
    ) AS url,

    COALESCE(
        ministry_logo.thumbnail_url,
        ministry_logo.url,
        CONCAT('/images/Kementerian/', ministry.code, '.png')
    ) AS thumbnail_url,

    CONCAT(
        program.name,
        ' - ',
        slot.title_template,
        ' ',
        slot.slot_no
    ) AS name,

    CONCAT(
        slot.title_template,
        ' ',
        program.name,
        ' dari ',
        ministry.name
    ) AS alt_text,

    CONCAT(
        slot.caption_template,
        ' ',
        program.name,
        ' - ',
        ministry.name,
        '.'
    ) AS caption,

    COALESCE(ministry_logo.mime_type, 'image/png') AS mime_type,
    COALESCE(ministry_logo.size_bytes, 0) AS size_bytes,
    ministry_logo.width,
    ministry_logo.height,
    'WORK_PROGRAM_DOCUMENTATION_DEMO' AS purpose,
    'ACTIVE' AS status,
    NOW(),
    NOW()

FROM seed_demo_program_documentation_targets AS target

INNER JOIN ministries AS ministry
    ON ministry.code COLLATE utf8mb4_unicode_ci
     = target.unit_code COLLATE utf8mb4_unicode_ci

INNER JOIN work_programs AS program
    ON program.ministry_id = ministry.id
   AND program.slug COLLATE utf8mb4_unicode_ci
     = target.slug COLLATE utf8mb4_unicode_ci

CROSS JOIN seed_demo_program_documentation_slots AS slot

LEFT JOIN media_assets AS ministry_logo
    ON ministry_logo.id = ministry.logo_media_id

WHERE
    ministry.unit_type = 'KEMENTERIAN'
    AND ministry.is_active = TRUE

ON DUPLICATE KEY UPDATE
    file_path = VALUES(file_path),
    url = VALUES(url),
    thumbnail_url = VALUES(thumbnail_url),
    name = VALUES(name),
    alt_text = VALUES(alt_text),
    caption = VALUES(caption),
    mime_type = VALUES(mime_type),
    size_bytes = VALUES(size_bytes),
    width = VALUES(width),
    height = VALUES(height),
    purpose = VALUES(purpose),
    status = VALUES(status),
    updated_at = NOW();

-- ============================================================
-- 4. LINK MEDIA ASSETS TO WORK PROGRAM DOCUMENTATIONS
-- ============================================================

INSERT INTO work_program_documentations (
    work_program_id,
    media_asset_id,
    title,
    caption,
    taken_at,
    display_order,
    is_cover,
    created_at,
    updated_at
)
SELECT
    program.id AS work_program_id,
    documentation_media.id AS media_asset_id,

    CONCAT(
        slot.title_template,
        ' ',
        program.name
    ) AS title,

    CONCAT(
        slot.caption_template,
        ' Program ',
        program.name,
        ' oleh ',
        ministry.name,
        '.'
    ) AS caption,

    TIMESTAMP(
        DATE_ADD(
            COALESCE(program.start_date, '2026-01-01'),
            INTERVAL (slot.slot_no - 1) DAY
        ),
        '09:00:00'
    ) AS taken_at,

    slot.slot_no AS display_order,
    (slot.slot_no = 1) AS is_cover,
    NOW(),
    NOW()

FROM seed_demo_program_documentation_targets AS target

INNER JOIN ministries AS ministry
    ON ministry.code COLLATE utf8mb4_unicode_ci
     = target.unit_code COLLATE utf8mb4_unicode_ci

INNER JOIN work_programs AS program
    ON program.ministry_id = ministry.id
   AND program.slug COLLATE utf8mb4_unicode_ci
     = target.slug COLLATE utf8mb4_unicode_ci

CROSS JOIN seed_demo_program_documentation_slots AS slot

INNER JOIN media_assets AS documentation_media
    ON documentation_media.imagekit_file_id = CONCAT(
        'demo-work-doc-',
        MD5(CONCAT(program.slug, ':', slot.slot_no))
    )

WHERE
    ministry.unit_type = 'KEMENTERIAN'
    AND ministry.is_active = TRUE

ON DUPLICATE KEY UPDATE
    title = VALUES(title),
    caption = VALUES(caption),
    taken_at = VALUES(taken_at),
    display_order = VALUES(display_order),
    is_cover = VALUES(is_cover),
    updated_at = NOW();

-- ============================================================
-- 5. SET FIRST DOCUMENTATION AS PROGRAM COVER
--
-- Existing real cover is preserved.
-- Only programs whose cover_media_id is still NULL are updated.
-- ============================================================

UPDATE work_programs AS program

INNER JOIN ministries AS ministry
    ON ministry.id = program.ministry_id

INNER JOIN seed_demo_program_documentation_targets AS target
    ON target.unit_code COLLATE utf8mb4_unicode_ci
     = ministry.code COLLATE utf8mb4_unicode_ci
   AND target.slug COLLATE utf8mb4_unicode_ci
     = program.slug COLLATE utf8mb4_unicode_ci

INNER JOIN media_assets AS documentation_cover
    ON documentation_cover.imagekit_file_id = CONCAT(
        'demo-work-doc-',
        MD5(CONCAT(program.slug, ':1'))
    )

SET
    program.cover_media_id = documentation_cover.id,
    program.updated_at = NOW()

WHERE
    program.cover_media_id IS NULL;

-- ============================================================
-- 6. CLEANUP
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_demo_program_documentation_slots;
DROP TEMPORARY TABLE IF EXISTS seed_demo_program_documentation_targets;

COMMIT;

-- ============================================================
-- 7. VERIFICATION
-- ============================================================

SELECT
    COUNT(*) AS total_demo_documentations
FROM work_program_documentations AS documentation
INNER JOIN media_assets AS media
    ON media.id = documentation.media_asset_id
WHERE media.purpose = 'WORK_PROGRAM_DOCUMENTATION_DEMO';

SELECT
    ministry.code AS ministry_code,
    ministry.name AS ministry_name,
    COUNT(DISTINCT program.id) AS total_programs,
    COUNT(documentation.id) AS total_documentations
FROM ministries AS ministry
INNER JOIN work_programs AS program
    ON program.ministry_id = ministry.id
LEFT JOIN work_program_documentations AS documentation
    ON documentation.work_program_id = program.id
WHERE
    (ministry.code, program.slug) IN (
        SELECT
            target.unit_code,
            target.slug
        FROM (
            SELECT
                ministry_inner.code AS unit_code,
                program_inner.slug AS slug
            FROM ministries AS ministry_inner
            INNER JOIN work_programs AS program_inner
                ON program_inner.ministry_id = ministry_inner.id
            INNER JOIN work_program_documentations AS documentation_inner
                ON documentation_inner.work_program_id = program_inner.id
            INNER JOIN media_assets AS media_inner
                ON media_inner.id = documentation_inner.media_asset_id
               AND media_inner.purpose = 'WORK_PROGRAM_DOCUMENTATION_DEMO'
            GROUP BY
                ministry_inner.code,
                program_inner.slug
        ) AS target
    )
GROUP BY
    ministry.id,
    ministry.code,
    ministry.name,
    ministry.display_order
ORDER BY
    ministry.parent_id ASC,
    ministry.display_order ASC;

SELECT
    program.slug,
    program.name,
    COUNT(documentation.id) AS documentation_count,
    SUM(documentation.is_cover = TRUE) AS cover_count
FROM work_programs AS program
INNER JOIN work_program_documentations AS documentation
    ON documentation.work_program_id = program.id
INNER JOIN media_assets AS media
    ON media.id = documentation.media_asset_id
   AND media.purpose = 'WORK_PROGRAM_DOCUMENTATION_DEMO'
GROUP BY
    program.id,
    program.slug,
    program.name
HAVING
    documentation_count <> 3
    OR cover_count <> 1
ORDER BY
    program.slug ASC;
