package main

import (
	"fmt"
)

func main() {

	// passing values as params (the value of username will not change from stacey to ted since we're creating/coping the username variable value which creates a new copy of the name)

	greet := "Hello"
	username := "Stacey"
	sayGreet(greet, username)
	fmt.Println(username)

	// OUTPUT:
	/*
		Hello Stacey
		Ted
		Stacey
	*/

	fmt.Println("-----------")

	// passing pointers as params (the value of name will change from stacey to ted since we're using pointers to point to the variables memory address)
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	// OUTPUT:
	/*
		Hello Stacey
		Ted
		Ted
	*/

}

func sayGreet(greet, username string) {
	fmt.Println(greet, username)
	username = "Ted"
	fmt.Println(username)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }
