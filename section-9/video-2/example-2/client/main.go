package main

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
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
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * (time.Duration(rand.Intn(10) + 5)))
		go func(index int) {
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
