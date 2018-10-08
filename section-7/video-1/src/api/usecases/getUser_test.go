package usecases

import (
	"testing"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/utils/appErrors"
)

func Test_GetUser_InCache(t *testing.T) {
	uc := &GetUserImpl{
		CacheRepo: &MockUsersCacheRepository{},
		Repo:      &MockUsersRepository{},
	}

	userDTO, err := uc.GetUser("Test1")
	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}

	if userDTO.Username != "Test1" {
		t.Errorf("Test failed. Wrong Username.")
	}
	if userDTO.Email != "test1@example.com" {
		t.Errorf("Test failed. Wrong Email.")
	}
	if userDTO.AccountType != entities.GoldAccount {
		t.Errorf("Test failed. Expected Gold Account.")
	}

}

func Test_GetUser_NotInCache(t *testing.T) {

	cacheRepo := &MockUsersCacheRepository{}

	uc := &GetUserImpl{
		CacheRepo: cacheRepo,
		Repo:      &MockUsersRepository{},
	}

	userDTO, err := uc.GetUser("Test2")
	if err != nil {
		t.Errorf("Test failed. Unexpected error.")
	}

	if userDTO.Username != "Test2" {
		t.Errorf("Test failed. Wrong Username.")
	}
	if userDTO.Email != "test2@example.com" {
		t.Errorf("Test failed. Wrong Email.")
	}
	if userDTO.AccountType != entities.SilverAccount {
		t.Errorf("Test failed. Expected Gold Account.")
	}
	if !cacheRepo.WasSetCalled {
		t.Errorf("Test failed. Expected Cache Set.")
	}

}

func Test_GetUser_NotInCacheOrDB(t *testing.T) {
	cacheRepo := &MockUsersCacheRepository{}

	uc := &GetUserImpl{
		CacheRepo: cacheRepo,
		Repo:      &MockUsersRepository{},
	}

	_, err := uc.GetUser("Test3")
	if err != appErrors.ErrorNotFound {
		t.Errorf("Test failed. Expected not Found Error.")
	}

}

type MockUsersRepository struct{}

func (m *MockUsersRepository) GetUserByUsername(username string) (*entities.User, error) {
	if username == "Test2" {
		return &entities.User{
			Username: "Test2",
			Account:  1600,
			Email:    "test2@example.com",
		}, nil
	} else {
		return nil, appErrors.ErrorNotFound
	}
}

func (m *MockUsersRepository) GetUserByID(userID uint32) (*entities.User, error) {
	//Not Used on this example tests.
	return nil, nil
}

func (m *MockUsersRepository) UpdateUser(user *entities.User) error {
	return nil
}

type MockUsersCacheRepository struct {
	WasSetCalled bool
}

func (m *MockUsersCacheRepository) GetUser(username string) (*entities.User, error) {
	if username == "Test1" {
		return &entities.User{
			Username: "Test1",
			Account:  3600,
			Email:    "test1@example.com",
		}, nil
	} else {
		return nil, appErrors.ErrorNotFoundOnDB
	}
}

func (m *MockUsersCacheRepository) SetUser(username string, User *entities.User) error {

	m.WasSetCalled = true
	return nil
}
