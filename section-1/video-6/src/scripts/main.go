package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/globalsign/mgo"
)

type timeZoneConvertion struct {
	TimeZone       string `bson:timeZone json:timeZone`
	TimeDifference string `bson:timeDifference json:timeDifference`
	Name           string `bson:name json:name`
}

type tzs []timeZoneConvertion

func main() {
	client, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB")
	}
	db := client.DB("packt")
	err = db.DropDatabase()
	if err != nil {
		log.Fatal("Couldn't drop MongoDB database. Err: " + err.Error())
	}

	coll := db.C("timeZones")

	data, err := ioutil.ReadFile("all_timezones.json")
	if err != nil {
		log.Fatal("Couldn't open file")
	}

	var timeZones tzs
	err = json.Unmarshal(data, &timeZones)
	if err != nil {
		log.Fatal("Couldn't unmarshall Json")
	}

	for _, v := range timeZones {
		coll.Insert(&v)
	}
}
