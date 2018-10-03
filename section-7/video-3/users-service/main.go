package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/users-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/users-service/repositories"
	"github.com/gorilla/mux"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewMySQLUserRepository(),
	}
	defer handler.Repo.Close()

	r := mux.NewRouter()
	r.HandleFunc("/user/by/username/{username}", handler.GetUserByUsernameHandler).Methods("GET")
	r.HandleFunc("/user/by/id/{userID}", handler.GetUserByIDHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
