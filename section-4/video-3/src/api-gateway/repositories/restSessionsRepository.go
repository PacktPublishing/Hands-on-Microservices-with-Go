package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/api-gateway/entities"
)

var Err404OnSessionRequest = errors.New("404 Not Found for Session Key")

type RestSessionsRepository struct{}

func (repo *RestSessionsRepository) GetSession(key string) (*entities.Session, error) {

	resp, err := http.Get("http://127.0.0.1:8001/session/" + url.PathEscape(key))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 404 {
		return nil, Err404OnSessionRequest
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) == "" {
		return nil, err
	}

	var session entities.Session
	err = json.Unmarshal(body, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil

}

func (repo *RestSessionsRepository) SetSession(key string, session *entities.Session) error {
	jsonSession, err := json.Marshal(session)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(jsonSession)
	url := "http://127.0.0.1:8001/session/" + url.PathEscape(key)
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodPut, url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return errors.New("Bad Response from Server.")
	}

	return nil
}
