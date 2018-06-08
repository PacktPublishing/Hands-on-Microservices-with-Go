package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	certPath := "server.pem"
	keyPath := "server.key"
	clientCertPath := "client.pem"

	clientCaCert, err := ioutil.ReadFile(clientCertPath)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(clientCaCert)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", myHandlerFunc)

	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr:      ":8443",
		Handler:   mux,
		TLSConfig: cfg,
	}

	err = srv.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {

	//Write Status Code
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Hello world from HTTPS.")
}
