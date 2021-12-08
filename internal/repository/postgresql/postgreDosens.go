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
