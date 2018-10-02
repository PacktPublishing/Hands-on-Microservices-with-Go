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

type Manager struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Account uint32 `json:"account"`
}
