package usecases

import (
	"errors"
	"sort"
	"sync"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/api-gateway-1/src/repositories"
)

var ErrNoVideosForUser = errors.New("No Videos for User")

type PlayerDTO struct {
	Player *entities.Player
}

type playerWithErrorDTO struct {
	Player *entities.Player
	Err    error
	Index  int
}

type playersWithErrorDTOList []*playerWithErrorDTO

func (a playersWithErrorDTOList) Len() int           { return len(a) }
func (a playersWithErrorDTOList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a playersWithErrorDTOList) Less(i, j int) bool { return a[i].Index < a[j].Index }

type VideoDTO struct {
	Video   *entities.Video
	Player1 *entities.Player
}

type AllUserVideosDTO struct {
	*entities.User
	Videos []*VideoDTO
}

type GetAllUserVideos struct {
	UsersRepo  repositories.RestUsersRepository
	VideosRepo repositories.RestVideosRepository
	WTARepo    repositories.RestWTARepository
}

func (uc GetAllUserVideos) GetAllVideosFromUser(userID uint32) (*AllUserVideosDTO, error) {

	var user *entities.User
	var userErr error
	var videos []*entities.Video
	var videosErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup, userID uint32) {
		defer wg.Done()
		user, userErr = uc.UsersRepo.GetUserByUserID(userID)
	}(&wg, userID)

	go func(wg *sync.WaitGroup, userID uint32) {
		defer wg.Done()
		videos, videosErr = uc.VideosRepo.GetAllVideosByUserID(userID)
	}(&wg, userID)

	wg.Wait()

	if userErr != nil {
		return nil, userErr
	}
	if videosErr != nil {
		return nil, videosErr
	}

	ch := make(chan *playerWithErrorDTO)

	count := len(videos)
	pwe := make(playersWithErrorDTOList, 0, count)

	if count == 0 {
		return nil, ErrNoVideosForUser
	}

	for i := 0; i < count; i++ {
		go func(playerID uint32, index int) {
			player, err := uc.WTARepo.GetPlayerByPlayerID(playerID)

			res := playerWithErrorDTO{
				Player: player,
				Err:    err,
				Index:  index,
			}
			ch <- &res
		}(videos[i].Player1ID, i)
	}

	for i := 0; i < count; i++ {
		pwe = append(pwe, <-ch)
	}

	sort.Sort(pwe)

	errsStr := ""

	for i := 0; i < count; i++ {
		if pwe[i].Err != nil {
			errsStr += pwe[i].Err.Error()
		}
	}

	if errsStr != "" {
		return nil, errors.New(errsStr)
	}

	videoDTOs := make([]*VideoDTO, 0, count)
	for i := 0; i < count; i++ {
		videoDTO := VideoDTO{}
		videoDTO.Video = videos[i]
		videoDTO.Player1 = pwe[i].Player

		videoDTOs = append(videoDTOs, &videoDTO)
	}

	allUserVideos := &AllUserVideosDTO{
		User:   user,
		Videos: videoDTOs,
	}

	return allUserVideos, nil

}
