package main

import (
	"fmt"
	"log"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	"tiny-letter/content/pkg/grpc/client"
	"tiny-letter/content/pkg/handlers"
	"tiny-letter/content/pkg/mq"

	"github.com/gin-gonic/gin"
)

func main() {
	config := app.GetConfig()
	client, err := client.ConnectSubscriptionClient(&config.GRPC)
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := db.Init(&config.DB)
	if err != nil {
		log.Fatal(err)
	}

	producer, err := mq.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}

	h := handlers.GetHandler(db, producer, client)

	r := gin.Default()
	r.POST("/publications", h.HandleCreatePublication)
	r.POST("/posts", h.HandleCreatePost)

	addr := fmt.Sprintf("%s:%d", config.REST.Domain, config.REST.Port)
	r.Run(addr)
}
