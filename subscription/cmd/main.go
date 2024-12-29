package main

import (
	"sync"
	grpc_server "tiny-letter/subscription/cmd/grpc/server"
	rest_server "tiny-letter/subscription/cmd/rest/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go grpc_server.Serve(&wg)
	go rest_server.Serve(&wg)
	wg.Wait()
}
