INSERT INTO users (id, name, email, password_hash, role, ministry, phone) VALUES
  (1, 'Admin BEM UNAIR', 'admin@bem.unair.ac.id', '$2a$10$rJQd5vARFtcOtuF8v38Me.4VIphTlfhmXLvkKZv.2X0tO2J3v4E6K', 'ADMIN', NULL, '6281111111111'),
  (2, 'Mentri Medinfo', 'mentri.medinfo@bem.unair.ac.id', '$2a$10$rJQd5vARFtcOtuF8v38Me.4VIphTlfhmXLvkKZv.2X0tO2J3v4E6K', 'MENTRI', 'MEDINFO', '6281222222222'),
  (3, 'Mentri PSDM', 'mentri.psdm@bem.unair.ac.id', '$2a$10$rJQd5vARFtcOtuF8v38Me.4VIphTlfhmXLvkKZv.2X0tO2J3v4E6K', 'MENTRI', 'PSDM', '6281333333333'),
  (4, 'Mentri Sosmas', 'mentri.sosmas@bem.unair.ac.id', '$2a$10$rJQd5vARFtcOtuF8v38Me.4VIphTlfhmXLvkKZv.2X0tO2J3v4E6K', 'MENTRI', 'SOSMAS', '6281444444444')
ON DUPLICATE KEY UPDATE name = VALUES(name), role = VALUES(role), ministry = VALUES(ministry), phone = VALUES(phone);

INSERT INTO medinfo_pj_queues (user_id, position, is_current) VALUES
  (2, 1, TRUE),
  (3, 2, FALSE),
  (4, 3, FALSE)
ON DUPLICATE KEY UPDATE position = VALUES(position), is_current = VALUES(is_current);
