package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", myHandler)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist("metonymie.com"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go func() {
		log.Println("Starting https server on port: 80")
		err := http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Println("Starting https server on port: 443")
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Recieved request: " + r.RequestURI + "\n")
		h.ServeHTTP(w, r)
	})
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Recieved request: " + r.RequestURI + "\n")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello HTTPS World")
}
