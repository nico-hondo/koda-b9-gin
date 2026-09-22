package service

import (
	"errors"

	"github.com/nico-hondo/internal/dto"
)

type PingService struct{}

func NewPingService() *PingService {
	return &PingService{}
}

func (p *PingService) EmptyValidation(body dto.Body) error {
	if body.Name == "" && body.Age == 0 {
		return errors.New("body kosong")
	}
	return nil
}
