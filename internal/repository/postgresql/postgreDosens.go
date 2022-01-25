package repository

import (
	"fmt"
	"log"

	"github.com/ifty123/kampus_service/internal/models"

	mhsErrors "github.com/ifty123/kampus_service/pkg/errors"
)

const (
	SaveDosens          = `INSERT INTO kampus.dosens (nama, nidn, created_at) VALUES ($1, $2, now())`
	SaveDosenWithReturn = `INSERT INTO kampus.dosens(nama, nidn, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveDosensAlamat    = `INSERT INTO kampus.dosen_alamats (jalan, no_rumah, created_at, id_dosens) VALUES ($1, $2, now(), $3)`
	UpdateDosens        = `UPDATE kampus.dosens SET nama=$1, nidn=$2, updated_at=now() WHERE id=$3`
	UpdateDosenAlamat   = `UPDATE kampus.dosen_alamats SET jalan=$1, no_rumah=$2 updated_at=now() WHERE id=$3`
	TampilDosen         = `SELECT id, nama, nidn FROM kampus.dosens`
	TampilDosenByID     = `SELECT id, nidn, nama FROM kampus.dosens WHERE id = $1`
	TampilDosenAlamat   = `SELECT a.id, a.nidn, a.nama, b.jalan, b.no_rumah, b.id_dosens FROM kampus.dosens a JOIN kampus.dosen_alamats b
							ON a.id = b.id_dosens ORDER BY a.id`
)

//func simpan dosens
func (p *PostgreSQLRepo) SaveDosens(dosen *models.DosensModels) error {
	_, err := statement.saveDosens.Exec(dosen.Name, dosen.Nidn)
	//cek
	if err != nil {
		log.Println("Failed Query Save Dosens", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	return nil
}

//simpan alamat
func (p *PostgreSQLRepo) SaveAlamatsDosen(dosen *models.DosenAlamatsModels) error {
	fmt.Println("isi model alamat dosen", dosen)
	_, err := statement.saveDosenAlamat.Exec(dosen.Jalan, dosen.No_rumah, dosen.IdDosen)

	if err != nil {
		log.Println("Failed to query Save Dosen Alamat", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	return nil
}

//save data dosen & alamat
func (p *PostgreSQLRepo) SaveDosensAndAlamats(dataDosen *models.DosensModels, dataAlamat []*models.DosenAndAlamatsModel) error {
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed to begin Tx Save dosen And Alamat", err.Error())
		return fmt.Errorf("Error at", err.Error())
	}
	//inisialisasi returning
	var id_return int64
	//save dosens
	err = tx.QueryRow(SaveDosenWithReturn, dataDosen.Name, dataDosen.Nidn).Scan(&id_return)
	if err != nil {
		tx.Rollback()
		log.Println("failed to save dosen with Return", err.Error())
		return fmt.Errorf("Failed : ", err.Error(), mhsErrors.ErrorDB)
	}

	//save alamat
	for _, valAlamat := range dataAlamat {
		_, err = tx.Exec(SaveDosensAlamat, valAlamat.Jalan, valAlamat.No_rumah, id_return)
		//cek
		if err != nil {
			tx.Rollback()
			log.Println("Failed to save Dosen Alamat", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}
	return tx.Commit()
}

func (p *PostgreSQLRepo) UpdateDosens(dataDosen *models.DosensModels) error {
	_, err := statement.updateDosens.Exec(dataDosen.Name, dataDosen.Nidn, dataDosen.Id)
	if err != nil {
		log.Println("failed to Update dosen", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}
	return nil
}

func (p *PostgreSQLRepo) TampilDosens() ([]*models.DosensModels, error) {

	//tampung data
	data := []*models.DosensModels{}
	err := statement.tampilDosen.Select(&data)
	if err != nil {
		log.Println("Failed to Tampil Dosen", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(data) < 1 {
		return nil, fmt.Errorf("Data nor found", mhsErrors.ErrNotFound)
	}
	fmt.Println(data)
	return data, nil
}

func (p *PostgreSQLRepo) TampilDosenByID(dataID int64) ([]*models.DosensModels, error) {
	dataDosen := []*models.DosensModels{}
	err := statement.tampilDosenByID.Select(&dataDosen, dataID)
	if err != nil {
		log.Println("failed to Tampil dosen by Id", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(dataDosen) < 1 {
		log.Println("Failed to Tampil dosen By Id", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}

	return dataDosen, nil
}

func (p *PostgreSQLRepo) TampilDosenAlamat() ([]*models.DosenAndAlamatsModel, error) {
	dataDosen := []*models.DosenAndAlamatsModel{}
	err := statement.tampilDosenAlamat.Select(&dataDosen)
	if err != nil {
		log.Println("Failed to query Tampil Dosen Alamat", err.Error())
		return nil, fmt.Errorf(mhsErrors.ErrorDB)
	}

	if len(dataDosen) < 1 {
		log.Println("Failed to tampil Dosen Alam")
		return nil, fmt.Errorf(mhsErrors.ErrorDataNotFound)
	}

	return dataDosen, nil
}

func (p *PostgreSQLRepo) UpdateDosenAlamat(dataDosen *models.DosensModels, dataAlamat []*models.DosenAndAlamatsModel) error {
	tx, err := p.Conn.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx SaveMahasiswa : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	_, err = tx.Exec(UpdateDosens, dataDosen.Name, dataDosen.Nidn, dataDosen.Id)
	if err != nil {
		tx.Rollback()
		log.Println("Failed Begin Tx SaveMahasiswa : ", err.Error())
		return fmt.Errorf(mhsErrors.ErrorDB)
	}

	for _, valAlamat := range dataAlamat {
		_, err := tx.Exec(UpdateDosenAlamat, valAlamat.Jalan, valAlamat.No_rumah, valAlamat.Id)
		if err != nil {
			tx.Rollback()
			log.Println("Failed Begin Tx SaveMahasiswa : ", err.Error())
			return fmt.Errorf(mhsErrors.ErrorDB)
		}
	}

	return tx.Commit()
}
