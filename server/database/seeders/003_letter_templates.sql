INSERT INTO letter_templates (id, name, type, subject, body) VALUES
  (
    900001,
    'Undangan Kegiatan Resmi',
    'UNDANGAN',
    'Undangan Kegiatan BEM Universitas Airlangga',
    'Yth. Bapak/Ibu/Saudara\nDi tempat\n\nDengan hormat,\n\nSehubungan dengan pelaksanaan kegiatan BEM Universitas Airlangga, kami bermaksud mengundang Bapak/Ibu/Saudara untuk hadir pada kegiatan tersebut.\n\nDemikian undangan ini kami sampaikan. Atas perhatian dan kehadirannya, kami ucapkan terima kasih.'
  ),
  (
    900002,
    'Permohonan Peminjaman Tempat',
    'PEMINJAMAN',
    'Permohonan Peminjaman Tempat Kegiatan',
    'Yth. Pimpinan Unit Terkait\nUniversitas Airlangga\n\nDengan hormat,\n\nDalam rangka mendukung pelaksanaan program kerja BEM Universitas Airlangga, kami mengajukan permohonan peminjaman tempat untuk kegiatan yang akan dilaksanakan.\n\nDemikian permohonan ini kami sampaikan. Atas izin dan kerja samanya, kami ucapkan terima kasih.'
  ),
  (
    900003,
    'Permohonan Kerja Sama',
    'KERJA_SAMA',
    'Permohonan Kerja Sama Program BEM UNAIR',
    'Yth. Pimpinan Mitra\nDi tempat\n\nDengan hormat,\n\nBEM Universitas Airlangga bermaksud mengajukan kerja sama dalam pelaksanaan program yang berorientasi pada pengembangan mahasiswa dan kontribusi sosial.\n\nBesar harapan kami agar kerja sama ini dapat terjalin dengan baik.'
  )
ON DUPLICATE KEY UPDATE
  name = VALUES(name),
  type = VALUES(type),
  subject = VALUES(subject),
  body = VALUES(body);
