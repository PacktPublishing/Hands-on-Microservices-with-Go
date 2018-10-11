package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/didip/tollbooth"
)

func main() {

	limiter := tollbooth.NewLimiter(50, nil)
	limiter.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})

	handlerWithLimiter := tollbooth.LimitFuncHandler(limiter, handler)
	http.Handle("/", handlerWithLimiter)

	log.Println("Starting Server on Port 9000.")
	http.ListenAndServe("localhost:9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Rate-Limit-Limit", "50")
	w.Header().Add("X-Rate-Limit-Duration", "1")
	w.Header().Add("X-Rate-Limit-Request-Forwarded-For", r.Header.Get("X-Forwarded-For"))
	w.Header().Add("X-Rate-Limit-Request-Remote-Addr", r.RemoteAddr)

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Hello!")
}
