package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	circuit "github.com/rubyist/circuitbreaker"
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
	cb := circuit.NewRateBreaker(0.50, 100)

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		time.Sleep(time.Millisecond * 5)
		go func(index int) {
			defer wg.Done()
			doGet(client, cb, index)
		}(i)
	}
	wg.Wait()

	log.Println("Final CircuitBreaker Error Rate: ", cb.ErrorRate())
}

func doGet(client *http.Client, cb *circuit.Breaker, index int) {
	cbErr := cb.Call(func() error {
		res, err := client.Get("http://localhost:9000")
		if err != nil || res.StatusCode != http.StatusOK {
			log.Println("Index: ", index, " - Request Error")
			cb.Fail()
			return errors.New("Error on HTTP Request")
		}
		log.Println("Index: ", index, " - Request Success")
		cb.Success()
		return nil
	}, time.Second)

	if cb.Tripped() {
		log.Println("Index: ", index, " - Circuit Breaker Open")
	}

	if cbErr == circuit.ErrBreakerTimeout {
		log.Println("Index: ", index, " - Circuit Breaker Timed Out")
	}
	return
}
