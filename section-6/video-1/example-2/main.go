package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	totalRequests := 0
	port := os.Getenv("APPLICATION_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		totalRequests++

		log.Printf("Recieved Request on uri: %s.", r.RequestURI)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Recieved Request on Port: %s\n", port)
		fmt.Fprintf(w, "Total Requests on this Port: %d\n", totalRequests)

	})

	fmt.Printf("Starting server on Port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
