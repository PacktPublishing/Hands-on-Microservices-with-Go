package repositories

import (
	"database/sql"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-1/agents-service/src/api/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBAgentsRepository struct {
	db *sql.DB
}

func NewMariaDBAgentsRepository() *MariaDBAgentsRepository {

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
		return nil, err
	}
	return m, nil
}
