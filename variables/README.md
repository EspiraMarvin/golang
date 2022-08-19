#### Go Variables

##### VARIABLE TYPES: int, string, float32, float 32
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
a.  var i int = 42
b. var seasonNumber int = 11

- ACRONYMS should ALL be in UPPERCASE
eg.
var theURL= "https:///ww.google.com"
var theHTTP= "https:///ww.twitter.com"
var theHTTPRequest= "https:///ww.twitter.com"

##### 5. Type Conversion
###### RULES:
a. You have to explicity convert one type to another
