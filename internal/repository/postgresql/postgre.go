package repository

import (
	"fmt"
	"log"

	"github.com/ifty123/kampus_service/internal/models"
	"github.com/ifty123/kampus_service/internal/repository"

	mhsErrors "github.com/ifty123/kampus_service/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const (
	SaveMahasiswaNoAlamat = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now())`
	SaveMahasiswa         = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat   = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	UpdateMahasiswaNama   = `UPDATE kampus.mahasiswas SET nama = $1, updated_at = now() where id = $2`
	TampilMahasiswa       = `SELECT id,nama,nim FROM kampus.mahasiswas ORDER BY id ASC`
	TampilMahasiswaAlamat = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah, b.id_mahasiswas
							FROM kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b ON a.id = b.id_mahasiswas ORDER BY a.id ASC`
	AmbilIDMahasiswa      = `SELECT id from kampus.mahasiswas WHERE id = $1`
	UpdateAlamatMahasiswa = `UPDATE kampus.mahasiswa_alamats SET jalan = $1, no_rumah = $2, updated_at = now() where id = $3`
	GetMahasiswaByID      = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah, b.id_mahasiswas FROM kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b
							ON a.id = b.id_mahasiswas WHERE a.id = $1`
	GetMahasiswaByName = `SELECT a.id, a.nama, a.nim, b.jalan, b.no_rumah, b.id_mahasiswas FROM kampus.mahasiswas a JOIN kampus.mahasiswa_alamats b
							ON a.id = b.id_mahasiswas WHERE a.nama = $1`
	GetOnlyAlamat   = `SELECT DISTINCT no_rumah, jalan from kampus.mahasiswa_alamats ORDER BY no_rumah`
	GetOnlyIdIfNull = `SELECT id, nim, nama FROM kampus."mahasiswas" WHERE %s`
)

var statement PreparedStatement

//persiapan mengambil data dari database
type PreparedStatement struct {
	updateMahasiswaNama   *sqlx.Stmt
	tampilMahasiswa       *sqlx.Stmt
	tampilMahasiswaAlamat *sqlx.Stmt
	ambilIDMahasiswa      *sqlx.Stmt
	saveMahasiswaAlamat   *sqlx.Stmt
	updateAlamatMahasiswa *sqlx.Stmt
	getMahasiswaByID      *sqlx.Stmt
	getMahasiswaByName    *sqlx.Stmt
	getOnlyAlamat         *sqlx.Stmt
	saveDosens            *sqlx.Stmt
	saveDosenAlamat       *sqlx.Stmt
	updateDosens          *sqlx.Stmt
	tampilDosen           *sqlx.Stmt
	tampilDosenByID       *sqlx.Stmt
	tampilDosenAlamat     *sqlx.Stmt
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

//prepare isi InitPrepare dengan isian Conn, func diberi isian
func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

//mengeksekusi query dan mengembalikan dalam bentuk sqlx.stmt
func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

//karna m struct dan nempel di method preparex, maka method preparex diambil dari m.
//ambil data dari database
func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		updateMahasiswaNama:   m.Preparex(UpdateMahasiswaNama),
		tampilMahasiswa:       m.Preparex(TampilMahasiswa),
		tampilMahasiswaAlamat: m.Preparex(TampilMahasiswaAlamat),
		ambilIDMahasiswa:      m.Preparex(AmbilIDMahasiswa),
		saveMahasiswaAlamat:   m.Preparex(SaveMahasiswaAlamat),
		updateAlamatMahasiswa: m.Preparex(UpdateAlamatMahasiswa),
		getMahasiswaByID:      m.Preparex(GetMahasiswaByID),
		getMahasiswaByName:    m.Preparex(GetMahasiswaByName),
		getOnlyAlamat:         m.Preparex(GetOnlyAlamat),
		saveDosens:            m.Preparex(SaveDosens),
		saveDosenAlamat:       m.Preparex(SaveDosensAlamat),
		updateDosens:          m.Preparex(UpdateDosens),
		tampilDosen:           m.Preparex(TampilDosen),
		tampilDosenByID:       m.Preparex(TampilDosenByID),
		tampilDosenAlamat:     m.Preparex(TampilDosenAlamat),
	}
}

//simpan mahasiswa
func (p *PostgreSQLRepo) SaveMahasiswa(dataMahasiswa *models.MahasiswaModelNoID) error {

	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	_, err = tx.Exec(SaveMahasiswa, dataMahasiswa.Name, dataMahasiswa.Nim)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveMahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	return tx.Commit()
}

//simpan alamat dengan ID sekian
func (p *PostgreSQLRepo) SaveAlamatFromMahasiswa(dataAlamat *models.AlamatsModel) error {

	_, err := statement.saveMahasiswaAlamat.Exec(dataAlamat.Jalan, dataAlamat.NoRumah, dataAlamat.IDMahasiswas)

	if err != nil {
		log.Println("Failed Query Save Alamat From Mahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	return nil

}

//simpan data mahasiswa dan alamat
func (p *PostgreSQLRepo) SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatsModels) error {

	//ini ada tx dan commit, jika ada satu error di tabel anak, tabel induk ga ke input data nya
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa Alamat : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	var idMahasiswa int64
	err = tx.QueryRow(SaveMahasiswa, dataMahasiswa.Name, dataMahasiswa.Nim).Scan(&idMahasiswa)

	if err != nil {
		tx.Rollback()
		log.Println("Failed Query SaveMahasiswa: ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	for _, val := range dataAlamat {
		_, err = tx.Exec(SaveMahasiswaAlamat, val.Jalan, val.NoRumah, idMahasiswa)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Query SaveMahasiswaAlamat : ", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}

	return tx.Commit()
}

func (p *PostgreSQLRepo) UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error {

	result, err := statement.updateMahasiswaNama.Exec(dataMahasiswa.Name, dataMahasiswa.ID)

	if err != nil {
		log.Println("Failed Query UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd UpdateMahasiswaNama : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateMahasiswaNama: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}

//update alamats mahasiswa
func (p *PostgreSQLRepo) UpdateMahasiswaAlamat(dataMahasiswa *models.AlamatsModel) error {

	result, err := statement.updateAlamatMahasiswa.Exec(dataMahasiswa.Jalan, dataMahasiswa.NoRumah, dataMahasiswa.ID)

	if err != nil {
		log.Println("Failed Query UpdateAlamatsMahasiswa : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd UpdateAlamatMahasiswa : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateAlamatMahasiswa: No Data Changed")
		return fmt.Errorf(mhsErrors.ErrorNoDataChange)
	}

	return nil
}

func (p *PostgreSQLRepo) TampilMahasiswa() ([]*models.MahasiswaModels, error) {
	data := []*models.MahasiswaModels{}
	err := statement.tampilMahasiswa.Select(&data)

	if err != nil {
		log.Println("Failed Query tampilMahasiswa : ", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	fmt.Println(data)
	return data, nil
}

func (p *PostgreSQLRepo) TampilMahasiswaAlamat() ([]*models.MahasiswaAlamatsModels, error) {

	dataMaha := []*models.MahasiswaAlamatsModels{}
	err := statement.tampilMahasiswaAlamat.Select(&dataMaha)

	if err != nil {
		log.Println("Failed Query UpdateMahasiswaNama : ", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	return dataMaha, nil
}

func (p *PostgreSQLRepo) GetMahasiswaByID(id int64) ([]*models.MahasiswaAlamatsModels, error) {

	//fmt.Println(id, "here")
	dataMaha := []*models.MahasiswaAlamatsModels{}
	err := statement.getMahasiswaByID.Select(&dataMaha, id)

	fmt.Println("panjang dataMaha ", len(dataMaha))

	if err != nil {
		log.Println("Failed Query GetMAhasiswa By ID : ", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}
	if len(dataMaha) < 1 {
		return nil, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}
	return dataMaha, nil
}

//fungsi get mahasiswa by name
func (p *PostgreSQLRepo) GetMahasiswaByName(name string) ([]*models.MahasiswaAlamatsModels, error) {
	dataMahasiswa := []*models.MahasiswaAlamatsModels{}
	err := statement.getMahasiswaByName.Select(&dataMahasiswa, name)

	if err != nil {
		log.Println("Failed Query GetMahasiswaByName : ", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(dataMahasiswa) < 1 {
		return nil, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}
	return dataMahasiswa, nil
}

func (p *PostgreSQLRepo) GetOnlyAlamat() ([]*models.AlamatsOnlyModel, error) {
	dataAlamat := []*models.AlamatsOnlyModel{}
	err := statement.getOnlyAlamat.Select(&dataAlamat)

	if err != nil {
		log.Println("Failed Query Get Only Alamat", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(dataAlamat) < 1 {
		log.Println("Data Not Found", fmt.Errorf(mhsErrors.ErrorDataNotFound))
	}

	return dataAlamat, nil
}

func (p *PostgreSQLRepo) GetOnlyIdIfNull(param *models.MahasiswaModels) ([]*models.MahasiswaModels, error) {
	dataNama := []*models.MahasiswaModels{}
	//var AddStatement string
	var err error

	if len(param.Name) != 0 {
		StatementNama := `"nama" = '%s'`
		AddStatement := fmt.Sprintf(StatementNama, param.Name)
		fmt.Println(AddStatement)
		err = p.Conn.Select(&dataNama, fmt.Sprintf(GetOnlyIdIfNull, AddStatement))

	} else {
		err = p.Conn.Select(&dataNama, (fmt.Sprintf(GetOnlyIdIfNull, "1=1")))
	}

	if err != nil {
		log.Println("Failed Query Get Only ID", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(dataNama) < 1 {
		log.Println("Data Not Found", fmt.Errorf(mhsErrors.ErrorDataNotFound))
	}

	return dataNama, nil
}
