package main

import (
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

	svc := service.ManagersServiceImpl{}
	svc.Repo = repo

	insertManagerPlayerEndpoint := endpoints.MakeInsertManagerPlayerEndpoint(svc)

	r.Methods("POST").Path("/manager-player/").Handler(httptransport.NewServer(
		insertManagerPlayerEndpoint,
		endpoints.DecodeInsertManagerPlayerRequest,
		endpoints.EncodeResponse,
	))

	getManagerByIDEndpoint := endpoints.MakeInsertManagerPlayerEndpoint(svc)

	r.Methods("GET").Path("/manager/{id}").Handler(httptransport.NewServer(
		getManagerByIDEndpoint,
		endpoints.DecodeGetManagerByIDRequest,
		endpoints.EncodeResponse,
	))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	repo.Close()
}
