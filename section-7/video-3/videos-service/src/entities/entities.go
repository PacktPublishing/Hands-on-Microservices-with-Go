package entities

type Video struct {
	ID        uint32 `json:"id"`
	MatchID   uint32 `json:"match_id"`
	Player1ID uint32 `json:"player1_id"`
	Player2ID uint32 `json:"player2_id"`
	Duration  uint32 `json:"duration"`
	Price     uint32 `json:"price"`
}

type BoughtVideo struct {
	VideoID uint32
	UserID  uint32
}
