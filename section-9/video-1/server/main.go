package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
)

//Run with go run server/main.go -failProb 89
var failProb uint16

func main() {
	fp := flag.Int("failProb", 0, "Failures per 100 Requests.")
	flag.Parse()
	failProb = uint16(*fp)

	http.HandleFunc("/", handler)
	log.Println("failProb:", failProb)
	log.Println("Starting Server on Port 9000.")
	http.ListenAndServe("localhost:9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fail := uint16(rand.Intn(100))
	if fail < failProb {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
