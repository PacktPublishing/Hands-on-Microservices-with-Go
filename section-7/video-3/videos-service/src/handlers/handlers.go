package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/usecases"
	"github.com/gorilla/mux"
)

type Handlers struct {
	GetAllUserVideosUsecase usecases.GetAllUserVideos
}

func (handler *Handlers) GetAllUserVideos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userIDstr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "UserID parameter is required.")
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	videos, err := handler.GetAllUserVideosUsecase.GetAllUserVideos(uint32(userID))
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	videosJSON, err := json.Marshal(videos)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(videosJSON))
}
