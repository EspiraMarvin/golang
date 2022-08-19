package main

import (
	"fmt"
)

func main() {

	// bit operators (AND, OR, Exclusive OR, AND-NOT/Bitclear operators respectively)

	m := 10             // 1010
	n := 3              // 0011
	fmt.Println(m & n)  // 0010 = 2
	fmt.Println(m | n)  // 1011 = 11
	fmt.Println(m ^ n)  // 1001 = 9          Exclusive OR means either one has a bit set but not both
	fmt.Println(m &^ n) // 1000 = 8          AND-NOT - m &^ n is like literally doing m & (^n). You're AND-ing the first operand with the negation of the second operand.

	/*
		1010  AND-NOT - m &^ n is like literally doing m & (^n). You're AND-ing the first operand with the negation of the second operand.
		0011
		----
		1000  = 8

	*/
}
