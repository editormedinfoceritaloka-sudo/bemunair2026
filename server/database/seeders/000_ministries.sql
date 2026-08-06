-- ============================================================
-- SEEDER KABINET CERITA LOKA
--
-- Isi:
-- 1 CabinetTerm
-- 5 Menko
-- 1 unit induk BPII
-- 17 Kementerian
-- Logo kementerian dari /images/Kementerian/
-- Foto Menko dari /images/Menko/
-- Foto Menteri, Dirjen, Deputi, Kepala, dan Kabiro
--
-- Target: MySQL 8
-- ============================================================

START TRANSACTION;

-- ============================================================
-- 1. CABINET TERM: KABINET CERITA LOKA
-- ============================================================

INSERT INTO cabinet_terms (
    name,
    slug,
    tagline,
    description,
    logo_media_id,
    hero_media_id,
    period_start,
    period_end,
    is_active,
    is_published,
    published_at,
    meta_title,
    meta_description,
    created_at,
    updated_at
) VALUES (
    'Kabinet Cerita Loka',
    'cerita-loka-2026',
    'Mengukir Cerita, Memberi Makna untuk Loka',
    'Kabinet Cerita Loka merupakan kabinet organisasi mahasiswa yang bergerak melalui kolaborasi, pelayanan, pengembangan, dan kebermanfaatan.',
    NULL,
    NULL,
    '2026-01-01',
    '2026-12-31',
    TRUE,
    TRUE,
    NOW(),
    'Kabinet Cerita Loka',
    'Profil, struktur organisasi, kementerian, anggota, dan program kerja Kabinet Cerita Loka.',
    NOW(),
    NOW()
)
ON DUPLICATE KEY UPDATE
    id = LAST_INSERT_ID(id),
    name = VALUES(name),
    tagline = VALUES(tagline),
    description = VALUES(description),
    period_start = VALUES(period_start),
    period_end = VALUES(period_end),
    is_active = VALUES(is_active),
    is_published = VALUES(is_published),
    published_at = VALUES(published_at),
    meta_title = VALUES(meta_title),
    meta_description = VALUES(meta_description),
    updated_at = NOW();

SET @cabinet_term_id = LAST_INSERT_ID();

-- ============================================================
-- 2. TEMPORARY DATA UNIT ORGANISASI
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_organization_units;

CREATE TEMPORARY TABLE seed_organization_units (
    code              VARCHAR(50)  NOT NULL PRIMARY KEY,
    name              VARCHAR(120) NOT NULL,
    short_name        VARCHAR(80)  NULL,
    slug              VARCHAR(160) NOT NULL,
    unit_type         VARCHAR(20)  NOT NULL,
    parent_code       VARCHAR(50)  NULL,
    display_order     INT UNSIGNED NOT NULL DEFAULT 0,
    logo_asset_key    VARCHAR(100) NULL,
    logo_path         VARCHAR(1000) NULL
);

INSERT INTO seed_organization_units (
    code,
    name,
    short_name,
    slug,
    unit_type,
    parent_code,
    display_order,
    logo_asset_key,
    logo_path
) VALUES

-- ============================================================
-- 5 MENKO
-- Foto Menko juga dipakai sebagai LogoMedia unit Menko.
-- ============================================================

(
    'MENKO_KEMAHASISWAAN',
    'Kementerian Koordinator Kemahasiswaan',
    'Menko Kemahasiswaan',
    'menko-kemahasiswaan',
    'MENKO',
    NULL,
    1,
    'local-menko-kemahasiswaan',
    '/images/Menko/menko_kemahasiswaan_johanes_richard_darmawan.png'
),
(
    'MENKO_KEMASYARAKATAN',
    'Kementerian Koordinator Kemasyarakatan',
    'Menko Kemasyarakatan',
    'menko-kemasyarakatan',
    'MENKO',
    NULL,
    2,
    'local-menko-kemasyarakatan',
    '/images/Menko/menko_kemasyarakatan_muhammad_ziyaad_difa_ul_haq.png'
),
(
    'MENKO_KOMINFO',
    'Kementerian Koordinator Komunikasi dan Informasi',
    'Menko Kominfo',
    'menko-komunikasi-dan-informasi',
    'MENKO',
    NULL,
    3,
    'local-menko-kominfo',
    '/images/Menko/menko_kominfo_andrean_marcellino.png'
),
(
    'MENKO_PERGERAKAN',
    'Kementerian Koordinator Pergerakan',
    'Menko Pergerakan',
    'menko-pergerakan',
    'MENKO',
    NULL,
    4,
    'local-menko-pergerakan',
    '/images/Menko/menko_pergerakan_dewa_pranata_putra_pratama.png'
),
(
    'MENKO_PPK',
    'Kementerian Koordinator PPK',
    'Menko PPK',
    'menko-ppk',
    'MENKO',
    NULL,
    5,
    'local-menko-ppk',
    '/images/Menko/menko_ppk_nasua_okta_kusuma_dewi_rofiq.png'
),

-- ============================================================
-- INDUK BPII
-- Tidak ada foto Menko BPII pada daftar file.
-- ============================================================

(
    'BPII',
    'BPII',
    'BPII',
    'bpii',
    'BPII',
    NULL,
    6,
    NULL,
    NULL
),

-- ============================================================
-- KEMENTERIAN DI BAWAH MENKO KEMAHASISWAAN
-- ============================================================

