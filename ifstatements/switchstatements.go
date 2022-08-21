package main

import (
	"fmt"
)

func main() {
	switch 3 {
	case 1, 5, 10: // falling through/ testings multiple cases
		fmt.Println("one, five or ten")
	case 2, 6, 7:
		fmt.Println("two, six or seven")
	default:
		fmt.Println("not one or two")
	}

	// another example with tags/ another name for a variable eg. i
	switch i := 2 + 3; i {
	case 1, 5, 10: // falling through/ testings multiple cases
		fmt.Println("one, five or ten")
	case 2, 6, 7:
		fmt.Println("two, six or seven")
	default:
		fmt.Println("not one or two")
	}

	// another example with tagless syntax
	a := 10
	switch {
	case a <= 10:
		fmt.Println("less than or equal to ten")
		// fallthrough
	case a <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// type switch
	var i interface{} = 1 // integer
	// var i interface{} = 1.58   // float64
	// var i interface{} = "wick" // string
	// var i interface{} = [3]int{}   // array of int with length 3
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break // breaking out early
		fmt.Println("This will not print")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is an array of int with len 3")
	default:
		fmt.Println("i is another type")

	}

}
