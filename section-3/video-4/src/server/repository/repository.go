package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-3/video-4/src/server/entities"
	_ "github.com/lib/pq"
)

type PsqlRepo struct {
	db *sql.DB
}

func NewWTARepository() *PsqlRepo {
	connStr := "postgres://packt:packt@localhost/wta?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	repo := &PsqlRepo{}
	repo.db = db
	return repo
}

func (repo *PsqlRepo) CloseWTARepository() {
	repo.db.Close()
}

func (repo *PsqlRepo) GetPlayerPlusHighestRanking(playerID uint32) (*entities.PlayerWithRanking, error) {
	row := repo.db.QueryRow(`select players.player_id, players.first_name, players.last_name, players.isRightHanded, players.birth_date, players.country_code, rankings.ranking_date, rankings.ranking, rankings.ranking_points from players, rankings where players.player_id = rankings.player_id and players.player_id = $1 order by rankings.ranking, rankings.ranking_date limit 1;`, playerID)

	p := entities.PlayerWithRanking{}

	//, &p.RankingDate, &p.RankingNumber, &p.RankingPoints

	//var dat time.Time
	//var inte uint64
	//var fl float64

	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.IsRightHanded, &p.BirthDate, &p.CountryCode, &p.RankingDate, &p.RankingNumber, &p.RankingPoints)
	if err != nil {
		return nil, errors.New("[GetPlayerPlusHighestRanking] Error on Query to Db:" + err.Error())
	}

	return &p, nil
}

func (repo *PsqlRepo) GetRankingsByPlayerID(PlayerID uint32) ([]*entities.Ranking, error) {
	var results []*entities.Ranking

	rows, err := repo.db.Query(`Select rankings.ranking_date, rankings.ranking, rankings.ranking_points from rankings where rankings.player_id=$1 order by rankings.ranking_date asc;`, PlayerID)
	if err != nil {
		return nil, errors.New("[GetRankingsByPlayerID] Error on Query to Db:" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		r := &entities.Ranking{}

		err := rows.Scan(&r.RankingDate, &r.RankingNumber, &r.RankingPoints)
		if err != nil {
			return nil, errors.New("[GetRankingsByPlayerID] Error on Query to Db:" + err.Error())
		}

		results = append(results, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, errors.New("[GetRankingsByPlayerID] Error on Query to Db:" + err.Error())
	}
	return results, nil
}

func (repo *PsqlRepo) GetPlayer(PlayerID uint32) (*entities.Player, error) {
	var p entities.Player

	row := repo.db.QueryRow(`select players.player_id, players.first_name, players.last_name, players.isRightHanded, players.birth_date, players.country_code from players where players.player_id=$1;`, PlayerID)

	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.IsRightHanded, &p.BirthDate, &p.CountryCode)
	if err != nil {
		return nil, errors.New("[GetPlayerPlusHighestRanking] Error on Query to Db:" + err.Error())
	}

	return &p, nil
}