(
    'ADKESMA',
    'Kementerian Advokasi dan Kesejahteraan Mahasiswa',
    'ADKESMA',
    'advokasi-dan-kesejahteraan-mahasiswa',
    'KEMENTERIAN',
    'MENKO_KEMAHASISWAAN',
    1,
    'local-logo-adkesma',
    '/images/Kementerian/ADKESMA.png'
),
(
    'PENGPROF',
    'Kementerian Pengembangan Profesi',
    'PENGPROF',
    'pengembangan-profesi',
    'KEMENTERIAN',
    'MENKO_KEMAHASISWAAN',
    2,
    'local-logo-pengprof',
    '/images/Kementerian/PENGPROF.png'
),
(
    'PSDM',
    'Kementerian Pengembangan Sumber Daya Mahasiswa',
    'PSDM',
    'pengembangan-sumber-daya-mahasiswa',
    'KEMENTERIAN',
    'MENKO_KEMAHASISWAAN',
    3,
    'local-logo-psdm',
    '/images/Kementerian/PSDM.png'
),

-- ============================================================
-- KEMENTERIAN DI BAWAH MENKO KEMASYARAKATAN
-- ============================================================

(
    'LH',
    'Kementerian Lingkungan Hidup',
    'LH',
    'lingkungan-hidup',
    'KEMENTERIAN',
    'MENKO_KEMASYARAKATAN',
    1,
    'local-logo-lh',
    '/images/Kementerian/LH.png'
),
(
    'KEMENKES',
    'Kementerian Kesehatan',
    'KEMENKES',
    'kementerian-kesehatan',
    'KEMENTERIAN',
    'MENKO_KEMASYARAKATAN',
    2,
    'local-logo-kemenkes',
    '/images/Kementerian/KEMENKES.png'
),
(
    'PENGMAS',
    'Kementerian Pengabdian Masyarakat',
    'PENGMAS',
    'pengabdian-masyarakat',
    'KEMENTERIAN',
    'MENKO_KEMASYARAKATAN',
    3,
    'local-logo-pengmas',
    '/images/Kementerian/PENGMAS.png'
),

-- ============================================================
-- KEMENTERIAN DI BAWAH MENKO KOMINFO
-- ============================================================

(
    'HUBLU',
    'Kementerian Hubungan Luar',
    'HUBLU',
    'hubungan-luar',
    'KEMENTERIAN',
    'MENKO_KOMINFO',
    1,
    'local-logo-hublu',
    '/images/Kementerian/HUBLU.png'
),
(
    'MEDINFO',
    'Kementerian Media dan Informasi',
    'MEDINFO',
    'media-dan-informasi',
    'KEMENTERIAN',
    'MENKO_KOMINFO',
    2,
    'local-logo-medinfo',
    '/images/Kementerian/MEDINFO.png'
),
(
    'SINEMA',
    'Kementerian Sinergitas Mahasiswa',
    'SINEMA',
    'sinergitas-mahasiswa',
    'KEMENTERIAN',
    'MENKO_KOMINFO',
    3,
    'local-logo-sinema',
    '/images/Kementerian/SINEMA.png'
),

-- ============================================================
-- KEMENTERIAN DI BAWAH MENKO PERGERAKAN
-- ============================================================

(
    'KKG',
    'Kementerian Keadilan dan Kesetaraan Gender',
    'KKG',
    'keadilan-dan-kesetaraan-gender',
    'KEMENTERIAN',
    'MENKO_PERGERAKAN',
    1,
    'local-logo-kkg',
    '/images/Kementerian/KKG.png'
),
(
    'SOSPOL',
    'Kementerian Sosial dan Politik',
    'SOSPOL',
    'sosial-dan-politik',
    'KEMENTERIAN',
    'MENKO_PERGERAKAN',
    2,
    'local-logo-sospol',
    '/images/Kementerian/SOSPOL.png'
),

-- ============================================================
-- KEMENTERIAN DI BAWAH MENKO PPK
--
-- EDIGPRO digabung menjadi satu kementerian karena repositori
-- memiliki satu folder dan satu logo EDIGPRO.
-- ============================================================

(
    'EDIGPRO',
    'Kementerian Ekonomi Digital dan Produk Kreatif',
    'EDIGPRO',
    'ekonomi-digital-dan-produk-kreatif',
    'KEMENTERIAN',
    'MENKO_PPK',
    1,
    'local-logo-edigpro',
    '/images/Kementerian/EDIGPRO.png'
),
(
    'RISKEL',
    'Kementerian Riset dan Keilmuan',
    'RISKEL',
    'riset-dan-keilmuan',
    'KEMENTERIAN',
    'MENKO_PPK',
    2,
    'local-logo-riskel',
    '/images/Kementerian/RISKEL.png'
),
(
    'SENIIORA',
    'Kementerian Seni dan Olahraga',
    'SENIIORA',
    'seni-dan-olahraga',
    'KEMENTERIAN',
    'MENKO_PPK',
    3,
    'local-logo-seniiora',
    '/images/Kementerian/SENIIORA.png'
),

-- ============================================================
-- UNIT DI BAWAH BPII
-- ============================================================

(
    'BENKAB',
    'Kementerian Bendahara Kabinet',
    'BENKAB',
    'bendahara-kabinet',
    'KEMENTERIAN',
    'BPII',
    1,
    'local-logo-benkab',
    '/images/Kementerian/BENKAB.png'
),
(
    'PAK',
    'Pendayagunaan Aparatur Kabinet',
    'PAK',
    'pendayagunaan-aparatur-kabinet',
    'KEMENTERIAN',
    'BPII',
    2,
    'local-logo-pak',
    '/images/Kementerian/PAK.png'
),
(
    'SESKAB',
    'Kementerian Sekretaris Kabinet',
    'SESKAB',
    'sekretaris-kabinet',
    'KEMENTERIAN',
    'BPII',
    3,
    'local-logo-seskab',
    '/images/Kementerian/SESKAB.png'
);

