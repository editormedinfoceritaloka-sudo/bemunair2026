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

SET NAMES utf8mb4 COLLATE utf8mb4_unicode_ci;

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
)
ENGINE = InnoDB
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/RICHIE_MENKO_KEMAHASISWAAN'
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/ZIYAAD_MENKO_KEMASYARAKATAN'
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENKO_KOMINFO_ANDREAN'
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DEWA_MENKO_PERGERAKAN'
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/NASUA_OKTA_MENKO_PPK'
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
    'Bendahara Kabinet',
    'Bendahara Kabinet',
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
    'Sekretaris Kabinet',
    'Sekretaris Kabinet',
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
                ' merupakan unit koordinator yang menaungi kementerian dalam bidangnya. ',
                'Melalui koordinasi, sinergi, dan evaluasi yang berkelanjutan, ',
                COALESCE(unit.short_name, unit.name),
                ' menyatukan arah gerak kementerian-kementerian di bawahnya untuk menghadirkan program kerja yang berdampak bagi Airlangga dan Indonesia.'
            )
        ELSE
            'BPII merupakan unit induk yang menaungi Bendahara Kabinet, Pendayagunaan Aparatur Kabinet, dan Sekretaris Kabinet. BPII memastikan tata kelola kabinet berjalan tertib, transparan, dan mendukung kebermanfaatan bersama.'
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
    ON media.imagekit_file_id COLLATE utf8mb4_unicode_ci
     = unit.logo_asset_key COLLATE utf8mb4_unicode_ci
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
        '. ',
        COALESCE(unit.short_name, unit.name),
        ' berkolaborasi dengan kementerian dan unit lain untuk menyusun, melaksanakan, dan mengevaluasi program kerja yang berdampak bagi mahasiswa serta masyarakat Airlangga.'
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
    ON parent.code COLLATE utf8mb4_unicode_ci
     = unit.parent_code COLLATE utf8mb4_unicode_ci
LEFT JOIN media_assets AS logo
    ON logo.imagekit_file_id COLLATE utf8mb4_unicode_ci
     = unit.logo_asset_key COLLATE utf8mb4_unicode_ci
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
-- 5B. DESKRIPSI RESMI MENKO DAN KEMENTERIAN
-- ============================================================

