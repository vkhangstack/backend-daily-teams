package repository

import (
	"context"
	"github.com/vkhangstack/dlt/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		logger.Log.Errorf("Error connecting to Mongo: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Log.Errorf("Error pinging Mongo: %v", err)
	}

	logger.Log.Infoln("")
	return client, nil
}
