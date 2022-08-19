package main

import (
	"fmt"
)

func main() {

	/* TEXT TYPE (2Types)
	1. string type - A string in Go stands for any UTF-8 character - it can't encode any type of characters available
	2. Rune type (UTF-32 characters) / though any x-cter less than UTF32 xcters can be a Rune too in Go - declared with single quotes as opposed to double quotes for strings
	*/

	// strings in go are actually aliases for bytes
	// strings are immutable, you can't manipulate a value in a string

	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))

	// string concatenation
	s1 := "this is a string"
	s2 := "this is a string"
	fmt.Printf("%v, %T\n", s1+" "+s2, s+s1+s2)

	// convert strings to collection of bytes/ a slice of bytes
	str := "this is a string"
	b := []byte(str)
	fmt.Printf("%v, %T\n", b, b)

	// RUNE
	var r rune = 'a' // r := 'a'
	fmt.Printf("%v, %T\n", r, r)

}
