INSERT INTO media_assets (id, uploaded_by, imagekit_file_id, file_path, url, thumbnail_url, name, alt_text, mime_type, size_bytes, width, height, purpose, status) VALUES
  (910100, 1, 'demo-cabinet-logo', '/bemunair/cabinet/logo', 'https://ik.imagekit.io/demo/default-image.jpg', 'https://ik.imagekit.io/demo/default-image.jpg', 'Logo kabinet demo', 'Logo Kabinet Cerita Loka demo', 'image/jpeg', 12000, 800, 800, 'cabinet', 'ACTIVE'),
  (910101, 1, 'demo-unit-logo', '/bemunair/organizations/logo', 'https://ik.imagekit.io/demo/default-image.jpg', 'https://ik.imagekit.io/demo/default-image.jpg', 'Logo unit demo', 'Logo unit organisasi demo', 'image/jpeg', 12000, 800, 800, 'organization', 'ACTIVE'),
  (910102, 1, 'demo-program-cover', '/bemunair/programs/cover', 'https://ik.imagekit.io/demo/default-image.jpg', 'https://ik.imagekit.io/demo/default-image.jpg', 'Cover program demo', 'Cover program kerja demo', 'image/jpeg', 12000, 1200, 800, 'program', 'ACTIVE'),
  (910103, 1, 'demo-documentation-1', '/bemunair/documentations/1', 'https://ik.imagekit.io/demo/default-image.jpg', 'https://ik.imagekit.io/demo/default-image.jpg', 'Dokumentasi demo satu', 'Dokumentasi program kerja demo satu', 'image/jpeg', 12000, 1200, 800, 'documentation', 'ACTIVE'),
  (910104, 1, 'demo-documentation-2', '/bemunair/documentations/2', 'https://ik.imagekit.io/demo/default-image.jpg', 'https://ik.imagekit.io/demo/default-image.jpg', 'Dokumentasi demo dua', 'Dokumentasi program kerja demo dua', 'image/jpeg', 12000, 1200, 800, 'documentation', 'ACTIVE'),
  (910105, 1, 'demo-kemenko-internal', '/bemunair/organizations/kemenko-internal/logo', 'https://ik.imagekit.io/demo/kemenko-internal.jpg', 'https://ik.imagekit.io/demo/kemenko-internal.jpg', 'Logo Kemenko Internal demo', 'Logo Kemenko Internal Kabinet Cerita Loka', 'image/jpeg', 12000, 800, 800, 'organization', 'ACTIVE'),
  (910106, 1, 'demo-kemenko-pengabdian', '/bemunair/organizations/kemenko-pengabdian/logo', 'https://ik.imagekit.io/demo/kemenko-pengabdian.jpg', 'https://ik.imagekit.io/demo/kemenko-pengabdian.jpg', 'Logo Kemenko Pengabdian demo', 'Logo Kemenko Pengabdian Kabinet Cerita Loka', 'image/jpeg', 12000, 800, 800, 'organization', 'ACTIVE'),
  (910107, 1, 'demo-letter-undangan', '/bemunair/letter-templates/undangan.pdf', 'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf', NULL, 'Template Undangan demo.pdf', 'Template surat undangan demo dalam format PDF', 'application/pdf', 25000, NULL, NULL, 'letter_template', 'ACTIVE'),
  (910108, 1, 'demo-letter-peminjaman', '/bemunair/letter-templates/peminjaman.pdf', 'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf', NULL, 'Template Peminjaman demo.pdf', 'Template surat peminjaman demo dalam format PDF', 'application/pdf', 25000, NULL, NULL, 'letter_template', 'ACTIVE'),
  (910109, 1, 'demo-letter-kerja-sama', '/bemunair/letter-templates/kerja-sama.pdf', 'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf', NULL, 'Template Kerja Sama demo.pdf', 'Template surat kerja sama demo dalam format PDF', 'application/pdf', 25000, NULL, NULL, 'letter_template', 'ACTIVE')
