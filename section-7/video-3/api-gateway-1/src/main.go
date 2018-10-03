package main

import (
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/usecases"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h := &handlers.Handler{
		SessionsRepo: &repositories.RestSessionsRepository{},
		UsersRepo:    &repositories.RestUsersRepository{},
		GetAllUserVideosUC: &usecases.GetAllUserVideos{
			UsersRepo:  repositories.RestUsersRepository{},
			VideosRepo: repositories.RestVideosRepository{},
			WTARepo:    repositories.RestWTARepository{},
		},
	}

	r.HandleFunc("/authorize", h.Authorize).Methods("POST")

	r.HandleFunc("/user/videos", handlers.VerifyJWT(h.AddSessionData(h.GetAllUserVideos)))

	certPath := "/app/server.pem"
	keyPath := "/app/server.key"

	err := http.ListenAndServeTLS(":8443", certPath, keyPath, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
