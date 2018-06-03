package handlers

import (
	"fmt"
	"net/http"
)

func ok200(w http.ResponseWriter, body string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, body)
}

func error400(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Error 400: "+msg)
}

func error500(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: "+err.Error())
}