ON DUPLICATE KEY UPDATE url = VALUES(url), thumbnail_url = VALUES(thumbnail_url), alt_text = VALUES(alt_text), name = VALUES(name), mime_type = VALUES(mime_type), purpose = VALUES(purpose), status = VALUES(status);

UPDATE letter_templates SET media_asset_id = CASE type WHEN 'UNDANGAN' THEN 910107 WHEN 'PEMINJAMAN' THEN 910108 WHEN 'KERJA_SAMA' THEN 910109 ELSE media_asset_id END, is_active = TRUE WHERE id IN (900001, 900002, 900003);

INSERT INTO cabinet_terms (id, name, slug, tagline, description, logo_media_id, is_active, is_published, published_at, meta_title, meta_description) VALUES
  (910000, 'Kabinet Cerita Loka', 'cerita-loka', 'Bertumbuh, bercerita, berdampak.', 'Kabinet demo untuk pengembangan halaman publik BEM UNAIR.', 910100, TRUE, TRUE, CURRENT_TIMESTAMP, 'Kabinet Cerita Loka', 'Profil dan program kerja Kabinet Cerita Loka.')
ON DUPLICATE KEY UPDATE name = VALUES(name), tagline = VALUES(tagline), description = VALUES(description), logo_media_id = VALUES(logo_media_id), is_active = VALUES(is_active), is_published = VALUES(is_published), published_at = VALUES(published_at);

INSERT INTO ministries (id, code, name, cabinet_term_id, unit_type, slug, short_name, description, display_order, is_active, is_published, logo_media_id) VALUES
  (910001, 'KEMENKO_INTERNAL', 'Kemenko Internal', 910000, 'KEMENKOAN', 'kemenko-internal', 'Internal', 'Kemenkoan yang mengoordinasikan pengembangan internal organisasi.', 1, TRUE, TRUE, 910105),
  (910002, 'KEMENKO_PENGABDIAN', 'Kemenko Pengabdian', 910000, 'KEMENKOAN', 'kemenko-pengabdian', 'Pengabdian', 'Kemenkoan yang mengoordinasikan kerja pengabdian dan dampak sosial.', 2, TRUE, TRUE, 910106)
ON DUPLICATE KEY UPDATE name = VALUES(name), cabinet_term_id = VALUES(cabinet_term_id), unit_type = VALUES(unit_type), slug = VALUES(slug), is_active = VALUES(is_active), is_published = VALUES(is_published), logo_media_id = VALUES(logo_media_id);

UPDATE ministries SET cabinet_term_id = 910000, unit_type = 'KEMENTERIAN', is_active = TRUE, is_published = TRUE, display_order = CASE code WHEN 'MEDINFO' THEN 1 WHEN 'PSDM' THEN 2 WHEN 'KESMA' THEN 3 WHEN 'SOSMAS' THEN 1 ELSE display_order END, slug = CASE code WHEN 'MEDINFO' THEN 'media-dan-informasi' WHEN 'PSDM' THEN 'pengembangan-sumber-daya-mahasiswa' WHEN 'KESMA' THEN 'kesejahteraan-mahasiswa' WHEN 'SOSMAS' THEN 'sosial-masyarakat' ELSE slug END, logo_media_id = 910101 WHERE code IN ('MEDINFO', 'PSDM', 'KESMA', 'SOSMAS');
UPDATE ministries SET parent_id = 910001 WHERE code IN ('MEDINFO', 'PSDM', 'KESMA');
UPDATE ministries SET parent_id = 910002 WHERE code = 'SOSMAS';

