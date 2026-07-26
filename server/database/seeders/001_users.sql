-- Password seluruh akun development: password
INSERT INTO users (id, name, email, password_hash, role, ministry, ministry_id, phone) VALUES
  (1, 'Admin BEM UNAIR', 'admin@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN_MEDINFO', NULL, NULL, '6281111111111'),
  (2, 'Koordinator Medinfo', 'mentri.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN_MEDINFO', 'MEDINFO', (SELECT id FROM ministries WHERE code = 'MEDINFO'), '6281222222222'),
  (3, 'Menteri PSDM', 'mentri.psdm@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN', 'PSDM', (SELECT id FROM ministries WHERE code = 'PSDM'), '6281333333333'),
  (4, 'Menteri Sosmas', 'mentri.sosmas@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN', 'SOSMAS', (SELECT id FROM ministries WHERE code = 'SOSMAS'), '6281444444444'),
  (5, 'Editor Medinfo', 'editor.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN_MEDINFO', 'MEDINFO', (SELECT id FROM ministries WHERE code = 'MEDINFO'), '6281555555555'),
  (6, 'Staf Medinfo', 'staf.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN_MEDINFO', 'MEDINFO', (SELECT id FROM ministries WHERE code = 'MEDINFO'), '6281666666666'),
  (7, 'Menteri Kesma', 'mentri.kesma@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN', 'KESMA', (SELECT id FROM ministries WHERE code = 'KESMA'), '6281777777777')
ON DUPLICATE KEY UPDATE
  name = VALUES(name),
  password_hash = VALUES(password_hash),
  role = VALUES(role),
  ministry = VALUES(ministry),
  ministry_id = VALUES(ministry_id),
  phone = VALUES(phone);
