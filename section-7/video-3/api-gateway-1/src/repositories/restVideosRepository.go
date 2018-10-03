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

type RestVideosRepository struct{}

var Err404OnVideoRequest = errors.New("404 Not Found on Video Request")

func (repo *RestVideosRepository) GetAllVideosByUserID(userID uint32) ([]*entities.Video, error) {

	videoIDStr := strconv.Itoa(int(userID))

	resp, err := http.Get("http://videos-service:8080/video/" + url.PathEscape(videoIDStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnVideoRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var videos []*entities.Video
	err = json.Unmarshal(body, &videos)
	if err != nil {
		return nil, err
	}

	return videos, nil
}
