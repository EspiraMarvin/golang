### Collection types - Arrays and Slices

## Arrays
- In GO, arrays are considered values and not references/pointers types to the arrays (they don't refer to the same underlying data 
) unlike slices
- You can perform slice operations
- They have a  fixed size

### A. Creation/ Declaration styles (3)
----
        i.   a := [3]int{1,2,3}
        ii.  a := [...]int{1,2,3}
        iii. var a [3]int
-----
#### B. Built-in functions 
  ###### 1. length(len(a)) - returns the size of the array
  ###### 2.slice,
 working with Arrays
- accessed via zero-based index eg a := [3]int{1,3,5} //a[1] === 3
- copies refer to different underlying data
### Slices - backed by arrays
- In GO, slice are references/pointers types, they refer to the same underlying data 
- You can perform slice operations
 ##### Creation Declaration styles (3)
 - Slice existing array or slice
  -----------
 -----------
 - Literal styles
 -----------
    m := []int{} // we leave the 3 dots out of the [] sign
 -----------
 - Via make functions
 -------
    a := make([]int, 10) // create slice with capacity and length == 10

    a := make([]int, 10, 100) // slice with length == 10 and capacity == 100


 -------
 ##### Built-in functions -
   ###### 1. length(len(a)) - returns length
   ###### 2. capacity(cap(a)) - eturns capacity
   ###### 3. slice,
   ###### 4. make, 
   ###### 5. append- add element to the last      // stack operations
   ###### 6. shift -add element at the beginning // stack operations eg end := n[:len(n)-1]
 ##### working with Slices
- Copies refer to the same underlying array