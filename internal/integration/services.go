package integration

import "github.com/ifty123/kampus_service/pkg/dto"

type IntegServices interface {
	GetRandomDadJokes(req *dto.GetDadJokesInternalReqDTO) (*dto.GetDadJokesRandomRespDTO, error)
}
