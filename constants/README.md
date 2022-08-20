#### Constants

##### NOTES:

1. Constants are immutable but can be shadowed - you can change the const at the inner scoped by changing its value and even its type
2. You can't set a constant equal to something that has to be executed at compile time/
- they're replaceable by the compiler at compile time, so the value must be calculable, 
 so we're able to access functions or CMD line arguments in order to determine value of constants in our app, but we're able to do some simple calculations like in Enumerated constants

##### 1. Naming convention
###### RULES:
1. PascalCase - for exported constants
2. camelCase - for internal constants

##### 2. Type constants
- They work like immutable variables
- Can interoperate only with similar types

###### RULES:
1. You can't set a constant equal to something that has to be executed at compile time 
------
    e.g
        package main
        import (
            "fmt"
            "math"
        )
        func main() {
            const myConst2 float64 = math.Sin(1.57) // error (You can't set a constant equal to something that has to be executed at compile time)
        }
------
2. Arrays are always variable types and not const types
3. Constants can be shadowed just like variables
------
    e.g
        package main
        import (
            "fmt"
            )
        // const shadowing
            const m int16 = 27

        func main() {
            const m int = 14
	        fmt.Printf("%v, %T\n", m, m)    // Output 14
        }
------
##### 3. Untype constants
- They work like literals /We can implicitly convert types unlike in variables, eg s is implicitly converted to int16
-------
      package main
        import (
            "fmt"
        )
        func main() {
            const s = 42
	        const t int16 = 27
	        fmt.Printf("%v, %T\n", s+t, s+t) // implicit conversion of types, s is implicitly converted to int16
        }
-------
- Can interoperate with similar types
##### 4. Enumerated constants
###### RULES:
- special symbol iota allows related constants to be created easily
- iota starts at 0 in each const block and increments by one
- wach out of constant values that match zero values for variables
-------
      package main
        import (
            "fmt"
        )

        // enumerated constants
        const (
            a = iota //  const a = iota 
            b            const b = iota
            c            const c = iota   
        )

        const  (
            a2 = iota // iota is scoped to that constant block, so this a2 will start from 0
        )

        func main() {
           const s = 42
	       fmt.Printf("%v\n", a)  // 0
           fmt.Printf("%v\n", b)  // 1
           fmt.Printf("%v\n", c)  // 2
           fmt.Printf("%v\n", a2) // 0
        }
-------

##### 5. Enumeration expressions
###### RULES:
-  Operations that can be determined at compile time are allowed
  * Arithmetic
  * Bitwise operations
  * Bitshifting


