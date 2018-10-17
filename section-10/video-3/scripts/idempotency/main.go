package main

import "fmt"

func main() {
	i := 10

	add5(&i)
	fmt.Println(i)
	add5(&i)
	fmt.Println(i)

	assign5(&i)
	fmt.Println(i)
	assign5(&i)
	fmt.Println(i)
}

//Not idempotent
func add5(i *int) {
	*i += 5
}

//Idempotent
func assign5(i *int) {
	*i = 5
}
