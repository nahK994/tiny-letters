package rest_server

import (
	handler "tiny-letter/coordinator/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(addr string) {
	h := handler.NewHandler()

	r := gin.Default()
	r.Group("/publisher")
	{
		r.POST("/subscribe", h.HandleConfirmPublisherSubscription)
		r.POST("/unsubscribe", h.HandleRevokePublisherSubscription)
		r.POST("/change-plan", h.HandleChangePublisherSubscription)
	}

	r.Group("/subscriber")
	{
		r.POST("/join-publication", h.HandleJoinPublication)
		r.POST("/leave-publication", h.HandleLeavePublication)
		r.POST("/change-plan", h.ChangeSubscriberSubscription)
	}

	r.Run(addr)
}
