package database

import (
	"context"
	"os"

	"github.com/matheuszzy/crud-basico-go/src/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	MONGO_URI      = "MONGO_URI"
	MONGO_DATABASE = "MONGO_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	MONGO_URI := os.Getenv(MONGO_URI)
	MONGO_DATABASE := os.Getenv(MONGO_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Info("Error trying to connect to MongoDB", zap.String("Journey", "StartConnectionMongoDB"))
		return nil, err
	} else {
		logger.Info("Connected to MongoDB", zap.String("Journey", "StartConnectionMongoDB"))
	}

	return client.Database(MONGO_DATABASE), nil
}
