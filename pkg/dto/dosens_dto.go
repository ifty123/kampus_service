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

type DosensWithIDReqDTO struct {
	Nama    string `json:"nama" valid:"required" validname:"nama"`
	Nidn    string `json:"nidn" valid:"required" validname:"nidn"`
	IdDosen int64  `json:"id_dosens" valid:"required,integer,non_zero" validname:"id_dosens"`
}

//input atau output perlu validasi
func (dto *DosensWithIDReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type DosensAlamatReqDTO struct {
	Jalan   string `json:"jalan" valid:"required" validname:"jalan"`
	NoRumah string `json:"no_rumah" valid:"required" validname:"no_rumah"`
	IdDosen int64  `json:"id_dosens" valid:"required,integer,non_zero" validname:"id_dosens"`
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

type DosenRespDTO struct {
	IdDosen int64  `json:"Id_dosens"`
	Nama    string `json:"nama"`
	Nidn    string `json:"nidn"`
}

type DosenIdDTO struct {
	Id_dosens int64 `json:"id_dosens" valid:"required,integer,non_zero" validname:"id_dosens"`
}

func (dto *DosenIdDTO) Validate() error {
	//ngisi validator nya
	v := validator.NewValidate((dto))

	//ini fungsi validasi
	return v.Validate()
}

type DosenAndAlamatRespDTO struct {
	Id      int64                  `json:"id"`
	Nama    string                 `json:"nama"`
	Nidn    string                 `json:"nidn"`
	Alamats []*DosensAlamatRespDTO `json:"alamat"`
}

type DosensAlamatRespDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
	IdDosen int64  `json:"id_dosens, omitempty"`
}

type DosenAndAlamatReqWithIDDTO struct {
	Id_dosens int64                        `json:"id_dosens" valid:"required,integer,non_zero" validname:"id_dosens"`
	Nama      string                       `json:"nama" valid:"required" validname:"nama"`
	Nidn      string                       `json:"nidn" valid:"required" validname:"nidn"`
	Alamats   []*DosensAlamatRespWithIdDTO `json:"alamat" valid:"required" `
}

func (dto *DosenAndAlamatReqWithIDDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type DosensAlamatRespWithIdDTO struct {
	Id      int64  `json:"id"`
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
	IdDosen int64  `json:"id_dosens, omitempty"`
}
