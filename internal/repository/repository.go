package repository

import "github.com/ifty123/kampus_service/internal/models"

type Repository interface {
	SaveMahasiswaAlamat(dataMahasiswa *models.MahasiswaModels, dataAlamat []*models.MahasiswaAlamatsModels) error
	UpdateMahasiswaNama(dataMahasiswa *models.MahasiswaModels) error
	TampilMahasiswa() ([]*models.MahasiswaModels, error)
	TampilMahasiswaAlamat() ([]*models.MahasiswaAlamatsModels, error)
	SaveMahasiswa(dataMahasiswa *models.MahasiswaModelNoID) error
	SaveAlamatFromMahasiswa(dataAlamat *models.AlamatsModel) error
	UpdateMahasiswaAlamat(dataMahasiswa *models.AlamatsModel) error
	GetMahasiswaByID(id int64) ([]*models.MahasiswaAlamatsModels, error)
	GetMahasiswaByName(name string) ([]*models.MahasiswaAlamatsModels, error)
	GetOnlyAlamat() ([]*models.AlamatsOnlyModel, error)
	GetOnlyIdIfNull(param *models.MahasiswaModels) ([]*models.MahasiswaModels, error)
	SaveDosens(dosen *models.DosensModels) error
	SaveAlamatsDosen(dosen *models.DosenAlamatsModels) error
	SaveDosensAndAlamats(dataDosen *models.DosensModels, dataAlamat []*models.DosenAndAlamatsModel) error
}
