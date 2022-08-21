### Control flow

- GO applications runs from the beginning to the bottom of the file

#### Defer - invoking a fun but delay its execution to some future point in time
- useful to group "open" and "close" functions together (be careful using in loops to prevent opening and leaving to many resources)
- Defer statements runs in  LIFO order(Last In First Out)
- Defer statements doesn't take a function, it takes a function call, you'll need to invoke that function otherwise it won't work
- Arguments evaluated at time defer is executed, not at tome of called function execution 
eg.
--------
        package main

        import (
            "fmt"
            "io"
            "log"
            "net/http"
        )

        func main() {
            a := "start"
            defer fmt.Println(a)
            a = "end"
            // from above the output of a will be start, REASON: defer takes the argument "start" at the time the defer is called and not the time the called function is executed
        }

        OUTPUT: // start
    --------

    #### Panic - Occurs when program cannot continue at all/How a Go app can enter a state where it can no longer continue to run, & how the go runtime can trigger that and also how we can trigger that on our own
    ###### Design:
    - Panic happens after the deferred statements are executed
    ###### Rules/ bes t usage:
    - Don't use when file can't be opened, unles its critical
    - Use for unrecoverable events - cannot obtain TCP port for web server
    eg
    -----------
        package main
        import "net/http"
        func main() {
            // Use for unrecoverable events - cannot obtain TCP port for web server
            // generates a panic when port 8080 is already in use
            http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Write([]byte("Hello Go!"))
            })
            err := http.ListenAndServe(":8080", nil)
            if err != nil {
                panic(err.Error())
            }  
        }
    -----------
    - Functions will stop executing when it reaches a panic() func
    - Deferred function will still fire
    - If nothing handles the panic, program will exit

    #### Recover - Used to recover from panics /someway to save ur program when your app starts to panic so that it doesn't bail out completely
    - Only use in deferred functions - because when a program panics, it will stop executing rest of the function but the only deferred functions with recover keyword
    - Current functions will not attempt to continue, but higher functions in call stack will

    eg implementation.
    --------

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


    OUTPUT
    /*
    start
    about to panick
    Error: Something bad happened
    end
   */

    /*
    STEPS in this program execution
    1. the code will run & print out "start"
    2. the run the panicker() fn
    3. print out "about to panic"
    4. skip the defer fun
    5. reach the panic keyword and stop executing
    6. go back  and run the deferred func with recover() keyword
    7. print out the panic error message
    8. the finally print out the "end" word in main func
    */

--------