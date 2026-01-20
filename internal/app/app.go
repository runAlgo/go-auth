package app

import (
	"context"
	"fmt"
	"time"

	"github.com/runAlgo/go-auth/internal/config"
	"github.com/runAlgo/go-auth/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	// Load env first
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	// Do the DB connection second
	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}
	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("mongo disconnect failed: %w", err)
	}

	return nil
}
