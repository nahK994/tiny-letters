package rest_server

import (
	"fmt"
	"tiny-letter/coordinator/pkg/app"
	handler "tiny-letter/coordinator/pkg/handlers"
	mq_producer "tiny-letter/coordinator/pkg/mq"

	"github.com/gin-gonic/gin"
)

func Serve(commConfig *app.CommConfig, producer *mq_producer.Producer) {
	h := handler.NewHandler(producer)

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
		r.POST("/change-plan", h.HandleChangeSubscriberSubscription)
	}

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
