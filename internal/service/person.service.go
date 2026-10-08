package service

import (
	"context"
	"log"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/repo"
)

type PersonService struct {
	pr *repo.PersonRepo
}

func NewPersonService(pr *repo.PersonRepo) *PersonService {
	return &PersonService{
		pr: pr,
	}
}

func (p *PersonService) GetAllPerson(ctx context.Context) ([]dto.Person, error) {
	repo, err := p.pr.GetAllPerson(ctx)

	data := make([]dto.Person, 0, len(repo))
	for _, v := range repo {
		data = append(data, dto.Person{
			Id:            v.Id,
			Nama:          v.Nama,
			Alamat:        v.Alamat,
			Jenis_Kelamin: v.Jenis_Kelamin,
			Email:         v.Email,
			No_Telp:       v.No_Telp,
		})
	}
	log.Println(data)
	return data, err
}
