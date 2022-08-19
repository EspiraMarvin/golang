package main

import (
	"fmt"
)

/* floating point numbers
TYPES:
1. 32bit floating point numbers (-1.18E to 3.4E38)
2. 64bit floating point numbers (2.23E-308 to 1.8E308)
*/

func main() {
	/*	var n float64 = 3.14 // or n := 3.14
		n = 13.7e2
		n = 2.1e14

		fmt.Printf("%v, %T", n, n)
	*/

	// arithmetic operations
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fmt.Printf("--------------\n")

	// complex numbers (2 TYPES: complex64 & complex128 numbers)

	// var n complex128 = 1 + 2i // (2i) 2 imaginary numbers
	var n complex128 = complex(5, 12) // (2i) 2 first no. is the real no and the 2nd no is the imaginary number
	fmt.Printf("%v, %T\n", n, n)

	fmt.Printf("%v, %T\n", real(n), real(n)) // pulls out the real numbers
	fmt.Printf("%v, %T\n", imag(n), imag(n)) // pulls out the imaginary numbers

}
