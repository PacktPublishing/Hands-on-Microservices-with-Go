package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/repositories"
	"github.com/gorilla/mux"
)

type ctxKey int

const (
	AgentID ctxKey = iota
)

type Handler struct {
	AgentsRepo repositories.RestAgentsRepository
	WTARepo    repositories.RestWTARepository
}

type AgentPlayersDTO struct {
	Agent   *entities.Agent
	Players []*entities.Player
}

func (h *Handler) GetAgentPlayers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	agentIDInt, err := strconv.Atoi(vars["id"])
	agentID := uint32(agentIDInt)

	var agent *entities.Agent
	var agentErr error
	var playerIDs *repositories.AgentPlayerIDsDTO
	var playerIDsErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup, agentID uint32) {
		defer wg.Done()
		log.Println("Agent API call")
		agent, agentErr = h.AgentsRepo.GetAgentByAgentID(agentID)
	}(&wg, agentID)

	go func(wg *sync.WaitGroup, agentID uint32) {
		defer wg.Done()
		log.Println("AgentPlayers API call")
		playerIDs, playerIDsErr = h.AgentsRepo.GetAgentPlayers(agentID)
	}(&wg, agentID)

	wg.Wait()

	if agentErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("A" + agentErr.Error()))
		return
	}

	if playerIDsErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("B" + playerIDsErr.Error()))
		return
	}

	type playerDTO struct {
		Player *entities.Player
		Err    error
	}

	ch := make(chan *playerDTO)

	count := len(playerIDs.PlayerIDs)
	playerDTOs := make([]*playerDTO, 0, count)

	if count == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No Players for this Agent."))
		return
	}

	for i := 0; i < count; i++ {
		go func(playerID uint32) {
			player, err := h.WTARepo.GetPlayerByPlayerID(playerID)

			playerDTO := &playerDTO{
				Player: player,
				Err:    err,
			}
			log.Println("Player API CALL: ", playerID)
			ch <- playerDTO
		}(playerIDs.PlayerIDs[i])
	}

	for i := 0; i < count; i++ {
		playerDTOs = append(playerDTOs, <-ch)
	}

	errsStr := ""

	for i := 0; i < count; i++ {
		if playerDTOs[i].Err != nil {
			errsStr += playerDTOs[i].Err.Error()
		}
	}

	if errsStr != "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errsStr))
		return
	}

	res := &AgentPlayersDTO{
		Agent: agent,
	}

	for i := 0; i < count; i++ {
		res.Players = append(res.Players, playerDTOs[i].Player)
	}

	json, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
}
