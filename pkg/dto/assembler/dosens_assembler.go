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
