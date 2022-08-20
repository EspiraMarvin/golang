package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	var students [3]string
	fmt.Printf("Grades: %v\n", grades)

	students[0] = "marvin"
	students[1] = "arnold"
	students[2] = "robertson"

	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Students #1: %v\n", students[1])

	// get size of arrays

	fmt.Printf("Students length: %v\n", len(students))

	// arrays of arrays
	var indentityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// OR
	var indentityMatrix2 [3][3]int
	indentityMatrix2[0] = [3]int{1, 0, 0}
	indentityMatrix2[1] = [3]int{0, 1, 0}
	indentityMatrix2[2] = [3]int{0, 0, 1}

	fmt.Println("indentityMatrix", indentityMatrix)
	fmt.Println("indentityMatrix2", indentityMatrix2)

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println("a", a)
	fmt.Println("b", b)

	fmt.Println("------------")

	// using pointers- point d to c, so if be changes the value of array c also changes

	c := [...]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println("c", c)
	fmt.Println("d", d)

	a1 := [...]int{1, 2, 3}
	b1 := a1
	b1[1] = 5
	fmt.Println("a1", a1)
	fmt.Println("b1", b1)
	fmt.Println("Length:", len(a1))
	fmt.Println("Capacity:", cap(a1)) // the no of elements in the doesnt necessarily match the size of backing array bcoz a slice is a projection of the underlying arrays

	o := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := o[:]   // slice of all elements of a
	q := o[3:]  // slice from 4th element to end
	r := o[:6]  // slice first 6 elements
	s := o[3:6] // slice the 4th, 5th and 6th element
	o[5] = 42
	fmt.Println(p)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
}
