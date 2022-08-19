package main

import (
	"fmt"
)

func main() {
	// var n bool = true
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	var z bool // without initializing a variable, Go sets the variable to 0 & zero equals to false
	fmt.Printf("%v, %T\n", z, z)

	/*
	 //integer types
	*/

	// var n uint16 = 42
	// fmt.Printf("%v, %T\n", n, n)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	fmt.Printf("--------------\n")

	var c int = 10
	var d int8 = 3
	fmt.Println(c + int(d))

}
