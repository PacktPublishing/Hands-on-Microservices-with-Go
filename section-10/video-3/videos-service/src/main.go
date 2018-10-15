package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/videos-service/src/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/videos-service/src/repositories"
	"github.com/gorilla/mux"
)

func main() {
	repo := repositories.NewMariaDBVideosRepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		Repo: repo,
	}

	r := mux.NewRouter()

	r.HandleFunc("/bought-video", handler.InsertBoughtVideo).Methods("POST")
	r.HandleFunc("/bought-video", handler.DeleteBoughtVideo).Methods("DELETE")

	fmt.Println("Starting server on Port: 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
