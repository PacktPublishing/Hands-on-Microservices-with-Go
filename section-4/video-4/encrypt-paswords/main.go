package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "CorrectPassword"
	wrongPassword := "WrongPassword"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Hashed password is: " + string(hashedPass))

	//Verify password is correct
	//will fail cause we are using wrong password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(wrongPassword))
	if err != nil {
		fmt.Println(wrongPassword + " is the Wrong Password")
	} else {
		fmt.Println(wrongPassword + " is the Correct Password")
	}

	//Verify password is correct
	//will not fail cause we are using correct password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		fmt.Println(password + " is the Wrong Password")
	} else {
		fmt.Println(password + " is the Correct Password")
	}

}
