package rest_server

import (
	"fmt"
	"sync"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()

	config := app.GetConfig()
	h := handlers.NewHandler()

	r := gin.Default()
	r.Group("/publisher")
	{
		r.POST("/subscribe", h.HandlePublisherSubscription)
		r.POST("/unsubscribe", h.HandlePublisherUnsubscription)
		r.POST("/change-plan", h.HandleChangePublisherPlan)
	}

	r.Group("/subscriber")
	{
		r.POST("/join-publication", h.HandleJoinPublication)
		r.POST("/leave-publication", h.HandleLeavePublication)
		r.POST("/change-plan", h.HandleChangeSubscriberPlan)
	}

	addr := fmt.Sprintf("%s:%d", config.Domain, config.Port)
	r.Run(addr)
}
