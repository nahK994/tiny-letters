package rest_server

import "sync"

func Serve(wg *sync.WaitGroup) {
	defer wg.Done()
}
