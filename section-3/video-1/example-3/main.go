package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Song struct {
	Band          string    `json:"band"`
	Song          string    `json:"song"`
	ReleaseDate   time.Time `json:"release_date"`
	Label         string    `json:"label,omitempty"` //omitempty goes inside the ""
	ChartPosition int       `json:"chart_position"`
	Producer      string    `json:"-"` //always omited
}

func main() {
	song := &Song{
		Band:        "Blur",
		Song:        "Coffe & Tv",
		ReleaseDate: time.Date(1999, 6, 28, 0, 0, 0, 0, time.UTC),
		Producer:    "William Orbit",
	}

	jsonSong, _ := json.Marshal(song)
	fmt.Println(string(jsonSong))
}
