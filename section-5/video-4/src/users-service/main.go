package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-4/src/users-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-4/src/users-service/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-4/src/users-service/usecases"
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

	//BASIC PPROF HANDLERS
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)

	// Manually add support for paths linked to by index page at /debug/pprof/
	r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/debug/pprof/block", pprof.Handler("block"))

	r.HandleFunc("/user/{username}", handler.GetUserByUsernameHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