UPDATE ministries
SET
    description = CASE code
        WHEN 'MENKO_KOMINFO' THEN
            'Kemenko Komunikasi dan Informasi merupakan bidang yang bertanggung jawab mengoordinasikan komunikasi internal dan eksternal BEM UNAIR serta membangun hubungan dengan berbagai mitra. Kemenkoan ini juga memastikan penyampaian informasi dan hubungan antarkementerian berjalan dengan baik.'
        WHEN 'MENKO_PPK' THEN
            'Kemenko Pengembangan Potensi dan Karya merupakan bidang yang berfokus pada pengembangan serta optimalisasi potensi mahasiswa dalam bidang akademik, penelitian, seni, olahraga, dan ekonomi kreatif. Kemenkoan ini bertujuan menciptakan ruang yang mendukung kreativitas, inovasi, dan pencapaian prestasi mahasiswa.'
        WHEN 'MENKO_PERGERAKAN' THEN
            'Kemenko Pergerakan merupakan bidang yang bertanggung jawab mengoordinasikan gerakan mahasiswa dalam mengawal isu-isu sosial, politik, keadilan, dan kesetaraan gender. Kemenkoan ini mendorong mahasiswa untuk memiliki kepedulian kritis serta terlibat dalam advokasi dan gerakan yang berlandaskan kajian intelektual.'
        WHEN 'MENKO_KEMAHASISWAAN' THEN
            'Kemenko Kemahasiswaan merupakan bidang yang berfokus pada pengembangan kualitas, kapasitas, dan kesejahteraan mahasiswa. Kemenkoan ini mengoordinasikan program kaderisasi, pengembangan kemampuan mahasiswa, advokasi hak mahasiswa, serta persiapan mahasiswa dalam menghadapi dunia profesional dan pascakampus.'
        WHEN 'MENKO_KEMASYARAKATAN' THEN
            'Kemenko Kemasyarakatan merupakan bidang yang bertanggung jawab mengoordinasikan program-program yang berorientasi pada kebermanfaatan masyarakat. Fokusnya meliputi pemberdayaan masyarakat, pelestarian lingkungan, serta peningkatan kesehatan dan kesejahteraan melalui kegiatan yang berkelanjutan.'
        WHEN 'MEDINFO' THEN
            'Kementerian Media dan Informasi merupakan kementerian yang bertanggung jawab mengelola dan menyebarluaskan informasi mengenai kegiatan serta program kerja BEM UNAIR. Kementerian ini juga berperan membangun citra dan identitas organisasi melalui publikasi yang informatif, edukatif, dan kreatif.'
        WHEN 'SINEMA' THEN
            'Kementerian Sinergitas Mahasiswa merupakan kementerian yang berfokus membangun hubungan dan kolaborasi antarorganisasi mahasiswa, komunitas, serta mitra internal Universitas Airlangga. Kementerian ini menjadi wadah untuk memperkuat kerja sama dan menciptakan hubungan yang harmonis antarelemen mahasiswa.'
        WHEN 'HUBLU' THEN
            'Kementerian Hubungan Luar merupakan kementerian yang bertanggung jawab menjalin dan mengembangkan hubungan kemitraan dengan pihak eksternal. Kementerian ini berperan memperluas jaringan kerja sama dengan organisasi, lembaga, maupun institusi di tingkat regional, nasional, dan internasional.'
        WHEN 'EDIGPRO' THEN
            'Kementerian Ekonomi Digital dan Produk Kreatif merupakan kementerian yang berfokus mengembangkan potensi mahasiswa dalam bidang ekonomi kreatif, kewirausahaan, dan ekonomi digital. Kementerian ini mendorong inovasi, kolaborasi, serta peningkatan kemampuan mahasiswa dalam menghasilkan produk dan mengembangkan peluang usaha.'
        WHEN 'SENIIORA' THEN
            'Kementerian Seni dan Olahraga merupakan kementerian yang bertanggung jawab mengembangkan minat dan bakat mahasiswa dalam bidang seni, budaya, dan olahraga. Kementerian ini menyediakan ruang aktualisasi kreativitas, pelestarian budaya, serta peningkatan kemampuan mahasiswa untuk meraih prestasi.'
        WHEN 'RISKEL' THEN
            'Kementerian Riset dan Keilmuan merupakan kementerian yang berfokus mendukung pengembangan kemampuan akademik, penelitian, dan keilmuan mahasiswa. Kementerian ini berperan menyediakan fasilitas, pendampingan, serta kesempatan bagi mahasiswa untuk mengembangkan penelitian, mengikuti kompetisi, dan menghasilkan karya ilmiah.'
        WHEN 'SOSPOL' THEN
            'Kementerian Sosial dan Politik merupakan kementerian yang bertanggung jawab mengkaji, mengawal, dan merespons berbagai isu sosial, politik, serta kebijakan publik yang berdampak terhadap mahasiswa dan masyarakat. Kementerian ini mendorong terbentuknya kesadaran kritis mahasiswa melalui kajian, diskusi, edukasi, dan gerakan berbasis intelektual.'
        WHEN 'KKG' THEN
            'Kementerian Keadilan dan Kesetaraan Gender merupakan kementerian yang berfokus memperjuangkan keadilan dan kesetaraan bagi seluruh mahasiswa tanpa diskriminasi. Kementerian ini berperan meningkatkan kesadaran mengenai isu gender, melakukan advokasi, serta mendorong terciptanya lingkungan kampus yang aman, nyaman, dan inklusif.'
        WHEN 'PSDM' THEN
            'Kementerian Pengembangan Sumber Daya Mahasiswa merupakan kementerian yang bertanggung jawab meningkatkan kualitas dan kapasitas mahasiswa melalui proses kaderisasi serta pengembangan potensi diri. Kementerian ini berfokus membentuk mahasiswa yang berintegritas, memiliki jiwa kepemimpinan, serta mampu mengembangkan keterampilan akademik maupun nonakademik.'
        WHEN 'ADKESMA' THEN
            'Kementerian Advokasi dan Kesejahteraan Mahasiswa merupakan kementerian yang bertanggung jawab mengawal dan memperjuangkan hak serta aspirasi mahasiswa, baik dalam bidang akademik maupun nonakademik. Kementerian ini juga berperan menyediakan pendampingan, bantuan, dan fasilitas pendukung untuk meningkatkan kesejahteraan mahasiswa.'
        WHEN 'PENGPROF' THEN
            'Kementerian Pengembangan Profesi merupakan kementerian yang berfokus mempersiapkan mahasiswa menghadapi dunia kerja dan kehidupan profesional. Kementerian ini mendukung peningkatan kompetensi melalui pelatihan, seminar, sertifikasi, serta pengembangan jejaring karier yang sesuai dengan kebutuhan dunia profesional.'
        WHEN 'PENGMAS' THEN
            'Kementerian Pengabdian Masyarakat merupakan kementerian yang bertanggung jawab merancang dan melaksanakan kegiatan pemberdayaan masyarakat secara berkelanjutan. Kementerian ini mendorong keterlibatan mahasiswa dalam membantu menyelesaikan permasalahan masyarakat melalui edukasi, pelatihan, penelitian, dan kolaborasi dengan berbagai pihak.'
        WHEN 'LH' THEN
            'Kementerian Lingkungan Hidup merupakan kementerian yang berfokus meningkatkan kepedulian mahasiswa terhadap pelestarian lingkungan dan keberlanjutan ekologi. Kementerian ini mendorong penerapan gaya hidup ramah lingkungan melalui edukasi, kampanye, pengelolaan sampah, daur ulang, serta berbagai aksi pelestarian lingkungan.'
        WHEN 'KEMENKES' THEN
            'Kementerian Kesehatan merupakan kementerian yang bertanggung jawab meningkatkan kesadaran dan kepedulian mahasiswa terhadap kesehatan fisik maupun mental. Kementerian ini menjalankan kegiatan promotif, preventif, dan responsif melalui edukasi kesehatan, kampanye, serta program yang mendukung terciptanya lingkungan kampus yang sehat dan suportif.'
        WHEN 'SESKAB' THEN
            'Sekretaris Kabinet merupakan unsur organisasi yang bertanggung jawab dalam mengelola administrasi dan kesekretariatan kabinet serta memastikan seluruh proses administrasi berjalan secara tertib, sistematis, dan terkoordinasi. Sekretaris Kabinet juga berperan dalam menyusun dan mengelola Standar Operasional Prosedur (SOP), mengatur alur birokrasi organisasi, serta memastikan kelancaran komunikasi dan koordinasi administratif antarunsur kabinet.'
        WHEN 'BENKAB' THEN
            'Bendahara Kabinet merupakan unsur organisasi yang bertanggung jawab atas pengelolaan keuangan kabinet, mulai dari perencanaan anggaran, pencatatan pemasukan dan pengeluaran, hingga pengawasan serta pelaporan keuangan. Bendahara Kabinet berperan memastikan pengelolaan dana organisasi berjalan secara transparan, akuntabel, tertib, dan sesuai dengan kebutuhan serta perencanaan program kerja kabinet.'
        WHEN 'PAK' THEN
            'Pendayagunaan Aparatur Kabinet (PAK) merupakan unsur organisasi yang bertugas melakukan evaluasi, pengawasan, audit, dan pengendalian terhadap kinerja kementerian serta seluruh fungsionaris BEM. PAK berperan memastikan setiap unsur kabinet menjalankan tugas dan tanggung jawabnya sesuai dengan fungsi, target, dan ketentuan organisasi. Selain itu, PAK menjadi instrumen evaluasi untuk mengidentifikasi kendala, meningkatkan efektivitas kerja, serta mendorong optimalisasi kinerja aparatur kabinet.'
        ELSE description
    END,
    updated_at = NOW()
