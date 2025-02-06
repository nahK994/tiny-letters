package rest_server

import (
	"fmt"
	"tiny-letter/content/pkg/app"
	"tiny-letter/content/pkg/db"
	"tiny-letter/content/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(db *db.Repository, commConfig *app.CommConfig) {
	h := handlers.GetHandler(db)

	r := gin.Default()
	r.POST("/publications", h.HandleCreatePublication)
	r.POST("/posts", h.HandleCreatePost)

	addr := fmt.Sprintf("%s:%d", commConfig.Domain, commConfig.Port)
	r.Run(addr)
}
