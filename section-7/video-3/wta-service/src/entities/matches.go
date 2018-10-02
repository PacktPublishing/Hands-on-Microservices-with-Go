package entities

import "time"

type Match struct {
	ID       uint32    `json:"id"`
	WinnerID uint32    `json:"winner_id"`
	LoserID  uint32    `json:"loser_id"`
	Date     time.Time `json:"date"`
}
