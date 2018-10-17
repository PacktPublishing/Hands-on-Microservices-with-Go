package repositories

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

type RestVideosRepository struct{}

func (repo *RestVideosRepository) InsertBoughtVideo(bvsDTO *BuyVideoSagaDTO) error {

	bvDTO := &BoughtVideoDTO{}
	bvDTO.VideoID = bvsDTO.VideoID
	bvDTO.UserID = bvsDTO.UserID

	url := "http://127.0.0.1:8081/bought-video"
	jsonBytes, err := json.Marshal(bvDTO)
	if err != nil {
		return err
	}
	bodyReader := bufio.NewReader(strings.NewReader(string(jsonBytes)))

	req, err := http.NewRequest("POST", url, bodyReader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 400 {
		return Err400OnRestRequest
	}
	if resp.StatusCode == 500 {
		return Err500OnRestRequest
	}
	//For idempotency
	if resp.StatusCode == http.StatusConflict {
		return nil
	}
	if resp.StatusCode != 200 {
		return ErrOnRestRequest
	}

	return nil
}

func (repo *RestVideosRepository) DeleteBoughtVideo(bvsDTO *BuyVideoSagaDTO) error {

	url := "http://127.0.0.1:8081/bought-video"
	jsonBytes, err := json.Marshal(bvsDTO)
	if err != nil {
		return err
	}
	bodyReader := bufio.NewReader(strings.NewReader(string(jsonBytes)))

	req, err := http.NewRequest("DELETE", url, bodyReader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 400 {
		return Err400OnRestRequest
	}
	if resp.StatusCode == 500 {
		return Err500OnRestRequest
	}
	//For idempotency
	if resp.StatusCode == http.StatusConflict {
		return nil
	}
	if resp.StatusCode != 200 {
		return ErrOnRestRequest
	}

	return nil
}
