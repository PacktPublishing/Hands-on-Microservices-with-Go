package usecases

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/repositories"
)

type UpdateUserUsecase struct {
	CacheRepo *repositories.RedisUsersRepository
	Repo      *repositories.MySQLUsersRepository
}

func (uc *UpdateUserUsecase) UpdateUser(user *entities.User) error {

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
