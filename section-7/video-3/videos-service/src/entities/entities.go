package entities

type Video struct {
	ID        uint32
	MatchID   uint32
	Player1ID uint32
	Player2ID uint32
	Duration  uint32
	Price     uint32
}

type BoughtVideo struct {
	VideoID uint32
	UserID  uint32
}
