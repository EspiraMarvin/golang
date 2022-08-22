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
##### Anonymous FUNCTIONS

##### Functions as types

##### Methods

