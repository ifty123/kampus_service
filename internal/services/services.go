package services

import (
	"github.com/ifty123/kampus_service/pkg/dto"
)

type Services interface {
	//menerima dari mahasiswaDTO
	SaveMahasiswaAlamat(req *dto.MahasiswaAlamatReqDTO) error
	UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error
	TampilMahasiswa() ([]*dto.MahasiswaRespDTO, error)
	TampilMahasiswaAlamats() ([]*dto.MahasiswaAlamatRespDTO, error)
	SaveMahasiswa(req *dto.MahasiswaReqDTO) error
	SaveAlamatFromMahasiswa(req *dto.AlamatReqWithIDDTO) error
	UpdateAlamatMahasiswa(req *dto.UpdateAlamatMahasiswaReq) error
	GetMahasiswaByID(req *dto.GetMahasiswaByIDReqDTO) (*dto.MahasiswaAlamatRespDTO, error)
	GetMahasiswaByName(req *dto.GetMahasiswaByNamaReqDTO) ([]*dto.MahasiswaAlamatRespDTO, error)
	GetOnlyAlamat() ([]*dto.AlamatReqDTO, error)
	GetAllIfNull(req *dto.GetNamaIfNotNull) ([]*dto.MahasiswaRespDTO, error)
	SaveDosens(req *dto.DosensReqDTO) error
	SaveAlamatsDosen(req *dto.DosensAlamatReqDTO) error
	SaveDosenAndAlamat(data *dto.DosenAndAlamatReqDTO) error
	UpdateDosens(dataDosen *dto.DosensWithIDReqDTO) error
	TampilDosen() ([]*dto.DosenRespDTO, error)
	TampilDosenByID(dataID *dto.DosenIdDTO) (*dto.DosenRespDTO, error)
	TampilDosenAlamat() ([]*dto.DosenAndAlamatRespDTO, error)
	UpdateDosenAlamat(dataDosen *dto.DosenAndAlamatReqWithIDDTO) error

	// with integration
	GetIntegDadJoke(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error)
}
