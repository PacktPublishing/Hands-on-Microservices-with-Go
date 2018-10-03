package repositories

import (
	"database/sql"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/videos-service/src/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBVideosRepository struct {
	db *sql.DB
}

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

func (repo *MariaDBVideosRepository) GetAllUserVideos(userID uint32) ([]*entities.Video, error) {
	rows, err := repo.db.Query(`select videos.id, videos.match_id, videos.player1_id, videos.player2_id, videos.duration from bought_videos, videos where bought_videos.video_id = videos.id and bought_videos.user_id = ?;`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	videos := make([]*entities.Video, 0, 4)

	for rows.Next() {
		var video entities.Video
		if err := rows.Scan(&video.ID, &video.MatchID, &video.Player1ID, &video.Player2ID, &video.Duration); err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}
	return videos, nil
}

func (repo *MariaDBVideosRepository) GetVideoByVideoID(videoID uint32) (*entities.Video, error) {
	video := &entities.Video{}
	row := repo.db.QueryRow(`select id, match_id, player1_id, player2_id, duration, price from videos where videos_id = ?;`)
	err := row.Scan(&video.ID, &video.MatchID, &video.Player1ID, &video.Player2ID, &video.Duration, &video.Price)
	if err != nil {
		return nil, err
	}
	return video, nil
}
