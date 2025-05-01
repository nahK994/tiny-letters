package server

import (
	"fmt"
	"sync"
	"tiny-letter/subscription/pkg/app"
	"tiny-letter/subscription/pkg/db"
	"tiny-letter/subscription/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func ServeREST(wg *sync.WaitGroup, db *db.Repository, commConfig *app.CommConfig) {
	defer wg.Done()
	h := handlers.GetREST_Handler(db)

	r := gin.Default()
	r.POST("/join-publication", h.JoinPublication)
	r.POST("/leave-publication", h.LeavePublication)
	r.POST("/change-subscriber-subscription", h.ChangeSubscriberSubscription)
	r.POST("/change-publisher-subscription", h.ChangePublisherSubscription)

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
