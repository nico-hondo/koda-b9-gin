package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/model"
)

type PersonRepo struct {
	db *pgxpool.Pool
}

func NewPersonRepo(db *pgxpool.Pool) *PersonRepo {
	return &PersonRepo{
		db: db,
	}
}

func (p *PersonRepo) GetAllPerson(ctx context.Context) ([]model.Person, error) {
	sql := "SELECT id, nama, alamat, jenis_kelamin, email, no_telp FROM person"
	rows, err := p.db.Query(ctx, sql)

	if err != nil {
		return nil, err
	}

	var persons []model.Person

	for rows.Next() {
		var person model.Person
		if err := rows.Scan(&person.Id, &person.Nama, &person.Alamat, &person.Jenis_Kelamin, &person.Email, &person.No_Telp); err != nil {
			log.Println(err)
			return nil, err
		}
		persons = append(persons, person)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	log.Println(persons)
	return persons, nil
}
