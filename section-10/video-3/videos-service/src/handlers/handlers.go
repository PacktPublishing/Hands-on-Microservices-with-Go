package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/videos-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-10/video-3/videos-service/src/repositories"
	//	"github.com/gorilla/mux"
)

type Handlers struct {
	Repo *repositories.MariaDBVideosRepository
}

func (h *Handlers) InsertBoughtVideo(w http.ResponseWriter, r *http.Request) {
	boughtVideo := &entities.BoughtVideo{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, boughtVideo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = h.Repo.InsertBoughtVideo(boughtVideo.UserID, boughtVideo.VideoID)

	if err != nil {
		if err == repositories.ErrBoughtVideoAlreadyExists {
			w.WriteHeader(http.StatusConflict)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) DeleteBoughtVideo(w http.ResponseWriter, r *http.Request) {
	boughtVideo := &entities.BoughtVideo{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(body, boughtVideo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	err = h.Repo.RollBackBoughtVideo(boughtVideo.UserID, boughtVideo.VideoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
