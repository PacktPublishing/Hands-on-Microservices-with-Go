package usecases

import (
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/users-service/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/users-service/utils/appErrors"
)

//The difference between data transfer objects and business objects or data access objects is that a DTO does not have any behavior except for storage, retrieval, serialization and deserialization of its own data (mutators, accessors, parsers and serializers). In other words, DTOs are simple objects that should not contain any business logic but may contain serialization and deserialization mechanisms for transferring data over the wire..
type UserDTO struct {
	entities.User
	AccountType entities.UserAccountType
}

type GetUser interface {
	GetUser(username string) (*UserDTO, error)
}

type GetUserImpl struct {
	CacheRepo repositories.UsersCacheRepository
	Repo      repositories.UsersRepository
}

func (uc *GetUserImpl) GetUser(username string) (*UserDTO, error) {
	//Look In Cache
	user, err := uc.CacheRepo.GetUser(username)

	//CHECK FOR NOT IN CACHE
	if err == nil {
		//It was in cache, return it
		userDTO := &UserDTO{
			User:        *user,
			AccountType: user.GetAccountType(),
		}

		return userDTO, nil
	}

	if err != appErrors.ErrorNotFoundOnCache {
		//There was an error different than it not being on cache.
		log.Println("Error on Cache.", err.Error())
	}

	//Not in cache
	user, err = uc.Repo.GetUserByUsername(username)

	if err != nil {
		if err == appErrors.ErrorNotFoundOnDB {
			return nil, appErrors.ErrorNotFound
		}
		return nil, err
	}

	//Update cache for future requets
	err = uc.CacheRepo.SetUser(username, user)

	if err != nil {
		//There was an error while working with the cache
		log.Println("Error on Cache.", err.Error())
	}

	userDTO := &UserDTO{
		User:        *user,
		AccountType: user.GetAccountType(),
	}

	return userDTO, nil
}
