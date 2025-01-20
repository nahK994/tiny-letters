package main

import (
	"sync"
	"tiny-letter/subscription/cmd/grpc/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go server.Serve(&wg)
	wg.Wait()
}
