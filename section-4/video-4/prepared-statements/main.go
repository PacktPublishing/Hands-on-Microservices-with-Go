package main

import "database/sql"

func main() {

	username := "jeff"

	db, _ := sql.Open("mysql", "packt:packt@/users?parseTime=true")

	//Correct: Use placeholders
	_ = db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where username=?", username)

	//Wrong: Concatenate Strings
	_ = db.QueryRow("Select id, username, first_name, last_name, email, birthdate, added, account, password from users where username=" + username)

}
