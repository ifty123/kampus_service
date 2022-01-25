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

//update dosen
func (s *service) UpdateDosens(dataDosen *dto.DosensWithIDReqDTO) error {
	//assembler
	dtDosen := assembler.ToUpdateDosens(dataDosen)

	err := s.repo.UpdateDosens(dtDosen)
	if err != nil {
		log.Println("Failed to Repo Update Dosens")
		return err
	}
	return nil
}

//tampil dosen
func (s *service) TampilDosen() ([]*dto.DosenRespDTO, error) {
	dtdosen, err := s.repo.TampilDosens()
	if err != nil {
		log.Println("failed to tampil dosen")
		return nil, err
	}

	dtdosen1 := assembler.ToTampilDosen(dtdosen)
	for _, val := range dtdosen1 {
		fmt.Println(val)
	}
	return dtdosen1, nil

}

func (s *service) TampilDosenByID(dataID *dto.DosenIdDTO) (*dto.DosenRespDTO, error) {
	dtdosen, err := s.repo.TampilDosenByID(dataID.Id_dosens)
	var dataSave *dto.DosenRespDTO

	if err != nil {
		return nil, err
	}

	for _, data := range dtdosen {
		dataSave = &dto.DosenRespDTO{
			Nama:    data.Name,
			Nidn:    data.Nidn,
			IdDosen: data.Id,
		}
	}

	return dataSave, nil
}

func (s *service) TampilDosenAlamat() ([]*dto.DosenAndAlamatRespDTO, error) {
	dtDosen, err := s.repo.TampilDosenAlamat()
	if err != nil {
		log.Println("Failed to repo Tampil dosen", err.Error())
		return nil, err
	}

	dataDosen := assembler.ToTampilDosensMerge(dtDosen)
	dataAlamat := assembler.ToTampilDosenAlamatsMerge(dtDosen)

	for _, data1 := range dataDosen {
		for _, data2 := range dataAlamat {
			if data1.Id == data2.IdDosen {
				data1.Alamats = append(data1.Alamats, data2)
			}
		}
	}

	return dataDosen, nil
}

func (s *service) UpdateDosenAlamat(dataDosen *dto.DosenAndAlamatReqWithIDDTO) error {
	//bikin assembler untuk di masukan ke repo
	dtDosen := assembler.ToUpdateDosensWithAlamatDTO(dataDosen)
	dtAlamat := assembler.ToUpdateDosensAndAlamatsMerge(dataDosen.Alamats)

	//masuk ke repo
	err := s.repo.UpdateDosenAlamat(dtDosen, dtAlamat)
	if err != nil {
		return err
	}
	return nil
}
