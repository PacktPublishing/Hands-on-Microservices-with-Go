package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/sessions-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/sessions-service/repositories"
	"github.com/gorilla/mux"
)

type Handlers struct {
	Repo *repositories.RedisSessionsRepository
}

func (handler *Handlers) GetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key, ok := vars["key"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	session, err := handler.Repo.GetSession(key)
	if err == repositories.ErrRespIsNil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	sessionJSON, err := json.Marshal(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(sessionJSON))
}

func (handler *Handlers) SetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key, ok := vars["key"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	strSession, err := ioutil.ReadAll(r.Body)
	if err != nil || string(strSession) == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	var session *entities.Session
	err = json.Unmarshal(strSession, &session)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	err = handler.Repo.SetSession(key, session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, session)
}
