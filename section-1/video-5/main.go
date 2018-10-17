package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var conversionMap = map[string]string{
	"ASR": "-3h",     //North America Atlantinc Standard Time
	"EST": "-5h",     //North America Eastern Standard Time
	"BST": "+1h",     //British Summer Time
	"IST": "+5h30m",  //Indian Standard Time
	"HKT": "+8h",     //Hong Kong Time
	"ART": "-3h",     //Argentina Time
	"RET": "MMMMNUN", //WILL RAISE AN ERROR
}

type timeZoneConvertion struct {
	TimeZone    string
	CurrentTime string
}

func main() {
	http.HandleFunc("/convert", loggingMiddleware(handler))
	http.HandleFunc("/", loggingMiddleware(notFoundHandler))
	http.ListenAndServe("localhost:8080", nil)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func loggingMiddleware(handler handlerFunc) handlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", time.Now().Format("2018-11-25 14:32:58"), r.Method, r.URL.String())
		handler(w, r)
	}
	return fn
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404: The requested URL does not exist.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")
	if timeZone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error 400: tz query parameter is required.")
		return
	}

	timeDifference, ok := conversionMap[timeZone]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `Error 404: the tz value: "%s" does not correspond to an existing tz value.`, timeZone)
		return
	}
	currentTimeConverted, err := getcurrentTimeByTimeDifference(timeDifference)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server Error.")
		return
	}

	w.WriteHeader(http.StatusOK)

	tzc := new(timeZoneConvertion)
	tzc.CurrentTime = currentTimeConverted
	tzc.TimeZone = timeZone

	jsonResponse, err := json.Marshal(tzc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server Error.")
		return
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func getcurrentTimeByTimeDifference(timeDifference string) (string, error) {
	now := time.Now().UTC()
	difference, err := time.ParseDuration(timeDifference)
	if err != nil {
		return "", err
	}
	now = now.Add(difference)
	return now.Format("15:04:05"), nil
}
