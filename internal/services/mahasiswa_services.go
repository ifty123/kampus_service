package services

import (
	"fmt"

	integ "github.com/ifty123/kampus_service/internal/integration"
	"github.com/ifty123/kampus_service/internal/repository"
	"github.com/ifty123/kampus_service/pkg/dto"
	"github.com/ifty123/kampus_service/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
}

func (s *service) SaveMahasiswaAlamat(req *dto.MahasiswaAlamatReqDTO) error {

	dtAlamat := assembler.ToSaveMahasiswaAlamats(req.Alamats)
	dtMahasiswa := assembler.ToSaveMahasiswa(req)

	err := s.repo.SaveMahasiswaAlamat(dtMahasiswa, dtAlamat)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) SaveAlamatFromMahasiswa(req *dto.AlamatReqWithIDDTO) error {

	dtAlamat := assembler.ToSaveAlamatWithID(req)

	err := s.repo.SaveAlamatFromMahasiswa(dtAlamat)
	if err != nil {
		return err
	}
	fmt.Println("Data masuk")
	return nil
}

//simpan mahasiswa
func (s *service) SaveMahasiswa(req *dto.MahasiswaReqDTO) error {

	dtMahasiswa := assembler.ToSaveMahasiswaNoAlamat(req)

	err := s.repo.SaveMahasiswa(dtMahasiswa)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateMahasiswaNama(req *dto.UpadeMahasiswaNamaReqDTO) error {

	dtMhsiswa := assembler.ToUpdateMahasiswaNama(req)

	err := s.repo.UpdateMahasiswaNama(dtMhsiswa)
	if err != nil {
		return err
	}

	return nil
}

//update alamatMahasiswa
func (s *service) UpdateAlamatMahasiswa(req *dto.UpdateAlamatMahasiswaReq) error {

	dtMhsiswa := assembler.ToUpdateAlamatMahasiswa(req)

	err := s.repo.UpdateMahasiswaAlamat(dtMhsiswa)
	if err != nil {
		return err
	}

	return nil
}

//tampil semua data
func (s *service) TampilMahasiswa() ([]*dto.MahasiswaRespDTO, error) {
	data, err := s.repo.TampilMahasiswa()

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToMahasiswas(data)
	return data1, nil
}

//tampil semua data
func (s *service) TampilMahasiswaAlamats() ([]*dto.MahasiswaAlamatRespDTO, error) {
	data, err := s.repo.TampilMahasiswaAlamat()

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToMhs(data)
	data2 := assembler.ToAlamats1(data)

	for _, val := range data1 {
		for _, val1 := range data2 {
			if val.ID == val1.IDMahasiswas {
				val.Alamats = append(val.Alamats, val1)
			}
		}
	}
	return data1, nil
}

//dibenerin dulu DTO nya
func (s *service) GetMahasiswaByID(req *dto.GetMahasiswaByIDReqDTO) (*dto.MahasiswaAlamatRespDTO, error) {

	data, err := s.repo.GetMahasiswaByID(req.ID)

	if err != nil {
		return nil, err
	}

	dataMap := make(map[int64]*dto.MahasiswaAlamatRespDTO)
	for _, val := range data {
		//karna val bentuknya array dan return nya 2, dan ini pengecekan id
		//jika ada id yg sama, maka false
		if _, ok := dataMap[val.ID]; !ok {
			dataMap[val.ID] = &dto.MahasiswaAlamatRespDTO{
				ID:   val.ID,
				Nama: val.Name,
				NIM:  val.Nim,
			}

			dataMap[val.ID].Alamats = append(dataMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		} else {
			dataMap[val.ID].Alamats = append(dataMap[val.ID].Alamats, &dto.AlamatRespDTO{
				Jalan:   val.Jalan,
				NoRumah: val.NoRumah,
			})
		}
	}

	var dataDTO *dto.MahasiswaAlamatRespDTO

	for _, val := range dataMap {
		dataDTO = val
	}

	return dataDTO, nil

}

func (s *service) GetMahasiswaByName(req *dto.GetMahasiswaByNamaReqDTO) ([]*dto.MahasiswaAlamatRespDTO, error) {

	data, err := s.repo.GetMahasiswaByName(req.Name)

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToMhs(data)
	dataAlamat := assembler.ToAlamats1(data)

	fmt.Println("data GetMahasiswaBy Name ", data1)

	for _, val := range data1 {
		for _, val1 := range dataAlamat {
			if val.ID == val1.IDMahasiswas {
				val.Alamats = append(val.Alamats, val1)
			}
		}
	}

	return data1, nil
}

func (s *service) GetAllIfNull(req *dto.GetNamaIfNotNull) ([]*dto.MahasiswaRespDTO, error) {

	NamaAsemb := assembler.ToAssembNama(req)

	data, err := s.repo.GetOnlyIdIfNull(NamaAsemb)

	if err != nil {
		return nil, err
	}

	data1 := assembler.ToMahasiswas(data)

	fmt.Println("data GetMahasiswaBy Name ", data1)

	return data1, nil
}

func (s *service) GetOnlyAlamat() ([]*dto.AlamatReqDTO, error) {
	dataAlamats, err := s.repo.GetOnlyAlamat()

	if err != nil {
		return nil, err
	}

	dataAlamatAssembler := assembler.ToGetOnlyAlamatsAll(dataAlamats)

	return dataAlamatAssembler, nil
}

func (s *service) GetIntegDadJoke(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error) {
	var resp *dto.GetDadJokesRandomRespDTO

	resp, err := s.IntegServ.GetRandomDadJokes(req)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
