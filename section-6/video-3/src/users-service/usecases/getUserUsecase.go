package usecases

import (
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/entities"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-6/video-3/src/users-service/repositories"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	cacheHits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_hits",
			Help: "Number of cache hits.",
		},
		[]string{"cache"},
	)
	dbHits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_hits",
			Help: "Number of db gets.",
		},
		[]string{"db"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cacheHits)
	prometheus.MustRegister(dbHits)
}

type GetUserUsecase struct {
	CacheRepo *repositories.RedisUsersRepository
	Repo      *repositories.MySQLUsersRepository
}

func (uc *GetUserUsecase) GetUser(username string) (*entities.User, error) {
	//Look In Cache
	user, err := uc.CacheRepo.GetUser(username)
	//CHECK FOR NOT IN CACHE
	if err == nil {
		cacheHits.With(prometheus.Labels{"cache": "redis"}).Inc()
		//It was in cache, return it
		return user, nil
	}

	//Not in cache
	user, err = uc.Repo.GetUserByUsername(username)
	//CHECK FOR NOT IN DB
	if err != nil {
		return nil, err
	}

	dbHits.With(prometheus.Labels{"db": "mariaDB"}).Inc()

	//Update cache for future requets
	err = uc.CacheRepo.SetUser(username, user)
	//Ignore Error?

	return user, nil
}
