package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/usecases"
	"github.com/gorilla/mux"
)

func main() {
	cacheRepo := repositories.NewRedisUsersRepository()
	repo := repositories.NewMySQLUsersRepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		GetUserUsecase: &usecases.GetUserImpl{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
		UpdateUserUsecase: &usecases.UpdateUserImpl{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/user/{username}", handler.GetUserByUsername).Methods("GET")

	fmt.Println("Starting server on Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
