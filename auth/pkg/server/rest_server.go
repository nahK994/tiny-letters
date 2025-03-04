package server

import (
	"fmt"
	"sync"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
	"tiny-letter/auth/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func ServeREST(wg *sync.WaitGroup, db *db.Repository, commConfig *app.CommConfig) {
	defer wg.Done()
	h := handlers.GetREST_Handlers(db)

	r := gin.Default()
	r.POST("/login", h.Login)
	r.POST("/register", h.HandleUserRegistration)

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
