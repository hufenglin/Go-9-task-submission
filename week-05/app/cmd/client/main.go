package main

import (
	"hystrix-demo/app/client"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ { //客户端单次发其并发请求100次
		wg.Add(1)
		go func() {
			client.Client()
			wg.Done()
		}()
	}
	wg.Wait()
}
