package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	type Author struct {
		Name     string `json:"composer_name"`
		LastName string `json:"composer_last_name"`
	}
	type Song struct {
		*Author
		Band          string    `json:"band"`
		Song          string    `json:"song"`
		ReleaseDate   time.Time `json:"release_date"`
		Label         string    `json:"label,omitempty"` //This is wierd the go inside the ""
		ChartPosition int       `json:"chart_position"`
		Producer      string    `json:"-"` //allways omited
	}
	song := &Song{
		Author: &Author{Name: "Bob",
			LastName: "Dylan"},
		Band:        "The Rolling Stones",
		Song:        "Like a rolling stone",
		ReleaseDate: time.Date(1995, 7, 13, 0, 0, 0, 0, time.UTC),
	}
	jsonSong, _ := json.Marshal(song)
	fmt.Println(string(jsonSong))
}
