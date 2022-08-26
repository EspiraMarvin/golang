/* Use buffered channels when send and receiver have assymmetric loading /if we can generate messages faster than we can receive them in channels */

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // buffer channel with 50 to tell the channel to store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
