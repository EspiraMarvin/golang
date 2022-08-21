package main

import (
	"fmt"
)

func main() {
	// example 1 by compiler
	a, b := 1, 0
	ans := a / b // Because in Go we cant divide 1 by 0 the compiler throws a panic
	fmt.Println(ans)

	// example 2 throw panic yourself
	fmt.Println("Start")
	panic("something bad happened") // something bad happened printed out
	fmt.Println("end")

}
