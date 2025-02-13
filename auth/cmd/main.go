package main

import (
	"log"
	"sync"
	grpc_server "tiny-letter/auth/cmd/grpc/server"
	rest_server "tiny-letter/auth/cmd/rest/server"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
)

func main() {
	config := app.GetConfig()
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go grpc_server.Serve(&wg, db, &config.GRPC)
	go rest_server.Serve(&wg, db, &config.REST)
	wg.Wait()
}
