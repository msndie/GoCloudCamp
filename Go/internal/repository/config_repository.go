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

func NewConfigRepository(database *mongo.Database) *ConfigRepository {
	return &ConfigRepository{database: database}
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

func (c *ConfigRepository) InsertInCollection(config *ConfigService.Config) (bool, error) {
	collection := c.database.Collection(config.Service)
	if _, err := collection.InsertOne(context.TODO(), config); err != nil {
		return false, err
	}
	return true, nil
}

func (c *ConfigRepository) FindLastVersionInCollection(name string) (*ConfigService.Config, error) {
	collection := c.database.Collection(name)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []*ConfigService.Config
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	l := len(result)
	if l == 0 {
		return nil, nil
	} else {
		return result[l-1], nil
	}
}

func (c *ConfigRepository) FindAllInCollection(name string) ([]*ConfigService.Config, error) {
	collection := c.database.Collection(name)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []*ConfigService.Config
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ConfigRepository) FindAllLastEntities() ([]*ConfigService.Config, error) {
	names, err := c.database.ListCollectionNames(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return nil, nil
	}
	configs := make([]*ConfigService.Config, len(names))
	for i, name := range names {
		if configs[i], err = c.FindLastVersionInCollection(name); err != nil {
			return nil, err
		}
	}
	return configs, nil
}

func (c *ConfigRepository) DeleteCollection(name string) error {
	err := c.database.Collection(name).Drop(context.TODO())
	return err
}
