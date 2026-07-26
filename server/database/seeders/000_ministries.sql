INSERT INTO ministries (code, name, is_active) VALUES
  ('MEDINFO', 'Media dan Informasi', TRUE),
  ('PSDM', 'Pengembangan Sumber Daya Mahasiswa', TRUE),
  ('SOSMAS', 'Sosial Masyarakat', TRUE),
  ('KESMA', 'Kesejahteraan Mahasiswa', TRUE)
ON DUPLICATE KEY UPDATE
  name = VALUES(name),
  is_active = VALUES(is_active);
