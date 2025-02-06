package main

import (
	"log"
	"sync"
	grpc_server "tiny-letter/subscription/cmd/grpc/server"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go grpc_server.Serve(&wg, db, &config.GRPC)
	wg.Wait()
}
