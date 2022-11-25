package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}
