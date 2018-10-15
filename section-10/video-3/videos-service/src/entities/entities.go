package entities

type BoughtVideo struct {
	VideoID uint32 `json:"video_id"`
	UserID  uint32 `json:"user_id"`
}
