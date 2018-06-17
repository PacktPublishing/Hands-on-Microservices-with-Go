package repositories

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-4/video-3/src/sessions-service/entities"
	"github.com/mediocregopher/radix.v2/pool"
)

var ErrRespIsNil = errors.New("Response is Nil")

type RedisSessionsRepository struct {
	pool *pool.Pool
}

func NewRedisSessionsRepository() *RedisSessionsRepository {
	pool, err := pool.New("tcp", "localhost:6379", 100)
	if err != nil {
		log.Fatal(err)
	}
	return &RedisSessionsRepository{
		pool: pool,
	}
}

func (repo *RedisSessionsRepository) GetSession(key string) (*entities.Session, error) {
	conn, err := repo.pool.Get()
	if err != nil {
		return nil, err
	}

	resp, err := conn.Cmd("GET", key).Str()
	if resp == "" {
		return nil, ErrRespIsNil
	}
	if err != nil {
		return nil, err
	}
	var session *entities.Session
	err = json.Unmarshal([]byte(resp), &session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (repo *RedisSessionsRepository) SetSession(key string, session *entities.Session) error {
	jsonBytes, err := json.Marshal(session)
	if err != nil {
		return err
	}
	conn, err := repo.pool.Get()
	if err != nil {
		return err
	}
	resp := conn.Cmd("SET", key, string(jsonBytes))
	if resp.Err != nil {
		return resp.Err
	}
	//Expire in 12 Hours
	resp = conn.Cmd("EXPIRE", key, 43200)
	if resp.Err != nil {
		return resp.Err
	}
	return nil
}
