package entities

import "time"

type Video struct {
	ID        uint32 `json:"id"`
	MatchID   uint32 `json:"match_id"`
	Player1ID uint32 `json:"player1_id"`
	Player2ID uint32 `json:"player2_id"`
	Duration  uint32 `json:"duration"`
	Price     uint32 `json:"price"`
}

type Match struct {
	ID       uint32    `json:"id"`
	WinnerID uint32    `json:"winner_id"`
	LoserID  uint32    `json:"loser_id"`
	Date     time.Time `json:"date"`
}

type Player struct {
	ID            uint32    `json:"player_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	IsRightHanded bool      `json:"is_rignt_handed"`
	BirthDate     time.Time `json:"birth_date"`
	CountryCode   string    `json:"country_code"`
}

type User struct {
	ID        uint32    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email, omitempty"`
	BirthDate time.Time `json:"birth_date, omitempty"`
	Added     time.Time `json:"added,  omitempty"`
	Account   uint32    `json:"account"`
	Password  string    `json:"password"`
}
