//  Channels can be cast into send-only or receive only channels
// send-only and receive-only channels

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // receiver-only - receives data from the channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // sender-only - sends data to the channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}
