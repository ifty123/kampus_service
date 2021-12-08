package services

import (
	"fmt"
	"log"

	"github.com/ifty123/kampus_service/pkg/dto"
	"github.com/ifty123/kampus_service/pkg/dto/assembler"
)

//save dosen
func (s *service) SaveDosens(req *dto.DosensReqDTO) error {
	dataDosenAssembler := assembler.ToSaveDosensOnly(req)

	//ke repo
	err := s.repo.SaveDosens(dataDosenAssembler)

	if err != nil {
		return err
	}
	return nil
}

//save dosen
func (s *service) SaveAlamatsDosen(req *dto.DosensAlamatReqDTO) error {

	fmt.Println("ini isi DTO service :", req)

	dataDosenAlAssemb := assembler.ToSaveDosensAlamatOnly(req)

	//ke repo
	err := s.repo.SaveAlamatsDosen(dataDosenAlAssemb)

	if err != nil {
		return err
	}
	return nil
}

//save dosen and alamat
func (s *service) SaveDosenAndAlamat(data *dto.DosenAndAlamatReqDTO) error {
	dtAlamat := assembler.ToSaveDosenAndAlamat(data.Alamats)
	dataDosen := assembler.ToSaveDosens(data)

	err := s.repo.SaveDosensAndAlamats(dataDosen, dtAlamat)

	if err != nil {
		log.Println("failed to Repo Save dosen Alamat")
		return err
	}
	return nil
}
