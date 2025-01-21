package main

import (
	"fmt"
	"log"
	"sync"
	grpc_server "tiny-letter/auth/cmd/grpc/server"
	rest_server "tiny-letter/auth/cmd/rest/server"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
)

func main() {
	config := app.GetConfig()
	grpcAddr := fmt.Sprintf("%s:%d", config.App.GRPC.Domain, config.App.GRPC.Port)
	restAddr := fmt.Sprintf("%s:%d", config.App.REST.Domain, config.App.REST.Port)
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go grpc_server.Serve(&wg, db, grpcAddr)
	go rest_server.Serve(&wg, db, restAddr)
	wg.Wait()
}
