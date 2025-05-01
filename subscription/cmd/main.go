package main

import (
	"log"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(2)
	go server.ServeGRPC(&wg, db, &config.GRPC)
	go server.ServeREST(&wg, db, &config.REST)
	wg.Wait()
}
