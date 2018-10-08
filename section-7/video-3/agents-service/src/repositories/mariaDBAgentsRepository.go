package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/agents-service/src/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBAgentsRepository struct {
	db *sql.DB
}

func NewMariaDBAgentsRepository() *MariaDBAgentsRepository {

	// Create the database handle, confirm driver is present
	//	db, err := sql.Open("mysql", "packt:packt@tcp(agents-mariadb:3306)/users?parseTime=true")
	db, err := sql.Open("mysql", "root:root-password@tcp(agents-mariadb:3306)/managers?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MariaDBAgentsRepository{
		db,
	}

	return repo
}

func (repo *MariaDBAgentsRepository) Close() {
	repo.db.Close()
}

func (repo *MariaDBAgentsRepository) InsertAgentPlayer(agentID uint32, playerID uint32) error {
	rows, err := repo.db.Query("Insert into manager_player(manager_id, player_id) values(?, ?);", agentID, playerID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (repo *MariaDBAgentsRepository) GetAgentByID(agentID uint32) (*entities.Agent, error) {
	m := &entities.Agent{}
	row := repo.db.QueryRow("Select id, manager, account from manager where id=?", agentID)
	err := row.Scan(&m.ID, &m.Name, &m.Account)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return m, nil
}

func (repo *MariaDBAgentsRepository) GetAgentPlayers(agentID uint32) ([]uint32, error) {
	playerIDs := make([]uint32, 0, 4)
	rows, err := repo.db.Query("Select player_id from manager_player where manager_id=?", agentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var playerID uint32
		if err := rows.Scan(&playerID); err != nil {
			return nil, err
		}
		playerIDs = append(playerIDs, playerID)
	}

	return playerIDs, nil
}
