### Maps & Structs

#### Maps
- What they are
- Creation
- Manipulation 
    - adding to a map,
    - deletion 
    - selecting item from a map
    - finding how aman elements are their in a map len(map)

##### Notes:
- bool, strings,all the numeric types, pointers, interfaces, channels structs, arrays can be tested for equality checking (string type == another string type) except slices, maps and other functions
- An array is a valid key type but not a slice
- Arrays and slice returns items in the order we ordered them, but a map doesn't guarantee an order in which we created

#### Structs
- Creation
- Naming conventions
- Embedding
##### Notes:
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

- Tags

##### Notes:
- What they are - unlike arrays which store same type of datatype, also slices and maps(keys always have the same type, values always have the same type), structs can be used to store different data types
- unlike maps and slices, structs refer to independent data sets / when you pass a struct around in ur app, you're actually passing copies of the same data around.