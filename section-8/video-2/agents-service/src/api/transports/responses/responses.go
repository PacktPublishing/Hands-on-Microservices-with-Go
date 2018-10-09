package responses

import "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/entities"

type InsertAgentPlayerResponse struct {
	Err string `json:"error,omitempty`
}

type GetAgentByIDResponse struct {
	Agent *entities.Agent `json: agent,omitempty`
	Err   string          `json:"error,omitempty"`
}