WHERE code IN (
    'MENKO_KOMINFO',
    'MENKO_PPK',
    'MENKO_PERGERAKAN',
    'MENKO_KEMAHASISWAAN',
    'MENKO_KEMASYARAKATAN',
    'MEDINFO',
    'SINEMA',
    'HUBLU',
    'EDIGPRO',
    'SENIIORA',
    'RISKEL',
    'SOSPOL',
    'KKG',
    'PSDM',
    'ADKESMA',
    'PENGPROF',
    'PENGMAS',
    'LH',
    'KEMENKES',
    'SESKAB',
    'BENKAB',
    'PAK'
);

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
    PRIMARY KEY (unit_code, name)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/RICHIE_MENKO_KEMAHASISWAAN',
    1,
    TRUE
),
(
    'MENKO_KEMASYARAKATAN',
    'Muhammad Ziyaad Difaa’ul Haq',
    'Menteri Koordinator Kemasyarakatan',
    'MENKO',
    'local-menko-kemasyarakatan',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/ZIYAAD_MENKO_KEMASYARAKATAN',
    1,
    TRUE
),
(
    'MENKO_KOMINFO',
    'Andrean Marcellino',
    'Menteri Koordinator Komunikasi dan Informasi',
    'MENKO',
    'local-menko-kominfo',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENKO_KOMINFO_ANDREAN',
    1,
    TRUE
),
(
    'MENKO_PERGERAKAN',
    'Dewa Pranata Putra Pratama',
    'Menteri Koordinator Pergerakan',
    'MENKO',
    'local-menko-pergerakan',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DEWA_MENKO_PERGERAKAN',
    1,
    TRUE
),
(
    'MENKO_PPK',
    'Nasua Okta Kusuma Dewi Rofiq',
    'Menteri Koordinator PPK',
    'MENKO',
    'local-menko-ppk',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/NASUA_OKTA_MENKO_PPK',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/menteri_adkesma_elsa_ADKESMA',
    1,
    TRUE
),
(
    'ADKESMA',
    'Yustia Khuspriyanti Feby Martha',
    'Direktur Jenderal Advokasi Kebijakan Kampus',
    'DIRJEN',
    'local-member-adkesma-yustia-khuspriyanti-feby',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/Direktur_JenderalYustia_Khuspriyanti_Feby_ADKESMA',
    2,
    FALSE
),
(
    'ADKESMA',
    'Nabila Afifah Abida',
    'Direktur Jenderal Kesejahteraan Mahasiswa',
    'DIRJEN',
    'local-member-adkesma-nabila-afifah-abida',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/Direktur_JenderalNabila_Afifah_Abida_ADKESMA',
    3,
    FALSE
),

