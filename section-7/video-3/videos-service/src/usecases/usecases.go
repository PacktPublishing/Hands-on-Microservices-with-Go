package usecases

import (
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/repositories"
)

type GetAllUserVideos struct {
	Repo *repositories.MariaDBVideosRepository
}

func (uc *GetAllUserVideos) GetAllUserVideos(userID uint32) ([]*entities.Video, error) {
	videos, err := uc.Repo.GetAllUserVideos(userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return videos, err
}
