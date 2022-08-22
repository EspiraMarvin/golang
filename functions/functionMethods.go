package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // g greeter gets a copy of the struct
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// OUTPUT
// Hello Go
// The new name is: Go

// with pointer

/*
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

// OUTPUT
// Hello Go
// The new name is:

*/
