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

func (c *ConfigRepository) InsertInCollection(config *ConfigService.Config) error {
	collection := c.database.Collection(config.Service)
	_, err := collection.InsertOne(context.TODO(), config)
	return err
}

func (c *ConfigRepository) FindLastVersionInCollection(name string) (*ConfigService.Config, error) {
	collection := c.database.Collection(name)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	l := len(result)
	if l == 0 {
		return nil, nil
	} else {
		bs, err := bson.Marshal(result[l-1])
		if err != nil {
			return nil, err
		}
		conf := &ConfigService.Config{}
		err = bson.Unmarshal(bs, &conf)
		if err != nil {
			return nil, err
		}
		return conf, nil
	}
}

//func (c *ConfigRepository) FindAllInCollection(name string) ([]*ConfigService.Config, error) {
//	collection := c.database.Collection(name)
//}
