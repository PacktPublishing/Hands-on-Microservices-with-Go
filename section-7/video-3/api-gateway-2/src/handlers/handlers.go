package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-2/src/repositories"
)

type ctxKey int

const (
	ManagerID ctxKey = iota
)

type Handler struct {
	ManagersRepo repositories.RestManagersRepository
	WTARepo      repositories.RestWTARepository
}

type ManagerPlayersDTO struct {
	Manager *entities.Manager
	Players []*entities.Player
}

func (h *Handler) GetManagerPlayers(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	managerID := ctx.Value(ManagerID).(uint32)

	var manager *entities.Manager
	var managerErr error
	var playerIDs *repositories.ManagerPlayerIDsDTO
	var playerIDsErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup, managerID uint32) {
		defer wg.Done()
		manager, managerErr = h.ManagersRepo.GetManagerByManagerID(managerID)
	}(&wg, managerID)

	go func(wg *sync.WaitGroup, managerID uint32) {
		defer wg.Done()
		playerIDs, playerIDsErr = h.ManagersRepo.GetManagerPlayers(managerID)
	}(&wg, managerID)

	wg.Wait()

	if managerErr != nil || playerIDsErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
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
		w.Write([]byte("No Players for this Manager."))
		return
	}

	for i := 0; i < count; i++ {
		go func(playerID uint32) {
			player, err := h.WTARepo.GetPlayerByPlayerID(playerID)

			playerDTO := &playerDTO{
				Player: player,
				Err:    err,
			}
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
		w.Write([]byte("No Players for this Manager."))
		return
	}

	res := &ManagerPlayersDTO{
		Manager: manager,
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
