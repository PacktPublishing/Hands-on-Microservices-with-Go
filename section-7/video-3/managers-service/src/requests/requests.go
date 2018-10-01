package requests

type InsertManagerPlayerRequest struct {
	ManagerID uint32
	PlayerID  uint32
}

type GetManagerByIDRequest struct {
	ManagerID uint32
}
