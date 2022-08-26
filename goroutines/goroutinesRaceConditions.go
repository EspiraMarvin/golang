// detect race condition with below command
// go run -race filename
package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
