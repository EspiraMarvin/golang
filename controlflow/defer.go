package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	/*
		defer fmt.Println("start")
		defer fmt.Println("middle")
		defer fmt.Println("end")
	*/

	a := "start"
	defer fmt.Println(a)
	a = "end"
	// from above the output of a will be start, REASON: defer takes the argument "start" at the time the defer is called and not the time the called function is executed

	res, err := http.Get("https://www.google.com/robots.txt")
	// commentsRes, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := io.ReadAll(res.Body)
	// comments, err := io.ReadAll(commentsRes.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	// fmt.Printf("%s, %T", comments, comments)

}
