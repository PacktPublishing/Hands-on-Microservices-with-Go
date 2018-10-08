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

type RestAgentsRepository struct{}

var Err404OnAgentRequest = errors.New("404 Not Found on Agent Request")

type AgentPlayerIDsDTO struct {
	PlayerIDs []uint32 `json:"PlayerIDs,omitempty"`
	Err       string   `json:"error,omitempty"`
}

func (repo *RestAgentsRepository) GetAgentByAgentID(agentID uint32) (*entities.Agent, error) {

	agentIDStr := strconv.Itoa(int(agentID))
	url := "http://agents-service:8080/agent/" + url.PathEscape(agentIDStr)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnAgentRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var agent entities.Agent
	err = json.Unmarshal(body, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

func (repo *RestAgentsRepository) GetAgentPlayers(agentID uint32) (*AgentPlayerIDsDTO, error) {

	agentIDStr := strconv.Itoa(int(agentID))

	resp, err := http.Get("http://agents-service:8080/agent/players/" + url.PathEscape(agentIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnAgentRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	playerIDs := AgentPlayerIDsDTO{}
	err = json.Unmarshal(body, &playerIDs)
	if err != nil {
		return nil, err
	}

	return &playerIDs, nil
}
