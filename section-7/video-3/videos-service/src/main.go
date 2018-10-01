package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/usecases"
	"github.com/gorilla/mux"
)

func main() {
	repo := repositories.NewMariaDBVideosRepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		GetAllUserVideosUsecase: usecases.GetAllUserVideos{
			Repo: repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/videos/{id}", handler.GetAllUserVideos).Methods("GET")

	fmt.Println("Starting server on Port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
