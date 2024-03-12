# Declaration and initialization

This guide covers the fundamental concepts of arrays and slices in Go, including their declaration, initialization, and usage.

## Topics:

- Introduction to Arrays and Slices
- Declaration and Initialization of Arrays
- Working with Slices

## Introduction to Arrays and Slices

In many programming languages, the array, denoted by `[]`, serves as a foundational data structure.
Go's approach to arrays shares similarities with other languages but introduces unique characteristics, particularly in its handling of arrays and the introduction of slices.

While arrays in Go are fixed in size and less flexible compared to languages like C, slices offer a dynamic and powerful alternative, built on top of Go's array type.
Understanding both arrays and slices is crucial for effective programming in Go, as slices are prevalent due to their versatility and convenience.

## Arrays in Go

An **array** in Go is a sequence of elements of a specific type, characterized by its numbered and fixed-length nature.
The length of an array is part of its type, making the declaration of arrays with different sizes distinct in type.
For example, `[5]int` and `[10]int` are considered different types due to their differing lengths.

### Key Characteristics

- **Homogeneous Data Structure:** All elements must be of the same type.
- **Fixed Length:** The size of the array must be defined at compile time and cannot change.
- **Index-Based Access:** Elements are accessed and modified using their index, starting from `0`.

### Declaration Syntax

The syntax for declaring an array in Go is as follows:

```go
var identifier [len]type
```

For instance, to declare an array of 5 integers:

```go
var arr1 [5]int
```

## Initialization and Usage Example

Below is an example demonstrating the declaration, initialization, and iteration over an array in Go:

```go
package main

import "fmt"

func main() {
    var arr1 [5]int // Declaration of an array
    // Initialization of array elements
    for i := 0; i < len(arr1); i++ {
        arr1[i] = i * 2
    }
    // Iteration over array elements
    for i, value := range arr1 {
        fmt.Printf("Item at index %d is %d\n", i, value)
    }
}
```

This example illustrates basic array operations, including setting and accessing elements by their indices.
Notice the use of `len(arr1)` to determine the array's length and the range keyword for iteration, which simplifies accessing the index and value of each element.

> **Note**: Arrays in Go have a maximum size limit of 2Gb, but in practice, slices are often used for their flexibility and dynamic resizing capabilities.

---

## Slices in Go

A **slice** is a dynamically-sized, flexible view of an array's elements.
Unlike arrays, slices do not have a fixed size, making them more versatile and commonly used in Go applications.
Slices maintain a reference to an underlying array, providing the ability to modify its elements and dynamically change its size.

### Declaration

A slice is declared without including its length:

```go
var identifier []type
```

For example, to declare a slice of integers:

```go
var slice1 []int
```

This declares a slice named `slice1` of type `int`, but it doesn't allocate any space for the slice's elements.
At this point, `slice1` is `nil`.

### Initialization

Slices can be initialized in several ways:

- Using the `make` function: This is the most common way to initialize a slice because it allows you to specify the initial length and capacity.
  The `make` function creates a slice of a specified type, length, and optionally capacity.

  ```go
  slice1 := make([]int, 5) // A slice of 5 integers, initialized to zero.
  ```
  
- **Slice literals**: You can also initialize a slice with a slice literal, which specifies the values for its elements.

  ```go
  slice2 := []int{1, 2, 3, 4, 5} // A slice containing 1, 2, 3, 4, and 5.
  ```
  
- **From an array or another slice**: Slices can be created by "slicing" an array or another slice using the `low : high` expression syntax.

  ```go
  arr := [5]int{0, 1, 2, 3, 4}
  slice3 := arr[1:4] // A slice including elements 1 through 3 of arr.
  ```
  
## Example

```go
package main

import "fmt"

func main() {
    // Creating a slice using make
    slice1 := make([]int, 5)
    for i := range slice1 {
        slice1[i] = i * 10
    }

    // Creating a slice using a slice literal
    slice2 := []string{"hello", "world"}

    // Creating a slice from an array
    arr := [5]int{0, 1, 2, 3, 4}
    slice3 := arr[1:4] // includes elements 1, 2, and 3

    // Printing slices
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)
}
```
