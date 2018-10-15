package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/agents-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/agents-service/repositories"
	"github.com/gorilla/mux"
)

func main() {
	handler := &handlers.Handlers{
		Repo: repositories.NewMariaDBAgentsRepository(),
	}
	defer handler.Repo.Close()

	r := mux.NewRouter()
	r.HandleFunc("/agent/ammount/update", handler.UpdateAgentAccount).Methods("PATCH")
	r.HandleFunc("/agent/ammount/rollback", handler.RollbackUpdateAgentAccount).Methods("PATCH")

	log.Println("Starting server on Port: 8083.")
	log.Fatal(http.ListenAndServe(":8083", r))
}
