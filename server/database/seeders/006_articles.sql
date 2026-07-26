INSERT INTO articles (
  id, slug, title, excerpt, body, cover_image, author_id, status, published_at
) VALUES
  (
    900001,
    'mengenal-bem-unair-ruang-tumbuh-dan-berdampak',
    'Mengenal BEM UNAIR: Ruang Tumbuh dan Berdampak',
    'Mengenal peran BEM Universitas Airlangga sebagai ruang kolaborasi, pengembangan diri, dan kontribusi mahasiswa.',
    '<h2>Ruang kolaborasi mahasiswa</h2><p>BEM Universitas Airlangga menjadi ruang bagi mahasiswa untuk belajar, bekerja sama, dan menghadirkan manfaat yang dapat dirasakan oleh sivitas akademika maupun masyarakat.</p><blockquote>Setiap gagasan bertumbuh melalui kerja bersama.</blockquote><h2>Bergerak dengan dampak</h2><p>Program disusun berdasarkan kebutuhan mahasiswa, data lapangan, dan semangat pengabdian.</p>',
    'https://ik.imagekit.io/demo/default-image.jpg',
    (SELECT id FROM users WHERE email = 'admin@bem.unair.ac.id'),
    'PUBLISHED',
    DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 2 DAY)
  ),
  (
    900002,
    'panduan-mengakses-layanan-mahasiswa',
    'Panduan Mengakses Layanan Mahasiswa',
    'Rangkuman kanal layanan akademik dan kesejahteraan yang penting diketahui mahasiswa.',
    '<h2>Mulai dari kebutuhanmu</h2><p>Kenali jenis bantuan yang diperlukan, siapkan data pendukung, lalu hubungi kanal resmi yang sesuai.</p><h3>Simpan nomor penting</h3><ul><li>Layanan akademik</li><li>Layanan kesejahteraan mahasiswa</li><li>Kanal aspirasi BEM UNAIR</li></ul>',
    NULL,
    (SELECT id FROM users WHERE email = 'editor.medinfo@bem.unair.ac.id'),
    'DRAFT',
    NULL
  ),
  (
    900003,
    'cerita-di-balik-program-pengabdian',
    'Cerita di Balik Program Pengabdian Mahasiswa',
    'Catatan proses kolaborasi mahasiswa dan masyarakat dalam merancang program yang berkelanjutan.',
    '<h2>Mendengar sebelum bergerak</h2><p>Program pengabdian yang baik dimulai dengan mendengar kebutuhan masyarakat dan menyusun tujuan bersama.</p><h2>Menjaga keberlanjutan</h2><p>Evaluasi dan dokumentasi memastikan manfaat program tidak berhenti ketika kegiatan selesai.</p>',
    NULL,
    (SELECT id FROM users WHERE email = 'mentri.sosmas@bem.unair.ac.id'),
    'PUBLISHED',
    DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 7 DAY)
  )
ON DUPLICATE KEY UPDATE
  slug = VALUES(slug),
  title = VALUES(title),
  excerpt = VALUES(excerpt),
  body = VALUES(body),
  cover_image = VALUES(cover_image),
  author_id = VALUES(author_id),
  status = VALUES(status),
  published_at = VALUES(published_at);
