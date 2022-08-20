package main

import "fmt"

// slices are reference data types
func main() {
	a1 := [...]int{1, 2, 3}
	b1 := a1
	b1[1] = 5
	fmt.Println("a1", a1)
	fmt.Println("b1", b1)
	fmt.Println("Length:", len(a1))
	fmt.Println("Capacity:", cap(a1)) // the no of elements in the doesnt necessarily match the size of backing array bcoz a slice is a projection of the underlying arrays

	fmt.Println("------------")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements of a
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th and 6th element
	a[5] = 42
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)

	fmt.Println("------------")

	// m := make([]int, 3, 100)
	// fmt.Println("m", m)
	// fmt.Printf("Length %v\n", len(m))
	// fmt.Printf("Capacity %v\n", cap(m))

	// fmt.Println("------------")

	m := []int{}
	fmt.Println(m)
	fmt.Printf("Length %v\n", len(m))
	fmt.Printf("Capacity %v\n", cap(m))
	m = append(m, 1)
	fmt.Println(m)
	fmt.Printf("Length %v\n", len(m))
	fmt.Printf("Capacity %v\n", cap(m))
	m = append(m, 2, 3, 4, 5)
	// spread operator
	m = append(m, []int{6, 7, 8, 9}...) // this works the same way as append method
	fmt.Println(m)
	fmt.Printf("Length %v\n", len(m))
	fmt.Printf("Capacity %v\n", cap(m))

	fmt.Println("----------")

	// stack operations- remove element at beginning, remove element at the end and at the middle
	n := []int{1, 2, 3, 4, 5}
	beginning := n[1:]  // remove at the beginning
	end := n[:len(n)-1] // remove at the end
	middle := append(n[:2], n[3:]...)
	fmt.Println(beginning)
	fmt.Println(middle)
	fmt.Println(end)
	fmt.Println(n)
}
