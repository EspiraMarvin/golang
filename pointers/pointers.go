package main

import (
	"fmt"
)

func main() {
	a := 42
	b := a
	fmt.Println(a, b) //OUTPUT a = 42, b = 42 - b copies the value of a, b is not pointed to the same memory as a
	a = 27
	fmt.Println(a, b) //OUTPUT a = 27, b = 42 - b remains 42 while a becomes 27 this proves b is not pointed to the same memory as a

	// WE CAN CHANGE THIS BEHAVIOUR USING POINTERS TO make b point to the same memory as a here
	// here we use c & d, we make d point to the same memory as c,
	// using address and operator(&) to get the memory address where the variable data is stored
	// and dereferencing operators(*) to get the data store in the memory address

	var c int = 42
	var d *int = &c
	fmt.Println(c, *d) //OUTPUT c = 42, d = 42
	c = 27
	fmt.Println(c, *d) //OUTPUT c = 27, d = 27
	*d = 14
	fmt.Println(c, *d) //OUTPUT c = 14, d = 14

}
