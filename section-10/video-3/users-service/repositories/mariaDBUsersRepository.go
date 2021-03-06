package repositories

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MariaDBUsersRepository struct {
	db *sql.DB
}

var ErrNothingToRollback = errors.New("No receipt on db.")
var ErrReceiptAlreadyExists = errors.New("Receipt Already Exists.")

func NewMariaDBUsersRepository() *MariaDBUsersRepository {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "root:root-password@tcp(users-mariadb:3306)/users?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MariaDBUsersRepository{
		db,
	}

	return repo
}

func (repo *MariaDBUsersRepository) Close() {
	repo.db.Close()
}

func (repo *MariaDBUsersRepository) UpdateUserAccount(userID uint32, videoID uint32, ammount uint32) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`update users set account=account-? where id=?;`, ammount, userID)
	if err != nil {
		//If there is a Rollback error
		//We might want to respond in some way to that erro
		//but we are not caring about that for this example
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`insert into receipts(user_id, video_id, date) values(?,?, NOW());`, userID, videoID)
	if err != nil {
		tx.Rollback()
		//Error 1062: Duplicate entry '1-1' for key 'PRIMARY'
		if err.Error()[0:10] == "Error 1062" {
			return ErrReceiptAlreadyExists
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *MariaDBUsersRepository) RollbackUpdateUserAccount(userID uint32, videoID uint32, ammount uint32) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(`delete from receipts where user_id=? and video_id=?;`, userID, videoID)
	if err != nil {
		tx.Rollback()
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	//We need to see that the delete of the receipt
	//Actually deleted one receipt
	//If there was no receipt either there was never
	//an update for this user for this video
	//or we have already rollbacked the update account
	if rowsAffected < 1 {
		tx.Rollback()
		return ErrNothingToRollback
	} else {
		_, err = tx.Exec(`update users set account=account+? where id=?;`, ammount, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
