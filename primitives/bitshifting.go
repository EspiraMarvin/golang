package main

import (
	"fmt"
)

/* Bit Shifting	 */

func main() {
	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 -> 64     // bit shift a left 3 places  - we take the our no.8 we multiple to the power we're shifting with 2^3 * 2^3 = 2^6 this equals to 64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 -> 1      // bit shift a right 3 places - we take our no.8 we divide to the power we're shfiting with 2^3 / 2^3 = 2^0 this equals to 1

}
