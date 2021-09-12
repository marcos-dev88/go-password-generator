package persistence

import (
	"context"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository interface {
	GetPasswordGen(password string) (*entity.PasswordGen, error)
	SavePasswordGen(*entity.PasswordGen) (*entity.PasswordGen, error)
	PasswordExists(password string) bool
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

func (r *repository) PasswordExists(password string) bool {
	_, table, cntx, err := r.mongodb.GetConn()

	if err != nil {
		panic(err)
	}

	cursor, err := table.Find(cntx, bson.M{"password": password})

	if err != nil {
		panic(err)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error to close context -> %v", err)
		}
	}(cursor, cntx)

	if cursor.Next(cntx) {
		return true
	}

	return false
}
