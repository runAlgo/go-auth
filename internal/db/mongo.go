package db

import (
	"context"
	"fmt"
	"time"

	"github.com/runAlgo/go-auth/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client // connects to the server.
	DB *mongo.Database // is the chosen database, from DB you get collections to work on.
}


func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {

	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOpt := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(connectCtx, clientOpt)
	if err != nil {
		return nil, fmt.Errorf("Mongo connection failed: %w", err)
	}

	database := client.Database(cfg.MongoDBName)

	return &Mongo{
		Client: client,
		DB: database,
	}, nil
}