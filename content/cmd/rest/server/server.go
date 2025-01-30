package rest_server

import (
	"tiny-letter/content/pkg/db"
	"tiny-letter/content/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(db *db.Repository, addr string) {
	h := handlers.GetHandler(db)

	r := gin.Default()
	r.POST("/publications", h.HandleCreatePublication)
	r.POST("/posts", h.HandleCreatePost)

	r.Run(addr)
}
