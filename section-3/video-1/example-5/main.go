package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//Unknown type
	var unk interface{}

	jsonData := []byte(`
	{
		"integer":18,
		"float":9.998,
		"string":"Hello World",
		"array":[1,2,3,4,5,6,7,8,9],
		"substruct":{
			"integer":98,
			"array":[9,8,7,6,5,4,3,2,1]
		},
		"substruct2":{
			"integer":98,
			"array":[9,8,7,6,5,4,3,2,1]
		}		
	}`)
	err := json.Unmarshal(jsonData, &unk)
	if err != nil {
		log.Println("Error Unmarshalling: " + err.Error())
	}

	fmt.Printf("%+v\n", unk)
}
