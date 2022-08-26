#### Go Primitive Types

####  1.Boolean types

1. - values are true or false
2. - Not an alias for other types (eg. int) 
3. - Zero value is false
 
 #### 2.Numeric types (Integers, Floating point, complex numbers)

 ###### A. Integer types
 
1. Signed integer - have a + & -  and can't store numbers so large
  - int type has varying size, but min 32 bits
  - 8bit(int8) through 64bit(int64)
  -----
    1. An int is atleast 32 bits(the default), but it could be 64 bits, or 128 bits depending on the system
    Other types
    2. 8 bit integers ranging from -128 to 127
    3. 16 bit ranging from -32,768 to 32,767
    4. 32 bit ranging from  -2 billion to +(positive) 2 billion
    5. 64 bit ranging from -9 quintillion to +9 quintillion
  ------
2. Unsigned integer - store larger numbers & they can only be positive numbers 
  -  8bit(byte and uint8) through 32bit(uint32)
  ------
    1. btye - alias for an 8bit uint(unassigned integer)
    1. uint8 - 0-255
    2. uint16 - 0-65,535
    3. uint32 - 0-4,294,967,295

  ------

  ###### Arithmetic Operations
  - Addition+, subtraction-, multiplication*, division/, remainder%
  ###### Arithmetic Operations
  - AND (&), OR (|), XOR/exclusive or (^), AND NOT (&^)

   ------------
    About bit shifting, if you have a number 9 (1001 in binary) and bit shift by 2 to the left (9 << 2) you just add two zeros at the end of number 9 (100100) and you get a number 36. Similar for the >> operator, you just add n zeros to the beginning.
   -------------

###### MORE NOTES:
- Zero value for any integer type is 0 - when you initialize an integer type and don't assign a value, that's will be equivalent to 0
-  You can't mix types in same family (uint16 + uint32 = error)

 ###### B. Floating  types
 - Follow IEEE-754 standard
 - Zero value is 0
 - 32 and 64 bit versions
 - several literation styles:
  1. Decimal (3.14)
  2. Exponential (13e18 or 2E10)
  3. Mixed(13.7e12)

- Arithmetic Operations
  - Addition+, subtraction-, multiplication*, division/
 ###### B. Complex Number  types
 - Zero value is (0+0i)
 - 64 and 128 bit versions
 - Built-in functions
  - complex -male complex number from two floats
   - real - get real part as float
   - imag - get imaginary part as float


 #### 3. Text types
1. strings - alias for an utf8
- Collection of UTF-8 x-characters/ made up of sequence of bytes
- They're immutable
- Can be concatenated with plus(+) operator
- Can be converted to []byte

2. Rune - represents any utf32 characters/ store code that represetns unitcode x-cters
- alias for int32
- special methods normally required to process eg strings.Reader#ReadRune



####