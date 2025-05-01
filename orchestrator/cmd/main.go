package main

import (
	"fmt"
	"log"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/grpc/client"
	"tiny-letter/orchestrator/pkg/handlers"
	"tiny-letter/orchestrator/pkg/mq"

	"github.com/gin-gonic/gin"
)

func main() {
	config := app.GetConfig()

	if err := client.ConnectGRPC(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}

	producer, err := mq.NewProducer(&config.MQ)
	if err != nil {
		log.Fatalf("Failed to connect to MQ: %v", err)
	}
	h := handlers.NewHandler(producer)

	r := gin.Default()
	r.Group("/registration")
	{
		r.POST("/publisher", h.HandlerPublisherRegistration)
	}

	addr := fmt.Sprintf("%s:%d", config.REST.Domain, config.REST.Port)
	r.Run(addr)
}
