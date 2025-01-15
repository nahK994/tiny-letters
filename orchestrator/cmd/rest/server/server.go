package rest_server

import (
	handler "tiny-letter/orchestrator/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(addr string) {
	h := handler.NewHandler()

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

	r.Run(addr)
}
