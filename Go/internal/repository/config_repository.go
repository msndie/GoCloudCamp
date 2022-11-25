package repository

import (
	ConfigService "cmd/proto"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigRepository struct {
	database *mongo.Database
}

func NewConfigRepository(client *mongo.Client) *ConfigRepository {
	return &ConfigRepository{client.Database("GoCloud")}
}

func (c *ConfigRepository) CollectionExists(name string) (bool, error) {
	names, err := c.database.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return false, err
	}
	for _, s := range names {
		if s == name {
			return true, nil
		}
	}
	return false, nil
}

func (c *ConfigRepository) InsertInCollection(config ConfigService.Config) (bool, error) {
	collection := c.database.Collection(config.Service)
	//one, err := collection.InsertOne()
	if err != nil {
		return false, err
	}
}
