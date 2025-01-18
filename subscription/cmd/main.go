package main

import (
	"sync"
	"tiny-letter/subscription/cmd/grpc/listener/content"
	"tiny-letter/subscription/cmd/grpc/listener/coordinator"
	"tiny-letter/subscription/pkg/app"
)

func main() {
	config := app.GetConfig()
	var wg sync.WaitGroup
	go coordinator.Listen(&wg, &config.App.Coordinator)
	go content.Listen(&wg, &config.App.Content)
	wg.Wait()
}