INSERT INTO organization_members (id, ministry_id, name, position, position_type, photo_media_id, biography, display_order, is_leader, is_active) VALUES
  (910200, (SELECT id FROM ministries WHERE code = 'MEDINFO'), 'Menteri Media Demo', 'Menteri', 'MINISTER', 910101, 'Profil demo Menteri Media dan Informasi.', 1, TRUE, TRUE),
  (910201, (SELECT id FROM ministries WHERE code = 'MEDINFO'), 'Dirjen Media Demo', 'Direktur Jenderal', 'DIRECTOR_GENERAL', 910101, 'Profil demo Direktur Jenderal Media dan Informasi.', 2, TRUE, TRUE),
  (910202, (SELECT id FROM ministries WHERE code = 'PSDM'), 'Menteri PSDM Demo', 'Menteri', 'MINISTER', 910101, 'Profil demo Menteri PSDM.', 1, TRUE, TRUE),
  (910203, (SELECT id FROM ministries WHERE code = 'PSDM'), 'Dirjen PSDM Demo', 'Direktur Jenderal', 'DIRECTOR_GENERAL', 910101, 'Profil demo Direktur Jenderal PSDM.', 2, TRUE, TRUE),
  (910204, (SELECT id FROM ministries WHERE code = 'KESMA'), 'Menteri Kesma Demo', 'Menteri', 'MINISTER', 910101, 'Profil demo Menteri Kesma.', 1, TRUE, TRUE),
  (910205, (SELECT id FROM ministries WHERE code = 'KESMA'), 'Dirjen Kesma Demo', 'Direktur Jenderal', 'DIRECTOR_GENERAL', 910101, 'Profil demo Direktur Jenderal Kesma.', 2, TRUE, TRUE),
  (910206, (SELECT id FROM ministries WHERE code = 'SOSMAS'), 'Menteri Sosmas Demo', 'Menteri', 'MINISTER', 910101, 'Profil demo Menteri Sosmas.', 1, TRUE, TRUE),
  (910207, (SELECT id FROM ministries WHERE code = 'SOSMAS'), 'Dirjen Sosmas Demo', 'Direktur Jenderal', 'DIRECTOR_GENERAL', 910101, 'Profil demo Direktur Jenderal Sosmas.', 2, TRUE, TRUE)
ON DUPLICATE KEY UPDATE name = VALUES(name), position = VALUES(position), photo_media_id = VALUES(photo_media_id), biography = VALUES(biography), is_active = VALUES(is_active);

