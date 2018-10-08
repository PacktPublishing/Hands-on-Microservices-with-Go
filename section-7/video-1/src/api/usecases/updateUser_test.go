package usecases

import (
	"testing"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
)

func Test_UpdateUser(t *testing.T) {

	cacheRepo := &MockUsersCacheRepository{}

	uc := &UpdateUserImpl{
		CacheRepo: cacheRepo,
		Repo:      &MockUsersRepository{},
	}

	user := &entities.User{}
	user.Username = "Pepe"
	user.Email = "pepe@example.com"
	user.FirstName = "Pepe"
	user.LastName = "L.J. Smith"
	user.Password = "Hashed"

	err := uc.UpdateUser(user)

	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}
	if !cacheRepo.WasSetCalled {
		t.Errorf("Test failed. Expected Cache Set.")
	}
}
