package rest_server

import (
	"fmt"
	"log"
	"sync"
	"tiny-letter/auth/pkg/app"
	"tiny-letter/auth/pkg/db"
	"tiny-letter/auth/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
	config := app.GetConfig()
	db, err := db.Init(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewHandler(db)

	r := gin.Default()
	r.POST("/login", h.Login)
	r.POST("/register-subscriber", h.HandleSubscriberRegistration)
	r.POST("/register-publisher", h.HandlePublisherRegistration)

	addr := fmt.Sprintf("%s:%s", config.App.REST.Domain, config.App.REST.Port)
	r.Run(addr)
}
