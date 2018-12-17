package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

func main() {

	port := flag.Int("port", 0, "Port to use on this server.")
	flag.Parse()
	if *port == 0 {
		fmt.Println("You must specify a Port for the server to run.")
		os.Exit(0)
	}

	totalRequests := uint64(0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Using atomic to prevent Race Conditions
		atomic.AddUint64(&totalRequests, 1)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Recieved Request on Port: %d\n", *port)
		fmt.Fprintf(w, "Total Requests on this Port: %d\n", totalRequests)

	})

	portStr := strconv.Itoa(*port)

	fmt.Printf("Starting server on Port: %s\n", portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
