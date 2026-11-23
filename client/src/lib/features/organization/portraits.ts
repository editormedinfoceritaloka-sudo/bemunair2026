const portraitBase = 'https://res.cloudinary.com/rvn0bdia/image/upload/e_background_removal,f_auto,q_auto/';

const legacyPortraits: Partial<Record<string, string>> = {
  "/images/Menko/menko_kemahasiswaan_johanes_richard_darmawan":
    "MENKO_KEMAHASISWAAAN_RICHIE",
  "/images/Menko/menko_kemasyarakatan_muhammad_ziyaad_difa_ul_haq":
    "ZIYAAD_MENKO_KEMASYARAKATAN",
  "/images/Menko/menko_kominfo_andrean_marcellino":
    "MENKO_KOMINFO_ANDREAN",
  "/images/Menko/menko_pergerakan_dewa_pranata_putra_pratama":
    "DEWA_MENKO_PERGERAKAN",
  "/images/Menko/menko_ppk_nasua_okta_kusuma_dewi_rofiq":
    "NASUA_OKTA_MENKO_PPK",
  "/images/ADKESMA/dirjen_adkesma_advokasi_yustia_khuspriyanti_feby":
    "Direktur_JenderalYustia_Khuspriyanti_Feby_ADKESMA",
  "/images/ADKESMA/dirjen_adkesma_kesejahteraan_nabila_afifah_abida":
    "Direktur_JenderalNabila_Afifah_Abida_ADKESMA",
  "/images/PENGPROF/dirjen_pengprof_keahlian_raizatya_qoriananda":
    "DIRJEN_RAIS_PENGPROF",
  "/images/PENGPROF/dirjen_pengprof_relasi_deva_zia_ul_haq":
    "DIRJEN_PENGPROFDeva_Zia_Ul_Haq",
  "/images/PSDM/dirjen_psdm_akselerasi_ferdian_amartha":
    "DIRJEN_PSDM_FERDIAN",
  "/images/PSDM/dirjen_psdm_kaderisasi_qheisya_luna_alifa":
    "DIRJENPSDM_LUNA",
  "/images/LH/dirjen_lh_advokasi_kajian_strategis":
    "DIRJEN_LINGKUNGAN_HIDUP_DAFFA",
  "/images/LH/dirjen_lh_pencegahan_pengendalian":
    "DIRJEN_LINGKUNGAN_HIDUP_FEBRIA",
  "/images/LH/dirjen_lh_pengembangan_preservasi":
    "DIRJEN_LINGKUNGANHIDUP_AZIZAH",
  "/images/Menkes/dirjen_menkes_kesehatan_masyarakat":
    "DIRJEN_KESEHATAN_ZEIN",
  "/images/Menkes/dirjen_menkes_psikososial_inklusivitas":
    "DIRJEN_KESEHATAN_CYINTIA",
  "/images/Pengabdian_Masyarakat/dirjen_pengmas_fahrur_rosi":
    "DIRJEN_PENGABDIAN_MASYARAKAT_ROSI",
  "/images/Pengabdian_Masyarakat/dirjen_pengmas_m_wahyu":
    "DIRJEN_PENGMAS_WAHYU",
  "/images/Hubungan_Luar/dirjen_hublu_billal_syahdan_arrafah":
    "DIRJEN_HUBLU_BILLAL",
  "/images/Hubungan_Luar/dirjen_hublu_hafiz_naufal_utomo":
    "DIRJEN_HUBLU_Hafiz_Naufal_Utomo",
  "/images/Medinfo/dirjen_medinfo_publikasi_ocha_della_fitriani":
    "2",
  "/images/Medinfo/dirjen_medinfo_visual_febiola_al_jannah":
    "FEBIOLA_DIRJEN_VIDKRE_MEDIA_DAN_INFORMASI",
  "/images/Sinergitas_Mahasiswa/dirjen_sinema_bina_vicentius_nolan":
    "DIRJEN_NOLAN_SINEMA",
  "/images/Sinergitas_Mahasiswa/dirjen_sinema_korlaborasi_achmad_rafi_syahputra":
    "DIRJEN_RAFI_SINEMA",
  "/images/KKG/dirjen_kkg_edukasi_annisa_rahmawati":
    "DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_ANNISA_RAHMA",
  "/images/KKG/dirjen_kkg_pemberdayaan_titis_nastiti_dwi":
    "DIRJEN_KEMENTERIAN_KESTARAAN_GENDER_TITIS",
  "/images/Sosial_dan_Politik/dirjen_sospol_aksi_moch_amir_nazaruddin":
    "MENTERISOSPOL_AMIR",
  "/images/Sosial_dan_Politik/dirjen_sospol_analisis_ahmad_raihan_fadhillah":
    "MENTERISOSPOLRAIHAN",
  "/images/EDIGPRO/dirjen_edigpro_argi":
    "DIRJENEDIGPRO_ARGI_WICAKSANA",
  "/images/EDIGPRO/dirjen_edigpro_yudhistira":
    "DIRJEN_EDIGPRO_YUDHISTIRA",
  "/images/Riset_dan_Keilmuan/dirjen_riset_ata":
    "DIRJENRISKEL_ATA",
  "/images/Riset_dan_Keilmuan/dirjen_riset_naswa":
    "DIRJENNASYWA_RISKEL",
  "/images/Seni_dan_Olahraga/dirjen_seniora_azzam":
    "ABDULLAH_AZZAM_DIRJEN_SENIORA",
  "/images/Seni_dan_Olahraga/dirjen_seniora_m_fadhil":
    "PUPUT_DIRJEN_SENIORA"
};

export function resolveOrganizationPortrait(url: string): string {
  const portrait = legacyPortraits[url.replace(/\.(png|webp)$/, '')];
  return portrait ? portraitBase + portrait : url;
}
