package database

import (
	"context"
	"log"
	"time"

	"github.com/tonrock01/fantasia-shop/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("Error: Connect to database error: %s", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error: Pinging to database error: %s", err.Error())
	}

	return client
}
