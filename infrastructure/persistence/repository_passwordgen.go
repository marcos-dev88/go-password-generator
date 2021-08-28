package persistence

import (
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	SavePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
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
func (r *repository) SavePasswordGen(password *entity.PasswordGen) (*entity.PasswordGen, error) {
	_, table, cntx, err := r.mongodb.GetConn()

	if err != nil {
		return nil, err
	}

	_, err = table.InsertOne(cntx, bson.D{
		{Key: "uuid", Value: password.Uuid},
		{Key: "length", Value: password.Length},
		{Key: "password", Value: password.Password},
	})

	if err != nil {
		return nil, err
	}

	return password, nil
}
func (r *repository) PasswordExists(entity.PasswordGen) bool {
	return false
}
