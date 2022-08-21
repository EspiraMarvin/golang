package main

import (
	"fmt"
	"reflect"
)

// USECASE: data/form input validation to pass data validation messages
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
