package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

func main() {

	type MyStruct struct {
		Int       int     `xml:"integer,attr"`
		Float     float64 `xml:"float"`
		Array     []int   `xml:"array"`
		FirstName string  `xml:"name>firstName"`
		LastName  string  `xml:"name>lastName"`
		Substruct struct {
			Attr1 string   `xml:"attr1,attr"`
			Attr2 string   `xml:"attr2,attr"`
			Int   int      `xml:"integer>sub"`
			Array []string `xml:"array>value"`
		} `xml:"substruct"`
	}

	myStruct := &MyStruct{
		Int:       20,
		Float:     1.234,
		Array:     []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31},
		FirstName: "Bruce",
		LastName:  "Banner",
		Substruct: struct {
			Attr1 string   `xml:"attr1,attr"`
			Attr2 string   `xml:"attr2,attr"`
			Int   int      `xml:"integer>sub"`
			Array []string `xml:"array>value"`
		}{
			Attr1: "This is Attribute 1 on substruct",
			Int:   67,
			Array: []string{"un", "deux", "trois", "quatre"},
			Attr2: "This is Attribute 2 on integer",
		},
	}
	xmlBytes, err := xml.MarshalIndent(myStruct, "", "      ")
	if err != nil {
		log.Println("Error Marshalling: " + err.Error())
	}
	fmt.Println(string(xmlBytes))
}
