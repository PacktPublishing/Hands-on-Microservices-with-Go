package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MyStruct struct {
	Int       int     `json:"integer"`
	Float     float64 `json:"float"`
	Array     []int   `json:"array"`
	Substruct struct {
		Int   int   `json:"integer"`
		Array []int `json:"array"`
	} `json:"substruct"`
	Map map[string]interface{} `json:"substruct2"`
}

func main() {
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

	var myStruct *MyStruct
	err := json.Unmarshal(jsonData, &myStruct)
	if err != nil {
		log.Println("Error Unmarshalling: " + err.Error())
	}

	fmt.Printf("%+v\n", myStruct)
}
