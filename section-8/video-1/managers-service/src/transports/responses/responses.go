package responses

import "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/entities"

type InsertManagerPlayerResponse struct {
	Err string `json:"error,omitempty`
}

type GetManagerByIDResponse struct {
	Manager *entities.Manager `json: manager,omitempty`
	Err     string            `json:"error,omitempty"`
}
