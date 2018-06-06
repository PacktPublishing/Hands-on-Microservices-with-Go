package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", myHandlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")

	source := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(source)
	n := r1.Intn(10)

	fmt.Fprintf(w, "%d", n)
}
