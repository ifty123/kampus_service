package assembler

import (
	"fmt"

	"github.com/ifty123/kampus_service/internal/models"
	"github.com/ifty123/kampus_service/pkg/dto"
)

//save mahasiswa
func ToSaveMahasiswaNoAlamat(d *dto.MahasiswaReqDTO) *models.MahasiswaModelNoID {
	return &models.MahasiswaModelNoID{
		Name: d.Nama,
		Nim:  d.Nim,
	}
}

func ToAssembNama(d *dto.GetNamaIfNotNull) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
	}
}

//save alamat with ID
func ToSaveAlamatWithID(d *dto.AlamatReqWithIDDTO) *models.AlamatsModel {
	return &models.AlamatsModel{
		Jalan:        d.Jalan,
		NoRumah:      d.NoRumah,
		IDMahasiswas: d.IDMahasiswas,
	}
}

//cuma simpan mahasiswa sebelum alamat
func ToSaveMahasiswa(d *dto.MahasiswaAlamatReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		Nim:  d.Nim,
	}
}

//simpan mahasiswa dengan alamat
func ToSaveMahasiswaAlamat(d *dto.AlamatReqDTO) *models.MahasiswaAlamatsModels {
	return &models.MahasiswaAlamatsModels{
		Jalan:   d.Jalan,
		NoRumah: d.NoRumah,
	}
}

func ToSaveMahasiswaAlamats(datas []dto.AlamatReqDTO) []*models.MahasiswaAlamatsModels {
	var mds []*models.MahasiswaAlamatsModels
	for _, m := range datas {
		mds = append(mds, ToSaveMahasiswaAlamat(&m))
	}
	return mds
}

func ToUpdateMahasiswaNama(d *dto.UpadeMahasiswaNamaReqDTO) *models.MahasiswaModels {
	return &models.MahasiswaModels{
		Name: d.Nama,
		ID:   d.ID,
	}
}

//update alamats
func ToUpdateAlamatMahasiswa(d *dto.UpdateAlamatMahasiswaReq) *models.AlamatsModel {
	return &models.AlamatsModel{
		Jalan:   d.Jalan,
		NoRumah: d.NoRumah,
		ID:      d.ID,
	}
}

func ToGetMahasiswa(m *models.MahasiswaModels) *dto.MahasiswaRespDTO {
	return &dto.MahasiswaRespDTO{
		ID:   m.ID,
		Nama: m.Name,
		NIM:  m.Nim,
	}
}

func ToMahasiswas(datas []*models.MahasiswaModels) []*dto.MahasiswaRespDTO {
	var ds []*dto.MahasiswaRespDTO
	for _, m := range datas {
		ds = append(ds, ToGetMahasiswa(m))
	}
	return ds
}

func ToGetMahasiswaAlamats(m *models.MahasiswaAlamatsModels) *dto.MahasiswaAlamatRespDTO {
	return &dto.MahasiswaAlamatRespDTO{
		ID:   m.ID,
		Nama: m.Name,
		NIM:  m.Nim,
	}
}

//dari DB,jumlah model nya udah 2 - sesuai jumlah alamats di ID yg disimpan
func ToMhs(datas []*models.MahasiswaAlamatsModels) []*dto.MahasiswaAlamatRespDTO {
	var data []*dto.MahasiswaAlamatRespDTO

	for _, m := range datas {
		//fmt.Println("ini data ID ke ", i, " ", m.ID)
		data = append(data, ToGetMahasiswaAlamats(m))
	}

	fmt.Println("panjang data di Asembler : ", len(data))

	return data
}

func ToGetAlamats(m *models.MahasiswaAlamatsModels) *dto.AlamatRespDTO {
	return &dto.AlamatRespDTO{
		IDMahasiswas: m.IDMahasiswas,
		Jalan:        m.Jalan,
		NoRumah:      m.NoRumah,
	}
}

func ToAlamats1(datas []*models.MahasiswaAlamatsModels) []*dto.AlamatRespDTO {
	var data []*dto.AlamatRespDTO
	for _, m := range datas {
		data = append(data, ToGetAlamats(m))
	}
	return data
}

//untuk menampilkan alamat saja

func ToGetOnlyAlamats(data *models.AlamatsOnlyModel) *dto.AlamatReqDTO {
	return &dto.AlamatReqDTO{
		Jalan:   data.Jalan,
		NoRumah: data.NoRumah,
	}
}

func ToGetOnlyAlamatsAll(data []*models.AlamatsOnlyModel) []*dto.AlamatReqDTO {
	var dataDTO []*dto.AlamatReqDTO

	for _, valModel := range data {
		dataDTO = append(dataDTO, ToGetOnlyAlamats(valModel))
	}

	return dataDTO
}
