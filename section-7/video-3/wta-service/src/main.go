package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/repository"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/usecases"
	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewWTARepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		GetPlayerUsecase: usecases.GetPlayer{
			Repo: repo,
		},
		GetMatchUsecase: usecases.GetMatch{
			Repo: repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/player/{id}", handler.GetPlayer).Methods("GET")
	r.HandleFunc("/match/{id}", handler.GetMatch).Methods("GET")

	fmt.Println("Starting server on Port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
