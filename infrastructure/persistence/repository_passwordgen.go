package persistence

import "github.com/marcos-dev88/go-password-generator/domain/entity"

type Repository interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	CreatePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(entity.PasswordGen) bool
}

type repository struct {
	mongodb MongoDB
}

func NewRepository(mongodb MongoDB) *repository {
	return &repository{mongodb: mongodb}
}

func (r *repository) GetPasswordGen(password string) (*entity.PasswordGen, error) {
	return nil, nil
}
func (r *repository) CreatePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error) {
	return nil, nil
}
func (r *repository) PasswordExists(entity.PasswordGen) bool {
	return false
}
