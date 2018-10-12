package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/panic", panic)

	log.Println("Starting Server on Port 9000.")
	http.ListenAndServe("localhost:9000", nil)
}

//Pong will still be working even after panic has been called
func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

//Simulate a Panic.
//Panic does not cross boundaries.
//It won't bring down the server.
func panic(w http.ResponseWriter, r *http.Request) {
	panic("Forced Panic")
	w.WriteHeader(http.StatusOK)
}
