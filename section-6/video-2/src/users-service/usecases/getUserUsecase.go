package usecases

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-2/src/users-service/repositories"
)

type GetUserUsecase struct {
	CacheRepo *repositories.RedisUsersRepository
	Repo      *repositories.MySQLUsersRepository
}

func (uc *GetUserUsecase) GetUser(username string) (*entities.User, error) {
	//Look In Cache
	user, err := uc.CacheRepo.GetUser(username)
	//CHECK FOR NOT IN CACHE
	if err == nil {
		//It was in cache, return it
		return user, nil
	}

	//Not in cache
	user, err = uc.Repo.GetUserByUsername(username)
	//CHECK FOR NOT IN DB
	if err != nil {
		return nil, err
	}

	//Update cache for future requets
	err = uc.CacheRepo.SetUser(username, user)
	//Ignore Error?

	return user, nil
}
