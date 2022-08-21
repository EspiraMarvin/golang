### Pointers

#### Creating Pointers
- Pointer types use an asterisk(*) as a prefix to type pointed  to 
-----
    *int - a pointer to an integer

       	var c int = 42
	    var d *int = &c
-----
- Use the addressof operator(&) to get address of variable in memory

#### Dereferencing Pointers
-using a pointer to get an underlying data - done by preceding an asterisk(*)
------
        var c int = 42
        var d *int = &c
        fmt.Println(c, *d) //OUTPUT c = 42, d = 42
------
- Complex types (eg. structs) are automatically dereferenced

### Create pointers to objects
- can use the addressof operator(&) if value type already exists
-------
    ms := myStruct{foo: 42}
    p := &ms
--------
- Use addressof operator before initializer
---------
    &myStruct{foo: 42}
---------
- with the new keyword
  - cant initialize fields at the same time
-------
    ar ms *myStruct
	ms = new(myStruct)
    ms.foo = 42
	fmt.Println(ms)     // memory pointer
	fmt.Println(ms.foo) // underlying value
--------

#### Working with nil 

#### Types with internal pointers
- All assignments operations in Go are copy operations
- Slices & Maps contain internal pointers, so copies point to same underlying data
                                                                                                                                                                                                                                                                                                                                                                                                                                                               