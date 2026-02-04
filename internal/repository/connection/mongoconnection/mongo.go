package mongoconnection

import (
	"context"
	"fmt"
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
