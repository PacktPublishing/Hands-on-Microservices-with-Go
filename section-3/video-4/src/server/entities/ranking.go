package entities

import "time"

type Ranking struct {
	PlayerID      uint32
	RankingDate   time.Time
	RankingNumber uint16
	RankingPoints float64
}
