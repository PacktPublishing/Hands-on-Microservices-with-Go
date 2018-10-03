package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/service"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/transports/endpoints"
)

func main() {
	r := mux.NewRouter()

	repo := repositories.NewMariaDBManagersRepository()
	defer repo.Close()

	svc := service.ManagersServiceImpl{}
	svc.Repo = repo

	insertManagerPlayerEndpoint := endpoints.MakeInsertManagerPlayerEndpoint(svc)

	r.Methods("POST").Path("/manager-player").Handler(httptransport.NewServer(
		insertManagerPlayerEndpoint,
		endpoints.DecodeInsertManagerPlayerRequest,
		endpoints.EncodeInsertManagerPlayerResponse,
	))

	getManagerPlayerIDs := endpoints.MakeGetManagerPlayerIDsEndpoint(svc)

	r.Methods("GET").Path("/manager/players/{id}").Handler(httptransport.NewServer(
		getManagerPlayerIDs,
		endpoints.DecodeGetManagerPlayerIDsRequest,
		endpoints.EncodeGetManagerPlayerIDsResponse,
	))

	getManagerByIDEndpoint := endpoints.MakeGetManagerByIDEndpoint(svc)

	r.Methods("GET").Path("/manager/{id}").Handler(httptransport.NewServer(
		getManagerByIDEndpoint,
		endpoints.DecodeGetManagerByIDRequest,
		endpoints.EncodeGetManagerByIDResponse,
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
