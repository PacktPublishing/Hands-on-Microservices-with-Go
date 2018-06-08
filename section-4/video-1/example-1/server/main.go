package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	certPath := "server.pem"
	keyPath := "server.key"

	http.HandleFunc("/hello", myHandlerFunc)

	err := http.ListenAndServeTLS(":8443", certPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {

	//Write Status Code
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Hello world from HTTPS.")
}
