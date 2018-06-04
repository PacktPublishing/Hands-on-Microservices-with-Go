package handlers

import (
	"fmt"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprint(w, "Request was succesful")
}
