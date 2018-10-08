package usecases

import (
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/repositories"
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

	//If there was an error when setting cache, we log it but we do not return it
	if err != nil {
		//There was an error different than it not being on cache.
		log.Println("Error on Cache.", err.Error())
	}

	return nil
}
