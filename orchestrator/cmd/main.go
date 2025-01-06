package main

import (
	"sync"
	rest_server "tiny-letter/orchestrator/cmd/rest/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go rest_server.Serve(&wg)
	wg.Wait()
}
