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
-- Password: 53066444
-- ============================================================

(
    'Admin BEM UNAIR',
    'admin@bem.unair.ac.id',
    '$2a$10$ixhu.4fm5xE4cnTD2u0Z0ufQDc0JpTYmJhx/RASEi5Pv/5JvF6UMG',
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
-- Password: 51975298
-- ============================================================

(
    'Johanes Richard Darmawan',
    'menko.kemahasiswaan@bem.unair.ac.id',
    '$2a$10$0VnW1WGgWrjC3Esjr0x7m.sBy76GhomOCuqs2sp3yR44eBl4emFou',
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
-- Password: 42038708
-- ============================================================

(
    'Muhammad Ziyaad Difa Ul Haq',
    'menko.kemasyarakatan@bem.unair.ac.id',
    '$2a$10$CRGHA5/IHNdHXlDJ72xGB.jOft7PELu5m95O.KE.45Gj1qDxtg1NO',
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
-- Password: 26770499
-- ============================================================

(
    'Andrean Marcellino',
    'menko.kominfo@bem.unair.ac.id',
    '$2a$10$HLFWEEd3p6yNLPbmonrCk.k/e9Uta9X5gV7P0/9pUGmGavuB6zjby',
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
-- Password: 33605240
-- ============================================================

(
    'Dewa Pranata Putra Pratama',
    'menko.pergerakan@bem.unair.ac.id',
    '$2a$10$N1y9JpkjQdBFV5Ey5o987uI42KaAqQRtDYKhbkl/69f7hyFfIZeu6',
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
-- Password: 62270982
-- ============================================================

(
    'Nasua Okta Kusuma Dewi Rofiq',
    'menko.ppk@bem.unair.ac.id',
    '$2a$10$LqHfhO7euPvBodvg/KfF6OQEcjvc5Gm/tKEvpD4f4VlD8QGHnc0hG',
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
-- Password: 66400846
-- ============================================================

(
    'Elsa Putri Aprizalni',
    'menteri.adkesma@bem.unair.ac.id',
    '$2a$10$3CXsAlxvagYIjTNYJgiw3OTzyIB30HZZX/fynScMXTNggs43nwgeG',
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
-- Password: 30654122
-- ============================================================

(
    'M. Rifqy Yusuf',
    'menteri.pengprof@bem.unair.ac.id',
    '$2a$10$4b7dcS8UMdOJ6tXIDdZgzuZspgv44hjwEjZPnm0TDutEMnhZzn2Ne',
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
-- Password: 17576656
-- ============================================================

(
    'Labib Azka Asysyafiq',
    'menteri.psdm@bem.unair.ac.id',
    '$2a$10$DH147wtZAInQWDv1r7rXz.JQ2C86DA/fHMTo6MRUj5gziqYLh1yfG',
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
-- Password: 43526893
-- ============================================================

(
    'Zaskia Darojah',
    'menteri.lh@bem.unair.ac.id',
    '$2a$10$wZ0DRlig0tiTZHUqL9AX1uw4QcGxsOWOToLAT98npAxvCjiir3xqe',
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
-- Password: 91823126
-- ============================================================

(
    'An Nadhofah Adlin',
    'menteri.kemenkes@bem.unair.ac.id',
    '$2a$10$awANxxNJcVGmQuRc83t9fORNHNLqN2AdTBJ1IH8VNCjdr1w1eSx3.',
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
-- Password: 75669751
-- ============================================================

(
    'Sahrul Efendi',
    'menteri.pengmas@bem.unair.ac.id',
    '$2a$10$lc8no1gy2BT.4EuBl.Mp8uvBtqs0IoTQ8CdgcrO3muxM8MaP4DhSi',
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
-- Password: 57911111
-- ============================================================

(
    'Zalfaa Putri Arfiliesia',
    'menteri.hublu@bem.unair.ac.id',
    '$2a$10$nCKezeNBhdUO4zhR5IrjduyJdhof4dd2VFOkNY3.Mr3xDTRP8cqcO',
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
-- Password: 62218191
-- ============================================================

(
    'Ganesya Intantalia (Caca)',
    'caca.medinfo@bem.unair.ac.id',
    '$2a$10$rKzrvdy6kxPxylEdXXHjceCjJpYP0hRyA3D1DmU6TKQ2PT.4U0bNK',
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
-- Password: 86586573
-- ============================================================

(
    'Ocha Della Fitriani',
    'ocha.medinfo@bem.unair.ac.id',
    '$2a$10$R.l.8UYcVP.yvNeC6NfhH.t7cGhiuTueebeqgWHF.fuGiwUL/SU5S',
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
-- Password: 63545998
-- ============================================================

(
    'Febiola Al Jannah',
    'febi.medinfo@bem.unair.ac.id',
    '$2a$10$jpW7zCkocSxty1uOdvGZrud7OPg8dqJoP547unh2TvBPcfEjl8cHW',
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
-- Password: 23671369
-- ============================================================

(
    'Eugenius Nathaniel',
    'menteri.sinema@bem.unair.ac.id',
    '$2a$10$XRgOwS3JpYBRHkMghiVK0eeGF/MRuNkD7tuP81allFlelz0dWV9Vu',
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
-- Password: 57023964
-- ============================================================

(
    'Farrel Ardisto Samai',
    'menteri.kkg@bem.unair.ac.id',
    '$2a$10$8WIQmyc7Khv52wHs29qEwulkm8IrYT3dCKdgIUmnlNyd6lTTW1b9O',
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
-- Password: 52746829
-- ============================================================

(
    'Daniel Nikon Martua Situmorang',
    'menteri.sospol@bem.unair.ac.id',
    '$2a$10$qrXu/n0LeJerzkm/ZJ6Nn.LDU86Bcwy.t4BqRQ8XMHJ7vJiTKBn/S',
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
-- Password: 87287946
-- ============================================================

(
    'Farhan Kamil',
    'menteri.edigpro@bem.unair.ac.id',
    '$2a$10$DRqi3JrKyOvZ1y4SRoiq7u/2GT30fL39j8g6.MgfJwXMYK1XeY6jG',
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
-- Password: 35608695
-- ============================================================

(
    'Farrel Bisma Abyakta',
    'menteri.riskel@bem.unair.ac.id',
    '$2a$10$zHl5b7933VhIVqgYvDy3c.NdAoz8R8PM8ACSSYSHit3kClCuyBYOu',
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
-- Password: 91646330
-- ============================================================

(
    'Alaina Atwa Awali Harahap',
    'menteri.seniiora@bem.unair.ac.id',
    '$2a$10$h3nSH06JoY5BqLoHvDUQOeehSUCaOySPkHLJEXZCrsogtog6cbd5m',
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
-- Password: 71318325
-- ============================================================

(
    'Inez Faradina Kasih',
    'menteri.benkab@bem.unair.ac.id',
    '$2a$10$SN8TJDvcpNOS3CrYGbkufuTjuN7RcOVYFc6PdasX1WZSmiECGKv1G',
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
-- Password: 36909827
-- ============================================================

(
    'Deri Bayu Setiawan',
    'kepala.pak@bem.unair.ac.id',
    '$2a$10$tiVA1VGOx9D.mEoZxevJDOnfo2O2r6MjMcNuH2uVR4wq/HRViUIcy',
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
-- Password: 49594396
-- ============================================================

(
    'Gheahaq Danty El Zahra',
    'menteri.seskab@bem.unair.ac.id',
    '$2a$10$A2xTO3gNuvNO9qf.MXHKkuS8.3B.iRhYdsSM.oJq2ubiae.v6310m',
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