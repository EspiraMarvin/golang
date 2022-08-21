package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{ // key type string // value of type int
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371   // add key with population value
	fmt.Println(statePopulations["Georgia"]) // get value of Georgia
	fmt.Println(statePopulations["Ohio"])    // get value of Ohio
	fmt.Println("Final statePopulations ", statePopulations)

	// delete from a map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations) // delete Georgia

	// okay syntax to check if the value is deleted/available in a map
	popGerorgia, ok := statePopulations["Georgia"]
	fmt.Println(popGerorgia, ok) // returns 0 & false since it has been deleted

	// another example
	popOhio, ok := statePopulations["Ohio"]
	fmt.Println(popOhio, ok) // returns Ohio pop & true since its present in the map

	// find no of elements in the map
	fmt.Println("statePopulations length", len(statePopulations))

	// manipulation a map may over effect on any other place the map is used
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
