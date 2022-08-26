
## GOROUTINES

### Creating goroutines

- By adding `go` keyword in front of function call
- when using anonymous functions, pass data as local variables - this prevents race conditions, although by using WaitGroups we can go around this concept
eg 
-----------
    package main
    import (
        "fmt"
        "time"
    )
    func main() {
        var msg = "Hello"
        go func(msg string) {
            fmt.Println(msg)
        }(msg)
        msg = "Goodbye"
        time.Sleep(100 * time.Millisecond)
    }
------------


### Synchronization
- Use `sync.WaitGroup` to wait for groups of goroutins to complete
- `Wait` method blocks the goroutine that is called on until the wait groups is completed
- `Done` method lets the WaitGroup know that one of the goroutine is completed its work
#### WaitGroups 
- WaitGroups synchronize goroutines together, to make sure that our main routine to wait all of our other to goroutines finish
------------
        package main
        import (
            "fmt"
            "runtime"
            "sync"
        )

        var wg = sync.WaitGroup{}
        var counter = 0
        var m = sync.RWMutex{}

        func main() {
            runtime.GOMAXPROCS(100) // no of threads to run on
            for i := 0; i < 10; i++ {
                wg.Add((2))
                m.RLock()
                go sayHello()
                m.Lock()
                go increment()
            }
            wg.Wait()
        }

        func sayHello() {
            fmt.Printf("Hello #%v\n", counter)
            m.RUnlock()
            wg.Done()
        }

        func increment() {
            counter++
            m.Unlock()
            wg.Done()

        }
-------------
#### Mutexes
- Use `sync.Mutex` and `sync.RWMutex` to protect data access

- creation
--------
        var m = sync.RWMutex{}
---------
- usage - lock and unlock
--------
            m.Lock() & m.Unlock()
            m.RLock() & m.RUnlock()
--------
### Parallelism
- By default, Go will use CPU threads equal to available cores
- We can changes no of threads with `runtime.GOMAXPROCS`
- More threads can increase performance, but too many can slow it down

-----------
package main

import (
	"fmt"
	"runtime"
)
func main() {
    runtime.GOMAXPROCS(1) // 1 core/ 1 thread of OS
    fmt.Println("threads: %v\n", runtime.GOMAXPROCS(1))
    fmt.Println("threads: %v\n", runtime.GOMAXPROCS(-1)) - to get no of threads our computer has

}
-----------

### Best practices
- Don't create goroutines in libraries
  let consumer of ur library control concurrency

- When creating a goroutine, know how it will end
  avoid subtle memory leaks 

- Check for race conditions at compile time

----------
    go run -race filename
-----------

