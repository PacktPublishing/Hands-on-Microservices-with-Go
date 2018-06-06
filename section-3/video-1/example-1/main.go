package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//Marshal -- take an element and transform it into a json formatted []byte
	ex1, err := json.Marshal([]int{1, 2, 3})
	if err != nil {
		log.Println("Error Marshalling: " + err.Error())
	}
	fmt.Println(string(ex1))

	//Unmarshall -- take a json formatted []byte and return an element (go object, map, primitive, etc)
	var ex1Dat []int
	err = json.Unmarshal(ex1, &ex1Dat)
	if err != nil {
		log.Println("Error Unmarshalling: " + err.Error())
	}
	fmt.Println(ex1Dat[1])

}
