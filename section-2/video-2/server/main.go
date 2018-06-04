package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/example", myHandler)
	http.HandleFunc("/exampleHttpError", myErrorHandler)

	http.HandleFunc("/forPostClient", myPostHandler)

	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)

	//Writing Other Headers
	header := w.Header()
	header.Set("Content-Type", "application/text")

	n, err := w.Write([]byte("Hi There"))
	if err != nil {
		log.Fatal("Error Writing.")
	}
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "The number of bytes on the last operation were: %d", n)

}

func myErrorHandler(w http.ResponseWriter, r *http.Request) {
	//Simple way to write errors
	http.Error(w, "My error MSG that will go in the Body", http.StatusInternalServerError)
}

func myPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	fmt.Fprintln(w, "Server:: The Request Body was:")
	fmt.Fprintln(w, string(body))
}
