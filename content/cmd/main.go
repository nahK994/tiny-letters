package main

import (
	"log"
	grpc_client "tiny-letter/content/cmd/grpc/client"
	rest_server "tiny-letter/content/cmd/rest/server"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
)

func main() {
	config := app.GetConfig()

	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpc_client.IsGRPC_ClientAvailable(&config.GRPC); err != nil {
		log.Fatal(err)
	}

	rest_server.Serve(db, &config.REST)
}
