package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/usecases"
	"github.com/gorilla/mux"
)

type Handlers struct {
	GetPlayerUsecase usecases.GetPlayer
	GetMatchUsecase  usecases.GetMatch
}

func (handler *Handlers) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerIDstr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "PlayerID parameter is required.")
		return
	}
	playerID, err := strconv.Atoi(playerIDstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	player, err := handler.GetPlayerUsecase.GetPlayer(uint32(playerID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	playerJSON, err := json.Marshal(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(playerJSON))
}

func (handler *Handlers) GetMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	matchIDstr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "MatchID parameter is required.")
		return
	}
	matchID, err := strconv.Atoi(matchIDstr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	match, err := handler.GetMatchUsecase.GetMatch(uint32(matchID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	matchJSON, err := json.Marshal(match)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(matchJSON))
}