-- ============================================================
-- PENGEMBANGAN PROFESI
-- ============================================================

(
    'PENGPROF',
    'Muhammad Rifqy Yusuf',
    'Menteri Pengembangan Profesi',
    'MENTERI',
    'local-member-pengprof-m-rifqy-yusuf',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/Menteri_PENGPROFM._Rifqy_Yusuf',
    1,
    TRUE
),
(
    'PENGPROF',
    'Raizsatya Qoriananda Hutama',
    'Direktur Jenderal Relasi Korporasi dan Jaringan Alumni',
    'DIRJEN',
    'local-member-pengprof-raizatya-qoriananda',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_RAIS_PENGPROF',
    2,
    FALSE
),
(
    'PENGPROF',
    'Deva Zia''ul Haq Khoirudin',
    'Direktur Jenderal Pengembangan Keahlian',
    'DIRJEN',
    'local-member-pengprof-deva-zia-ul-haq',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_PENGPROFDeva_Zia_Ul_Haq',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERIPSDM_LABIB',
    1,
    TRUE
),
(
    'PSDM',
    'Ferdian Amartha',
    'Direktur Jenderal Akselerasi Potensi Mahasiswa',
    'DIRJEN',
    'local-member-psdm-ferdian-amartha',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_PSDM_FERDIAN',
    2,
    FALSE
),
(
    'PSDM',
    'Qheisya Luna Alifa',
    'Direktur Jenderal Kaderisasi',
    'DIRJEN',
    'local-member-psdm-qheisya-luna-alifa',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJENPSDM_LUNA',
    3,
    FALSE
),

-- ============================================================
-- LINGKUNGAN HIDUP
--
-- Nama Dirjen Lingkungan Hidup: Dhafa Rhofiq Firdaus, Febria Rizqi Aulia S, Rifqah Azizah.
-- ============================================================

