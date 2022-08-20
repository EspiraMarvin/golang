package main

import (
	"fmt"
)

// const shadowing
const m int16 = 27

// enumerated constants
const (
	i = iota // or const i = iota const j = iota  const k = iota
	j
	k
)

const (
	a2 = iota
)

func main() {
	const a int = 42
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	fmt.Println("-----------")

	const m int = 14
	fmt.Printf("%v, %T\n", m, m)

	fmt.Println("-----------")

	const x int = 42
	const y int = 27
	fmt.Printf("%v, %T\n", x+y, x+y)

	fmt.Println("-----------")

	// Untype constants
	const s = 42
	const t int16 = 27
	fmt.Printf("%v, %T\n", s+t, s+t) // implicit conversion of types, s is implicitly converted to int16

	// enumerated constants
	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
	fmt.Printf("%v\n", a2)

}
