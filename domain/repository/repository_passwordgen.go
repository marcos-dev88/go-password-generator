package repository

import (
	"github.com/marcos-dev88/go-password-generator/domain/entity"
)

type Repository interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	CreatePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(entity.PasswordGen) bool
}
