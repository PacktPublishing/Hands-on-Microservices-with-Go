package repositories

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDBVideosRepository struct {
	db *sql.DB
}

var ErrBoughtVideoAlreadyExists = errors.New("Bought Video Already Exists.")

func NewMariaDBVideosRepository() *MariaDBVideosRepository {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "root:root-password@tcp(videos-mariadb:3306)/videos?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	repo := &MariaDBVideosRepository{
		db,
	}

	return repo
}

func (repo *MariaDBVideosRepository) Close() {
	repo.db.Close()
}

func (repo *MariaDBVideosRepository) InsertBoughtVideo(videoID uint32, userID uint32) error {
	rows, err := repo.db.Query(`Insert into bought_videos(video_id, user_id) values(?, ?);`, videoID, userID)
	if err != nil {
		//Error 1062: Duplicate entry '1-1' for key 'PRIMARY'
		if err.Error()[0:10] == "Error 1062" {
			return ErrBoughtVideoAlreadyExists
		}
		return err
	}
	defer rows.Close()
	return nil
}

func (repo *MariaDBVideosRepository) RollBackBoughtVideo(videoID uint32, userID uint32) error {
	rows, err := repo.db.Query(`delete from bought_videos where video_id=? and user_id=?;`, videoID, userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
