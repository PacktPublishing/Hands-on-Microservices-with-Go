package usecases

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/repository"
)

type GetPlayer struct {
	Repo *repository.PsqlRepo
}

func (uc *GetPlayer) GetPlayer(playerID uint32) (*entities.Player, error) {
	player, err := uc.Repo.GetPlayer(playerID)
	if err != nil {
		return nil, err
	}
	return player, err
}

type GetMatch struct {
	Repo *repository.PsqlRepo
}

func (uc *GetMatch) GetMatch(matchID uint32) (*entities.Match, error) {
	match, err := uc.Repo.GetMatch(matchID)
	if err != nil {
		return nil, err
	}
	return match, err
}
