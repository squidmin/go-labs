# Overview of data types

This describes the data types in Go.

The following topics are covered:
- Types
- Type conversions

---

## Primitive types

`int`, `float`, `bool`, `string`

---

## Integer types

Go has architecture-dependent types such as `int`, `uint`, and `uintptr`.
They have the appropriate length for the machine on which a program runs.

An `int` is a default signed type, which means it takes a size of 32 bit (4 bytes) on a 32-bit machine and 64 bit (8 bytes) on a 64-bit machine, and the same goes for uint (`unsigned int`).
Meanwhile, `uintptr` is an unsigned integer large enough to store a bit pattern of any pointer.

The architecture independent types have a fixed size (in bits) indicated by their names.=

For _signed_ integers:

- `int8` (`-128` to `127`)
- `int16` (`-32768` to `32767`)
- `int32` (`−2,147,483,648` to `2,147,483,647`)
- `int64` (`−9,223,372,036,854,775,808` to `9,223,372,036,854,775,807`)

For _unsigned_ integers:

- `uint8` (with the alias byte, `0` to `255`)
- `uint16` (`0` to `65,535`)
- `uint32` (`0` to `4,294,967,295`)
- `uint64` (`0` to `18,446,744,073,709,551,615`)

#### Code demonstration

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

## Machine-dependent integer types

There are 3 machine-dependent (e.g., size occupied depends on architecture) integer types:

- `uint`
- `int`
- `uintptr`

#### Code demonstration

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

## Alias types

In Go, alias types are a way to create a new name for an existing type. They are useful for adding more semantic meaning to your code or for gradual code refactoring.
Here's a detailed guide on alias types in Go, with code examples.

### Introduction to Alias Types

An alias type in Go is declared using the `type` keyword, followed by the new type name and the original type.
The alias type is fully compatible with the original type, meaning values of the alias type can be used wherever the original type is expected and vice versa.

**Basic Syntax**

```
type AliasTypeName OriginalType
```

### The `byte` and `rune` special types

In Go, byte and rune are special types that represent specific kinds of integers, each with its unique purpose and use case.

#### `byte` keyword

The `byte` type is an alias for `uint8`, an unsigned 8-bit integer. It represents an ASCII character.

**Common Usage**

- Primarily used to manipulate individual bytes.
- Often seen in tasks related to handling raw binary data, byte streams in I/O operations, or processing ASCII character data.
- Example: Reading files byte by byte, handling binary protocols, or working with ASCII character sets.

#### `rune` keyword 

The `rune` type is an alias for `int32`, a signed 32-bit integer. It represents a Unicode code point.

**Common Usage**

- Used to handle Unicode characters, which may be more than one byte in size.
- Essential for processing text in various languages and symbols not covered in the ASCII range.
- Example: Parsing strings in multi-language applications, handling emojis, or working with any Unicode characters.

#### Code examples

```go
package main

import "fmt"

func main() {
    var aliasedUnsignedInt byte
    fmt.Println(aliasedUnsignedInt)
}
```

```go
package main

import "fmt"

func main() {
    var r rune = 'あ'
    fmt.Println(r) // Outputs: 12354 (Unicode code point for 'あ')
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

### Booleans

#### Code example

```go
package main

import "fmt"

func main() {
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

**Output**

```
true
false
true
true
false
```

---

### Composite types

`struct`, `array`, `slice`, `map`, `channel`

---

### Interfaces

TBD