-- ============================================================
-- 3. MEDIA ASSET UNTUK MENKO DAN LOGO KEMENTERIAN
--
-- ImageKitFileID tetap harus diisi karena field NOT NULL.
-- Nilainya menggunakan identifier lokal, bukan ID ImageKit.
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
    NULL,
    unit.logo_asset_key,
    unit.logo_path,
    unit.logo_path,
    unit.logo_path,

    CASE
        WHEN unit.unit_type = 'MENKO'
            THEN CONCAT('Foto ', unit.name)
        ELSE CONCAT('Logo ', unit.name)
    END,

    CASE
        WHEN unit.unit_type = 'MENKO'
            THEN CONCAT('Foto Menteri Koordinator ', unit.short_name)
        ELSE CONCAT('Logo resmi ', unit.name)
    END,

    CASE
        WHEN unit.unit_type = 'MENKO'
            THEN CONCAT('Foto profil pimpinan ', unit.name)
        ELSE CONCAT('Logo organisasi ', unit.name)
    END,

    'image/png',
    0,
    NULL,
    NULL,

    CASE
        WHEN unit.unit_type = 'MENKO' THEN 'MENKO_PHOTO'
        ELSE 'MINISTRY_LOGO'
    END,

    'ACTIVE',
    NOW(),
    NOW()
FROM seed_organization_units AS unit
WHERE
    unit.logo_asset_key IS NOT NULL
    AND unit.logo_path IS NOT NULL
ON DUPLICATE KEY UPDATE
    file_path = VALUES(file_path),
    url = VALUES(url),
    thumbnail_url = VALUES(thumbnail_url),
    name = VALUES(name),
    alt_text = VALUES(alt_text),
    caption = VALUES(caption),
    mime_type = VALUES(mime_type),
    purpose = VALUES(purpose),
    status = VALUES(status),
    updated_at = NOW();

-- ============================================================
-- 4. INSERT UNIT INDUK: MENKO DAN BPII
-- ============================================================

INSERT INTO ministries (
    code,
    name,
    cabinet_term_id,
    parent_id,
    unit_type,
    slug,
    short_name,
    description,
    vision,
    mission,
    logo_media_id,
    cover_media_id,
    display_order,
    is_active,
    is_published,
    published_at,
    created_at,
    updated_at
)
SELECT
    unit.code,
    unit.name,
    @cabinet_term_id,
    NULL,
    unit.unit_type,
    unit.slug,
    unit.short_name,

    CASE
        WHEN unit.unit_type = 'MENKO' THEN
            CONCAT(
                unit.name,
                ' merupakan unit koordinator yang menaungi kementerian dalam bidangnya.'
            )
        ELSE
            'BPII merupakan unit induk yang menaungi Bendahara Kabinet, Pendayagunaan Aparatur Kabinet, dan Sekretaris Kabinet.'
    END,

    'Mewujudkan organisasi yang kolaboratif, profesional, inklusif, dan memberikan dampak nyata.',

    'Menguatkan koordinasi, pengawasan, pengembangan, pelayanan, dan evaluasi organisasi secara berkelanjutan.',

    media.id,
    NULL,
    unit.display_order,
    TRUE,
    TRUE,
    NOW(),
    NOW(),
    NOW()
FROM seed_organization_units AS unit
LEFT JOIN media_assets AS media
    ON media.imagekit_file_id = unit.logo_asset_key
WHERE unit.parent_code IS NULL
ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    cabinet_term_id = VALUES(cabinet_term_id),
    parent_id = NULL,
    unit_type = VALUES(unit_type),
    slug = VALUES(slug),
    short_name = VALUES(short_name),
    description = VALUES(description),
    vision = VALUES(vision),
    mission = VALUES(mission),
    logo_media_id = VALUES(logo_media_id),
    cover_media_id = VALUES(cover_media_id),
    display_order = VALUES(display_order),
    is_active = VALUES(is_active),
    is_published = VALUES(is_published),
    published_at = VALUES(published_at),
    updated_at = NOW();

-- ============================================================
-- 5. INSERT SELURUH KEMENTERIAN
-- ============================================================

INSERT INTO ministries (
    code,
    name,
    cabinet_term_id,
    parent_id,
    unit_type,
    slug,
    short_name,
    description,
    vision,
    mission,
    logo_media_id,
    cover_media_id,
    display_order,
    is_active,
    is_published,
    published_at,
    created_at,
    updated_at
)
SELECT
    unit.code,
    unit.name,
    @cabinet_term_id,
    parent.id,
    unit.unit_type,
    unit.slug,
    unit.short_name,

    CONCAT(
        unit.name,
        ' merupakan unit organisasi yang berada di bawah ',
        parent.name,
        '.'
    ),

    CONCAT(
        'Mewujudkan ',
        unit.short_name,
        ' yang profesional, kolaboratif, responsif, dan berdampak.'
    ),

    'Menyusun, melaksanakan, mendokumentasikan, dan mengevaluasi program kerja sesuai bidang kementerian.',

    logo.id,
    NULL,
    unit.display_order,
    TRUE,
    TRUE,
    NOW(),
    NOW(),
    NOW()
FROM seed_organization_units AS unit
INNER JOIN ministries AS parent
    ON parent.code = unit.parent_code
LEFT JOIN media_assets AS logo
    ON logo.imagekit_file_id = unit.logo_asset_key
