package entities

import "time"

type Match struct {
	ID       uint32
	WinnerID uint32
	LoserID  uint32
	Date     time.Time
}
