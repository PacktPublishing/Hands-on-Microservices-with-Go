package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//Request 1
	fmt.Println("Request 1")
	jsonBytesMsg := []byte(`{"hello":"world"}`)
	buf := bytes.NewBuffer(jsonBytesMsg)
	resp, err := http.Post("http://localhost:8080/forPostClient", "application/json", buf)
	if err != nil {
		log.Fatal(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	defer resp.Body.Close()

	fmt.Println("Client2:: Received:")
	fmt.Println(string(body))

	//Request 2
	fmt.Println("Request 2")
	resp, err = http.PostForm("http://localhost:8080/forPostClient", url.Values{"test": {"OK"}, "anotherTest": {"AlsoOK"}})
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	defer resp.Body.Close()

	fmt.Println("Client2:: Received:")
	fmt.Println(string(body))

}
