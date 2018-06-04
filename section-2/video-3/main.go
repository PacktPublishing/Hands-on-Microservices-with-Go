package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/example", myHandlerFunc)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           mux,
		ReadTimeout:       30 * time.Second, //Reading all of the request including body
		ReadHeaderTimeout: 15 * time.Second, //For Request headers
		WriteTimeout:      30 * time.Second, //Timeout of writing the response
		IdleTimeout:       10 * time.Second, //Timeout for Idle connections, when keepalive is true
		//MaxHeaderBytes: Size of Request Headers accepted.
		ErrorLog: log.New(os.Stdout, "Server: ", log.Lshortfile), //Logger for Server errors
		//Some extra we will se later. TLS.
	}

	//Function that can be invoked when state changes on the connections
	server.ConnState = func(c net.Conn, cs http.ConnState) {
		log.Println("Connection from Address: " + c.RemoteAddr().String() + " - State: " + cs.String())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		log.Println("Server Starting")
		log.Fatal(server.ListenAndServe())
	}()
	<-c
	// Shut down gracefully, but wait no longer than 30 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("Closing down")
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-1; Request was succesful")
}
