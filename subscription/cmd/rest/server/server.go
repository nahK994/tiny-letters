package rest_server

import (
	"fmt"
	"log"
	"sync"
	"tiny-letter-subscription/pkg/app"
	"tiny-letter-subscription/pkg/db"
	"tiny-letter-subscription/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewHandler(db)

	r := gin.Default()
	r.POST("/subscribe-publisher", h.HandlerSubscribePublisher)
	r.POST("/subscribe-publication", h.HandlerSubscribePublication)
	r.POST("/change-subscriber-subscription-plan", h.HandlerChangeSubscriberSubscriptionPlan)
	r.POST("/change-publication-subscription-plan", h.HandlerChangePublisherSubscriptionPlan)

	addr := fmt.Sprintf("%s:%s", config.App.REST.Domain, config.App.REST.Port)
	r.Run(addr)
}
