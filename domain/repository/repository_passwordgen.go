package repository

import (
	"github.com/marcos-dev88/go-password-generator/domain/entity"
)

type Repository interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	SavePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(password string) bool
}
