package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/repositories"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h := &handlers.Handler{
		SessionsRepo: &repositories.RestSessionsRepository{},
		UsersRepo:    &repositories.RestUsersRepository{},
	}

	r.HandleFunc("/authorize", h.Authorize).Methods("POST")

	r.HandleFunc("/restricted/resource-1", handlers.VerifyJWT(h.AddSessionData(h.Restricted1)))
	r.HandleFunc("/restricted/resource-2", handlers.VerifyJWT(h.AddSessionData(h.Restricted2)))

	certPath := "server.pem"
	keyPath := "server.key"

	err := http.ListenAndServeTLS(":8443", certPath, keyPath, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
