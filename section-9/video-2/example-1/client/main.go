package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {

	rt := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{Transport: rt}

	fmt.Println("Starting at: ", time.Now())
	limiter := rate.NewLimiter(10, 5)
	//Empty Context
	ctx := context.Background()

	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(index int) {
			limiter.Wait(ctx)
			defer wg.Done()
			doGet(client, index)
			fmt.Println("Goroutine: ", index, " - Finished at: ", time.Now())
		}(i)
	}
	wg.Wait()
	fmt.Println("Finishing at: ", time.Now())
}

func doGet(client *http.Client, index int) {
	res, err := client.Get("http://localhost:9000")
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Index: ", index, " - Request Error")
	}
	//log.Println("Index: ", index, " - Request Success")
}
