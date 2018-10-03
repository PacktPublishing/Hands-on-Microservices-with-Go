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

type RestWTARepository struct{}

var Err404OnMatchRequest = errors.New("404 Not Found on Match Request")
var Err500OnMatchRequest = errors.New("500 on Match Request")

func (repo *RestWTARepository) GetMatchByMatchID(matchID uint32) (*entities.Match, error) {

	matchIDStr := strconv.Itoa(int(matchID))

	resp, err := http.Get("http://wta-service:8080/match/" + url.PathEscape(matchIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnMatchRequest
	}
	if resp.StatusCode == 500 {
		return nil, Err500OnMatchRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var match entities.Match
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}

var Err404OnPlayerRequest = errors.New("404 Not Found on Player Request")

func (repo *RestWTARepository) GetPlayerByPlayerID(playerID uint32) (*entities.Player, error) {

	playerIDStr := strconv.Itoa(int(playerID))

	resp, err := http.Get("http://wta-service:8080/player/" + url.PathEscape(playerIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnPlayerRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var player entities.Player
	err = json.Unmarshal(body, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}
