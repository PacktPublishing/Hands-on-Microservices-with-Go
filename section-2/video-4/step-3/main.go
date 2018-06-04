package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/example", HeaderDecorator(http.HandlerFunc(myHandlerFunc)))
	http.Handle("/example2", HeaderDecorator(http.HandlerFunc(myHandlerFunc)))

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

func HeaderDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Passed-Through-Decorator", "YES")
		h.ServeHTTP(w, r)
	}) //returns an HTTP Handler, http.HandlerFunc( returns an http handler
}

func LogDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.RequestURI+": We have a connection from %s", r.RemoteAddr)
		h.ServeHTTP(w, r)
	}) //returns an HTTP Handler, http.HandlerFunc( returns an http handler
}
