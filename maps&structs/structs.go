package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	episodes   []string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John Wick",
		episodes:  []string{},
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)     // get actor name
	fmt.Println(aDoctor.companions)    // get companions
	fmt.Println(aDoctor.companions[1]) // get second companion //Jo Grant

	// anonymous struct- we dont need to declare with type lawyer struct - type Lawyer struct {}
	bLawyer := struct{ name string }{name: "John Pertwee"}
	anotherLawyer := bLawyer
	anotherLawyer.name = "Tom Baker"
	fmt.Println(bLawyer)
	fmt.Println(anotherLawyer)

	fmt.Println("----------")

	// pointing to the same underlying data, we use address AND operator (&)
	cLawyer := struct{ name string }{name: "John Pertwee"}
	dLawyer := &cLawyer // prevents
	dLawyer.name = "Tom Baker"
	fmt.Println(cLawyer)
	fmt.Println(dLawyer)

	// EMBEDDING - Go uses COMPOSITION instead of inheritance (GO doesnt support traditional OOP principles)

}
