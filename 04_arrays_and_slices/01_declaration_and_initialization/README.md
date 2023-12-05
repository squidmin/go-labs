# Declaration and initialization

Describes the important concepts of an array, e.g., how to use, declare, and initialize them.

Topics:

- Introduction
- Concept

## Introduction

The array-type indicated by the `[]` notation is well-known in almost every programming language as the basic workhorse in applications.

The Go array is very much the same but has a few peculiarities.
It is not as dynamic as in C, but to compensate, Go has the `slice` type.
This is an abstraction built on top of Go’s array type.
So to understand slices, we must first understand arrays.
Arrays have their place, but they are a bit inflexible.
Therefore, you don't see them too often in Go code.
Slices, though, are everywhere, and they are built on arrays to provide greater power and convenience.

## Concept

An **array** is a _numbered_ and _fixed-length_ sequence of data items (elements) of the same type (it is a homogeneous data structure).
This type can be anything from primitive types like integers or strings to self-defined types.
The length must be a constant expression, that must evaluate to a non-negative integer value.
It is part of the type of the array, so arrays declared as `[5]int` and `[10]int` differ in type.
The items can be _accessed_ and _changed_ through their _index_ (their position).
The index starts at `0`, so the first element has index 0, the second has index 1, and so on (arrays are zero-based as usual in the C-family of languages).

The number of items, also called the length (`len`) or _size_ of the array, is fixed and must be given when declaring the array (length has to be determined at compile time in order to allocate the memory).

>**Remark**: The maximum array length is 2Gb.

The format of the declaration is:

```
var identifier [len]type
```

For example:

```
var arr1 [5]int
```

#### Example

```go
package main

import "fmt"

func main() {
	var arr1 [5]int  // declaring an array
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2  // initializing items of the array
    }
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Item at index %d is %d\n", i, arr1[i])  // printing items of the array
    }
}
```
