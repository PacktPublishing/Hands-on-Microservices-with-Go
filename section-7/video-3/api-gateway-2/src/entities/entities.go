package entities

import "time"

type Player struct {
	ID            uint32    `json:"player_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	IsRightHanded bool      `json:"is_rignt_handed"`
	BirthDate     time.Time `json:"birth_date"`
	CountryCode   string    `json:"country_code"`
}

type Agent struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Account uint32 `json:"account"`
}

type Match struct {
	ID       uint32    `json:"id"`
	WinnerID uint32    `json:"winner_id"`
	LoserID  uint32    `json:"loser_id"`
	Date     time.Time `json:"date"`
}
