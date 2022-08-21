package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}

	// initializer syntax
	statePopulations := map[string]int{ // key type string // value of type int
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	//  initializer  - generates a boolean eg pop, ok := statePopulations["Florida"] returns true - ok == true & pop is population no.
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}

	// simple number guessing game
	number := 50
	guess := 105

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	}

}

// short circuiting - if any part of an OR returns true Go doesn't execute any remaining part of the OR condition
/* eg.
 if guess < 1 || returnTrue() || guess > 100 {
	fmt.Println("The guess must be between 1 and 100!")
 }

 if guess is less than 1 say -5, Go stops executing the remaining OR conditions since it has gotten a one true condition

*/

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
