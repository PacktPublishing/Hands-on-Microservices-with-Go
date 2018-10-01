package service

import (
	"errors"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/entities"
)

// Service Interface.
type ManagersService interface {
	InsertManagerPlayer(managerID uint32, playerID uint32) error
	GetManagerByID(managerID uint32) (*entities.Manager, error)
}
 
//Service Implementation
type ManagersServiceImpl struct{}

func (ManagersServiceImpl) InsertManagerPlayer(managerID uint32, playerID uint32) error {

}

func (ManagersServiceImpl) GetManagerByID(managerID uint32) (*entities.Manager, error)
{
	
}


// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for ManagersService.
type ServiceMiddleware func(managersService) ManagersServiceImpl
