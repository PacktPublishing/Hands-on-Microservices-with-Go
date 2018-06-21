package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

var whitelist = []string{
	"127.0.0.1",
	"67.98.127.14",
	//...
}

func main() {
	http.HandleFunc("/", checkWhitelist(myHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello world")
}

func checkWhitelist(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		fmt.Println("IP: " + ip)

		inWhitelist := false

		for _, v := range whitelist {
			if ip == v {
				inWhitelist = true
				break
			}
		}

		if !inWhitelist {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		f(w, r)
	}
}
