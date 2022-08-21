package main

import (
	"fmt"
)

// EMBEDDING - Go uses COMPOSITION instead of inheritance (GO doesnt support traditional OOP principles)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // embedding Animal struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// 	b := Bird{}
	// 	b.Name = "Emu"
	// 	b.Origin = "Australia"
	// 	b.SpeedKPH = 48
	// 	b.CanFly = false
	// 	fmt.Println(b)
	// 	fmt.Println(b.Name)

	// OR

	b := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b.Name)

}

// NOTE:
// Bird is not of type Animal, it just inherts properties of Animal through composition.
// It just says Bird has Animal characteristics, but its not the same thing as an Animal and both Bird & Animal cannot be used interchangeably, unless we use interfaces
