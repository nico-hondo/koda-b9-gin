package service

import (
	"errors"
	"fmt"

	"github.com/nico-hondo/internal/dto"
)

type UserService struct{}

var Users = []dto.User{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) EmptyValidationUser(data dto.User) error {
	if data.Email == "" && data.Password == "" {
		return errors.New("body kosong")
	}
	return nil
}

func (u *UserService) AvailableUser(data dto.User) error {
	for _, val := range dto.Users {
		if data.Email == val.Email {
			return fmt.Errorf("email %s sudah tersedia", data.Email)
		}
	}

	return nil
}

func (u *UserService) Authentication(user dto.LoginInput) error {
	for _, val := range Users {
		if user.Email != val.Email || user.Password != val.Password {
			return errors.New("email atau password salah!")
		}
	}
	return nil
}
