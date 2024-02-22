package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return client, nil
}
