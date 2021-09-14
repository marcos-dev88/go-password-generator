package repository

import (
	"github.com/marcos-dev88/go-password-generator/domain/entity"
)

type Repository interface {
	GetLastTenPasswords() ([]*entity.PasswordGen, error)
	SavePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(password string) bool
}
