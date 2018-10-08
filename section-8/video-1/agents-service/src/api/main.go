package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/service"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/transports/endpoints"
)

func main() {
	r := mux.NewRouter()

	repo := repositories.NewMariaDBAgentsRepository()
	defer repo.Close()

	svc := service.AgentsServiceImpl{}
	svc.Repo = repo

	insertAgentPlayerEndpoint := endpoints.MakeInsertAgentPlayerEndpoint(svc)

	r.Methods("POST").Path("/agent-player").Handler(httptransport.NewServer(
		insertAgentPlayerEndpoint,
		endpoints.DecodeInsertAgentPlayerRequest,
		endpoints.EncodeInsertAgentPlayerResponse,
	))

	getAgentByIDEndpoint := endpoints.MakeGetAgentByIDEndpoint(svc)

	r.Methods("GET").Path("/agent/{id}").Handler(httptransport.NewServer(
		getAgentByIDEndpoint,
		endpoints.DecodeGetAgentByIDRequest,
		endpoints.EncodeGetAgentByIDResponse,
	))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on port: 8080")
	log.Fatal(srv.ListenAndServe())

}
