package responses

import "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/agents-service/src/entities"

type InsertAgentPlayerResponse struct {
	Err string `json:"error,omitempty`
}

type GetAgentByIDResponse struct {
	Agent *entities.Agent `json: agent,omitempty`
	Err   string          `json:"error,omitempty"`
}

type GetAgentPlayerIDsResponse struct {
	PlayerIDs []uint32 `json: player_ids,omitempty`
	Err       string   `json:"error,omitempty"`
}
