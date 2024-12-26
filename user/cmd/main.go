package main

import (
	"sync"
	rest_server "tiny-letter-user/cmd/rest/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go rest_server.Serve(&wg)
	wg.Wait()
}
