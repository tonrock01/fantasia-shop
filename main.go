package main

import (
	"context"
	"log"
	"os"

	"github.com/tonrock01/fantasia-shop/config"
	"github.com/tonrock01/fantasia-shop/pkg/database"
	"github.com/tonrock01/fantasia-shop/server"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	// Start Server
	server.Start(ctx, &cfg, db)
}
