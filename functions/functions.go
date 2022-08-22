package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello go")
	// sayMessage("its marvin")

	// for i := 0; i < 5; i++ {
	// 	sayMessage("Hello Go!", i)
	// }
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
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
