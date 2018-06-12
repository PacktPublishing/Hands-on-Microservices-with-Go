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
	mux.Handle("/", logDecorator(http.HandlerFunc(myHandler)))

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("certs"),
		HostPolicy: autocert.HostWhitelist("httpstest.metonymie.com"),
	}
	cfg := &tls.Config{
		GetCertificate:           certManager.GetCertificate,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   mux,
		TLSConfig: cfg,
	}

	go func() {
		log.Println("Starting http server on port: 80")
		err := http.ListenAndServe(":80", logDecorator(certManager.HTTPHandler(nil)))
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello HTTPS World")
}
