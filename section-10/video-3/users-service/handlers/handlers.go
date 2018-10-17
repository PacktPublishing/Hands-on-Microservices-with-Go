package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/users-service/repositories"
)

type Handlers struct {
	Repo *repositories.MariaDBUsersRepository
}

type UpdateUserAccountDTO struct {
	UserID  uint32 `json:"user_id"`
	VideoID uint32 `json:"video_id"`
	Ammount uint32 `json:"price"`
}

func (h *Handlers) UpdateUserAccount(w http.ResponseWriter, r *http.Request) {
	uua := &UpdateUserAccountDTO{}
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
	err = h.Repo.UpdateUserAccount(uua.UserID, uua.VideoID, uua.Ammount)
	if err == repositories.ErrReceiptAlreadyExists {
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

func (h *Handlers) RollbackUpdateUserAccount(w http.ResponseWriter, r *http.Request) {
	uua := &UpdateUserAccountDTO{}
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
	err = h.Repo.RollbackUpdateUserAccount(uua.UserID, uua.VideoID, uua.Ammount)
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
