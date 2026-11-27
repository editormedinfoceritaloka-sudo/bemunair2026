const cabinetCloud = 'https://res.cloudinary.com/rvn0bdia/image/upload/';
const portraitBase = cabinetCloud + 'e_background_removal,f_auto,q_auto/';

// Names and Cloudinary public IDs from the Cerita Loka cabinet seed.
const rosterPortraits = new Map<string, string>([
  ['Johanes Richard Darmawan', 'RICHIE_MENKO_KEMAHASISWAAN'],
  ['Muhammad Ziyaad Difaa’ul Haq', 'ZIYAAD_MENKO_KEMASYARAKATAN'],
  ['Andrean Marcellino', 'MENKO_KOMINFO_ANDREAN'],
  ['Dewa Pranata Putra Pratama', 'DEWA_MENKO_PERGERAKAN'],
  ['Nasua Okta Kusuma Dewi Rofiq', 'NASUA_OKTA_MENKO_PPK'],
  ['Elsa Putri Aprizalni', 'menteri_adkesma_elsa_ADKESMA'],
  ['Yustia Khuspriyanti Feby Martha', 'Direktur_JenderalYustia_Khuspriyanti_Feby_ADKESMA'],
  ['Nabila Afifah Abida', 'Direktur_JenderalNabila_Afifah_Abida_ADKESMA'],
  ['Muhammad Rifqy Yusuf', 'Menteri_PENGPROFM._Rifqy_Yusuf'],
  ['Raizsatya Qoriananda Hutama', 'DIRJEN_RAIS_PENGPROF'],
  ['Deva Zia\'ul Haq Khoirudin', 'DIRJEN_PENGPROFDeva_Zia_Ul_Haq'],
  ['Labib Azka Asysyafiq', 'MENTERIPSDM_LABIB'],
  ['Ferdian Amartha', 'DIRJEN_PSDM_FERDIAN'],
  ['Qheisya Luna Alifa', 'DIRJENPSDM_LUNA'],
  ['Zaskia Darojah', 'MENTERI_LINGKUNGAN_HIDUP_ZASKIA'],
  ['Dhafa Rofiq Firdaus', 'DIRJEN_LINGKUNGAN_HIDUP_DAFFA'],
  ['Febria Rizqi Aulia Sandy', 'DIRJEN_LINGKUNGAN_HIDUP_FEBRIA'],
  ['Rifqah Azizah', 'DIRJEN_LINGKUNGANHIDUP_AZIZAH'],
  ['An Nadhofah Adlin', 'MENTERIKESEHATAN_AN_NADHOFAH'],
  ['M. Zain Hisyam Al Fikri Zuhdi', 'DIRJEN_KESEHATAN_ZEIN'],
  ['Cyintia Uzliva Salsabilla', 'DIRJEN_KESEHATAN_CYINTIA'],
  ['Sahrul Efendi', 'MENTERI_OENGMAS_SAHRUL'],
  ['Fahrur Rosi', 'DIRJEN_PENGABDIAN_MASYARAKAT_ROSI'],
  ['Mohammad Wahyu Ramadani', 'DIRJEN_PENGMAS_WAHYU'],
  ['Zalfaa’ Putri Arfiliesia', 'MENTERI_HUBUNGAN_LUAR_ZALFA'],
  ['Billal Syahdan Arrafah', 'DIRJEN_HUBLU_BILLAL'],
  ['Hafihz Naufal Utomo', 'DIRJEN_HUBLU_Hafiz_Naufal_Utomo'],
  ['Ganesya Intantalia', '1'],
  ['Ocha Della Fitriani', '2'],
  ['Febiola Al Jannah', 'FEBIOLA_DIRJEN_VIDKRE_MEDIA_DAN_INFORMASI'],
  ['Eugenius Nathaniel Christoffany', 'menteri_sinema_nathan'],
  ['Vincentius Nolan Nathanael Selfianto', 'DIRJEN_NOLAN_SINEMA'],
  ['Achmad Rafi Syahputra', 'DIRJEN_RAFI_SINEMA'],
  ['Farrel Ardisto Sama\'i', 'MENTERI_KEMENTERIAN_KEADILAN_DAN_KESTARAAN_GENDER_FARREL'],
  ['Annisa Rahmawati', 'DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_ANNISA_RAHMA'],
  ['Titis Nastiti Dwi Lestari', 'DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_TITIS'],
  ['Daniel Nikon Martua Situmorang', 'MENTERI_SOSPOL_DANIEL'],
  ['Moch. Amir Nazaruddin', 'MENTERISOSPOL_AMIR'],
  ['Ahmad Raihan Fadhillah Yasin', 'MENTERISOSPOLRAIHAN'],
  ['Farhan Kamil', 'MENTERI_EDIGPRO_FARHAN_KAMIL'],
  ['Argi Wicaksana', 'DIRJENEDIGPRO_ARGI_WICAKSANA'],
  ['M. Yudhistira Setya Heri', 'DIRJEN_EDIGPRO_YUDHISTIRA'],
  ['Farrel Bisma Abyakta', 'MENTERIRISKEL_BISMA'],
  ['Abdul Wahid Tamimi', 'DIRJENRISKEL_ATA'],
  ['Naswa Ayudya Sasta', 'DIRJENNASYWA_RISKEL'],
  ['Alaina Atwa Awali Harahap', 'MENTERISENIORA_ALAINA'],
  ['Abdullah Azzam', 'ABDULLAH_AZZAM_DIRJEN_SENIORA'],
  ['Muhammad Fadhil Saputra', 'v1795411756/PUPUT_DIRJEN_SENIORA.png'],
  ['Inez Faradina Kasih', 'v1795426454/BENKAN_INEZ.png'],
  ['Ananda Magfira Nurochma', 'v1795426455/BIRO_ANANDA.png'],
  ['M. Salman Al Farisi Al Ghofsah', 'v1795426456/BIRO_SALMAN.png'],
  ['Deri Bayu Setiawan', 'v1795426456/KEPALA_PAK.png'],
  ['Mevlana El Fithra Abdullah', 'v1795426456/DEPUTI_MEVLANA.png'],
  ['Nayla Khansa Abbylia', 'v1795426456/DEPUTI_NAYLA.png'],
  ['Gheahaq Danty El Zahra', 'v1795426456/SESKAB_GHEA.png'],
  ['Fathimah Nuurun Najma As-Salman', 'v1795426455/BIRO_NUURUN.png'],
  ['Najwa Nur Hafidza', 'v1795426456/BIRO_NAJWA.png']
]);

