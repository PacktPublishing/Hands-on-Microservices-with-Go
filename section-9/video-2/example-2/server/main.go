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
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Hello!")
}
