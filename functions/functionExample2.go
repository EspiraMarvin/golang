package main

import (
	"fmt"
)

func main() {
	// anonymous function
	func() {
		fmt.Println("Hello GO")
	}() // self invoke an anonymous function

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}
