package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/8treenet/freedom/infra/requests"
)

func main() {
	round := 30
	for i := 0; i < round; i++ {
		press()
	}
}

func press() {
	requestCount := 10000
	group := sync.WaitGroup{}
	begin := time.Now()

	for i := 0; i < requestCount; i++ {
		group.Add(1)
		go func() {
			result, _ := requests.NewH2CRequest("http://127.0.0.1:8080/ping").Get().ToString()
			if result != "pong" {
				panic("ping error")
			}
			group.Done()
		}()
	}

	group.Wait()
	ms := time.Now().Sub(begin).Milliseconds()
	fmt.Printf("Number of requests: %d, total time: %d ms. \n", requestCount, ms)
}
