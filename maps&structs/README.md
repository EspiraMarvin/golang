### Maps & Structs

#### Maps
##### - What they are - collection of value types that are accessed via keys
- they are reference types/ meaning multiple assignments refer to the same underlying data
- Creation methods
1. via literals
-----
    statePopulations := map[[]int]string{}
-----
2. via the make function
------
	statePopulations := make(map[string]int)
------

##### - Manipulation 

- adding to a map,
- deletion 
- selecting item from a map 
    -----
        myMap["key"] = "value"
    -----
- finding how many elements are their in a map len(map)
- check for presence of a with `"value, ok"` form of result

##### Notes:
- bool, strings,all the numeric types, pointers, interfaces, channels structs, arrays can be tested for equality checking (string type == another string type) except slices, maps and other functions
- An array is a valid key type but not a slice
- Arrays and slice returns items in the order we ordered them, but a map doesn't guarantee an order in which we created
- 

#### Structs
##### - Creation - collection of different data types that describe a single concept
- structs are value types - if you assign a variable to an existing struct then all the struct values are copied over to create a brand new struct, so if you manipulate a struct in one location is not going to affect any other variables in ur app
1. created as types
-----
    type Doctor struct {
        number     int
        actorName  string
        episodes   []string
        companions []string
    }
-----
2. via anonymous structs
------
	bDoctor := struct{ name string }{name: "John Pertwee"} 
    
    // UseCase:
    Generating  a json response to a web service call, (usecases are mostly short-lived)
------

##### - Naming conventions
- They're keyed by name fields - if you capitalize the 1st letter they're exported if small letter they're local to the package

##### - Embedding
- it doesn't support inheritance but can use composition via embedding
------

            package main

            import (
                "fmt"
            )

            // EMBEDDING - Go uses COMPOSITION instead of inheritance (GO doesnt support traditional OOP principles)

            type Animal struct {
                Name   string
                Origin string
            }

            type Bird struct {
                Animal   // embedding Animal struct
                SpeedKPH float32
                CanFly   bool
            }

            func main() {
                b := Bird{}
                b.Name = "Emu"
                b.Origin = "Australia"
                b.SpeedKPH = 48
                b.CanFly = false
                fmt.Print(b)
                fmt.Print(b.Name)
            }

-------------
- Bird is not of type Animal, it just inherts properties of Animal through composition.
- It just says Bird has Animal characteristics, but its not the same thing as an Animal and both Bird & Animal cannot be used interchangeably, unless we use interfaces

##### Tags
- Tags can be added to struct fields to describe field
###### Notes:
- What they are - unlike arrays which store same type of datatype, also slices and maps(keys always have the same type, values always have the same type), structs can be used to store different data types
- THey're value types. Unlike maps and slices, structs refer to independent data sets 
if you assign a variable to an existing struct then all the struct values are copied over to create a brand new struct, so if you manipulate a struct in one location is not going to affect any other variables in ur app
/ when you pass a struct around in ur app, you're actually passing copies of the same data around.