// type switch example

package main

import (
	"fmt"
)

func main() {
	var i interface{} = 0 // 0 -integer // "0"-string // true - "I don't know what it is"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("I don't know what it is")
	}
}
