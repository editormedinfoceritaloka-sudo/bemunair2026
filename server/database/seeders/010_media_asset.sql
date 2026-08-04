-- ==========================================
-- 010_media_assets.sql
-- Part 1 - MENKO
-- ==========================================
DELETE FROM organization_members WHERE photo_media_id IN (SELECT id FROM media_assets WHERE purpose = 'ORGANIZATION_MEMBER');
DELETE FROM media_assets WHERE purpose = 'ORGANIZATION_MEMBER';
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
) VALUES

(
    NULL,
    '6a7065335c7cd75eb859d9df',
    NULL,
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Foto%20Menko/PPK.png?updatedAt=1785750836987',
    NULL,
    'NASUA OKTA KUSUMA DEWI ROFIQ',
    'Menko PPK',
    NULL,
    'image/png',
    1,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE',
    NOW(),
    NOW()
),

(
    NULL,
    '6a7065195c7cd75eb85859c5',
    NULL,
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Foto%20Menko/Kominfo.png?updatedAt=1785750810828',
    NULL,
    'ANDREAN MARCELLINO',
    'Menko Kominfo',
    NULL,
    'image/png',
    1,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE',
    NOW(),
    NOW()
),

(
    NULL,
    '6a7064e45c7cd75eb85513a5',
    NULL,
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Foto%20Menko/Pergerakan.png?updatedAt=1785750757458',
    NULL,
    'DEWA PRANATA PUTRA PRATAMA',
    'Menko Pergerakan',
    NULL,
    'image/png',
    1,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE',
    NOW(),
    NOW()
),

(
    NULL,
    '6a7065125c7cd75eb857f1f6',
    NULL,
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Foto%20Menko/Kemahasiswaan.png?updatedAt=1785750803110',
    NULL,
    'JOHANES RICHARD DARMAWAN',
    'Menko Kemahasiswaan',
    NULL,
    'image/png',
    1,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE',
    NOW(),
    NOW()
),

(
    NULL,
    '6a7064e55c7cd75eb8552c5c',
    NULL,
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Foto%20Menko/Kemasyarakatan.png?updatedAt=1785750758571',
    NULL,
    'MUHAMMAD ZIYAAD DIFAA’UL HAQ',
    'Menko Kemasyarakatan',
    NULL,
    'image/png',
    1,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE',
    NOW(),
    NOW()
);

-- ==========================================
-- Part 2 - MENTERI
-- ==========================================

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
) VALUES