WHERE unit.parent_code IS NOT NULL
ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    cabinet_term_id = VALUES(cabinet_term_id),
    parent_id = VALUES(parent_id),
    unit_type = VALUES(unit_type),
    slug = VALUES(slug),
    short_name = VALUES(short_name),
    description = VALUES(description),
    vision = VALUES(vision),
    mission = VALUES(mission),
    logo_media_id = VALUES(logo_media_id),
    cover_media_id = VALUES(cover_media_id),
    display_order = VALUES(display_order),
    is_active = VALUES(is_active),
    is_published = VALUES(is_published),
    published_at = VALUES(published_at),
    updated_at = NOW();

-- ============================================================
-- 6. TEMPORARY DATA ANGGOTA ORGANISASI
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_organization_members;

CREATE TEMPORARY TABLE seed_organization_members (
    unit_code         VARCHAR(50)  NOT NULL,
    name              VARCHAR(120) NOT NULL,
    position          VARCHAR(120) NOT NULL,
    position_type     VARCHAR(40)  NOT NULL,
    photo_asset_key   VARCHAR(100) NULL,
    photo_path        VARCHAR(1000) NULL,
    display_order     INT UNSIGNED NOT NULL DEFAULT 0,
    is_leader         BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (unit_code, position)
);

INSERT INTO seed_organization_members (
    unit_code,
    name,
    position,
    position_type,
    photo_asset_key,
    photo_path,
    display_order,
    is_leader
) VALUES

-- ============================================================
-- MENKO
-- ============================================================

(
    'MENKO_KEMAHASISWAAN',
    'Johanes Richard Darmawan',
    'Menteri Koordinator Kemahasiswaan',
    'MENKO',
    'local-menko-kemahasiswaan',
    '/images/Menko/menko_kemahasiswaan_johanes_richard_darmawan.png',
    1,
    TRUE
),
(
    'MENKO_KEMASYARAKATAN',
    'Muhammad Ziyaad Difa Ul Haq',
    'Menteri Koordinator Kemasyarakatan',
    'MENKO',
    'local-menko-kemasyarakatan',
    '/images/Menko/menko_kemasyarakatan_muhammad_ziyaad_difa_ul_haq.png',
    1,
    TRUE
),
(
    'MENKO_KOMINFO',
    'Andrean Marcellino',
    'Menteri Koordinator Komunikasi dan Informasi',
    'MENKO',
    'local-menko-kominfo',
    '/images/Menko/menko_kominfo_andrean_marcellino.png',
    1,
    TRUE
),
(
    'MENKO_PERGERAKAN',
    'Dewa Pranata Putra Pratama',
    'Menteri Koordinator Pergerakan',
    'MENKO',
    'local-menko-pergerakan',
    '/images/Menko/menko_pergerakan_dewa_pranata_putra_pratama.png',
    1,
    TRUE
),
(
    'MENKO_PPK',
    'Nasua Okta Kusuma Dewi Rofiq',
    'Menteri Koordinator PPK',
    'MENKO',
    'local-menko-ppk',
    '/images/Menko/menko_ppk_nasua_okta_kusuma_dewi_rofiq.png',
    1,
    TRUE
),

-- ============================================================
-- ADKESMA
-- ============================================================

(
    'ADKESMA',
    'Elsa Putri Aprizalni',
    'Menteri Advokasi dan Kesejahteraan Mahasiswa',
    'MENTERI',
    'local-member-adkesma-elsa-putri-aprizalni',
    '/images/ADKESMA/menteri_adkesma_elsa_putri_aprizalni.png',
    1,
    TRUE
),
(
    'ADKESMA',
    'Yustia Khuspriyanti Feby',
    'Direktur Jenderal Advokasi',
    'DIRJEN',
    'local-member-adkesma-yustia-khuspriyanti-feby',
    '/images/ADKESMA/dirjen_adkesma_advokasi_yustia_khuspriyanti_feby.png',
    2,
    FALSE
),
(
    'ADKESMA',
    'Nabila Afifah Abida',
    'Direktur Jenderal Kesejahteraan',
    'DIRJEN',
    'local-member-adkesma-nabila-afifah-abida',
    '/images/ADKESMA/dirjen_adkesma_kesejahteraan_nabila_afifah_abida.png',
    3,
    FALSE
),

-- ============================================================
-- PENGEMBANGAN PROFESI
-- ============================================================

(
    'PENGPROF',
    'M. Rifqy Yusuf',
    'Menteri Pengembangan Profesi',
    'MENTERI',
    'local-member-pengprof-m-rifqy-yusuf',
    '/images/PENGPROF/menteri_pengprof_m_rifqy_yusuf.png',
    1,
    TRUE
),
(
    'PENGPROF',
    'Raizatya Qoriananda',
    'Direktur Jenderal Keahlian',
    'DIRJEN',
    'local-member-pengprof-raizatya-qoriananda',
    '/images/PENGPROF/dirjen_pengprof_keahlian_raizatya_qoriananda.png',
    2,
    FALSE
),
(
    'PENGPROF',
    'Deva Zia Ul Haq',
    'Direktur Jenderal Relasi',
    'DIRJEN',
    'local-member-pengprof-deva-zia-ul-haq',
    '/images/PENGPROF/dirjen_pengprof_relasi_deva_zia_ul_haq.png',
    3,
    FALSE
),

-- ============================================================
-- PSDM
-- ============================================================

(
    'PSDM',
    'Labib Azka Asysyafiq',
    'Menteri Pengembangan Sumber Daya Mahasiswa',
    'MENTERI',
    'local-member-psdm-labib-azka-asysyafiq',
    '/images/PSDM/menteri_psdm_labib_azka_asysyafiq.png',
    1,
    TRUE
),
(
    'PSDM',
    'Ferdian Amartha',
    'Direktur Jenderal Akselerasi',
    'DIRJEN',
    'local-member-psdm-ferdian-amartha',
    '/images/PSDM/dirjen_psdm_akselerasi_ferdian_amartha.png',
    2,
    FALSE
),
(
    'PSDM',
    'Qheisya Luna Alifa',
    'Direktur Jenderal Kaderisasi',
    'DIRJEN',
    'local-member-psdm-qheisya-luna-alifa',
    '/images/PSDM/dirjen_psdm_kaderisasi_qheisya_luna_alifa.png',
    3,
    FALSE
),

