package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/agents-service/repositories"
)

type Handlers struct {
	Repo *repositories.MariaDBAgentsRepository
}

type UpdateAgentAccountDTO struct {
	AgentID uint32 `json:"agent_id"`
	UserID  uint32 `json:"user_id"`
	VideoID uint32 `json:"video_id"`
	Ammount uint32 `json:"price"`
}

func (h *Handlers) UpdateAgentAccount(w http.ResponseWriter, r *http.Request) {
	uua := &UpdateAgentAccountDTO{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, uua)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = h.Repo.UpdateAgentAccount(uua.AgentID, uua.UserID, uua.VideoID, uua.Ammount)
	if err != nil {
		if err == repositories.ErrReceiptAlreadyExists {
			w.WriteHeader(http.StatusConflict)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) RollbackUpdateAgentAccount(w http.ResponseWriter, r *http.Request) {
	uua := &UpdateAgentAccountDTO{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, uua)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = h.Repo.RollbackUpdateAgentAccount(uua.AgentID, uua.UserID, uua.VideoID, uua.Ammount)
	if err == repositories.ErrNothingToRollback {
		w.WriteHeader(http.StatusConflict)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
