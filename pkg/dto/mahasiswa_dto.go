package dto

import (
	"fmt"

	"github.com/ifty123/kampus_service/pkg/common/validator"
)

type MahasiswaAlamatReqDTO struct {
	Nama    string         `json:"nama" valid:"required" validname:"nama"`
	Nim     string         `json:"nim" valid:"required" validname:"nim"`
	Alamats []AlamatReqDTO `json:"alamat" valid:"required" `
}

//input atau output perlu validasi
func (dto *MahasiswaAlamatReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type MahasiswaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	Nim  string `json:"nim" valid:"required" validname:"nim"`
}

func (dto *MahasiswaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type AlamatReqDTO struct {
	Jalan   string `json:"jalan"`
	NoRumah string `json:"no_rumah"`
}

type UpadeMahasiswaNamaReqDTO struct {
	Nama string `json:"nama" valid:"required" validname:"nama"`
	ID   int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

type AlamatRespDTO struct {
	Jalan        string `json:"jalan"`
	NoRumah      string `json:"no_rumah"`
	IDMahasiswas int64  `json:"id_mahasiswas, omitempty"`
}

func (dto *UpadeMahasiswaNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type MahasiswaRespDTO struct {
	ID   int64  `json: "id"`
	Nama string `json:"nama"`
	NIM  string `json: "nim"`
}

type MahasiswaAlamatRespDTO struct {
	ID      int64            `json: "id"`
	Nama    string           `json:"nama"`
	NIM     string           `json: "nim"`
	Alamats []*AlamatRespDTO `json:"alamat"`
}

type AlamatReqWithIDDTO struct {
	Jalan        string `json:"jalan" valid:"required" validname:"jalan"`
	NoRumah      string `json:"no_rumah" valid:"required" validname:"no_rumah"`
	IDMahasiswas int64  `json:"id_mahasiswas" valid:"required,integer,non_zero" validname:"id_mahasiswas"`
}

func (dto *AlamatReqWithIDDTO) Validate() error {
	v := validator.NewValidate(dto)
	//fmt.Println("udah di validasi")

	return v.Validate()
}

//untuk update jalan
type UpdateAlamatMahasiswaReq struct {
	ID      int64  `json:"id" valid:"required,integer,non_zero" validname:"id"`
	Jalan   string `json:"jalan" valid:"required" validname:"jalan"`
	NoRumah string `json:"no_rumah" valid:"required" validname:"no_rumah"`
}

func (dto *UpdateAlamatMahasiswaReq) Validate() error {
	v := validator.NewValidate(dto)
	fmt.Println("udah di validasi")

	return v.Validate()
}

type GetMahasiswaByIDReqDTO struct {
	ID int64 `json:"id" valid:"required,integer,non_zero" validname:"id"`
}

func (dto *GetMahasiswaByIDReqDTO) Validate() error {
	v := validator.NewValidate(dto)
	fmt.Println("udah di validasi")

	return v.Validate()
}

//untuk ambil nama maha & validasi nill / not
type GetMahasiswaByNamaReqDTO struct {
	Name string `json:"nama" valid:"required" validname:"nama"`
}

func (dto *GetMahasiswaByNamaReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}

type GetNamaIfNotNull struct {
	Nama string `json:"nama"`
}