-- ============================================================
-- LINGKUNGAN HIDUP
--
-- Nama tiga Dirjen tidak tersedia dalam nama file.
-- Nama sementara dibuat berdasarkan bidang jabatannya.
-- ============================================================

(
    'LH',
    'Zaskia Darojah',
    'Menteri Lingkungan Hidup',
    'MENTERI',
    'local-member-lh-zaskia-darojah',
    '/images/LH/menteri_lh_zaskia_darojah.png',
    1,
    TRUE
),
(
    'LH',
    'Nama Belum Tersedia - Advokasi dan Kajian Strategis',
    'Direktur Jenderal Advokasi dan Kajian Strategis',
    'DIRJEN',
    'local-member-lh-advokasi-kajian-strategis',
    '/images/LH/dirjen_lh_advokasi_kajian_strategis.png',
    2,
    FALSE
),
(
    'LH',
    'Nama Belum Tersedia - Pencegahan dan Pengendalian',
    'Direktur Jenderal Pencegahan dan Pengendalian',
    'DIRJEN',
    'local-member-lh-pencegahan-pengendalian',
    '/images/LH/dirjen_lh_pencegahan_pengendalian.png',
    3,
    FALSE
),
(
    'LH',
    'Nama Belum Tersedia - Pengembangan dan Preservasi',
    'Direktur Jenderal Pengembangan dan Preservasi',
    'DIRJEN',
    'local-member-lh-pengembangan-preservasi',
    '/images/LH/dirjen_lh_pengembangan_preservasi.png',
    4,
    FALSE
),

-- ============================================================
-- KEMENTERIAN KESEHATAN
--
-- Nama dua Dirjen tidak tersedia dalam nama file.
-- ============================================================

(
    'KEMENKES',
    'An Nadhofah Adlin',
    'Menteri Kesehatan',
    'MENTERI',
    'local-member-kemenkes-an-nadhofah-adlin',
    '/images/Menkes/menteri_menkes_an_nadhofah_adlin.png',
    1,
    TRUE
),
(
    'KEMENKES',
    'Nama Belum Tersedia - Kesehatan Masyarakat',
    'Direktur Jenderal Kesehatan Masyarakat',
    'DIRJEN',
    'local-member-kemenkes-kesehatan-masyarakat',
    '/images/Menkes/dirjen_menkes_kesehatan_masyarakat.png',
    2,
    FALSE
),
(
    'KEMENKES',
    'Nama Belum Tersedia - Psikososial dan Inklusivitas',
    'Direktur Jenderal Psikososial dan Inklusivitas',
    'DIRJEN',
    'local-member-kemenkes-psikososial-inklusivitas',
    '/images/Menkes/dirjen_menkes_psikososial_inklusivitas.png',
    3,
    FALSE
),

-- ============================================================
-- PENGABDIAN MASYARAKAT
-- ============================================================

(
    'PENGMAS',
    'Sahrul Efendi',
    'Menteri Pengabdian Masyarakat',
    'MENTERI',
    'local-member-pengmas-sahrul-efendi',
    '/images/Pengabdian_Masyarakat/menteri_pengmas_sahrul_efendi.png',
    1,
    TRUE
),
(
    'PENGMAS',
    'Fahrur Rosi',
    'Direktur Jenderal Pengabdian',
    'DIRJEN',
    'local-member-pengmas-fahrur-rosi',
    '/images/Pengabdian_Masyarakat/dirjen_pengmas_fahrur_rosi.png',
    2,
    FALSE
),
(
    'PENGMAS',
    'M. Wahyu',
    'Direktur Jenderal Pemberdayaan Masyarakat',
    'DIRJEN',
    'local-member-pengmas-m-wahyu',
    '/images/Pengabdian_Masyarakat/dirjen_pengmas_m_wahyu.png',
    3,
    FALSE
),

-- ============================================================
-- HUBUNGAN LUAR
-- ============================================================

(
    'HUBLU',
    'Zalfaa Putri Arfiliesia',
    'Menteri Hubungan Luar',
    'MENTERI',
    'local-member-hublu-zalfaa-putri-arfiliesia',
    '/images/Hubungan_Luar/menteri_hublu_zalfaa_putri_arfiliesia.png',
    1,
    TRUE
),
(
    'HUBLU',
    'Billal Syahdan Arrafah',
    'Direktur Jenderal Hubungan Eksternal',
    'DIRJEN',
    'local-member-hublu-billal-syahdan-arrafah',
    '/images/Hubungan_Luar/dirjen_hublu_billal_syahdan_arrafah.png',
    2,
    FALSE
),
(
    'HUBLU',
    'Hafiz Naufal Utomo',
    'Direktur Jenderal Kemitraan',
    'DIRJEN',
    'local-member-hublu-hafiz-naufal-utomo',
    '/images/Hubungan_Luar/dirjen_hublu_hafiz_naufal_utomo.png',
    3,
    FALSE
),

-- ============================================================
-- MEDIA DAN INFORMASI
-- ============================================================

