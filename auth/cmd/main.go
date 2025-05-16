package main

import (
	"log"
	"sync"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
	"tiny-letter/auth/pkg/server"
)

func main() {
	config := app.GetConfig()
	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go server.ServeGRPC(&wg, db, &config.GRPC)
	go server.ServeREST(&wg, db, &config.REST)
	wg.Wait()
}
