package repositories

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

type RestAgentsRepository struct{}

func (repo *RestAgentsRepository) UpdateAgentAccount(bvsDTO *BuyVideoSagaDTO) error {

	url := "http://127.0.0.1:8083/agent/ammount/update"
	jsonBytes, err := json.Marshal(bvsDTO)
	if err != nil {
		return err
	}
	bodyReader := bufio.NewReader(strings.NewReader(string(jsonBytes)))

	req, err := http.NewRequest("PATCH", url, bodyReader)

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

func (repo *RestAgentsRepository) RollbackUpdateAgentAccount(bvsDTO *BuyVideoSagaDTO) error {

	url := "http://127.0.0.1:8083/agent/ammount/rollback"
	jsonBytes, err := json.Marshal(bvsDTO)
	if err != nil {
		return err
	}
	bodyReader := bufio.NewReader(strings.NewReader(string(jsonBytes)))

	req, err := http.NewRequest("PATCH", url, bodyReader)

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
	//This is the case when we try to delete a receipt
	//that does not exist.
	//We will return nil (no error) so that the saga can continue
	//But in real life you should log the error and investigate
	//what is happening since this will happen when you are processing
	//a rollback 2 times consecutively.
	//The place to investigate would be how you are inserting and
	//processing messages from the quee.
	if resp.StatusCode == http.StatusConflict {
		return nil
	}
	if resp.StatusCode != 200 {
		return ErrOnRestRequest
	}

	return nil
}
