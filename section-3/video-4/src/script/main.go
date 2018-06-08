package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type player struct {
	playerID    string
	firstName   string
	lastName    string
	hand        string
	birthDate   string
	countryCode string
}

type ranking struct {
	rankingDate string
	ranking     string
	playerID    string
	points      string
}

func main() {
	connStr := "postgres://packt:packt@localhost/wta?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	_ = db.QueryRow(`DROP TABLE IF EXISTS rankings;`)
	_ = db.QueryRow(`DROP TABLE IF EXISTS players;`)
	_ = db.QueryRow(`CREATE TABLE players (
			  player_id SERIAL PRIMARY KEY,
			  first_name VARCHAR(255),
			  last_name VARCHAR(255),
			  isRightHanded BOOLEAN,
			  birth_date DATE,
			  country_code char(3)
			);`)
	_ = db.QueryRow(`CREATE TABLE rankings (
			  player_id INTEGER,
			  ranking_date DATE,
			  ranking INTEGER CHECK(ranking > 0),
			  ranking_points FLOAT,
			  constraint fk_rankings_player_id
				 foreign key (player_id)
				 REFERENCES players (player_id)
			);`)

	//Insert Players
	csvFile, _ := os.Open("data-sources/players.csv")
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var players []player
	//First line
	_, _ = reader.Read()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		players = append(players, player{
			playerID:    line[0],
			firstName:   line[1],
			lastName:    line[2],
			hand:        line[3],
			birthDate:   line[4],
			countryCode: line[5],
		})
	}
	fmt.Println("Number of Players: ", len(players))

	for ind, player := range players {
		//fmt.Println("Inserting Player #", ind)

		isRightHanded := true
		if player.hand == "L" {
			isRightHanded = false
		}
		if player.birthDate == "" {
			continue
		}
		year, err := strconv.Atoi(player.birthDate[0:4])
		if err != nil {
			log.Fatal(err.Error())
		}
		month, err := strconv.Atoi(player.birthDate[4:6])
		if err != nil {
			log.Fatal(err.Error())
		}
		day, err := strconv.Atoi(player.birthDate[6:8])
		if err != nil {
			log.Fatal(err.Error())
		}
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		playerId, err := strconv.Atoi(player.playerID)
		if err != nil {
			log.Fatal(err.Error())
		}
		row := db.QueryRow(`Insert into players (player_id, first_name, last_name, isRightHanded, birth_date, country_code) values($1, $2, $3, $4, $5, $6) returning player_id;`, playerId, player.firstName, player.lastName, isRightHanded, date, player.countryCode)
		var i int
		err = row.Scan(&i)
		if err != nil {
			fmt.Printf("%d :: %s\n", ind, err.Error())
		}
		if i != playerId {
			fmt.Printf("%d :: bad value for i: got %d, want %d\n", ind, i, playerId)
		}
	}

	//Insert Rankings
	csvFile2, _ := os.Open("data-sources/rankings.csv")
	defer csvFile2.Close()

	reader = csv.NewReader(bufio.NewReader(csvFile2))
	var rankings []ranking
	//First line
	_, _ = reader.Read()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		rankings = append(rankings, ranking{
			rankingDate: line[0],
			ranking:     line[1],
			playerID:    line[2],
			points:      line[3],
		})
	}
	fmt.Println("Number of Rankings: ", len(rankings))

	for ind, ranking := range rankings {
		//		if ind%10000 == 0 {
		// fmt.Println("Inserting Ranking #", ind)
		//		}

		if ranking.rankingDate == "" {
			continue
		}
		year, err := strconv.Atoi(ranking.rankingDate[0:4])
		if err != nil {
			log.Fatal(err.Error())
		}
		month, err := strconv.Atoi(ranking.rankingDate[4:6])
		if err != nil {
			log.Fatal(err.Error())
		}
		day, err := strconv.Atoi(ranking.rankingDate[6:8])
		if err != nil {
			log.Fatal(err.Error())
		}
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		playerIdTmp, err := strconv.ParseFloat(ranking.playerID, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		playerId := int(playerIdTmp)
		rank, err := strconv.Atoi(ranking.ranking)
		if err != nil {
			log.Fatal(err.Error())
		}
		points, err := strconv.ParseFloat(ranking.points, 64)
		row := db.QueryRow(`Insert into rankings (player_id, ranking_date, ranking, ranking_points) values($1, $2, $3, $4) returning player_id;`, playerId, date, rank, points)
		var i int
		err = row.Scan(&i)
		if err != nil {
			fmt.Printf("%d :: %s\n", ind, err.Error())
		}
		if i != playerId {
			fmt.Printf("%d :: bad value for i: got %d, want %d\n", ind, i, playerId)
		}

	}
}
