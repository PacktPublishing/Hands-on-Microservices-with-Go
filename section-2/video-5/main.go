package main

import (
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-2/video-5/handlers"
)

func main() {
	http.HandleFunc("/example", handlers.MyHandler)
	http.ListenAndServe(":8080", nil)
}
