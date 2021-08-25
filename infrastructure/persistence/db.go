package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	GetConn() (*mongo.Database, *mongo.Collection, context.Context, error)
}

type mongoDB struct {
	dbName string
	dbTableName string
	dbURI string
}

func NewMongoDB(dbName, dbTableName, dbURI string) *mongoDB{
	return &mongoDB{
		dbName:      dbName,
		dbTableName: dbTableName,
		dbURI:       dbURI,
	}
}

func (m *mongoDB) GetConn() (*mongo.Database, *mongo.Collection, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.dbURI))

	if err != nil {
		return nil, nil, nil, err
	}

	cntx := context.Background()

	if err := client.Connect(cntx); err != nil {
		return nil, nil, nil, err
	}

	db := client.Database(m.dbName)

	filterTableName := bson.M{"name": m.dbTableName}

	tableListResult, err := db.ListCollectionNames(cntx, filterTableName)

	if err != nil {
		return nil, nil, nil, err
	}

	if len(tableListResult) == 0 {
		if err = db.CreateCollection(cntx, m.dbTableName); err != nil {
			return nil, nil, nil, err
		}
	}

	return db, db.Collection(m.dbTableName), cntx, nil
}
