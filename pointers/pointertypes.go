package main

import "fmt"

func main() {
	// var ms *myStruct
	// ms = &myStruct{foo: 42}
	// fmt.Println(ms)

	// dereferencing a pointer
	// var ms *myStruct
	// ms = new(myStruct)
	// (*ms).foo = 42
	// fmt.Println((*ms).foo)

	// dereferencing a pointer better syntax
	var ms *myStruct
	ms = new(myStruct)
	fmt.Println(ms.foo)
	ms.foo = 42
	fmt.Println(ms)     // with pointer
	fmt.Println(ms.foo) // underlying value
}

type myStruct struct {
	foo int
}
