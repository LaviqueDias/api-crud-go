package mongodb

import (
	"context"
	"os"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	
	if err != nil {
		return nil, err
	}
	
	if err := client.Ping(ctx, nil); err != nil{
		return nil, err
	}

	logger.Info("MongoDB connection established successfully", zap.String("journey", "mongodb_connection"))
	return client.Database(mongodb_database), nil


}

