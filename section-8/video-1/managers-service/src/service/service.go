package service

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/managers-service/src/repositories"
)

// Service Interface.
type ManagersService interface {
	InsertManagerPlayer(managerID uint32, playerID uint32) error
	GetManagerByID(managerID uint32) (*entities.Manager, error)
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

// ServiceMiddleware is a chainable behavior modifier for ManagersService.
type ServiceMiddleware func(ManagersService) ManagersServiceImpl
