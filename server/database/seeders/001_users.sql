-- ============================================================
-- USER SEEDER KABINET CERITA LOKA
--
-- Aturan:
-- - Semua akun selain Medinfo: role ADMIN
-- - Akun Medinfo: role ADMIN_MEDINFO
-- - Satu akun per kementerian
-- - Medinfo memiliki 3 akun: Caca, Ocha, Febi
-- - Dev-only (compose production hanya menjalankan `migrate`, bukan seeder
--   ini). Semua akun berbagi SATU password dev yang sama, diambil dari env
--   SEED_USER_PASSWORD (lihat server/.env.example) dan di-hash saat seeder
--   dijalankan oleh database/seeder.go. Tidak ada password — plaintext
--   maupun hash — yang pernah ditulis ke file ini.
-- ============================================================

START TRANSACTION;

INSERT INTO users (
    name,
    email,
    password_hash,
    role,
    ministry,
    ministry_id,
    phone,
    created_at,
    updated_at
) VALUES

-- ============================================================
-- ADMIN UTAMA
-- ============================================================

(
    'Admin BEM UNAIR',
    'admin@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN_MEDINFO',
    NULL,
    NULL,
    '6281200000001',
    NOW(),
    NOW()
),

-- ============================================================
-- MENKO KEMAHASISWAAN
-- Johanes Richard Darmawan
-- ============================================================

(
    'Johanes Richard Darmawan',
    'menko.kemahasiswaan@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'MENKO_KEMAHASISWAAN',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MENKO_KEMAHASISWAAN'
        LIMIT 1
    ),
    '6281200000002',
    NOW(),
    NOW()
),

-- ============================================================
-- MENKO KEMASYARAKATAN
-- Muhammad Ziyaad Difa Ul Haq
-- ============================================================

(
    'Muhammad Ziyaad Difa Ul Haq',
    'menko.kemasyarakatan@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'MENKO_KEMASYARAKATAN',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MENKO_KEMASYARAKATAN'
        LIMIT 1
    ),
    '6281200000003',
    NOW(),
    NOW()
),

-- ============================================================
-- MENKO KOMINFO
-- Andrean Marcellino
-- ============================================================

(
    'Andrean Marcellino',
    'menko.kominfo@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'MENKO_KOMINFO',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MENKO_KOMINFO'
        LIMIT 1
    ),
    '6281200000004',
    NOW(),
    NOW()
),

-- ============================================================
-- MENKO PERGERAKAN
-- Dewa Pranata Putra Pratama
-- ============================================================

(
    'Dewa Pranata Putra Pratama',
    'menko.pergerakan@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'MENKO_PERGERAKAN',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MENKO_PERGERAKAN'
        LIMIT 1
    ),
    '6281200000005',
    NOW(),
    NOW()
),

-- ============================================================
-- MENKO PPK
-- Nasua Okta Kusuma Dewi Rofiq
-- ============================================================

(
    'Nasua Okta Kusuma Dewi Rofiq',
    'menko.ppk@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'MENKO_PPK',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MENKO_PPK'
        LIMIT 1
    ),
    '6281200000006',
    NOW(),
    NOW()
),

-- ============================================================
-- ADKESMA
-- Elsa Putri Aprizalni
-- ============================================================

(
    'Elsa Putri Aprizalni',
    'menteri.adkesma@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'ADKESMA',
    (
        SELECT id
        FROM ministries
        WHERE code = 'ADKESMA'
        LIMIT 1
    ),
    '6281200000007',
    NOW(),
    NOW()
),

-- ============================================================
-- PENGEMBANGAN PROFESI
-- M. Rifqy Yusuf
-- ============================================================

(
    'M. Rifqy Yusuf',
    'menteri.pengprof@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'PENGPROF',
    (
        SELECT id
        FROM ministries
        WHERE code = 'PENGPROF'
        LIMIT 1
    ),
    '6281200000008',
    NOW(),
    NOW()
),

-- ============================================================
-- PSDM
-- Labib Azka Asysyafiq
-- ============================================================

(
    'Labib Azka Asysyafiq',
    'menteri.psdm@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'PSDM',
    (
        SELECT id
        FROM ministries
        WHERE code = 'PSDM'
        LIMIT 1
    ),
    '6281200000009',
    NOW(),
    NOW()
),

-- ============================================================
-- LINGKUNGAN HIDUP
-- Zaskia Darojah
-- ============================================================

(
    'Zaskia Darojah',
    'menteri.lh@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'LH',
    (
        SELECT id
        FROM ministries
        WHERE code = 'LH'
        LIMIT 1
    ),
    '6281200000010',
    NOW(),
    NOW()
),

-- ============================================================
-- KEMENTERIAN KESEHATAN
-- An Nadhofah Adlin
-- ============================================================

(
    'An Nadhofah Adlin',
    'menteri.kemenkes@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'KEMENKES',
    (
        SELECT id
        FROM ministries
        WHERE code = 'KEMENKES'
        LIMIT 1
    ),
    '6281200000011',
    NOW(),
    NOW()
),

-- ============================================================
-- PENGABDIAN MASYARAKAT
-- Sahrul Efendi
-- ============================================================

(
    'Sahrul Efendi',
    'menteri.pengmas@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'PENGMAS',
    (
        SELECT id
        FROM ministries
        WHERE code = 'PENGMAS'
        LIMIT 1
    ),
    '6281200000012',
    NOW(),
    NOW()
),

-- ============================================================
-- HUBUNGAN LUAR
-- Zalfaa Putri Arfiliesia
-- ============================================================

(
    'Zalfaa Putri Arfiliesia',
    'menteri.hublu@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'HUBLU',
    (
        SELECT id
        FROM ministries
        WHERE code = 'HUBLU'
        LIMIT 1
    ),
    '6281200000013',
    NOW(),
    NOW()
),

