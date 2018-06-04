package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	proto := r.Proto
	fmt.Fprintln(w, "Proto: "+proto)

	method := r.Method
	fmt.Fprintln(w, "Method: "+method)

	requestURI := r.RequestURI
	fmt.Fprintln(w, "requestUri: "+requestURI)

	url := r.URL
	//URL Has parts
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	fmt.Fprintln(w, "URL:", url.String())
	fmt.Fprintln(w, "URL.Scheme:", url.Scheme)
	fmt.Fprintln(w, "URL.Host:", url.Host)
	fmt.Fprintln(w, "URL.Path:", url.Path)
	fmt.Fprintln(w, "URL.RawQuery:", url.RawQuery)
	fmt.Fprintln(w, "URL.Fragment:", url.Fragment)

	/*
		Usage of Username and Password
		is deprecated and should not be used.

		fmt.Fprintln(w, "URL.User.Username", url.User.Username())
		userPass, ok := url.User.Password()
		if ok {
			fmt.Fprintln(w, "URL.User.Password", userPass)
		}
	*/

	//Host can be separated into 2 parts, Hostname and Port
	hostname := url.Hostname()
	port := url.Port()
	fmt.Fprintln(w, "Hostname: ", hostname, ", Port: ", port)

	//To Parse the Query into Values
	//type Values map[string][]string
	queryValues := url.Query()
	fmt.Fprintln(w, "Queryvalues:")
	for key, val := range queryValues {
		fmt.Fprintln(w, key+": "+strings.Join(val, ","))
	}

	headers := ""
	for k, v := range r.Header {
		headers += k + "=" + strings.Join(v, ",") + "\n"
	}
	fmt.Fprintln(w, "Headers: \n"+headers)

	myHeader := r.Header.Get("My-Header")
	fmt.Fprintln(w, "myHeader: "+myHeader)

	//There are some Headers that are so common
	//that there are special functions for them
	fmt.Fprintln(w, "Referer: "+r.Referer())     //Refering URL
	fmt.Fprintln(w, "UserAgent: "+r.UserAgent()) //Client

	remoteAddress := r.RemoteAddr
	fmt.Fprintln(w, "Remote Address: "+remoteAddress)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	fmt.Fprintln(w, "Body:")
	fmt.Fprintln(w, string(body))

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	//First we need to call ParseForm
	//Request Body Params
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ERROR PARSING FORM")
	}

	method := r.Method
	fmt.Fprintln(w, "Method: "+method)

	requestURI := r.RequestURI
	fmt.Fprintln(w, "requestUri: "+requestURI)

	fmt.Fprintln(w, "Form Params:")

	//Post and Patch values And Also Querystring
	for key, val := range r.Form {
		fmt.Fprintln(w, key+": "+strings.Join(val, ","))
	}

	fmt.Fprintln(w, "---------------------------------")
	fmt.Fprintln(w, "Individual Form Param:")

	//Get an individual value from form
	aValue := r.Form.Get("Yet-Another-Value")
	fmt.Fprintln(w, "YAV: "+aValue)
}
