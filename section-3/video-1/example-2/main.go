package main

import (
	"encoding/json"
	"fmt"
)

//Marshalling

func main() {
	//Primitives
	bol, _ := json.Marshal(true)
	fmt.Println(string(bol))
	integ, _ := json.Marshal(100)
	fmt.Println(string(integ))
	flt, _ := json.Marshal(3.14159265359)
	fmt.Println(string(flt))
	str, _ := json.Marshal("A string")
	fmt.Println(string(str))

	//Slices and Maps
	slc := []string{"Sunday", "Monday", "Tuesday"}
	jsonSlc, _ := json.Marshal(slc)
	fmt.Println(string(jsonSlc))
	constants := map[string]float64{"π": 3.14159, "e": 2.71828, "√2": 1.41421, "φ": 1.61803}
	jsonConstants, _ := json.Marshal(constants)
	fmt.Println(string(jsonConstants))

	//Structs
	struct1 := struct {
		A int
		B string
		C []string
		D map[int]string
	}{
		99,
		"Hello B",
		[]string{"Axl", "Slash", "Duff", "Izzy"},
		map[int]string{1: "Uno", 2: "Dos", 3: "Tres", 4: "Cuatro", 5: "Cinco"},
	}
	jsonStruct1, _ := json.Marshal(struct1)
	fmt.Println(string(jsonStruct1))

	struct2 := struct {
		a int
		A int
	}{
		10,
		20,
	}
	jsonStruct2, _ := json.Marshal(struct2)
	fmt.Println(string(jsonStruct2))

	type substruct struct {
		F1 int
		F2 string
		F3 []int
	}

	struct3 := struct {
		E int
		F *substruct
	}{
		E: 100,
		F: &substruct{
			F1: 100,
			F2: "SubStruct",
			F3: []int{1, 2, 4, 8, 16, 32, 64, 128, 256},
		},
	}
	jsonStruct3, _ := json.Marshal(struct3)
	fmt.Println(string(jsonStruct3))

}
