### FUNCTIONS

-  Starts with the `func` keyword

##### Basic syntax

----------
    func foo() {
        .....
    }
-----------

##### Parameters

-----------
    func foo(bar string, baz int)

// Parameters of the same type

    func foo(bar, buz  int)

-----------

###### variadic Parameters - used to send list of same types in a Function
 - must be the last Parameter 
 - is received as slice
 - syntax, has 1 parameter bar of type string and baz as the variadic parameter which is received as slice
 ------
    func foo(bar string, baz..int)
 ------

 ---------
        func foo(bar string, baz..int)

        eg.

        func sum2(values ...int) (result int) { 
            fmt.Println(values)
            for _, v := range values {
                result += v
            }
        return
}
 ---------

##### Return values
- Single return values just list type

-------
    func foo() int
-------

- multiple return value list types surrounded by parantheses

--------
    func foo() (int, error) // returns an integer and an error
--------

###### Named return values

- Initializes returned variable to 0
- Return using return keyword on its own

-------

    func sum2(values ...int) (result int) { // (result int) is a named return value
        fmt.Println(values)
        for _, v := range values {
            result += v
        }
        return
    }
-------

###### Can return addresses of local variables
- Automatically promoted from local memory(stack) to shared memory(heap)


##### Anonymous Functions

-Immediately invoked function syntax 

----------
    func() {
      ....
    }()
----------

-Assigned to a variable or passed as an argument to a function syntax 
----------
        a := func() {
           ....
        }
        a()

        In this syntax a function can only be invoked after it has been declared
----------

##### Functions as types
- Can assign functions to variables or use as arguments and return values in functions
- Type signature is like function signature, with no parameter names

---------
      var f func(string, string, int) (int, error)
---------

 eg.

  ---------
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
 ---------
    

##### Methods

-

