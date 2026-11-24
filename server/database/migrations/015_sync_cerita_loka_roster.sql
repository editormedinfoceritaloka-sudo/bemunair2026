-- Match previous names and positions in Kabinet Cerita Loka before updating the existing profiles.
UPDATE organization_members AS profile
JOIN ministries AS ministry ON ministry.id = profile.ministry_id
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
JOIN (
  SELECT 'MENKO_KEMASYARAKATAN' AS unit_code, 'Muhammad Ziyaad Difa Ul Haq' AS seed_name, 'Muhammad Ziyaad Difa Ul Haq' AS previous_public_name, 'Menteri Koordinator Kemasyarakatan' AS previous_position, 'Muhammad Ziyaad Difaa’ul Haq' AS new_name, 'Menteri Koordinator Kemasyarakatan' AS new_position, 1 AS display_order
  UNION ALL
  SELECT 'ADKESMA', 'Yustia Khuspriyanti Feby', 'Yustia Khuspriyanti Feby', 'Direktur Jenderal Advokasi', 'Yustia Khuspriyanti Feby Martha', 'Direktur Jenderal Advokasi Kebijakan Kampus', 2
  UNION ALL
  SELECT 'ADKESMA', 'Nabila Afifah Abida', 'Nabila Afifah Abida', 'Direktur Jenderal Kesejahteraan', 'Nabila Afifah Abida', 'Direktur Jenderal Kesejahteraan Mahasiswa', 3
  UNION ALL
  SELECT 'PENGPROF', 'M. Rifqy Yusuf', 'M. Rifqy Yusuf', 'Menteri Pengembangan Profesi', 'Muhammad Rifqy Yusuf', 'Menteri Pengembangan Profesi', 1
  UNION ALL
  SELECT 'PENGPROF', 'Raizatya Qoriananda', 'Raizatya Qoriananda', 'Direktur Jenderal Keahlian', 'Raizsatya Qoriananda Hutama', 'Direktur Jenderal Relasi Korporasi dan Jaringan Alumni', 2
  UNION ALL
  SELECT 'PENGPROF', 'Deva Zia Ul Haq', 'Deva Zia Ul Haq', 'Direktur Jenderal Relasi', 'Deva Zia''ul Haq Khoirudin', 'Direktur Jenderal Pengembangan Keahlian', 3
  UNION ALL
  SELECT 'PSDM', 'Ferdian Amartha', 'Ferdian Amartha', 'Direktur Jenderal Akselerasi', 'Ferdian Amartha', 'Direktur Jenderal Akselerasi Potensi Mahasiswa', 2
  UNION ALL
  SELECT 'LH', 'Dhafa Rhofiq Firdaus', 'Nama Belum Tersedia - Advokasi dan Kajian Strategis', 'Direktur Jenderal Advokasi dan Kajian Strategis', 'Dhafa Rofiq Firdaus', 'Direktur Jenderal Pencegahan dan Pengendalian Kerusakan Lingkungan', 2
  UNION ALL
  SELECT 'LH', 'Febria Rizqi Aulia S', 'Nama Belum Tersedia - Pencegahan dan Pengendalian', 'Direktur Jenderal Pencegahan dan Pengendalian', 'Febria Rizqi Aulia Sandy', 'Direktur Jenderal Pengembangan dan Preservasi Sumber Daya Alam', 3
  UNION ALL
  SELECT 'LH', 'Rifqah Azizah', 'Nama Belum Tersedia - Pengembangan dan Preservasi', 'Direktur Jenderal Pengembangan dan Preservasi', 'Rifqah Azizah', 'Direktur Jenderal Advokasi Kajian Strategis Lingkungan', 4
  UNION ALL
  SELECT 'KEMENKES', 'M. Zain Hisyam Al F. Z', 'Nama Belum Tersedia - Kesehatan Masyarakat', 'Direktur Jenderal Kesehatan Masyarakat', 'M. Zain Hisyam Al Fikri Zuhdi', 'Direktur Jenderal Kesejahteraan Psikososial, Inklusivitas, dan Penguatan Kesehatan Mahasiswa', 2
  UNION ALL
  SELECT 'KEMENKES', 'Cyintia Uzliva Salsabilla', 'Nama Belum Tersedia - Psikososial dan Inklusivitas', 'Direktur Jenderal Psikososial dan Inklusivitas', 'Cyintia Uzliva Salsabilla', 'Direktur Jenderal Kesehatan Masyarakat dan Pengendalian Penyakit', 3
  UNION ALL
  SELECT 'PENGMAS', 'Fahrur Rosi', 'Fahrur Rosi', 'Direktur Jenderal Pengabdian', 'Fahrur Rosi', 'Direktur Jenderal Pengembangan Desa Mitra', 2
  UNION ALL
  SELECT 'PENGMAS', 'M. Wahyu', 'M. Wahyu', 'Direktur Jenderal Pemberdayaan Masyarakat', 'Mohammad Wahyu Ramadani', 'Direktur Jenderal Sosial Masyarakat', 3
  UNION ALL
  SELECT 'HUBLU', 'Zalfaa Putri Arfiliesia', 'Zalfaa Putri Arfiliesia', 'Menteri Hubungan Luar', 'Zalfaa’ Putri Arfiliesia', 'Menteri Hubungan Luar', 1
  UNION ALL
  SELECT 'HUBLU', 'Billal Syahdan Arrafah', 'Billal Syahdan Arrafah', 'Direktur Jenderal Hubungan Eksternal', 'Billal Syahdan Arrafah', 'Direktur Jenderal Mitra Nasional', 2
  UNION ALL
  SELECT 'HUBLU', 'Hafiz Naufal Utomo', 'Hafiz Naufal Utomo', 'Direktur Jenderal Kemitraan', 'Hafihz Naufal Utomo', 'Direktur Jenderal Mitra Internasional', 3
  UNION ALL
  SELECT 'MEDINFO', 'Ocha Della Fitriani', 'Ocha Della Fitriani', 'Direktur Jenderal Publikasi', 'Ocha Della Fitriani', 'Direktur Jenderal Publikasi dan Pengembangan Media Digital', 2
  UNION ALL
  SELECT 'MEDINFO', 'Febiola Al Jannah', 'Febiola Al Jannah', 'Direktur Jenderal Visual', 'Febiola Al Jannah', 'Direktur Jenderal Visual Kreatif', 3
  UNION ALL
  SELECT 'SINEMA', 'Eugenius Nathaniel', 'Eugenius Nathaniel', 'Menteri Sinergitas Mahasiswa', 'Eugenius Nathaniel Christoffany', 'Menteri Sinergitas Mahasiswa', 1
  UNION ALL
  SELECT 'SINEMA', 'Vicentius Nolan', 'Vicentius Nolan', 'Direktur Jenderal Bina Sinergitas', 'Vincentius Nolan Nathanael Selfianto', 'Direktur Jenderal Bina Organisasi', 2
  UNION ALL
  SELECT 'SINEMA', 'Achmad Rafi Syahputra', 'Achmad Rafi Syahputra', 'Direktur Jenderal Kolaborasi', 'Achmad Rafi Syahputra', 'Direktur Jenderal Kolaborasi Mahasiswa', 3
  UNION ALL
  SELECT 'KKG', 'Farrel Ardisto Samai', 'Farrel Ardisto Samai', 'Menteri Keadilan dan Kesetaraan Gender', 'Farrel Ardisto Sama''i', 'Menteri Keadilan dan Kesetaraan Gender', 1
  UNION ALL
  SELECT 'KKG', 'Annisa Rahmawati', 'Annisa Rahmawati', 'Direktur Jenderal Edukasi', 'Annisa Rahmawati', 'Direktur Jenderal Edukasi Kritis Gender', 2
  UNION ALL
  SELECT 'KKG', 'Titis Nastiti Dwi', 'Titis Nastiti Dwi', 'Direktur Jenderal Pemberdayaan', 'Titis Nastiti Dwi Lestari', 'Direktur Jenderal Pemberdayaan Perempuan Progresif', 3
  UNION ALL
  SELECT 'SOSPOL', 'Moch. Amir Nazaruddin', 'Moch. Amir Nazaruddin', 'Direktur Jenderal Aksi', 'Moch. Amir Nazaruddin', 'Direktur Jenderal Aksi dan Propaganda', 2
  UNION ALL
  SELECT 'SOSPOL', 'Ahmad Raihan Fadhillah', 'Ahmad Raihan Fadhillah', 'Direktur Jenderal Analisis', 'Ahmad Raihan Fadhillah Yasin', 'Direktur Jenderal Analisis Isu Strategis', 3
  UNION ALL
  SELECT 'EDIGPRO', 'Argi', 'Argi', 'Direktur Jenderal Ekonomi Digital', 'Argi Wicaksana', 'Direktur Jenderal Produk Kreatif dan Investasi', 2
  UNION ALL
  SELECT 'EDIGPRO', 'Yudhistira', 'Yudhistira', 'Direktur Jenderal Produk Kreatif', 'M. Yudhistira Setya Heri', 'Direktur Jenderal UMKM dan Kewirausahaan', 3
  UNION ALL
  SELECT 'RISKEL', 'Abdul Wahid Tamimi', 'Ata', 'Direktur Jenderal Riset', 'Abdul Wahid Tamimi', 'Direktur Jenderal Riset dan Inovasi', 2
  UNION ALL
  SELECT 'RISKEL', 'Nasywa Ayudya S', 'Naswa', 'Direktur Jenderal Keilmuan', 'Naswa Ayudya Sasta', 'Direktur Jenderal Akademik dan Prestasi', 3
  UNION ALL
  SELECT 'SENIIORA', 'Abdullah Azzam', 'Azzam', 'Direktur Jenderal Seni', 'Abdullah Azzam', 'Direktur Jenderal Pemuda dan Olahraga', 3
  UNION ALL
  SELECT 'BENKAB', 'Ananda Magfira Nurochma', 'Ananda Magfira Nurochma', 'Kepala Biro Bendahara Kabinet I', 'Ananda Magfira Nurochma', 'Kepala Biro Administrasi dan Keuangan', 2
  UNION ALL
  SELECT 'BENKAB', 'M. Salman Al Farisi', 'M. Salman Al Farisi', 'Kepala Biro Bendahara Kabinet II', 'M. Salman Al Farisi Al Ghofsah', 'Kepala Biro Kebendaharaan', 3
  UNION ALL
  SELECT 'PAK', 'Mevlana El Fithra', 'Mevlana El Fithra', 'Deputi Pendayagunaan Aparatur Kabinet I', 'Mevlana El Fithra Abdullah', 'Deputi Tata Laksana Kerja', 2
  UNION ALL
  SELECT 'PAK', 'Nayla Khansa Abbylia', 'Nayla Khansa Abbylia', 'Deputi Pendayagunaan Aparatur Kabinet II', 'Nayla Khansa Abbylia', 'Deputi Audit dan Penjaminan Mutu', 3
  UNION ALL
  SELECT 'SESKAB', 'Fathimah Nuurun', 'Fathimah Nuurun', 'Kepala Biro Sekretaris Kabinet I', 'Fathimah Nuurun Najma As-Salman', 'Kepala Biro Administrasi Persuratan', 2
  UNION ALL
  SELECT 'SESKAB', 'Najwa Nur Hafidza', 'Najwa Nur Hafidza', 'Kepala Biro Sekretaris Kabinet II', 'Najwa Nur Hafidza', 'Kepala Biro Arsip dan Kesekretariatan', 3
) AS roster
  ON ministry.code = roster.unit_code COLLATE utf8mb4_unicode_ci
  AND profile.position = roster.previous_position COLLATE utf8mb4_unicode_ci
  AND profile.name IN (
    roster.seed_name COLLATE utf8mb4_unicode_ci,
    roster.previous_public_name COLLATE utf8mb4_unicode_ci
  )
