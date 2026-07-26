-- Password seluruh akun development: password
INSERT INTO users (id, name, email, password_hash, role, ministry, phone) VALUES
  (1, 'Admin BEM UNAIR', 'admin@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'ADMIN', NULL, '6281111111111'),
  (2, 'Koordinator Medinfo', 'mentri.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'MEDINFO', '6281222222222'),
  (3, 'Menteri PSDM', 'mentri.psdm@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'PSDM', '6281333333333'),
  (4, 'Menteri Sosmas', 'mentri.sosmas@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'SOSMAS', '6281444444444'),
  (5, 'Editor Medinfo', 'editor.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'MEDINFO', '6281555555555'),
  (6, 'Staf Medinfo', 'staf.medinfo@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'MEDINFO', '6281666666666'),
  (7, 'Menteri Kesma', 'mentri.kesma@bem.unair.ac.id', '$2a$10$vXtJZlWskizrDilAfKBkeur9V9nlcRSH/GTvOfcw2HHkSCadc77Mm', 'MENTRI', 'KESMA', '6281777777777')
ON DUPLICATE KEY UPDATE
  name = VALUES(name),
  password_hash = VALUES(password_hash),
  role = VALUES(role),
  ministry = VALUES(ministry),
  phone = VALUES(phone);
