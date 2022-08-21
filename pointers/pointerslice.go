package main

import (
	"fmt"
)

// reference types same as maps
func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}
