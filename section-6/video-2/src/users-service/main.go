package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/usecases"
	"github.com/gorilla/mux"
)

func main() {
	cacheRepo := repositories.NewRedisUsersRepository()
	repo := repositories.NewMySQLUsersRepository()
	defer repo.Close()

	handler := &handlers.Handlers{
		GetUserUsecase: &usecases.GetUserUsecase{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
		UpdateUserUsecase: &usecases.UpdateUserUsecase{
			CacheRepo: cacheRepo,
			Repo:      repo,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/user/{username}", handler.GetUserByUsernameHandler).Methods("GET")

	log.Printf("Starting server on port 8000.")
	log.Fatal(http.ListenAndServe(":8000", r))
}