INSERT INTO work_programs (id, ministry_id, name, slug, short_description, description, objectives, target_audience, execution_month, lifecycle_status, cover_media_id, display_order, is_featured, is_published, published_at) VALUES
  (910300, (SELECT id FROM ministries WHERE code = 'MEDINFO'), 'Ruang Cerita Digital', 'ruang-cerita-digital', 'Ruang publikasi dan dokumentasi digital.', 'Program demo untuk membangun ekosistem cerita digital.', 'Meningkatkan kualitas informasi publik.', 'Mahasiswa UNAIR', 'Januari', 'COMPLETED', 910102, 1, TRUE, TRUE, CURRENT_TIMESTAMP),
  (910301, (SELECT id FROM ministries WHERE code = 'MEDINFO'), 'Kelas Kreator BEM', 'kelas-kreator-bem', 'Kelas pengembangan kemampuan kreatif.', 'Program pelatihan kreator untuk kebutuhan komunikasi organisasi.', 'Membangun kapasitas kreator.', 'Anggota organisasi', 'Februari', 'ONGOING', 910102, 2, TRUE, TRUE, CURRENT_TIMESTAMP),
  (910302, (SELECT id FROM ministries WHERE code = 'PSDM'), 'Sekolah Penggerak', 'sekolah-penggerak', 'Pengembangan kapasitas pengurus.', 'Program pembelajaran kepemimpinan dan kolaborasi.', 'Membentuk penggerak yang adaptif.', 'Pengurus BEM', 'Maret', 'PLANNED', 910102, 1, TRUE, TRUE, CURRENT_TIMESTAMP),
  (910303, (SELECT id FROM ministries WHERE code = 'PSDM'), 'Mentoring Organisasi', 'mentoring-organisasi', 'Pendampingan organisasi mahasiswa.', 'Program mentoring lintas organisasi.', 'Meningkatkan kemampuan tata kelola.', 'Organisasi mahasiswa', 'April', 'PLANNED', 910102, 2, FALSE, TRUE, CURRENT_TIMESTAMP),
  (910304, (SELECT id FROM ministries WHERE code = 'KESMA'), 'Teman Sehat', 'teman-sehat', 'Pendampingan kesejahteraan mahasiswa.', 'Program dukungan dan edukasi kesejahteraan.', 'Memperluas akses dukungan mahasiswa.', 'Mahasiswa UNAIR', 'Mei', 'ONGOING', 910102, 1, TRUE, TRUE, CURRENT_TIMESTAMP),
  (910305, (SELECT id FROM ministries WHERE code = 'KESMA'), 'Ruang Aman Kampus', 'ruang-aman-kampus', 'Edukasi lingkungan kampus aman.', 'Program edukasi dan rujukan bantuan.', 'Membangun lingkungan kampus yang aman.', 'Mahasiswa UNAIR', 'Juni', 'PLANNED', 910102, 2, FALSE, TRUE, CURRENT_TIMESTAMP),
  (910306, (SELECT id FROM ministries WHERE code = 'SOSMAS'), 'Desa Binaan', 'desa-binaan', 'Kolaborasi pengabdian masyarakat.', 'Program kerja sama berkelanjutan dengan masyarakat.', 'Menciptakan dampak sosial terukur.', 'Masyarakat mitra', 'Juli', 'ONGOING', 910102, 1, TRUE, TRUE, CURRENT_TIMESTAMP),
  (910307, (SELECT id FROM ministries WHERE code = 'SOSMAS'), 'Aksi Lingkungan', 'aksi-lingkungan', 'Gerakan lingkungan kolaboratif.', 'Program aksi lingkungan berbasis komunitas.', 'Mendorong kepedulian lingkungan.', 'Komunitas mahasiswa', 'Agustus', 'PLANNED', 910102, 2, FALSE, TRUE, CURRENT_TIMESTAMP),
  (910308, (SELECT id FROM ministries WHERE code = 'SOSMAS'), 'Peta Dampak Sosial', 'peta-dampak-sosial', 'Pemetaan dampak kegiatan sosial.', 'Program dokumentasi dan evaluasi dampak sosial.', 'Memperkuat evaluasi berbasis data.', 'Mitra sosial', 'September', 'DRAFT', 910102, 3, FALSE, TRUE, CURRENT_TIMESTAMP)
ON DUPLICATE KEY UPDATE name = VALUES(name), short_description = VALUES(short_description), lifecycle_status = VALUES(lifecycle_status), is_featured = VALUES(is_featured), is_published = VALUES(is_published), published_at = VALUES(published_at);

INSERT INTO work_program_milestones (id, work_program_id, title, description, status, display_order) VALUES
  (910400, 910300, 'Riset kebutuhan', 'Pemetaan kebutuhan informasi.', 'COMPLETED', 1),
  (910401, 910300, 'Peluncuran platform', 'Publikasi platform program.', 'COMPLETED', 2),
  (910402, 910304, 'Konsultasi awal', 'Pembukaan kanal konsultasi.', 'ONGOING', 1)
ON DUPLICATE KEY UPDATE title = VALUES(title), description = VALUES(description), status = VALUES(status), display_order = VALUES(display_order);

INSERT INTO work_program_documentations (id, work_program_id, media_asset_id, title, caption, display_order, is_cover) VALUES
  (910500, 910300, 910103, 'Peluncuran program', 'Dokumentasi demo peluncuran.', 1, TRUE),
  (910501, 910300, 910104, 'Sesi kolaborasi', 'Dokumentasi demo sesi kolaborasi.', 2, FALSE)
ON DUPLICATE KEY UPDATE title = VALUES(title), caption = VALUES(caption), display_order = VALUES(display_order), is_cover = VALUES(is_cover);

INSERT INTO user_organization_roles (user_id, ministry_id, permission, is_active)
SELECT u.id, m.id, 'MANAGE_CONTENT', TRUE FROM users u JOIN ministries m ON m.code = u.ministry WHERE u.role = 'ADMIN'
ON DUPLICATE KEY UPDATE is_active = VALUES(is_active);
