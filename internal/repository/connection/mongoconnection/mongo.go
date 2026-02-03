package mongoconnection

import (
	"context"
	"fmt"
	"log"
	"usermanager/internal/configs"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoConnection(ctx context.Context) (*mongo.Client, error) {
	uri := configs.Env.MongoURI
	if uri == "" {
		return nil, fmt.Errorf("The URI was not found")
	}

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(client *mongo.Client, dbName, colName string) *mongo.Collection {
	if dbName == "" || colName == "" {
		log.Fatal("MongoDB database or collection name is empty")
	}

	return client.Database(dbName).Collection(colName)
}
