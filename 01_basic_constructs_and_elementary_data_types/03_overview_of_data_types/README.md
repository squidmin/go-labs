# Overview of data types


This describes the data types in Go.

The following topics are covered:
- Types
- Type conversions


---


### Primitive types

`int`, `float`, `bool`, `string`


---


### Integer types

- `uint8`, `uint16`, `uint32`, `uint64`
- `int8`, `int16`, `int32`, `int64`


#### Code example

```go
package main

import "fmt"

func main() {
    printIntegers()
}

func printIntegers() {
    var unsignedIntExample1 uint8
    fmt.Println(unsignedIntExample1)
    var unsignedIntExample2 uint16
    fmt.Println(unsignedIntExample2)
    var unsignedIntExample3 uint32
    fmt.Println(unsignedIntExample3)
    var unsignedIntExample4 uint64
    fmt.Println(unsignedIntExample4)
	
    var signedIntExample1 int8
    fmt.Println(signedIntExample1)
    var signedIntExample2 int16
    fmt.Println(signedIntExample2)
    var signedIntExample3 int32
    fmt.Println(signedIntExample3)
    var signedIntExample4 int64
    fmt.Println(signedIntExample4)
}
```


### Machine-dependent integer types


There are 3 machine-dependent (e.g., size occupied depends on architecture) integer types:

- `uint`
- `int`
- `uintptr`


#### Code example

```go
package main

import "fmt"

func main() {
    printMachineDependentIntegerTypes()
}

func printMachineDependentIntegerTypes() {
    var unsignedIntExample uint
    fmt.Println(unsignedIntExample)
    var signedIntExample int
    fmt.Println(signedIntExample)
    var unsignedIntPtrExample uintptr
    fmt.Println(unsignedIntPtrExample)
}
```


---


### Alias types

- `byte` - Same as `uint8`
- `rune` - Same as `int32`


#### Code example

```go
package main

import "fmt"

func main() {
    var aliasedUnsignedInt byte
    fmt.Println(aliasedUnsignedInt)
}
```


---


### Floating-point types

Go has two floating-point types:
- `float32`
- `float64`

also often referred to as single precision and double precision respectively.

Go also has two additional types for representing complex numbers (e.g., numbers with imaginary parts):
- `complex64`
- `complex128`

Generally you should use `float64` when working with floating-point numbers.

#### Code example

```go
package main

import "fmt"

func main() {
    var floatExample1 float32
    fmt.Println(floatExample1)
    var floatExample2 float64
    fmt.Println(floatExample2)
    var complexExample1 complex64
    fmt.Println(complexExample1)
    var complexExample2 complex128
    fmt.Println(complexExample2)
	
    fmt.Println("1.0 + 1.0 = ", 1.0 + 1.0)
    fmt.Println("1 + 1.0 = ", 1 + 1.0)
}
```


---


### Strings

Go strings are made up of individual bytes, usually one for each character.

(Characters from other languages like Chinese are represented by more than one byte.)

String literals can be created using double quotes ( " ) or back tick ( ` ) characters.

**Double-quoted** strings cannot contain newlines and they allow special escape sequences.
For example `\n` gets replaced with a newline and `\t` gets replaced with a tab character.


#### Code example

```go
package main

import "fmt"

func main() {
    stringOperations()
}

func stringOperations() {
    fmt.Println(len("Hello World")) // len() usage
    fmt.Println("Hello World"[1])   // String indexing
    fmt.Println("Hello " + "World") // String concatenation
}
```


---


### Composite types

`struct`, `array`, `slice`, `map`, `channel`


---


### Interfaces

TBD