// Pre-migration names that identify the same person in older cabinet responses.
const previousNames = new Map<string, string>([
  ['Muhammad Ziyaad Difa Ul Haq', 'Muhammad Ziyaad Difaa’ul Haq'],
  ['Yustia Khuspriyanti Feby', 'Yustia Khuspriyanti Feby Martha'],
  ['M. Rifqy Yusuf', 'Muhammad Rifqy Yusuf'],
  ['Raizatya Qoriananda', 'Raizsatya Qoriananda Hutama'],
  ['Deva Zia Ul Haq', 'Deva Zia\'ul Haq Khoirudin'],
  ['Dhafa Rhofiq Firdaus', 'Dhafa Rofiq Firdaus'],
  ['Febria Rizqi Aulia S', 'Febria Rizqi Aulia Sandy'],
  ['M. Zain Hisyam Al F. Z', 'M. Zain Hisyam Al Fikri Zuhdi'],
  ['M. Wahyu', 'Mohammad Wahyu Ramadani'],
  ['Zalfaa Putri Arfiliesia', 'Zalfaa’ Putri Arfiliesia'],
  ['Hafiz Naufal Utomo', 'Hafihz Naufal Utomo'],
  ['Eugenius Nathaniel', 'Eugenius Nathaniel Christoffany'],
  ['Vicentius Nolan', 'Vincentius Nolan Nathanael Selfianto'],
  ['Farrel Ardisto Samai', 'Farrel Ardisto Sama\'i'],
  ['Titis Nastiti Dwi', 'Titis Nastiti Dwi Lestari'],
  ['Ahmad Raihan Fadhillah', 'Ahmad Raihan Fadhillah Yasin'],
  ['Argi', 'Argi Wicaksana'],
  ['Yudhistira', 'M. Yudhistira Setya Heri'],
  ['Ata', 'Abdul Wahid Tamimi'],
  ['Nasywa Ayudya S', 'Naswa Ayudya Sasta'],
  ['Naswa', 'Naswa Ayudya Sasta'],
  ['Azzam', 'Abdullah Azzam'],
  ['M. Salman Al Farisi', 'M. Salman Al Farisi Al Ghofsah'],
  ['Mevlana El Fithra', 'Mevlana El Fithra Abdullah'],
  ['Fathimah Nuurun', 'Fathimah Nuurun Najma As-Salman']
]);

const seededIDs = new Set(
  [...rosterPortraits.values(), 'PUPUT_DIRJEN_SENIORA'].map((asset) =>
    asset.replace(/^v\d+\//, '').replace(/\.png$/i, '')
  )
);

interface PortraitMember {
  name: string;
  photo?: { url?: string | null; thumbnail_url?: string | null } | null;
}

export function resolveOrganizationPortrait(member?: PortraitMember | null): string | null {
  if (!member) return null;

  const name = member.name.trim();
  const asset = rosterPortraits.get(previousNames.get(name) ?? name);
  for (const candidate of [member.photo?.url, member.photo?.thumbnail_url]) {
    const url = candidate?.trim();
    if (!url || !/^https:\/\/res\.cloudinary\.com\/[\w-]+\/image\/upload\/[^\s?#]+(?:[?#].*)?$/.test(url)) continue;

    if (url.startsWith(cabinetCloud)) {
      const path = url.slice(cabinetCloud.length).split(/[?#]/)[0];
      const match = /^(?:[a-z]+_[^/]+\/)*(?:v\d+\/)?([^/]+)$/.exec(path);
      const id = match?.[1].replace(/\.(png|jpe?g|webp|avif)$/i, '');
      if (id && seededIDs.has(id)) continue;
    }

    // A separate Cloudinary upload belongs to the profile that selected it.
    return url;
  }

  return asset ? portraitBase + asset : null;
}
