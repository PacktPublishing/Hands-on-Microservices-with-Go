package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type timeZoneConvertion struct {
	TimeZone    string
	CurrentTime string
}

var conversionMap = map[string]string{
	"ASR": "-3h",    //North America Atlantinc Standard Time
	"EST": "-5h",    //North America Eastern Standard Time
	"BST": "+1h",    //British Summer Time
	"IST": "+5h30m", //Indian Standard Time
	"HKT": "+8h",    //Hong Kong Time
	"ART": "-3h",    //Argentina Time
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")

	timeDifference, _ := conversionMap[timeZone]

	currentTimeConverted, _ := getcurrentTimeByTimeDifference(timeDifference)

	tzc := new(timeZoneConvertion)
	tzc.CurrentTime = currentTimeConverted
	tzc.TimeZone = timeZone

	jsonResponse, _ := json.Marshal(tzc)

	w.WriteHeader(http.StatusOK)
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