SET profile.name = roster.new_name,
    profile.position = roster.new_position,
    profile.display_order = CASE WHEN ministry.code = 'SENIIORA' THEN roster.display_order ELSE profile.display_order END
WHERE cabinet.slug = 'cerita-loka-2026'
  AND profile.is_active = TRUE;

-- Retain the former profile and its media without showing it as a current appointment.
UPDATE organization_members AS profile
JOIN ministries AS ministry ON ministry.id = profile.ministry_id
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
SET profile.is_active = FALSE
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND profile.position_type = 'DIRJEN'
  AND profile.position = 'Direktur Jenderal Olahraga'
  AND profile.name IN ('Puput', 'M. Fadhil')
  AND profile.is_active = TRUE;

-- The new appointee has no confirmed portrait in the supplied roster.
INSERT INTO organization_members (ministry_id, name, position, position_type, display_order)
SELECT ministry.id, 'Nurnajmi Dwihesti', 'Direktur Jenderal Seni dan Kebudayaan', 'DIRJEN', 2
FROM ministries AS ministry
JOIN cabinet_terms AS cabinet ON cabinet.id = ministry.cabinet_term_id
WHERE cabinet.slug = 'cerita-loka-2026'
  AND ministry.code = 'SENIIORA'
  AND NOT EXISTS (
    SELECT 1
    FROM organization_members AS existing
    WHERE existing.ministry_id = ministry.id
      AND (existing.name = 'Nurnajmi Dwihesti'
        OR (existing.position = 'Direktur Jenderal Seni dan Kebudayaan' AND existing.is_active = TRUE))
  );
