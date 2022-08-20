package main

import (
	"fmt"
)

// USECASE: bitshifting
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	_ = iota + 5 // _ = iota equals to 0
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	isAdmin = 1 << iota // 1
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	fileSize := 4000000000
	fmt.Printf("%.2fGB", fileSize/GB)

	fmt.Println("-----------")

	fmt.Printf("%v", catSpecialist) // 6
	fmt.Printf("%v", dogSpecialist) // 7

	fmt.Println("-----------")

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("IS Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("IS HQ? %v\n", isHeadquarters&roles == isHeadquarters)

}