(
    'LH',
    'Zaskia Darojah',
    'Menteri Lingkungan Hidup',
    'MENTERI',
    'local-member-lh-zaskia-darojah',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_LINGKUNGAN_HIDUP_ZASKIA',
    1,
    TRUE
),
(
    'LH',
    'Dhafa Rofiq Firdaus',
    'Direktur Jenderal Pencegahan dan Pengendalian Kerusakan Lingkungan',
    'DIRJEN',
    'local-member-lh-advokasi-kajian-strategis',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_LINGKUNGAN_HIDUP_DAFFA',
    2,
    FALSE
),
(
    'LH',
    'Febria Rizqi Aulia Sandy',
    'Direktur Jenderal Pengembangan dan Preservasi Sumber Daya Alam',
    'DIRJEN',
    'local-member-lh-pencegahan-pengendalian',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_LINGKUNGAN_HIDUP_FEBRIA',
    3,
    FALSE
),
(
    'LH',
    'Rifqah Azizah',
    'Direktur Jenderal Advokasi Kajian Strategis Lingkungan',
    'DIRJEN',
    'local-member-lh-pengembangan-preservasi',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_LINGKUNGANHIDUP_AZIZAH',
    4,
    FALSE
),

-- ============================================================
-- KEMENTERIAN KESEHATAN
--
-- Nama Dirjen Kesehatan: M. Zain Hisyam Al F. Z, Cyintia Uzliva Salsabilla.
-- ============================================================

(
    'KEMENKES',
    'An Nadhofah Adlin',
    'Menteri Kesehatan',
    'MENTERI',
    'local-member-kemenkes-an-nadhofah-adlin',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERIKESEHATAN_AN_NADHOFAH',
    1,
    TRUE
),
(
    'KEMENKES',
    'M. Zain Hisyam Al Fikri Zuhdi',
    'Direktur Jenderal Kesejahteraan Psikososial, Inklusivitas, dan Penguatan Kesehatan Mahasiswa',
    'DIRJEN',
    'local-member-kemenkes-kesehatan-masyarakat',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_KESEHATAN_ZEIN',
    2,
    FALSE
),
(
    'KEMENKES',
    'Cyintia Uzliva Salsabilla',
    'Direktur Jenderal Kesehatan Masyarakat dan Pengendalian Penyakit',
    'DIRJEN',
    'local-member-kemenkes-psikososial-inklusivitas',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_KESEHATAN_CYINTIA',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_OENGMAS_SAHRUL',
    1,
    TRUE
),
(
    'PENGMAS',
    'Fahrur Rosi',
    'Direktur Jenderal Pengembangan Desa Mitra',
    'DIRJEN',
    'local-member-pengmas-fahrur-rosi',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_PENGABDIAN_MASYARAKAT_ROSI',
    2,
    FALSE
),
(
    'PENGMAS',
    'Mohammad Wahyu Ramadani',
    'Direktur Jenderal Sosial Masyarakat',
    'DIRJEN',
    'local-member-pengmas-m-wahyu',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_PENGMAS_WAHYU',
    3,
    FALSE
),

-- ============================================================
-- HUBUNGAN LUAR
-- ============================================================

(
    'HUBLU',
    'Zalfaa’ Putri Arfiliesia',
    'Menteri Hubungan Luar',
    'MENTERI',
    'local-member-hublu-zalfaa-putri-arfiliesia',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_HUBUNGAN_LUAR_ZALFA',
    1,
    TRUE
),
(
    'HUBLU',
    'Billal Syahdan Arrafah',
    'Direktur Jenderal Mitra Nasional',
    'DIRJEN',
    'local-member-hublu-billal-syahdan-arrafah',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_HUBLU_BILLAL',
    2,
    FALSE
),
(
    'HUBLU',
    'Hafihz Naufal Utomo',
    'Direktur Jenderal Mitra Internasional',
    'DIRJEN',
    'local-member-hublu-hafiz-naufal-utomo',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_HUBLU_Hafiz_Naufal_Utomo',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/1',
    1,
    TRUE
),
(
    'MEDINFO',
    'Ocha Della Fitriani',
    'Direktur Jenderal Publikasi dan Pengembangan Media Digital',
    'DIRJEN',
    'local-member-medinfo-ocha-della-fitriani',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/2',
    2,
    FALSE
),
(
    'MEDINFO',
    'Febiola Al Jannah',
    'Direktur Jenderal Visual Kreatif',
    'DIRJEN',
    'local-member-medinfo-febiola-al-jannah',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/FEBIOLA_DIRJEN_VIDKRE_MEDIA_DAN_INFORMASI',
    3,
    FALSE
),

