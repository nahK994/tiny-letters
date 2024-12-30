package rest_server

import (
	"fmt"
	"log"
	"sync"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/handlers"

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
	r.POST("/subscribe-publication", h.HandlerJoinPublication)
	r.POST("/change-subscriber-subscription-plan", h.HandlerChangePublicationSubscriptionPlan)
	r.POST("/change-publisher-subscription-plan", h.HandlerChangePublisherSubscriptionPlan)
	r.POST("/unsubscription-publication", h.HandleLeavePublication)
	r.POST("/unsubscribe-publisher-plan", h.HandleUnsubscribePublisher)

	addr := fmt.Sprintf("%s:%s", config.App.REST.Domain, config.App.REST.Port)
	r.Run(addr)
}
