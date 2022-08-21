package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ { // looping by adding 1
		fmt.Println(i)
	}

	fmt.Println("-------------")

	for i := 0; i < 5; i = i + 2 { // looping by adding 2
		fmt.Println(i)
	}

	fmt.Println("-------------")

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 { // looping 2 variables at once
		fmt.Println(i, j)
	}

	fmt.Println("-------------")

	// without an initializer in the for loop
	i := 0
	for ; i < 5; i++ { // Here the variable i is not scoped to the for-loop unlike the ones above, so we can print the value of i after the loop
		fmt.Println(i)
	}
	fmt.Println("value of i after for-loop", i) // prints 5 because 5 is the 1st value that fails the test

	fmt.Println("-------------")

	/// we also don't need an incrementor value at the begin of statement, we can put it inside the loop
	// SAME AS DO-WHILE LOOP
	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}

	fmt.Println("-------------")

	// INFINITE FOR LOOP // runs undetermined no of times
	m := 0
	for {
		fmt.Println(m)
		m++
		if m == 5 {
			break
		}
	}

	fmt.Println("-------------")

	// with CONTINUE keyword
	// eg. to get odd numbers
	for d := 0; d < 10; d++ {
		if d%2 == 0 { // means if no is even don't print continue to the next no
			continue
		}
		fmt.Println("odd numbers betwween 0 and 10", d)
	}

	fmt.Println("-------------")

	for even := 0; even < 10; even++ {
		if even%2 != 0 { // means if no is even don't print continue to the next no
			continue
		}
		fmt.Println("even numbers between 0 and 10", even)
	}

	fmt.Println("-------------")

	// NESTED LOOPS and breaking out of the outer loop
	// eg below loop through i & j and find the product of i*j id the product is >= 3 break out of the loop

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop // without Loop keyword indicating the outer loop, the outer loop will run once (& show 4) even after breaking the inner loop
			}
		}
	}

	fmt.Println("-------------")

	// COLLECTIONS WITH FOR-LOOPS
	// Eg. slices

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println("-------------")

	// Eg. Arrays
	ar := [3]int{1, 2, 3}
	for k, v := range ar {
		fmt.Println(k, v)
	}

	fmt.Println("-------------")

	// Eg. maps
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	fmt.Println("-------------")

	// print out all the keys only
	for k := range statePopulations {
		fmt.Println("keys only", k)
	}

	fmt.Println("-------------")

	// print out all the values only
	for _, v := range statePopulations {
		fmt.Println("values only", v)
	}

	fmt.Println("-------------")

	// Eg. strings
	text := "hello Go!"
	for k, v := range text {
		fmt.Println(k, string(v))
	}

}
