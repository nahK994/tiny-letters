package main

import (
	"fmt"
	"log"
	"tiny-letter-user/pkg/app"
	"tiny-letter-user/pkg/db"
	"tiny-letter-user/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := app.GetConfig()
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewHandler(db)

	r := gin.Default()
	r.POST("/login", h.Login)
	r.POST("/register-publisher", h.HandleRegisterPublisher)
	r.POST("/register-subscriber", h.HandleRegisterSubscriber)

	addr := fmt.Sprintf("%s:%s", config.App.Domain, config.App.Port)
	r.Run(addr)
}
