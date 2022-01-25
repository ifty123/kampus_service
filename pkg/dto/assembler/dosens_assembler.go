package assembler

import (
	"github.com/ifty123/kampus_service/internal/models"
	"github.com/ifty123/kampus_service/pkg/dto"
)

func ToSaveDosensOnly(dataDTO *dto.DosensReqDTO) *models.DosensModels {
	return &models.DosensModels{
		Name: dataDTO.Nama,
		Nidn: dataDTO.Nidn,
	}
}

func ToSaveDosensAlamatOnly(dataDTO *dto.DosensAlamatReqDTO) *models.DosenAlamatsModels {
	return &models.DosenAlamatsModels{
		Jalan:    dataDTO.Jalan,
		No_rumah: dataDTO.NoRumah,
		IdDosen:  dataDTO.IdDosen,
	}
}

func ToSaveDosens(dataDTO *dto.DosenAndAlamatReqDTO) *models.DosensModels {
	return &models.DosensModels{
		Name: dataDTO.Nama,
		Nidn: dataDTO.Nidn,
	}
}

func ToSaveDosensAlamat(dataDTO *dto.AlamatReqDTO) *models.DosenAndAlamatsModel {
	return &models.DosenAndAlamatsModel{
		Jalan:    dataDTO.Jalan,
		No_rumah: dataDTO.NoRumah,
	}
}

func ToSaveDosenAndAlamat(d []dto.AlamatReqDTO) []*models.DosenAndAlamatsModel {
	var mds []*models.DosenAndAlamatsModel
	for _, m := range d {
		mds = append(mds, ToSaveDosensAlamat(&m))
	}
	return mds
}

func ToUpdateDosens(data *dto.DosensWithIDReqDTO) *models.DosensModels {
	return &models.DosensModels{
		Name: data.Nama,
		Nidn: data.Nidn,
		Id:   data.IdDosen,
	}
}

func ToUpdateDosensWithAlamatDTO(data *dto.DosenAndAlamatReqWithIDDTO) *models.DosensModels {
	return &models.DosensModels{
		Name: data.Nama,
		Nidn: data.Nidn,
		Id:   data.Id_dosens,
	}
}

func ToTampil(data *models.DosensModels) *dto.DosenRespDTO {
	return &dto.DosenRespDTO{
		IdDosen: data.Id,
		Nama:    data.Name,
		Nidn:    data.Nidn,
	}
}

func ToTampilDosen(data []*models.DosensModels) []*dto.DosenRespDTO {
	var dosen []*dto.DosenRespDTO
	for _, val := range data {
		dosen = append(dosen, ToTampil(val))
	}
	return dosen
}

func ToTampilDosens(data *models.DosenAndAlamatsModel) *dto.DosenAndAlamatRespDTO {
	return &dto.DosenAndAlamatRespDTO{
		Id:   data.Id,
		Nama: data.Name,
		Nidn: data.Nidn,
	}
}

func ToTampilDosensMerge(data []*models.DosenAndAlamatsModel) []*dto.DosenAndAlamatRespDTO {
	dataDTO := []*dto.DosenAndAlamatRespDTO{}
	for _, val := range data {
		dataDTO = append(dataDTO, ToTampilDosens(val))
	}
	return dataDTO
}

func ToTampilDosenAlamat(data *models.DosenAndAlamatsModel) *dto.DosensAlamatRespDTO {
	return &dto.DosensAlamatRespDTO{
		IdDosen: data.IdDosen,
		Jalan:   data.Jalan,
		NoRumah: data.No_rumah,
	}
}

func ToTampilDosenAlamatsMerge(dataDosen []*models.DosenAndAlamatsModel) []*dto.DosensAlamatRespDTO {
	var dataDTO []*dto.DosensAlamatRespDTO
	for _, val := range dataDosen {
		dataDTO = append(dataDTO, ToTampilDosenAlamat(val))
	}
	return dataDTO
}

func ToUpdateDosensAndAlamats(data *dto.DosensAlamatRespWithIdDTO) *models.DosenAndAlamatsModel {
	return &models.DosenAndAlamatsModel{
		Id:       data.Id,
		No_rumah: data.NoRumah,
		Jalan:    data.Jalan,
		IdDosen:  data.IdDosen,
	}
}

func ToUpdateDosensAndAlamatsMerge(data []*dto.DosensAlamatRespWithIdDTO) []*models.DosenAndAlamatsModel {
	dataModel := []*models.DosenAndAlamatsModel{}
	for _, val := range data {
		dataModel = append(dataModel, ToUpdateDosensAndAlamats(val))
	}
	return dataModel
}
