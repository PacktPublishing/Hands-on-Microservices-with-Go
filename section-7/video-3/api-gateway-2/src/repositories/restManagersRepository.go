package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/entities"
)

type RestManagersRepository struct{}

var Err404OnManagerRequest = errors.New("404 Not Found on Manager Request")

type ManagerPlayerIDsDTO struct {
	PlayerIDs []uint32 `json:"PlayerIDs,omitempty"`
	Err       string   `json:"error,omitempty"`
}

func (repo *RestManagersRepository) GetManagerByManagerID(managerID uint32) (*entities.Manager, error) {

	managerIDStr := strconv.Itoa(int(managerID))
	url := "http://managers-service:8080/manager/" + url.PathEscape(managerIDStr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnManagerRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var manager entities.Manager
	err = json.Unmarshal(body, &manager)
	if err != nil {
		return nil, err
	}

	return &manager, nil
}

func (repo *RestManagersRepository) GetManagerPlayers(managerID uint32) (*ManagerPlayerIDsDTO, error) {

	managerIDStr := strconv.Itoa(int(managerID))

	resp, err := http.Get("http://managers-service:8080/manager/players/" + url.PathEscape(managerIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnManagerRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	playerIDs := ManagerPlayerIDsDTO{}
	err = json.Unmarshal(body, &playerIDs)
	if err != nil {
		return nil, err
	}

	return &playerIDs, nil
}
