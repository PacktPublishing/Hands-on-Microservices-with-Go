package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/wta-service/src/entities"
	_ "github.com/lib/pq"
)

type PsqlRepo struct {
	db *sql.DB
}

func NewWTARepository() *PsqlRepo {
	connStr := "postgres://packt:packt@wta-psql/wta?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	repo := &PsqlRepo{}
	repo.db = db
	return repo
}

func (repo *PsqlRepo) Close() {
	repo.db.Close()
}

func (repo *PsqlRepo) GetPlayer(PlayerID uint32) (*entities.Player, error) {
	var p entities.Player

	row := repo.db.QueryRow(`select players.player_id, players.first_name, players.last_name, players.isRightHanded, players.birth_date, players.country_code from players where players.player_id=$1;`, PlayerID)

	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.IsRightHanded, &p.BirthDate, &p.CountryCode)
	if err != nil {
		return nil, errors.New("[GetPlayer] Error on Query to Db:" + err.Error())
	}

	return &p, nil
}

func (repo *PsqlRepo) GetMatch(MatchID uint32) (*entities.Match, error) {
	var m entities.Match

	row := repo.db.QueryRow(`select matches.id, matches.winner_id, matches.loser_id, matches.match_date from matches where id=$1;`, MatchID)
	err := row.Scan(&m.ID, &m.WinnerID, &m.LoserID, &m.Date)
	if err != nil {
		return nil, errors.New("[GetMatch] Error on Query to Db:" + err.Error())
	}

	return &m, nil
}
