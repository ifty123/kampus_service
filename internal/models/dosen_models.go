package models

//dosens
type DosensModels struct {
	Id   int64  `db:"id"`
	Name string `db:"nama"`
	Nidn string `db:"nidn"`
}

type DosenAlamatsModels struct {
	Jalan    string `db:"jalan"`
	No_rumah string `db:"no_rumah"`
	IdDosen  int64  `db:"id_dosens"`
}

type DosenAndAlamatsModel struct {
	Id       int64  `db:"id"`
	Name     string `db:"nama"`
	Nidn     string `db:"nidn"`
	Jalan    string `db:"jalan"`
	No_rumah string `db:"no_rumah"`
	IdDosen  int64  `db:"id_dosens"`
}
