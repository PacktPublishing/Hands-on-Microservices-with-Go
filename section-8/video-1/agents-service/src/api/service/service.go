package service

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/repositories"
)

// Service Interface.
type AgentsService interface {
	InsertAgentPlayer(agentID uint32, playerID uint32) error
	GetAgentByID(agentID uint32) (*entities.Agent, error)
}

//Service Implementation
type AgentsServiceImpl struct {
	Repo *repositories.MariaDBAgentsRepository
}

func (srv AgentsServiceImpl) InsertAgentPlayer(agentID uint32, playerID uint32) error {
	err := srv.Repo.InsertAgentPlayer(agentID, playerID)
	return err
}

func (srv AgentsServiceImpl) GetAgentByID(agentID uint32) (*entities.Agent, error) {
	agent, err := srv.Repo.GetAgentByID(agentID)
	return agent, err
}

// ServiceMiddleware is a chainable behavior modifier for AgentsService.
type ServiceMiddleware func(AgentsService) AgentsServiceImpl