(
NULL,
'6a70652e5c7cd75eb859b9d0',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Riset%20dan%20Keilmuan/Menteri%20bismo.png?updatedAt=1785750832137',
NULL,
'Farrel Bisma Abyakta',
'Menteri Riset dan Keilmuan',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7065155c7cd75eb8583bf4',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Seni%20dan%20Olahraga/Menteri%20Alaina.png?updatedAt=1785750806652',
NULL,
'Alaina Atwa Awali Harahap',
'Menteri Seni dan Olahraga',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7065105c7cd75eb857d381',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Bendahara%20Kabinet/Menteri%20Benkab.png?updatedAt=1785750801977',
NULL,
'INEZ FARADINA KASIH',
'Menteri Bendahara Kabinet',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064da5c7cd75eb85454bd',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Sekretaris%20Kabinet/Menteri%20Gheahaq.png?updatedAt=1785750747086',
NULL,
'GHEAHAQ DANTY EL ZAHRA',
'Menteri Sekretaris Kabinet',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064be5c7cd75eb852f244',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Sosial%20dan%20Politik/MENTERI%20SOSPOL%20-%20DANIEL%20NIKON%20MARTUA%20SITUMORANG.png?updatedAt=1785750719564',
NULL,
'Daniel Nikon Martua Situmorang',
'Menteri Sosial dan Politik',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064b95c7cd75eb8529aac',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/KKG/MENTERI%20KKG%20-%20FARREL%20ARDISTO%20SAMA_I.png?updatedAt=1785750714764',
NULL,
'Farrel Ardisto Sama''i',
'Menteri Keadilan dan Kesetaraan Gender',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064c25c7cd75eb8530350',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/PENGPROF/MENTERI%20PENGPROF%20-%20M.%20RIFQY%20YUSUF.png?updatedAt=1785750722778',
NULL,
'Muhammad Rifqy Yusuf',
'Menteri Pengembangan Profesi',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064da5c7cd75eb854566e',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/PSDM/MENTERI%20PSDM%20-%20LABIB%20AZKA%20ASYSYAFIQ.png?updatedAt=1785750747167',
NULL,
'Labib Azka Asysyafiq',
'Menteri Pengembangan Sumber Daya Mahasiswa',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064eb5c7cd75eb85596d5',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/EDIGPRO/Menteri%20Kamil.png?updatedAt=1785750764175',
NULL,
'Farhan Kamil',
'Menteri Ekonomi Digital dan Produk Kreatif',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064d75c7cd75eb854498a',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Medinfo/MENTERI%20MEDINFO%20-%20GANESYA%20INTANTALIA.png?updatedAt=1785750744324',
NULL,
'Ganesya Intantalia',
'Menteri Media dan Informasi',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064a35c7cd75eb8513414',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Pengabdian%20Masyarakat/Sahrul-menteri.png?updatedAt=1785750692267',
NULL,
'Sahrul Efendi',
'Menteri Pengabdian Masyarakat',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064ce5c7cd75eb853c13f',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Hubungan%20Luar/MENTERI%20HUBLU%20-%20ZALFAA_%20PUTRI%20ARFILIESIA.png?updatedAt=1785750734746',
NULL,
'Zalfaa'' Putri Arfiliesia',
'Menteri Hubungan Luar',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064cd5c7cd75eb853afd1',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/ADKESMA/MENTERI%20ADKESMA%20-%20ELSA%20PUTRI%20APRIZALNI.png?updatedAt=1785750734091',
NULL,
'Elsa Putri Aprizalni',
'Menteri Advokasi dan Kesejahteraan Mahasiswa',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a7064cf5c7cd75eb853eddb',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Sinergitas%20Mahasiswa/MENTERI%20SINEMA%20-%20EUGENIUS%20NATHANIEL%20C.png?updatedAt=1785750736828',
NULL,
'Eugenius Nathaniel Christoffany',
'Menteri Sinergitas Mahasiswa',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a70dab85c7cd75eb823e327',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Lingkungan%20Hidup/WhatsApp_Image_2026-08-04_at_01.09.38-removebg-preview%20(1).png',
NULL,
'Zaskia Darojah',
'Menteri Lingkungan Hidup',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
),

(
NULL,
'6a70db665c7cd75eb829c28a',
NULL,
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/MENKES/WhatsApp%20Image%202026-08-04%20at%2001-remove-bg-io.png',
NULL,
'An Nadhofah Adlin',
'Menteri Kesehatan',
NULL,
'image/png',
1,
NULL,
NULL,
'ORGANIZATION_MEMBER',
'ACTIVE',
NOW(),
NOW()
);

-- =====================================================
-- BIRO BENDAHARA KABINET
-- =====================================================

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
    status
)
VALUES
(
    NULL,
    '6a7065325c7cd75eb859d1a5',
    '-',
    'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Bendahara%20Kabinet/Kabiro%20Ananda%20Magfira.png?updatedAt=1785750835559',
    '-',
    'Kabiro Administrasi dan Keuangan',
    'Ananda Magfira Nurochma',
    '-',
    'image/png',
    0,
    NULL,
    NULL,
    'ORGANIZATION_MEMBER',
    'ACTIVE'
);

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
    ended_at
)
VALUES (
    (SELECT id FROM ministries WHERE slug='bendahara-kabinet'),
    NULL,
    'Ananda Magfira Nurochma',
    'Kabiro Administrasi dan Keuangan',
    'KABIRO',
    '-',
    '-',
    '-',
    '-',
    LAST_INSERT_ID(),
    1,
    FALSE,
    TRUE,
    NULL,
    NULL
);

----------------------------------------------------------

