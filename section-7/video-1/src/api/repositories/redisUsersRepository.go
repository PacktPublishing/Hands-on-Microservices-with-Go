package repositories

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/utils/appErrors"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-1/src/api/entities"
)

type UsersCacheRepository interface {
	GetUser(username string) (*entities.User, error)
	SetUser(username string, User *entities.User) error
}

type RedisUsersRepository struct {
	pool *pool.Pool
}

const keySyntax = "user-%s"

func NewRedisUsersRepository() *RedisUsersRepository {
	pool, err := pool.New("tcp", "users-cache-redis:6379", 100)
	if err != nil {
		log.Fatal(err)
	}
	return &RedisUsersRepository{
		pool: pool,
	}
}

func (repo *RedisUsersRepository) GetUser(username string) (*entities.User, error) {
	conn, err := repo.pool.Get()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf(keySyntax, username)

	res, err := conn.Cmd("GET", key).Str()
	if err != nil {
		if err == redis.ErrRespNil {
			return nil, appErrors.ErrorNotFoundOnCache
		}
		return nil, err
	}
	var User *entities.User
	err = json.Unmarshal([]byte(res), &User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (repo *RedisUsersRepository) SetUser(username string, User *entities.User) error {
	conn, err := repo.pool.Get()
	if err != nil {
		return err
	}
	key := fmt.Sprintf(keySyntax, username)

	jsonBytes, err := json.Marshal(User)
	if err != nil {
		return err
	}
	_, err = conn.Cmd("SET", key, string(jsonBytes)).Str()
	if err != nil {
		return err
	}
	return nil
}
