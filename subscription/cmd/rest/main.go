package main

import (
	"fmt"
	"log"
	"tiny-letter-subscription/pkg/app"
	"tiny-letter-subscription/pkg/db"
	"tiny-letter-subscription/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := app.GetConfig()
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewHandler(db)

	r := gin.Default()
	r.POST("/subscribe-publisher", h.HandlerSubscribePublisherPlan)
	r.POST("/subscribe-subscriber", h.HandlerSubscribeSubscriberPlan)

	addr := fmt.Sprintf("%s:%s", config.App.Domain, config.App.Port)
	r.Run(addr)
}
