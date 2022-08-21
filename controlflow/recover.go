package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panick")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			// panic(err)
			// if you realize you can't handle recovery from a panic, rethrow the panic and defer the management of that panic higher up the callstack
			// which is the go runtime, the go runtime doesn't have a built-in func to handle panics so the program will exit
			// this will not print "end" at the main func, uncomment panic(err) to test it
		}
	}() //Defer statements doesn't take a function, it takes a function call, you'll need to invoke that function like this otherwise it work work
	panic("Something bad happened")
	fmt.Println("done panicking")

}

/*
STEPS

1. the code will run & print out "start"
2. the run the panicker() fn
3. print out "about to panic"
4. skip the defer fun
5. reach the panic keyword and stop executing
6. go back  and run the deferred func with recover() keyword
7. print out the panic error message
8. the finally print out the "end" word in main func
*/
