package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
)

//Run with go run server/main.go -failProb 50
var failProb uint16

func main() {
	fp := flag.Int("failProb", 0, "Probability of Failure per 100 Requests.")
	flag.Parse()
	failProb = uint16(*fp) - 1

	http.HandleFunc("/", handler)
	log.Println("failProb:", (failProb + 1))
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
