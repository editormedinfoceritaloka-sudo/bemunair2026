package database

type seedAccount struct {
	name, email, role, unit, phone string
	leader                         bool
}

var seedAccounts = []seedAccount{
	{"Admin BEM UNAIR", "admin@bem.unair.ac.id", "ADMIN_MEDINFO", "", "6281200000001", false},
	{"Johanes Richard Darmawan", "menko.kemahasiswaan@bem.unair.ac.id", "ADMIN", "MENKO_KEMAHASISWAAN", "6281200000002", false},
	{"Muhammad Ziyaad Difa Ul Haq", "menko.kemasyarakatan@bem.unair.ac.id", "ADMIN", "MENKO_KEMASYARAKATAN", "6281200000003", false},
	{"Andrean Marcellino", "menko.kominfo@bem.unair.ac.id", "ADMIN", "MENKO_KOMINFO", "6281200000004", false},
	{"Dewa Pranata Putra Pratama", "menko.pergerakan@bem.unair.ac.id", "ADMIN", "MENKO_PERGERAKAN", "6281200000005", false},
	{"Nasua Okta Kusuma Dewi Rofiq", "menko.ppk@bem.unair.ac.id", "ADMIN", "MENKO_PPK", "6281200000006", false},
	{"Elsa Putri Aprizalni", "menteri.adkesma@bem.unair.ac.id", "ADMIN", "ADKESMA", "6281200000007", true},
	{"M. Rifqy Yusuf", "menteri.pengprof@bem.unair.ac.id", "ADMIN", "PENGPROF", "6281200000008", true},
	{"Labib Azka Asysyafiq", "menteri.psdm@bem.unair.ac.id", "ADMIN", "PSDM", "6281200000009", true},
	{"Zaskia Darojah", "menteri.lh@bem.unair.ac.id", "ADMIN", "LH", "6281200000010", true},
	{"An Nadhofah Adlin", "menteri.kemenkes@bem.unair.ac.id", "ADMIN", "KEMENKES", "6281200000011", true},
	{"Sahrul Efendi", "menteri.pengmas@bem.unair.ac.id", "ADMIN", "PENGMAS", "6281200000012", true},
	{"Zalfaa Putri Arfiliesia", "menteri.hublu@bem.unair.ac.id", "ADMIN", "HUBLU", "6281200000013", true},
	{"Ganesya Intantalia (Caca)", "caca.medinfo@bem.unair.ac.id", "ADMIN_MEDINFO", "MEDINFO", "6281200000014", true},
	{"Ocha Della Fitriani", "ocha.medinfo@bem.unair.ac.id", "ADMIN_MEDINFO", "MEDINFO", "6281200000015", false},
	{"Febiola Al Jannah", "febi.medinfo@bem.unair.ac.id", "ADMIN_MEDINFO", "MEDINFO", "6281200000016", false},
	{"Eugenius Nathaniel", "menteri.sinema@bem.unair.ac.id", "ADMIN", "SINEMA", "6281200000017", true},
	{"Farrel Ardisto Samai", "menteri.kkg@bem.unair.ac.id", "ADMIN", "KKG", "6281200000018", true},
	{"Daniel Nikon Martua Situmorang", "menteri.sospol@bem.unair.ac.id", "ADMIN", "SOSPOL", "6281200000019", true},
	{"Farhan Kamil", "menteri.edigpro@bem.unair.ac.id", "ADMIN", "EDIGPRO", "6281200000020", true},
	{"Farrel Bisma Abyakta", "menteri.riskel@bem.unair.ac.id", "ADMIN", "RISKEL", "6281200000021", true},
	{"Alaina Atwa Awali Harahap", "menteri.seniiora@bem.unair.ac.id", "ADMIN", "SENIIORA", "6281200000022", true},
	{"Inez Faradina Kasih", "menteri.benkab@bem.unair.ac.id", "ADMIN", "BENKAB", "6281200000023", true},
	{"Deri Bayu Setiawan", "kepala.pak@bem.unair.ac.id", "ADMIN", "PAK", "6281200000024", true},
	{"Gheahaq Danty El Zahra", "menteri.seskab@bem.unair.ac.id", "ADMIN", "SESKAB", "6281200000025", true},
}
