package main

import (
	"fmt"
	"log"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/grpc/client"
	"tiny-letter/orchestrator/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := app.GetConfig()

	if err := client.ConnectGRPC(&config.GRPC); err != nil {
		log.Fatal(err.Error())
	}

	h := handlers.NewHandler()

	r := gin.Default()
	r.Group("/registration")
	{
		r.POST("/publisher", h.HandlerPublisherRegistration)
	}

	addr := fmt.Sprintf("%s:%d", config.REST.Domain, config.REST.Port)
	r.Run(addr)
}
