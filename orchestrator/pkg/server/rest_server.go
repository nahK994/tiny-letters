package server

import (
	"fmt"
	"tiny-letter/orchestrator/pkg/app"
	"tiny-letter/orchestrator/pkg/handlers"
	"tiny-letter/orchestrator/pkg/mq"

	"github.com/gin-gonic/gin"
)

func Serve(commConfig *app.CommConfig, producer *mq.Producer) {
	h := handlers.NewHandler(producer)

	r := gin.Default()
	r.Group("/registration")
	{
		r.POST("/subscriber", h.HandlerSubscriberRegistration)
		r.POST("/publisher", h.HandlerPublisherRegistration)
	}

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
		r.POST("/change-plan", h.HandleChangeSubscriberSubscription)
	}

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
