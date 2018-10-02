package service

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/repositories"
)

// Service Interface.
type ManagersService interface {
	InsertManagerPlayer(managerID uint32, playerID uint32) error
	GetManagerByID(managerID uint32) (*entities.Manager, error)
	GetManagerPlayerIDs(managerID uint32) ([]uint32, error)
}

//Service Implementation
type ManagersServiceImpl struct {
	Repo *repositories.MariaDBManagersRepository
}

func (srv ManagersServiceImpl) InsertManagerPlayer(managerID uint32, playerID uint32) error {
	err := srv.Repo.InsertManagerPlayer(managerID, playerID)
	return err
}

func (srv ManagersServiceImpl) GetManagerByID(managerID uint32) (*entities.Manager, error) {
	manager, err := srv.Repo.GetManagerByID(managerID)
	return manager, err
}

func (srv ManagersServiceImpl) GetManagerPlayerIDs(managerID uint32) ([]uint32, error) {
	playerIDs, err := srv.Repo.GetManagerPlayers(managerID)
	return playerIDs, err
}

// ServiceMiddleware is a chainable behavior modifier for ManagersService.
type ServiceMiddleware func(ManagersService) ManagersServiceImpl
