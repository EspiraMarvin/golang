### Interfaces
- they describe behaviour
#### Basics

eg. implementation
------------
    package main
    import "fmt"

    func main() {
        var w Writer = ConsoleWriter{}
        w.Write([]byte("Hello Go!"))
    }

    type Writer interface {
        Write([]byte) (int, error)
    }

    type ConsoleWriter struct{}

    func (cw ConsoleWriter) Write(data []byte) (int, error) {
        n, err := fmt.Println(string(data))
        return n, err
    }

-------------

#### Composing Interfaces
 e.g.
-------------
        type Writer interface{
            Write([]byte)(int, error)
        }

        type Closer interface{
            Close()(error)
        }

        type WriterCloser interface{
            Writer
            Close
        }
-------------- 
#### Type Conversion
###### 1. Empty Interface
 - is an Empty interface that has no methods in it
 e.g
---------
    var in interface{} = 0
---------
###### 2. Type switches
 
 eg.

 ----------
    
    package main
    import "fmt"

    func main() {
        var i interface{} = 0 // 0 -integer // "0"-string // true - "I don't know what it is"
        switch i.(type) {
        case int:
            fmt.Println("i is an integer")
        case string:
            fmt.Println("i is an string")
        default:
            fmt.Println("I don't know what it is")
        }
    }
 ----------

#### Implementing with values vs. pointers

#### Interfaces best practices
- use many small interfaces versus large monolithic ones
  - single method interfaces are some of the most powerful and flexible
    eg io.Writer, io.Reader, interface{}
- Don't export interfaces for types that will be consumed
- Do export interfaces for types that will be used by package
- Design functions and methods to receive interfaces whenever possible


