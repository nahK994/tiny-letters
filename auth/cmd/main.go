package main

import (
	"log"
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

	server.ServeGRPC(db, &config.GRPC)
}
