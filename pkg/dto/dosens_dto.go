package dto

import (
	"github.com/ifty123/kampus_service/pkg/common/validator"
)

type DosensReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	Nidn string `json:"nidn" valid:"required" validname:"nidn"`
}

//input atau output perlu validasi
func (dto *DosensReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type DosensAlamatReqDTO struct {
	Jalan   string `json:"jalan" valid:"required" validname:"jalan"`
	NoRumah string `json:"no_rumah" valid:"required" validname:"no_rumah"`
	IdDosen int64  `json:"id_dosen" valid:"required,integer,non_zero" validname:"id_dosens"`
}

//input atau output perlu validasi
func (dto *DosensAlamatReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type DosenAndAlamatReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nidn    string         `json:"nidn" valid:"required" validname:"nidn"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

func (dto *DosenAndAlamatReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
