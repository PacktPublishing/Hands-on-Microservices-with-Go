package main

import (
	"context"
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

	limiter := rate.NewLimiter(50, 10)
	//Empty Context
	ctx := context.Background()

	wg := sync.WaitGroup{}
	wg.Add(300)
	for i := 0; i < 300; i++ {
		go func(index int) {
			limiter.Wait(ctx)
			defer wg.Done()
			doGet(client, index)
		}(i)
	}
	wg.Wait()
}

func doGet(client *http.Client, index int) {
	res, err := client.Get("http://localhost:9000")
	if err != nil || res.StatusCode == http.StatusInternalServerError {
		log.Println("Index: ", index, " - Request Error")
		return
	}
	if res.StatusCode == http.StatusTooManyRequests {
		log.Println("Index: ", index, " - Too Many Requests")
		return
	}

	log.Println("Index: ", index, " - Request Success")
}
