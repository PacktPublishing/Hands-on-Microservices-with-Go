package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/entities"
)

type RestUsersRepository struct{}

var Err404OnUserRequest = errors.New("404 Not Found on User Request")
var Err500OnUserRequest = errors.New("500 on User Request")

func (repo *RestUsersRepository) GetUserByUsername(username string) (*entities.User, error) {

	resp, err := http.Get("http://users-service:8080/user/by/username/" + url.PathEscape(username))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnUserRequest
	}
	if resp.StatusCode == 500 {
		return nil, Err500OnUserRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var user entities.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RestUsersRepository) GetUserByUserID(userID uint32) (*entities.User, error) {

	userIDStr := strconv.Itoa(int(userID))

	resp, err := http.Get("http://users-service:8080/user/by/id/" + url.PathEscape(userIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnUserRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var user entities.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
