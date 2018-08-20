package usecases

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/users-service/repositories"
)

type UpdateUser interface {
	UpdateUser(user *entities.User) error
}

type UpdateUserImpl struct {
	CacheRepo repositories.UsersCacheRepository
	Repo      repositories.UsersRepository
}

func (uc *UpdateUserImpl) UpdateUser(user *entities.User) error {

	//Update User DB
	err := uc.Repo.UpdateUser(user)
	if err != nil {
		return err
	}

	//Update Cache
	err = uc.CacheRepo.SetUser(user.Username, user)
	//Importa error?
	if err != nil {
		return err
	}

	return nil
}
