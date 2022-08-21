package main

import (
	"fmt"
)

// reference types same as slices
func main() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "quix"
	fmt.Println(a, b)
}
