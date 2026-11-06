-- ============================================================
-- USER SEEDER KABINET CERITA LOKA
--
-- Aturan:
-- - Semua akun selain Medinfo: role ADMIN
-- - Akun Medinfo: role ADMIN_MEDINFO
-- - Satu akun per kementerian
-- - Medinfo memiliki 3 akun: Caca, Ocha, Febi
-- - Password berbeda dan terdiri dari 8 digit
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
    '$2a$10$r1pia3hWuIcj/mUOE4WYiuLQ.H9Nd5KwgR6xoxY/Nct3sEe1b.VXG',
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
    '$2a$10$Oef94oZusjkGYHbVjf6awuaMFenWPN/TiPqKFnsfqBB6Ylzyuk.R6',
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
    '$2a$10$qPX/06lAa93rbbxAu19e3.c6I3tgyYfcxdKizoUIlFDrwxGbW.3D.',
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
    '$2a$10$GntLR1rdqq/BI63lxRuY3u639ZiqEc8BBY6gkmcDupIBOfFdvFTnS',
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
    '$2a$10$IKXudAcmNi6W08WuWbAnW.Sc2kxuv1ur19d2e6mi6d4ew5./A4FNi',
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
    '$2a$10$iuVHVMdgzbZM96cFc/ymUOMfn2um8W9e4x1G4qlyenU7on21rs0xC',
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
    '$2a$10$A3hkE1tdTfYzh74p5mQliOmfpCK0tfhC2SXlxxLBj5W3EeU.92cGu',
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
    '$2a$10$RZe5.8ttuYvP/ShWwjxlmuphHUTLUE2XzhNeSH6uQDmG5cZ.GkKqG',
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
    '$2a$10$Y62KUut2z841yA3yHMqKUu2kzBJ26kZDbgVz9R5j5oTuTwKRbYGjK',
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
    '$2a$10$usU8RNR7etoS/Fzru9G9auA5Xb/ZwQVO4SOpIaHi76w3mgqVt6j32',
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
    '$2a$10$StDMNiWfJoIdgIGuXzkuIOzDRQC.n/JcF5vktLvP8BUKa/4m5HBmC',
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
    '$2a$10$WaPpjkgExDiCy3qa2fqWteEcAhw.QeSPjCcLeE1Jh3jVe5BH/3Bfi',
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
    '$2a$10$1BMQrrYZNy1MVkho9B9V1uiYzIlQre2Hs9uBKrPb6bb4iygySAjJ2',
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
    '$2a$10$70EjhwZ5Uj.aHmZkOmHMCe.3RDX2B14ib6VzmbXfAS70DNwO8qB22',
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
    '$2a$10$KWUt9zk1GjjGl0ao6b25oeviFQKW41UiI3bAHRA0XGzxWOQAIZXB6',
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
    '$2a$10$pRzCfPCEEd984QVcqtRfjuO5lJUF9rJg.FcCunJTxVhfijPlXpM7q',
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
    '$2a$10$sFQ1isIdimtF2Zp6SgoJYeeZinRAsw/XmfHrQjqyHplHr598t8Z3.',
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
    '$2a$10$K24gBVYvv4Rpxscp4Ou9HOLBbqacqAshCnmo7jppYvNQLjGwDVauG',
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
    '$2a$10$Ydb9OprRJA5NAgljUPvIx.5BmRzUBY8muLG3t6eV4uD4/dyp0SXo2',
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
    '$2a$10$8sj0DXJpkS/M0LYirsmode5pszF1Pvj0lLDHwRgquP/yP2qSnhQw6',
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
    '$2a$10$niEcMv2Ri4hW9zhkwKlQ8OwqtfcOMQMWs1T/zSYq8.B5w10vrdr/S',
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
    '$2a$10$xK3wmlQfLc64qfxEtdGWhuDCdenHiG8ZXS2avEY6TJ5RK6xgj3sRq',
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
    '$2a$10$RjwIRno/Pi2tCJeyb6tlUuKIqWyp41fbmWP4uklGBzkAOLrbR/qse',
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
    '$2a$10$tX6E8G95KQf4SumaeUTpXe8.s5smME7/nwb5Hsd2Y29lYA.wrnfxa',
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
    '$2a$10$0MvAsSSBOnHUys7bk6SSTOUo9qwhMSDJNy6P2424zURWeH1OFDCfO',
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