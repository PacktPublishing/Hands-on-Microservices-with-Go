package requests

type InsertAgentPlayerRequest struct {
	AgentID  uint32 `json:"agent_id"`
	PlayerID uint32 `json:"player_id"`
}

type GetAgentByIDRequest struct {
	AgentID uint32
}

type GetAgentPlayerIDsRequest struct {
	AgentID uint32
}
