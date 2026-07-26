INSERT INTO letter_submissions (
  id, submitter_id, ministry, letter_type, subject, body, deadline,
  assigned_pj_id, status, notes
) VALUES
  (
    900001,
    (SELECT id FROM users WHERE email = 'mentri.psdm@bem.unair.ac.id'),
    'PSDM',
    'UNDANGAN',
    'Undangan Narasumber Leadership Class',
    'Dengan hormat, kami mengundang Bapak/Ibu untuk menjadi narasumber dalam kegiatan Leadership Class BEM Universitas Airlangga.',
    DATE_ADD(CURRENT_TIMESTAMP, INTERVAL 3 DAY),
    (SELECT id FROM users WHERE email = 'mentri.medinfo@bem.unair.ac.id'),
    'PENDING',
    NULL
  ),
  (
    900002,
    (SELECT id FROM users WHERE email = 'mentri.sosmas@bem.unair.ac.id'),
    'SOSMAS',
    'PEMINJAMAN',
    'Permohonan Peminjaman Aula',
    'Dengan hormat, kami mengajukan permohonan peminjaman aula untuk pelaksanaan kegiatan pengabdian masyarakat.',
    DATE_ADD(CURRENT_TIMESTAMP, INTERVAL 5 DAY),
    (SELECT id FROM users WHERE email = 'editor.medinfo@bem.unair.ac.id'),
    'IN_REVIEW',
    'Jadwal penggunaan aula sedang dikonfirmasi.'
  ),
  (
    900003,
    (SELECT id FROM users WHERE email = 'mentri.kesma@bem.unair.ac.id'),
    'KESMA',
    'KERJA_SAMA',
    'Permohonan Kerja Sama Layanan Konseling',
    'Dengan hormat, kami bermaksud mengajukan kerja sama penyediaan informasi dan akses layanan konseling bagi mahasiswa.',
    DATE_ADD(CURRENT_TIMESTAMP, INTERVAL 7 DAY),
    (SELECT id FROM users WHERE email = 'staf.medinfo@bem.unair.ac.id'),
    'APPROVED',
    'Substansi surat telah disetujui.'
  ),
  (
    900004,
    (SELECT id FROM users WHERE email = 'mentri.psdm@bem.unair.ac.id'),
    'PSDM',
    'SURAT_TUGAS',
    'Surat Tugas Delegasi Forum Mahasiswa',
    'Dengan hormat, mohon diterbitkan surat tugas bagi delegasi BEM Universitas Airlangga.',
    DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 1 DAY),
    NULL,
    'REJECTED',
    'Daftar delegasi belum dilampirkan.'
  )
ON DUPLICATE KEY UPDATE
  submitter_id = VALUES(submitter_id),
  ministry = VALUES(ministry),
  letter_type = VALUES(letter_type),
  subject = VALUES(subject),
  body = VALUES(body),
  deadline = VALUES(deadline),
  assigned_pj_id = VALUES(assigned_pj_id),
  status = VALUES(status),
  notes = VALUES(notes);
