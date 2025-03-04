package main

import (
	"log"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	grpc_server "tiny-letter/subscription/pkg/grpc/server"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	grpc_server.Serve(db, &config.GRPC)
}