INSERT INTO media_assets (
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a70652c5c7cd75eb859ad4b',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Bendahara%20Kabinet/Kabiro%20M.%20Salman.png?updatedAt=1785750829742',
'-',
'Kabiro Kebendaharaan',
'M. Salman Al Farisi Al Ghofsah',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='bendahara-kabinet'),
NULL,
'M. Salman Al Farisi Al Ghofsah',
'Kabiro Kebendaharaan',
'KABIRO',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
2,
FALSE,
TRUE,
NULL,
NULL
);

-- =====================================================
-- SEKRETARIS KABINET
-- =====================================================

INSERT INTO media_assets(
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a70652d5c7cd75eb859b1f6',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Sekretaris%20Kabinet/Kabiro%20Fathimah.png?updatedAt=1785750830723',
'-',
'Kabiro Administrasi Persuratan',
'Fathimah Nuurun Najma As-Salman',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='sekretaris-kabinet'),
NULL,
'Fathimah Nuurun Najma As-Salman',
'Kabiro Administrasi Persuratan',
'KABIRO',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
1,
FALSE,
TRUE,
NULL,
NULL
);

----------------------------------------------------------

INSERT INTO media_assets(
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a7065155c7cd75eb858391e',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Sekretaris%20Kabinet/Kabiro%20Najwa%20Nur.Png?updatedAt=1785750805984',
'-',
'Kabiro Arsip dan Kesekretariatan',
'Najwa Nur Hafidza',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='sekretaris-kabinet'),
NULL,
'Najwa Nur Hafidza',
'Kabiro Arsip dan Kesekretariatan',
'KABIRO',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
2,
FALSE,
TRUE,
NULL,
NULL
);

-- =====================================================
-- PENDAYAGUNAAN APARATUR KABINET
-- =====================================================

INSERT INTO media_assets(
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a7065325c7cd75eb859d25e',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Pendayagunaan%20Aparatur%20Kabinet/Kepala_%20PAK.png?updatedAt=1785750835658',
'-',
'Kepala Pendayagunaan Aparatur Kabinet',
'DERI BAYU SETIAWAN',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='pendayagunaan-aparatur-kabinet'),
NULL,
'DERI BAYU SETIAWAN',
'Kepala Pendayagunaan Aparatur Kabinet',
'KEPALA',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
1,
TRUE,
TRUE,
NULL,
NULL
);

----------------------------------------------------------

INSERT INTO media_assets(
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a7065405c7cd75eb85b002a',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Pendayagunaan%20Aparatur%20Kabinet/Deputi%20Mevlana.png?updatedAt=1785750850373',
'-',
'Deputi Tata Laksana Kerja',
'Mevlana El Fithra Abdullah',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='pendayagunaan-aparatur-kabinet'),
NULL,
'Mevlana El Fithra Abdullah',
'Deputi Tata Laksana Kerja',
'DEPUTI',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
2,
FALSE,
TRUE,
NULL,
NULL
);

----------------------------------------------------------

INSERT INTO media_assets(
uploaded_by,imagekit_file_id,file_path,url,thumbnail_url,
name,alt_text,caption,mime_type,size_bytes,width,height,purpose,status
)
VALUES(
NULL,
'6a7065435c7cd75eb85b18f8',
'-',
'https://ik.imagekit.io/medinfoceritaloka/Kebutuhan%20Website/Pendayagunaan%20Aparatur%20Kabinet/Deputi%20Nayla%20Khansa.png?updatedAt=1785750853693',
'-',
'Deputi Audit dan Penjaminan Mutu',
'Nayla Khansa Abbylia',
'-',
'image/png',
0,NULL,NULL,
'ORGANIZATION_MEMBER',
'ACTIVE'
);

INSERT INTO organization_members(
ministry_id,user_id,name,position,position_type,email_public,
phone_public,biography,quote,photo_media_id,
display_order,is_leader,is_active,started_at,ended_at
)
VALUES(
(SELECT id FROM ministries WHERE slug='pendayagunaan-aparatur-kabinet'),
NULL,
'Nayla Khansa Abbylia',
'Deputi Audit dan Penjaminan Mutu',
'DEPUTI',
'-',
'-',
'-',
'-',
LAST_INSERT_ID(),
3,
FALSE,
TRUE,
NULL,
NULL
);