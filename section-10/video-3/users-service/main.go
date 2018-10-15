package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/users-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/users-service/repositories"
	"github.com/gorilla/mux"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewMariaDBUsersRepository(),
	}
	defer handler.Repo.Close()

	r := mux.NewRouter()
	r.HandleFunc("/user/ammount/update", handler.UpdateUserAccount).Methods("PATCH")
	r.HandleFunc("/user/ammount/rollback", handler.RollbackUpdateUserAccount).Methods("PATCH")

	log.Println("Starting server on Port: 8082.")
	log.Fatal(http.ListenAndServe(":8082", r))
}
