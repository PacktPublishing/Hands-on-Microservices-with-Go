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

type RestManagersRepository struct{}

var Err404OnManagerRequest = errors.New("404 Not Found on Manager Request")

func (repo *RestManagersRepository) GetManagerByManagerID(managerID uint32) (*entities.Manager, error) {

	managerIDStr := strconv.Itoa(int(managerID))

	resp, err := http.Get("http://127.0.0.1:8000/manager/" + url.PathEscape(managerIDStr))
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
