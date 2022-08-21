package main

import (
	"fmt"
)

func main() {
	// example 2 throw panic yourself
	fmt.Println("Start")
	defer fmt.Println("this was deferred") // this will be printed out before panic
	panic("something bad happened")        // something bad happened printed out
	fmt.Println("end")

}