(
    'MEDINFO',
    'Ganesya Intantalia',
    'Menteri Media dan Informasi',
    'MENTERI',
    'local-member-medinfo-ganesya-intantalia',
    '/images/Medinfo/menteri_medinfo_ganesya_intantalia.png',
    1,
    TRUE
),
(
    'MEDINFO',
    'Ocha Della Fitriani',
    'Direktur Jenderal Publikasi',
    'DIRJEN',
    'local-member-medinfo-ocha-della-fitriani',
    '/images/Medinfo/dirjen_medinfo_publikasi_ocha_della_fitriani.png',
    2,
    FALSE
),
(
    'MEDINFO',
    'Febiola Al Jannah',
    'Direktur Jenderal Visual',
    'DIRJEN',
    'local-member-medinfo-febiola-al-jannah',
    '/images/Medinfo/dirjen_medinfo_visual_febiola_al_jannah.png',
    3,
    FALSE
),

-- ============================================================
-- SINERGITAS MAHASISWA
-- ============================================================

(
    'SINEMA',
    'Eugenius Nathaniel',
    'Menteri Sinergitas Mahasiswa',
    'MENTERI',
    'local-member-sinema-eugenius-nathaniel',
    '/images/Sinergitas_Mahasiswa/menteri_sinema_eugenius_nathaniel.png',
    1,
    TRUE
),
(
    'SINEMA',
    'Vicentius Nolan',
    'Direktur Jenderal Bina Sinergitas',
    'DIRJEN',
    'local-member-sinema-vicentius-nolan',
    '/images/Sinergitas_Mahasiswa/dirjen_sinema_bina_vicentius_nolan.png',
    2,
    FALSE
),
(
    'SINEMA',
    'Achmad Rafi Syahputra',
    'Direktur Jenderal Kolaborasi',
    'DIRJEN',
    'local-member-sinema-achmad-rafi-syahputra',
    '/images/Sinergitas_Mahasiswa/dirjen_sinema_korlaborasi_achmad_rafi_syahputra.png',
    3,
    FALSE
),

-- ============================================================
-- KEADILAN DAN KESETARAAN GENDER
-- ============================================================

(
    'KKG',
    'Farrel Ardisto Samai',
    'Menteri Keadilan dan Kesetaraan Gender',
    'MENTERI',
    'local-member-kkg-farrel-ardisto-samai',
    '/images/KKG/menteri_kkg_farrel_ardisto_samai.png',
    1,
    TRUE
),
(
    'KKG',
    'Annisa Rahmawati',
    'Direktur Jenderal Edukasi',
    'DIRJEN',
    'local-member-kkg-annisa-rahmawati',
    '/images/KKG/dirjen_kkg_edukasi_annisa_rahmawati.png',
    2,
    FALSE
),
(
    'KKG',
    'Titis Nastiti Dwi',
    'Direktur Jenderal Pemberdayaan',
    'DIRJEN',
    'local-member-kkg-titis-nastiti-dwi',
    '/images/KKG/dirjen_kkg_pemberdayaan_titis_nastiti_dwi.png',
    3,
    FALSE
),

-- ============================================================
-- SOSIAL DAN POLITIK
-- ============================================================

(
    'SOSPOL',
    'Daniel Nikon Martua Situmorang',
    'Menteri Sosial dan Politik',
    'MENTERI',
    'local-member-sospol-daniel-nikon-martua-situmorang',
    '/images/Sosial_dan_Politik/menteri_sospol_daniel_nikon_martua_situmorang.png',
    1,
    TRUE
),
(
    'SOSPOL',
    'Moch. Amir Nazaruddin',
    'Direktur Jenderal Aksi',
    'DIRJEN',
    'local-member-sospol-moch-amir-nazaruddin',
    '/images/Sosial_dan_Politik/dirjen_sospol_aksi_moch_amir_nazaruddin.png',
    2,
    FALSE
),
(
    'SOSPOL',
    'Ahmad Raihan Fadhillah',
    'Direktur Jenderal Analisis',
    'DIRJEN',
    'local-member-sospol-ahmad-raihan-fadhillah',
    '/images/Sosial_dan_Politik/dirjen_sospol_analisis_ahmad_raihan_fadhillah.png',
    3,
    FALSE
),

-- ============================================================
-- EKONOMI DIGITAL DAN PRODUK KREATIF
-- ============================================================

(
    'EDIGPRO',
    'Farhan Kamil',
    'Menteri Ekonomi Digital dan Produk Kreatif',
    'MENTERI',
    'local-member-edigpro-farhan-kamil',
    '/images/EDIGPRO/menteri_edigpro_farhan_kamil.png',
    1,
    TRUE
),
(
    'EDIGPRO',
    'Argi',
    'Direktur Jenderal Ekonomi Digital',
    'DIRJEN',
    'local-member-edigpro-argi',
    '/images/EDIGPRO/dirjen_edigpro_argi.png',
    2,
    FALSE
),
(
    'EDIGPRO',
    'Yudhistira',
    'Direktur Jenderal Produk Kreatif',
    'DIRJEN',
    'local-member-edigpro-yudhistira',
    '/images/EDIGPRO/dirjen_edigpro_yudhistira.png',
    3,
    FALSE
),

-- ============================================================
-- RISET DAN KEILMUAN
-- ============================================================

(
    'RISKEL',
    'Farrel Bisma Abyakta',
    'Menteri Riset dan Keilmuan',
    'MENTERI',
    'local-member-riskel-farrel-bisma-abyakta',
    '/images/Riset_dan_Keilmuan/menteri_risel_farrel_bisma_abyakta.png',
    1,
    TRUE
),
(
    'RISKEL',
    'Ata',
    'Direktur Jenderal Riset',
    'DIRJEN',
    'local-member-riskel-ata',
    '/images/Riset_dan_Keilmuan/dirjen_riset_ata.png',
    2,
    FALSE
),
(
    'RISKEL',
    'Naswa',
    'Direktur Jenderal Keilmuan',
    'DIRJEN',
    'local-member-riskel-naswa',
    '/images/Riset_dan_Keilmuan/dirjen_riset_naswa.png',
    3,
    FALSE
),

