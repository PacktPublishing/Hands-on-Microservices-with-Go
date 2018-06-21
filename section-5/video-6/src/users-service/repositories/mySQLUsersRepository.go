package repositories

import (
	"database/sql"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-5/video-6/src/users-service/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUsersRepository struct {
	db *sql.DB
}

func NewMySQLUsersRepository() *MySQLUsersRepository {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "packt:packt@/users?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MySQLUsersRepository{
		db,
	}

	return repo
}

func (repo *MySQLUsersRepository) Close() {
	repo.db.Close()
}

func (repo *MySQLUsersRepository) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}
	row := repo.db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where username=?", username)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Added, &user.Account, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MySQLUsersRepository) GetUserByID(userID uint32) (*entities.User, error) {
	user := &entities.User{}
	row := repo.db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where id=?", userID)
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.Added, &user.Account, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MySQLUsersRepository) UpdateUser(user *entities.User) error {

	row := repo.db.QueryRow("Update users set first_name=?, last_name=?, email=?, birthdate=?, account=?, password=? from users where id=?", user.FirstName, user.LastName, user.Email, user.BirthDate, user.Account, user.Password)
	err := row.Scan()
	if err != nil {
		return err
	}
	return nil
}
