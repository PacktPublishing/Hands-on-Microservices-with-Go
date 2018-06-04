package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/item/{name}/page/{page}", myHandlerFunc)

	//Methods
	r.HandleFunc("/example", postHandlerFunc).Methods("POST")
	r.HandleFunc("/example", getHandlerFunc).Methods("GET")

	//Plenty of others see http://www.gorillatoolkit.org/pkg/mux

	http.ListenAndServe(":8080", r)

}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	name := vars["name"]
	page := vars["page"]

	//Vars must have URL valid
	//Always Validate Type!
	fmt.Fprintln(w, "name: "+name)
	fmt.Fprintln(w, "page: "+page)
}

func getHandlerFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintln(w, "Method: "+method)

	requestURI := r.RequestURI
	fmt.Fprintln(w, "requestUri: "+requestURI)
}

func postHandlerFunc(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintln(w, "Method: "+method)

	requestURI := r.RequestURI
	fmt.Fprintln(w, "requestUri: "+requestURI)
}
