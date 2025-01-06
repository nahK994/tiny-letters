package grpc_server

import "sync"

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
}
