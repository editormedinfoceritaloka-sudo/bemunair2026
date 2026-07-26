-- Bersihkan anggota antrean dari seeder development lama yang bukan MEDINFO.
DELETE queue
FROM medinfo_pj_queues AS queue
JOIN users AS user ON user.id = queue.user_id
WHERE user.email IN (
  'admin@bem.unair.ac.id',
  'mentri.psdm@bem.unair.ac.id',
  'mentri.sosmas@bem.unair.ac.id'
);

INSERT INTO medinfo_pj_queues (user_id, position, is_current) VALUES
  ((SELECT id FROM users WHERE email = 'mentri.medinfo@bem.unair.ac.id'), 1, TRUE),
  ((SELECT id FROM users WHERE email = 'editor.medinfo@bem.unair.ac.id'), 2, FALSE),
  ((SELECT id FROM users WHERE email = 'staf.medinfo@bem.unair.ac.id'), 3, FALSE)
ON DUPLICATE KEY UPDATE
  position = VALUES(position),
  is_current = VALUES(is_current);
