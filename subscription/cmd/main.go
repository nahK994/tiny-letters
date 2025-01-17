package main

import (
	"sync"
	"tiny-letter/subscription/cmd/grpc/listener/email"
	"tiny-letter/subscription/cmd/grpc/listener/orchestrator"
	"tiny-letter/subscription/pkg/app"
)

func main() {
	config := app.GetConfig()
	var wg sync.WaitGroup
	go orchestrator.Listen(&wg, &config.App.Orchestrator)
	go email.Listen(&wg, &config.App.Email)
	wg.Wait()
}