-- ============================================================
-- MEDINFO: CACA
-- Menteri Media dan Informasi
-- ============================================================

(
    'Ganesya Intantalia (Caca)',
    'caca.medinfo@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN_MEDINFO',
    'MEDINFO',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MEDINFO'
        LIMIT 1
    ),
    '6281200000014',
    NOW(),
    NOW()
),

-- ============================================================
-- MEDINFO: OCHA
-- Direktur Jenderal Publikasi
-- ============================================================

(
    'Ocha Della Fitriani',
    'ocha.medinfo@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN_MEDINFO',
    'MEDINFO',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MEDINFO'
        LIMIT 1
    ),
    '6281200000015',
    NOW(),
    NOW()
),

-- ============================================================
-- MEDINFO: FEBI
-- Direktur Jenderal Visual
-- ============================================================

(
    'Febiola Al Jannah',
    'febi.medinfo@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN_MEDINFO',
    'MEDINFO',
    (
        SELECT id
        FROM ministries
        WHERE code = 'MEDINFO'
        LIMIT 1
    ),
    '6281200000016',
    NOW(),
    NOW()
),

-- ============================================================
-- SINERGITAS MAHASISWA
-- Eugenius Nathaniel
-- ============================================================

(
    'Eugenius Nathaniel',
    'menteri.sinema@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'SINEMA',
    (
        SELECT id
        FROM ministries
        WHERE code = 'SINEMA'
        LIMIT 1
    ),
    '6281200000017',
    NOW(),
    NOW()
),

-- ============================================================
-- KEADILAN DAN KESETARAAN GENDER
-- Farrel Ardisto Samai
-- ============================================================

(
    'Farrel Ardisto Samai',
    'menteri.kkg@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'KKG',
    (
        SELECT id
        FROM ministries
        WHERE code = 'KKG'
        LIMIT 1
    ),
    '6281200000018',
    NOW(),
    NOW()
),

-- ============================================================
-- SOSIAL DAN POLITIK
-- Daniel Nikon Martua Situmorang
-- ============================================================

(
    'Daniel Nikon Martua Situmorang',
    'menteri.sospol@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'SOSPOL',
    (
        SELECT id
        FROM ministries
        WHERE code = 'SOSPOL'
        LIMIT 1
    ),
    '6281200000019',
    NOW(),
    NOW()
),

-- ============================================================
-- EKONOMI DIGITAL DAN PRODUK KREATIF
-- Farhan Kamil
-- ============================================================

(
    'Farhan Kamil',
    'menteri.edigpro@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'EDIGPRO',
    (
        SELECT id
        FROM ministries
        WHERE code = 'EDIGPRO'
        LIMIT 1
    ),
    '6281200000020',
    NOW(),
    NOW()
),

-- ============================================================
-- RISET DAN KEILMUAN
-- Farrel Bisma Abyakta
-- ============================================================

(
    'Farrel Bisma Abyakta',
    'menteri.riskel@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'RISKEL',
    (
        SELECT id
        FROM ministries
        WHERE code = 'RISKEL'
        LIMIT 1
    ),
    '6281200000021',
    NOW(),
    NOW()
),

-- ============================================================
-- SENI DAN OLAHRAGA
-- Alaina Atwa Awali Harahap
-- ============================================================

(
    'Alaina Atwa Awali Harahap',
    'menteri.seniiora@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'SENIIORA',
    (
        SELECT id
        FROM ministries
        WHERE code = 'SENIIORA'
        LIMIT 1
    ),
    '6281200000022',
    NOW(),
    NOW()
),

-- ============================================================
-- BENDAHARA KABINET
-- Inez Faradina Kasih
-- ============================================================

(
    'Inez Faradina Kasih',
    'menteri.benkab@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'BENKAB',
    (
        SELECT id
        FROM ministries
        WHERE code = 'BENKAB'
        LIMIT 1
    ),
    '6281200000023',
    NOW(),
    NOW()
),

-- ============================================================
-- PENDAYAGUNAAN APARATUR KABINET
-- Deri Bayu Setiawan
-- ============================================================

(
    'Deri Bayu Setiawan',
    'kepala.pak@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'PAK',
    (
        SELECT id
        FROM ministries
        WHERE code = 'PAK'
        LIMIT 1
    ),
    '6281200000024',
    NOW(),
    NOW()
),

-- ============================================================
-- SEKRETARIS KABINET
-- Gheahaq Danty El Zahra
-- ============================================================

(
    'Gheahaq Danty El Zahra',
    'menteri.seskab@bem.unair.ac.id',
    '__SEED_PASSWORD_HASH__',
    'ADMIN',
    'SESKAB',
    (
        SELECT id
        FROM ministries
        WHERE code = 'SESKAB'
        LIMIT 1
    ),
    '6281200000025',
    NOW(),
    NOW()
)

ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    password_hash = VALUES(password_hash),
    role = VALUES(role),
    ministry = VALUES(ministry),
    ministry_id = VALUES(ministry_id),
    phone = VALUES(phone),
    updated_at = NOW();

COMMIT;

-- ============================================================
-- VERIFIKASI HASIL USER SEEDER
-- ============================================================

SELECT
    users.id,
    users.name,
    users.email,
    users.role,
    users.ministry,
    ministries.name AS ministry_name,
    users.phone
FROM users
LEFT JOIN ministries
    ON ministries.id = users.ministry_id
ORDER BY
    CASE
        WHEN users.email = 'admin@bem.unair.ac.id' THEN 0
        WHEN ministries.unit_type = 'MENKO' THEN 1
        WHEN users.ministry = 'MEDINFO' THEN 2
        ELSE 3
    END,
    ministries.display_order ASC,
    users.name ASC;