-- ============================================================
-- SENI DAN OLAHRAGA
-- ============================================================

(
    'SENIIORA',
    'Alaina Atwa Awali Harahap',
    'Menteri Seni dan Olahraga',
    'MENTERI',
    'local-member-seniiora-alaina-atwa-awali-harahap',
    '/images/Seni_dan_Olahraga/menteri_seniora_alaina_atwa_awali_harahap.png',
    1,
    TRUE
),
(
    'SENIIORA',
    'Azzam',
    'Direktur Jenderal Seni',
    'DIRJEN',
    'local-member-seniiora-azzam',
    '/images/Seni_dan_Olahraga/dirjen_seniora_azzam.png',
    2,
    FALSE
),
(
    'SENIIORA',
    'M. Fadhil',
    'Direktur Jenderal Olahraga',
    'DIRJEN',
    'local-member-seniiora-m-fadhil',
    '/images/Seni_dan_Olahraga/dirjen_seniora_m_fadhil.png',
    3,
    FALSE
),

-- ============================================================
-- BENDAHARA KABINET
-- ============================================================

(
    'BENKAB',
    'Inez Faradina Kasih',
    'Menteri Bendahara Kabinet',
    'MENTERI',
    'local-member-benkab-inez-faradina-kasih',
    '/images/Bendahara_Kabinet/menteri_benkab_inez_faradina_kasih.png',
    1,
    TRUE
),
(
    'BENKAB',
    'Ananda Magfira Nurochma',
    'Kepala Biro Bendahara Kabinet I',
    'KABIRO',
    'local-member-benkab-ananda-magfira-nurochma',
    '/images/Bendahara_Kabinet/kabiro_benkab_ananda_magfira_nurochma.png',
    2,
    FALSE
),
(
    'BENKAB',
    'M. Salman Al Farisi',
    'Kepala Biro Bendahara Kabinet II',
    'KABIRO',
    'local-member-benkab-m-salman-al-farisi',
    '/images/Bendahara_Kabinet/kabiro_benkab_m_salman_al_farisi.png',
    3,
    FALSE
),

-- ============================================================
-- PENDAYAGUNAAN APARATUR KABINET
-- ============================================================

(
    'PAK',
    'Deri Bayu Setiawan',
    'Kepala Pendayagunaan Aparatur Kabinet',
    'KEPALA',
    'local-member-pak-deri-bayu-setiawan',
    '/images/Pendayagunaan_Aparatur_Kabinet/kepala_pak_deri_bayu_setiawan.png',
    1,
    TRUE
),
(
    'PAK',
    'Mevlana El Fithra',
    'Deputi Pendayagunaan Aparatur Kabinet I',
    'DEPUTI',
    'local-member-pak-mevlana-el-fithra',
    '/images/Pendayagunaan_Aparatur_Kabinet/deputi_pak_mevlana_el_fithra.png',
    2,
    FALSE
),
(
    'PAK',
    'Nayla Khansa Abbylia',
    'Deputi Pendayagunaan Aparatur Kabinet II',
    'DEPUTI',
    'local-member-pak-nayla-khansa-abbylia',
    '/images/Pendayagunaan_Aparatur_Kabinet/deputi_pak_nayla_khansa_abbylia.png',
    3,
    FALSE
),

-- ============================================================
-- SEKRETARIS KABINET
-- ============================================================

(
    'SESKAB',
    'Gheahaq Danty El Zahra',
    'Menteri Sekretaris Kabinet',
    'MENTERI',
    'local-member-seskab-gheahaq-danty-el-zahra',
    '/images/Sekretaris_Kabinet/menteri_seskab_gheahaq_danty_el_zahra.png',
    1,
    TRUE
),
(
    'SESKAB',
    'Fathimah Nuurun',
    'Kepala Biro Sekretaris Kabinet I',
    'KABIRO',
    'local-member-seskab-fathimah-nuurun',
    '/images/Sekretaris_Kabinet/kabiro_seskab_fathimah_nuurun.png',
    2,
    FALSE
),
(
    'SESKAB',
    'Najwa Nur Hafidza',
    'Kepala Biro Sekretaris Kabinet II',
    'KABIRO',
    'local-member-seskab-najwa-nur-hafidza',
    '/images/Sekretaris_Kabinet/kabiro_seskab_najwa_nur_hafidza.png',
    3,
    FALSE
);

-- ============================================================
-- 7. MEDIA ASSET FOTO ORGANIZATION MEMBER
--
-- Foto Menko dilewati di sini karena sudah dibuat pada bagian
-- MediaAsset unit Menko dan menggunakan asset key yang sama.
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
    NULL,
    member.photo_asset_key,
    member.photo_path,
    member.photo_path,
    member.photo_path,
    CONCAT('Foto ', member.name),
    CONCAT(member.name, ' - ', member.position),
    CONCAT('Foto profil ', member.name, ' sebagai ', member.position),
    'image/png',
    0,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER_PHOTO',
    'ACTIVE',
    NOW(),
    NOW()
FROM seed_organization_members AS member
WHERE
    member.photo_asset_key IS NOT NULL
    AND member.photo_path IS NOT NULL
    AND NOT EXISTS (
        SELECT 1
        FROM seed_organization_units AS unit
        WHERE unit.logo_asset_key = member.photo_asset_key
    )
