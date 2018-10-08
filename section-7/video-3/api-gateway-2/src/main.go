package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/repositories"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h := &handlers.Handler{
		AgentsRepo: repositories.RestAgentsRepository{},
		WTARepo:    repositories.RestWTARepository{},
	}

	r.HandleFunc("/agent/players/{id}", h.GetAgentPlayers)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	fmt.Println("Starting server on port: 8080")
	log.Fatal(srv.ListenAndServe())
}
