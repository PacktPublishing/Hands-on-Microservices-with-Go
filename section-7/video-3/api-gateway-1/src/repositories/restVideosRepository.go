package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/entities"
)

type RestVideosRepository struct{}

var Err404OnVideoRequest = errors.New("404 Not Found on Videos Request")
var Err500OnVideoRequest = errors.New("500 on Videos Request")

func (repo *RestVideosRepository) GetAllVideosByUserID(userID uint32) ([]*entities.Video, error) {

	userIDStr := strconv.Itoa(int(userID))

	resp, err := http.Get("http://videos-service:8080/videos/" + url.PathEscape(userIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnVideoRequest
	}
	if resp.StatusCode == 500 {
		return nil, Err500OnVideoRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var videos []*entities.Video
	err = json.Unmarshal(body, &videos)
	if err != nil {
		log.Println(string(body))
		return nil, err
	}

	return videos, nil
}
