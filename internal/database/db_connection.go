package database

import (
	"context"

	"github.com/Cr4z1k/MEDODS-test-task/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() (*mongo.Collection, error) {
	connectionConf, err := config.GetConnection()
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionConf.URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	collection := client.Database(connectionConf.DbName).Collection(connectionConf.CollectionName)

	return collection, nil
}
