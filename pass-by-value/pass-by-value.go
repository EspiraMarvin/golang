/**
Go is exclusively a pass-by-value language,
meaning, each function receives a copy of the value of what you passed in. No exceptions.
You can pass pointers too (you're technically passing the value of the pointer- the address)
And with Go's strong typing, you'll get type checking on the undelying pointer.
*/

package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func DoubleHeight(rect Rectangle) {
	rect.Height = rect.Height * 2

}

func DoubleHeightWithPointer(rect *Rectangle) {
	rect.Height = rect.Height * 2

}

func main() {
	rect := Rectangle{
		Width:  10,
		Height: 3,
	}

	// this won't actually modify rect
	DoubleHeight(rect)

	// this modifies rect
	DoubleHeightWithPointer(&rect)

	fmt.Println("Width", rect.Width)
	fmt.Println("Height", rect.Height)

	fmt.Println("Width modified", rect.Width)
	fmt.Println("Height modified", rect.Height)
}
