package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/entities"
)

var Err404OnUserRequest = errors.New("404 Not Found on User Request")

type RestUsersRepository struct{}

func (repo *RestUsersRepository) GetUserByUsername(username string) (*entities.User, error) {

	resp, err := http.Get("127.0.0.1:8000/user/" + username)
	if resp.StatusCode == 404 {
		return nil, Err404OnUserRequest
	}
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var user *entities.User
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

/*
func (repo *RestUsersRepository) GetUserByUserID(UserID uint32) (*entities.User, error) {

	strUserID := strconv.FormatInt(int64(UserID), 32)

	resp, err := http.Get("127.0.0.1:8000/user/" + strUserID)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var user *entities.User
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
*/
