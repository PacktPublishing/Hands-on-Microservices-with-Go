package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/example", logDecorator(myHandlerFunc))
	http.HandleFunc("/example2", logDecorator(myHandlerFunc2))

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-1; Request was succesful")
}

func myHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-2; Request was succesful")
}

func logDecorator(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.RequestURI+": We have a connection from %s", r.RemoteAddr)
		next(w, r)
	}
}
