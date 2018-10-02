package requests

type InsertManagerPlayerRequest struct {
	ManagerID uint32 `json:"manager_id"`
	PlayerID  uint32 `json:"player_id"`
}

type GetManagerByIDRequest struct {
	ManagerID uint32
}

type GetManagerPlayerIDsRequest struct {
	ManagerID uint32
}
