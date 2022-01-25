package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ifty123/kampus_service/pkg/dto"
	util "github.com/ifty123/kampus_service/pkg/utils"
)

type integService struct {
}

func NewService() IntegServices {
	return &integService{}
}

func (s *integService) GetRandomDadJokes(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error) {

	var response dto.GetDadJokesRandomRespDTO
	var url string

	if req.ID != "" {
		url = util.GetIntegURL("icanhazdadjoke", "byId")

		url = fmt.Sprintf(url, req.ID)

	} else {
		url = util.GetIntegURL("icanhazdadjoke", "random")
	}

	getReq, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to create http request icanhazdadjoke getRandom: %v", err)
	}

	//define header, karna mau ke API sebelah, dan required nya pake accept, maka header nya pake acept
	getReq.Header["Accept"] = []string{" application/json"}

	client := http.Client{
		Timeout: 15 * time.Second,
	}
	//bentuk respon masih buffer, belum marchall
	resp, err := client.Do(getReq)

	if err != nil {
		return nil, fmt.Errorf("failed to create http request icanhazdadjoke getRandom: %v", err)
	}
	log.Println("Success execute  : ", url)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode icanhazdadjoke getRandom response: %v", err)
	}

	return &response, nil
}
