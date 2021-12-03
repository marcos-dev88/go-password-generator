package persistence

import (
	"context"
	"encoding/json"
	"github.com/marcos-dev88/go-password-generator/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repository struct {
	mongodb MongoDB
}

func NewRepository(mongodb MongoDB) *repository {
	return &repository{mongodb: mongodb}
}

func (r *repository) GetLastTenPasswords() ([]*entity.PasswordGen, error){

	_, table, cntx, err := r.mongodb.GetConn()

	if err != nil {
		return nil, err
	}

	pipe := []bson.M{
		{
			"$limit": 10,
		},
	}

	cursor, err := table.Aggregate(cntx, pipe)

	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, cntx)

	var passwords []*entity.PasswordGen

	for cursor.Next(cntx) {
		passEntity := entity.NewPasswordGen("", "", 0, false, false, false)
		passBson := bson.M{}

		if err = cursor.Decode(&passBson); err != nil {
			return nil, err
		}

		passJson, _ := json.Marshal(passBson)
		if err := json.Unmarshal(passJson, &passEntity); err != nil {
			return nil, err
		}

		passwords = append(passwords, passEntity)
	}

	return passwords, nil
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
		{Key: "has_letter", Value: password.HasLetter},
		{Key: "has_number", Value: password.HasNumber},
		{Key: "has_special_char", Value: password.HasSpecialChar},
	})

	if err != nil {
		return nil, err
	}

	return password, nil
}

func (r *repository) PasswordExists(password string) (bool, error) {
	_, table, cntx, err := r.mongodb.GetConn()

	if err != nil {
		return false, err
	}

	cursor, err := table.Find(cntx, bson.M{"password": password})

	if err != nil {
		return false, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error to close context -> %v", err)
		}
	}(cursor, cntx)

	if cursor.Next(cntx) {
		return true, nil
	}

	return false, nil
}
