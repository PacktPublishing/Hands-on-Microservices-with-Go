package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	//More control..
	client := &http.Client{}

	//Transport can be reused by multiple goroutines
	//Cache connections bettween multiple requests
	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,  //Time for a connection to complete
			KeepAlive: 120 * time.Second, //Keep alive period for an active network connection
		}).Dial,
		DisableKeepAlives:     false, //prevent reuse of connections
		MaxIdleConns:          100,   //Sets size of connection pool to 100 connections. Accross all hosts.
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second, //Maximum time an idle conncection will remain idle before closing itself
		ResponseHeaderTimeout: 60 * time.Second, //Maximum time to wait for response headers after request has been completely sent
		//Other options. Proxy, etc. We'll see some in next sections.
	}

	client.Transport = tr
	client.Timeout = 60 * time.Second //Includes connection time, any redirects and reading response body.

	//We can do client.Get(), Post(), etc..

	//or we can Create a new Request
	jsonBytesMsg := []byte(`{"hello-again":"world"}`)
	buf := bytes.NewBuffer(jsonBytesMsg)
	req, err := http.NewRequest("GET", "http://localhost:8080/forPostClient", buf)

	//All of the elements we saw on the previous video
	//are accesible Here
	req.Header.Set("Accept", "text/html")

	//Does the Request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	fmt.Println("Client3:: Recieved:")
	fmt.Println(string(body))
}
