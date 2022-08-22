package main

import (
	"fmt"
)

func main() {
	sum("The sum is", 1, 2, 3, 4, 5)

	fmt.Println("---------")

	s1 := sum1(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s1)

	fmt.Println("---------")

	s2 := sum1(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s2)

	fmt.Println("---------")

	d, err := divide(5.0, 0.0) // cannot divide by zero
	// d, err := divide(5.0, 3.0) // 1.666666667
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func sum(msg string, values ...int) { // the 3 dots tells Go, to take all the last arguments passed-in & wrap them into a slice that has the name of the variable 'values'
	fmt.Println(values) // here msg is the 1st parameter and values is variatic parameter, which is always passed last
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

//  another syntax with return keyword

func sum1(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

//  another syntax with named return value // not widely used, in long functions it can be confusing

func sum2(values ...int) (result int) { // (result int) is a named return value
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// another syntax - you can return multiple return values from a function

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
