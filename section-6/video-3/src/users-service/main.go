package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/handlers"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/usecases"
	"github.com/gorilla/mux"
)

var (
	requestTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "api_request_time",
		Help:       "Time it takes to handle a Request.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	userRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests",
			Help: "Number of requests for the API.",
		},
		[]string{"request"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(requestTime)
	prometheus.MustRegister(userRequests)
}

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

	r.HandleFunc("/user/{username}", countRequests(measureRequests(handler.GetUserByUsernameHandler))).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting server on Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func countRequests(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRequests.With(prometheus.Labels{"request": "/user/username"}).Inc()
		f(w, r)
	}
}

func measureRequests(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		elapsed := time.Since(start)
		requestTime.Observe(float64(elapsed))

	}
}
