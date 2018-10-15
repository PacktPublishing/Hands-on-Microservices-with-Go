package repositories

import "errors"

var Err400OnRestRequest = errors.New("400 on Request")
var Err500OnRestRequest = errors.New("500 on Request")
var ErrOnRestRequest = errors.New("Non 200, 400 or 500 on Request")

type BuyVideoSagaDTO struct {
	AgentID uint32 `json:"agent_id"`
	UserID  uint32 `json:"user_id"`
	VideoID uint32 `json:"video_id"`
	Ammount uint32 `json:"price"`
}

type BoughtVideoDTO struct {
	VideoID uint32 `json:"video_id"`
	UserID  uint32 `json:"user_id"`
}
