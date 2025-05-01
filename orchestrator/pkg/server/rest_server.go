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
		r.POST("/publisher", h.HandlerPublisherRegistration)
	}

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
