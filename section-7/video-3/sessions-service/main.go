package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/sessions-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/sessions-service/repositories"
	"github.com/gorilla/mux"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewRedisSessionsRepository(),
	}

	r := mux.NewRouter()
	r.HandleFunc("/session/{key}", handler.GetSession).Methods("GET")
	r.HandleFunc("/session/{key}", handler.SetSession).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", r))
}
