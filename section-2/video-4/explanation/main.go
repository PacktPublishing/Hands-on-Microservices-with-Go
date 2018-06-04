package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	h := &MyHandler{}
	http.Handle("/fromHandler", h)

	http.HandleFunc("/fromHandlerFunc", myHandlerFunc)

	anotherHandler := http.HandlerFunc(myHandlerFunc)
	http.Handle("/fromAnotherHandler", anotherHandler)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

type MyHandler struct{}

func (j *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	myHandlerFunc(w, r)
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Request to "+r.RequestURI+" was succesful")
}
