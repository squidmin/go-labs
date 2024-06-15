# Overview of functions

This section explains how to write a simple function in Go.


The following topics are covered:
- Functions
  - Hello World
- Comments
- Naming things in Go


---


### Functions

The simplest function declaration has the format:

```go
func functionName()
```

### Syntax

```go
func exampleFunction(param_1 type_1, ..., param_n type_n) (ret_1 type_1, ..., ret_n type_n) {
    ...
}
```

### Function types

```go
func exampleFunction(a int, b bool) (c float32) {  // Function type is float32.
    return 2.0
}
```

Smaller functions can be written on one line like:

```go
func Sum(a, b int) int { return A + b }
```


---


### Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```


---


### Comments

Example:

```go
package main

import "fmt" // Package implementing formatted I/O.

func main() {
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
```

This illustrated the international character of Go by printing "Καλημέρα κόσμε; or こんにちは 世界".


---


### Naming things in Go

Clean readable code and simplicity are major goals of Go development.
As such, the names of things in Go should be short, concise, and evocative.
Long names with mixed caps and underscores (which are often seen in Java and Python code) sometimes hinder readability.

Names should not contain an indication of the package.

A method or function which returns an object is named as a noun; no `get...` is needed.
To change an object, use `set...`.

If necessary, Go uses`MixedCaps` or `mixedCaps` rather than underscores to write multi-word names.
