package models

type MahasiswaModelNoID struct {
	Name string `db:"nama"`
	Nim  string `db:"nim"`
}

type MahasiswaModels struct {
	ID   int64  `db:"id"`
	Name string `db:"nama"`
	Nim  string `db:"nim"`
}

type AlamatsModel struct {
	ID           int64  `db:"id"`
	Jalan        string `db:"jalan"`
	NoRumah      string `db:"no_rumah"`
	IDMahasiswas int64  `db:"id_mahasiswas"`
}

type MahasiswaAlamatsModels struct {
	ID           int64  `db:"id"`
	Name         string `db:"nama"`
	Nim          string `db:"nim"`
	Jalan        string `db:"jalan"`
	NoRumah      string `db:"no_rumah"`
	IDMahasiswas int64  `db:"id_mahasiswas"`
}

type AlamatsOnlyModel struct {
	Jalan   string `db:"jalan"`
	NoRumah string `db:"no_rumah"`
}
