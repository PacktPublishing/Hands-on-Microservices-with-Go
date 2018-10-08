package service

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/agents-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/agents-service/src/repositories"
)

// Service Interface.
type AgentsService interface {
	InsertAgentPlayer(agentID uint32, playerID uint32) error
	GetAgentByID(agentID uint32) (*entities.Agent, error)
	GetAgentPlayerIDs(agentID uint32) ([]uint32, error)
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

func (srv AgentsServiceImpl) GetAgentPlayerIDs(agentID uint32) ([]uint32, error) {
	playerIDs, err := srv.Repo.GetAgentPlayers(agentID)
	return playerIDs, err
}

// ServiceMiddleware is a chainable behavior modifier for AgentsService.
type ServiceMiddleware func(AgentsService) AgentsServiceImpl