ON DUPLICATE KEY UPDATE
    file_path = VALUES(file_path),
    url = VALUES(url),
    thumbnail_url = VALUES(thumbnail_url),
    name = VALUES(name),
    alt_text = VALUES(alt_text),
    caption = VALUES(caption),
    mime_type = VALUES(mime_type),
    purpose = VALUES(purpose),
    status = VALUES(status),
    updated_at = NOW();

-- ============================================================
-- 8. UPDATE MEMBER YANG SUDAH ADA
--
-- Pencocokan menggunakan ministry_id + position supaya seeder
-- dapat dijalankan ulang tanpa membuat duplikasi.
-- ============================================================

UPDATE organization_members AS organization_member
INNER JOIN ministries AS ministry
    ON ministry.id = organization_member.ministry_id
INNER JOIN seed_organization_members AS seed_member
    ON seed_member.unit_code = ministry.code
    AND seed_member.position = organization_member.position
LEFT JOIN media_assets AS photo
    ON photo.imagekit_file_id = seed_member.photo_asset_key
SET
    organization_member.name = seed_member.name,
    organization_member.position_type = seed_member.position_type,
    organization_member.photo_media_id = photo.id,
    organization_member.display_order = seed_member.display_order,
    organization_member.is_leader = seed_member.is_leader,
    organization_member.is_active = TRUE,
    organization_member.started_at = '2026-01-01',
    organization_member.ended_at = NULL,
    organization_member.updated_at = NOW();

-- ============================================================
-- 9. INSERT MEMBER YANG BELUM ADA
-- ============================================================

INSERT INTO organization_members (
    ministry_id,
    user_id,
    name,
    position,
    position_type,
    email_public,
    phone_public,
    biography,
    quote,
    photo_media_id,
    display_order,
    is_leader,
    is_active,
    started_at,
    ended_at,
    created_at,
    updated_at
)
SELECT
    ministry.id,
    NULL,
    seed_member.name,
    seed_member.position,
    seed_member.position_type,
    NULL,
    NULL,
    NULL,
    NULL,
    photo.id,
    seed_member.display_order,
    seed_member.is_leader,
    TRUE,
    '2026-01-01',
    NULL,
    NOW(),
    NOW()
FROM seed_organization_members AS seed_member
INNER JOIN ministries AS ministry
    ON ministry.code = seed_member.unit_code
LEFT JOIN media_assets AS photo
    ON photo.imagekit_file_id = seed_member.photo_asset_key
WHERE NOT EXISTS (
    SELECT 1
    FROM organization_members AS existing_member
    WHERE
        existing_member.ministry_id = ministry.id
        AND existing_member.position = seed_member.position
);

-- ============================================================
-- 10. PASTIKAN KABINET CERITA LOKA MENJADI KABINET AKTIF
-- ============================================================

UPDATE cabinet_terms
SET
    is_active = FALSE,
    updated_at = NOW()
WHERE id <> @cabinet_term_id;

UPDATE cabinet_terms
SET
    is_active = TRUE,
    is_published = TRUE,
    published_at = COALESCE(published_at, NOW()),
    updated_at = NOW()
WHERE id = @cabinet_term_id;

-- ============================================================
-- 11. HAPUS TEMPORARY TABLE
-- ============================================================

DROP TEMPORARY TABLE IF EXISTS seed_organization_members;
DROP TEMPORARY TABLE IF EXISTS seed_organization_units;

COMMIT;

-- ============================================================
-- 12. VERIFIKASI KABINET
-- ============================================================

SELECT
    cabinet.id,
    cabinet.name,
    cabinet.slug,
    cabinet.period_start,
    cabinet.period_end,
    cabinet.is_active,
    cabinet.is_published
FROM cabinet_terms AS cabinet
WHERE cabinet.id = @cabinet_term_id;

-- ============================================================
-- 13. VERIFIKASI STRUKTUR MENKO, BPII, DAN KEMENTERIAN
-- ============================================================

SELECT
    parent.id AS parent_id,
    parent.code AS parent_code,
    parent.name AS parent_name,
    parent.unit_type AS parent_unit_type,
    parent_media.url AS parent_media_url,

    child.id AS ministry_id,
    child.code AS ministry_code,
    child.name AS ministry_name,
    child.unit_type AS ministry_unit_type,
    ministry_logo.url AS ministry_logo_url

FROM ministries AS parent

LEFT JOIN media_assets AS parent_media
    ON parent_media.id = parent.logo_media_id

LEFT JOIN ministries AS child
    ON child.parent_id = parent.id

LEFT JOIN media_assets AS ministry_logo
    ON ministry_logo.id = child.logo_media_id

WHERE
    parent.cabinet_term_id = @cabinet_term_id
    AND parent.parent_id IS NULL
    AND parent.unit_type IN ('MENKO', 'BPII')

ORDER BY
    parent.display_order ASC,
    child.display_order ASC;

-- ============================================================
-- 14. VERIFIKASI ORGANIZATION MEMBER
-- ============================================================

SELECT
    parent.name AS parent_name,
    ministry.name AS ministry_name,
    member.name AS member_name,
    member.position,
    member.position_type,
    member.is_leader,
    member.display_order,
    photo.url AS photo_url

FROM organization_members AS member

INNER JOIN ministries AS ministry
    ON ministry.id = member.ministry_id

LEFT JOIN ministries AS parent
    ON parent.id = ministry.parent_id

LEFT JOIN media_assets AS photo
    ON photo.id = member.photo_media_id

WHERE
    ministry.cabinet_term_id = @cabinet_term_id
    AND member.is_active = TRUE

ORDER BY
    COALESCE(parent.display_order, ministry.display_order) ASC,
    ministry.display_order ASC,
    member.display_order ASC;