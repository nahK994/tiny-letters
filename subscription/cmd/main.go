package main

import (
	"log"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/server"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server.Serve(db, &config.GRPC)
}
