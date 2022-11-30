package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func GetMongoClient(uri string) (*mongo.Client, error) {
	log.Printf("Connecting to %s\n", uri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	log.Print("Pinging mongo")
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}