-- ============================================================
-- SINERGITAS MAHASISWA
-- ============================================================

(
    'SINEMA',
    'Eugenius Nathaniel Christoffany',
    'Menteri Sinergitas Mahasiswa',
    'MENTERI',
    'local-member-sinema-eugenius-nathaniel',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/menteri_sinema_nathan',
    1,
    TRUE
),
(
    'SINEMA',
    'Vincentius Nolan Nathanael Selfianto',
    'Direktur Jenderal Bina Organisasi',
    'DIRJEN',
    'local-member-sinema-vicentius-nolan',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_NOLAN_SINEMA',
    2,
    FALSE
),
(
    'SINEMA',
    'Achmad Rafi Syahputra',
    'Direktur Jenderal Kolaborasi Mahasiswa',
    'DIRJEN',
    'local-member-sinema-achmad-rafi-syahputra',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_RAFI_SINEMA',
    3,
    FALSE
),

-- ============================================================
-- KEADILAN DAN KESETARAAN GENDER
-- ============================================================

(
    'KKG',
    'Farrel Ardisto Sama''i',
    'Menteri Keadilan dan Kesetaraan Gender',
    'MENTERI',
    'local-member-kkg-farrel-ardisto-samai',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_KEMENTERIAN_KEADILAN_DAN_KESTARAAN_GENDER_FARREL',
    1,
    TRUE
),
(
    'KKG',
    'Annisa Rahmawati',
    'Direktur Jenderal Edukasi Kritis Gender',
    'DIRJEN',
    'local-member-kkg-annisa-rahmawati',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_ANNISA_RAHMA',
    2,
    FALSE
),
(
    'KKG',
    'Titis Nastiti Dwi Lestari',
    'Direktur Jenderal Pemberdayaan Perempuan Progresif',
    'DIRJEN',
    'local-member-kkg-titis-nastiti-dwi',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_TITIS',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_SOSPOL_DANIEL',
    1,
    TRUE
),
(
    'SOSPOL',
    'Moch. Amir Nazaruddin',
    'Direktur Jenderal Aksi dan Propaganda',
    'DIRJEN',
    'local-member-sospol-moch-amir-nazaruddin',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERISOSPOL_AMIR',
    2,
    FALSE
),
(
    'SOSPOL',
    'Ahmad Raihan Fadhillah Yasin',
    'Direktur Jenderal Analisis Isu Strategis',
    'DIRJEN',
    'local-member-sospol-ahmad-raihan-fadhillah',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERISOSPOLRAIHAN',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERI_EDIGPRO_FARHAN_KAMIL',
    1,
    TRUE
),
(
    'EDIGPRO',
    'Argi Wicaksana',
    'Direktur Jenderal Produk Kreatif dan Investasi',
    'DIRJEN',
    'local-member-edigpro-argi',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJENEDIGPRO_ARGI_WICAKSANA',
    2,
    FALSE
),
(
    'EDIGPRO',
    'M. Yudhistira Setya Heri',
    'Direktur Jenderal UMKM dan Kewirausahaan',
    'DIRJEN',
    'local-member-edigpro-yudhistira',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJEN_EDIGPRO_YUDHISTIRA',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERIRISKEL_BISMA',
    1,
    TRUE
),
(
    'RISKEL',
    'Abdul Wahid Tamimi',
    'Direktur Jenderal Riset dan Inovasi',
    'DIRJEN',
    'local-member-riskel-ata',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJENRISKEL_ATA',
    2,
    FALSE
),
(
    'RISKEL',
    'Naswa Ayudya Sasta',
    'Direktur Jenderal Akademik dan Prestasi',
    'DIRJEN',
    'local-member-riskel-naswa',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/DIRJENNASYWA_RISKEL',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/MENTERISENIORA_ALAINA',
    1,
    TRUE
),
(
    'SENIIORA',
    'Abdullah Azzam',
    'Direktur Jenderal Pemuda dan Olahraga',
    'DIRJEN',
    'local-member-seniiora-azzam',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/ABDULLAH_AZZAM_DIRJEN_SENIORA',
    3,
    FALSE
),
(
    'SENIIORA',
    'Muhammad Fadhil Saputra',
    'Direktur Jenderal Seni dan Olahraga',
    'DIRJEN',
    'local-member-seniiora-puput',
    'https://res.cloudinary.com/rvn0bdia/image/upload/v1795411756/PUPUT_DIRJEN_SENIORA.png',
    2,
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426454/BENKAN_INEZ.png',
    1,
    TRUE
),
(
    'BENKAB',
    'Ananda Magfira Nurochma',
    'Kepala Biro Administrasi dan Keuangan',
    'KABIRO',
    'local-member-benkab-ananda-magfira-nurochma',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426455/BIRO_ANANDA.png',
    2,
    FALSE
),
(
    'BENKAB',
    'M. Salman Al Farisi Al Ghofsah',
    'Kepala Biro Kebendaharaan',
    'KABIRO',
    'local-member-benkab-m-salman-al-farisi',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/BIRO_SALMAN.png',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/KEPALA_PAK.png',
    1,
    TRUE
),
(
    'PAK',
    'Mevlana El Fithra Abdullah',
    'Deputi Tata Laksana Kerja',
    'DEPUTI',
    'local-member-pak-mevlana-el-fithra',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/DEPUTI_MEVLANA.png',
    2,
    FALSE
),
(
    'PAK',
    'Nayla Khansa Abbylia',
    'Deputi Audit dan Penjaminan Mutu',
    'DEPUTI',
    'local-member-pak-nayla-khansa-abbylia',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/DEPUTI_NAYLA.png',
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
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/SESKAB_GHEA.png',
    1,
    TRUE
),
(
    'SESKAB',
    'Fathimah Nuurun Najma As-Salman',
    'Kepala Biro Administrasi Persuratan',
    'KABIRO',
    'local-member-seskab-fathimah-nuurun',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426455/BIRO_NUURUN.png',
    2,
    FALSE
),
(
    'SESKAB',
    'Najwa Nur Hafidza',
    'Kepala Biro Arsip dan Kesekretariatan',
    'KABIRO',
    'local-member-seskab-najwa-nur-hafidza',
    'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/v1795426456/BIRO_NAJWA.png',
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
        WHERE unit.logo_asset_key COLLATE utf8mb4_unicode_ci
              = member.photo_asset_key COLLATE utf8mb4_unicode_ci
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
-- Pencocokan menggunakan ministry_id + name supaya seeder
-- dapat dijalankan ulang tanpa membuat duplikasi.
-- ============================================================

UPDATE organization_members AS organization_member
INNER JOIN ministries AS ministry
    ON ministry.id = organization_member.ministry_id
INNER JOIN seed_organization_members AS seed_member
    ON seed_member.unit_code COLLATE utf8mb4_unicode_ci
     = ministry.code COLLATE utf8mb4_unicode_ci
    AND seed_member.name COLLATE utf8mb4_unicode_ci
     = organization_member.name COLLATE utf8mb4_unicode_ci
LEFT JOIN media_assets AS photo
    ON photo.imagekit_file_id COLLATE utf8mb4_unicode_ci
     = seed_member.photo_asset_key COLLATE utf8mb4_unicode_ci
SET
    organization_member.position = seed_member.position,
    organization_member.position_type = seed_member.position_type,
    organization_member.photo_media_id = COALESCE(organization_member.photo_media_id, photo.id),
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
    ON ministry.code COLLATE utf8mb4_unicode_ci
     = seed_member.unit_code COLLATE utf8mb4_unicode_ci
LEFT JOIN media_assets AS photo
    ON photo.imagekit_file_id COLLATE utf8mb4_unicode_ci
     = seed_member.photo_asset_key COLLATE utf8mb4_unicode_ci
WHERE NOT EXISTS (
    SELECT 1
    FROM organization_members AS existing_member
    WHERE
        existing_member.ministry_id = ministry.id
        AND existing_member.name COLLATE utf8mb4_unicode_ci
         = seed_member.name COLLATE utf8mb4_unicode_ci
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
