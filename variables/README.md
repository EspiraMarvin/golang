#### Go Variables

##### VARIABLE TYPES: bool int, string, byte/uint8, float32, float64, complex64, complex128

###### Int Types

-> int, int8, int16, int32/rune, int64,
uint,
uint8/byte (can stored 255 as the largest number),
uint16 (can stored 65,000 as the largest number),
uint32, uintptr

![Alt text](./variables.png?raw=true "Variables")

##### 1. Variable declaration

###### RULES:

a. Declaring variables at the package level you cannot use the colon(:) syntax
b. Examples ways to declare a variable
var i int - variable declare and not initialized
var i int = 25 - variable declare and initialized
i := 13 - shortform of declaring variables only done in the main() block scope
c. You can't assign declared variables twice in the same scope -the main() scope
d. All variables declared must be used

##### 2. Redeclaration and shadowing

###### RULES:

-you can;t declare variables twice in the same scope/ variable in the inner most scope takes precedence

##### 3. Visibility (3 levels)

- If variable name is lowercase & declared at the package level, Its scoped to the package, any file in the package can access that variable
- If variable name is uppercase & declared at the package level, Its exported from the package and its globally visible
- If variable is declared in the main() scope, that variable is blocked scoped and only accessible in that scope

##### 4. Naming conventions - (Pascal or camelCase)

###### RULES:

- the length of a variable name should reflect the life of the variable
  eg.
  a. var i int = 42
  b. var seasonNumber int = 11

- ACRONYMS should ALL be in UPPERCASE
  eg.
  var theURL= "https:///ww.google.com"
  var theHTTP= "https:///ww.twitter.com"
  var theHTTPRequest= "https:///ww.twitter.com"

##### 5. Type Conversion

###### RULES:

a. You have to explicity convert one type to another
