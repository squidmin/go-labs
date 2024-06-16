# Pointers

This section discusses how pointers are used in the Go language.

## Topics

- Introduction
- Pointers in Go

## Introduction

Go allows programmers to explicitly choose whether a data structure should be a pointer, unlike Java and .NET, offering more control over memory layout.
This control helps optimize the size, allocation count, and memory access patterns of data structures, crucial for high-performance system design.
However, Go does not support arithmetic operations on pointer values.

Pointers are crucial for optimizing performance and are essential for low-level systems programming involving the operating system and network.
They may be unfamiliar to modern object-oriented programmers, so they will be thoroughly explained in this and subsequent chapters.

## Pointers in Go

In Go, every value stored in memory is assigned to a memory block with a unique hexadecimal address.
The address-of operator `&` is used to retrieve the memory address of a variable.
For example:

```go
package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d; Its location in memory: %p\n", i1, &i1)
}
```

The output of this example can be different every time it's ran since the value `&i1` can be different will vary across program executions.

This memory address can be stored in a special data type called a _pointer_.
In the above case, it is a pointer to an _int_.
So `i1` is denoted by `*int`.
If we call that pointer `intP`, we can decalre it as:

```
var intP *int
```

Then, the following statement is true:

```
intP = &i1
```

So `intP` stores the memory address of `i1` which means its points to the physical memory location of `i1`.
In other words, it references the variable `i1`.

A pointer variable stores the memory address of another value, effectively pointing to it in memory.
Its size varies by machine architecture: 4 bytes on 32-bit and 8 bytes on 64-bit systems.
Pointers can reference any type of value, with the asterisk (`*`) prefix indicating the type of value being pointed to, acting as a type modifier.

Using a pointer to refer to a value is called _indirection_.
A newly declared pointer which has not been assigned to a variable has the `nil` value.
A pointer variable is often abbreviated as `ptr` in an expression like:

```
var P *type
```

Maintain clarity in code by inserting a space between a pointer's name and the asterisk (`*`) to avoid confusion with multiplication in complex expressions.
This convention helps with readability and prevents syntax errors in programming involving pointers.

The `*` symbol, when placed before a pointer such as `*intP`, is used to access or retrieve the value stored at the memory address the pointer is holding.
This process is known as dereferencing. It effectively "flattens" the pointer, allowing direct access to the value it points to.
Hence, for any variable `var`, dereferencing its address `*(&var)` yields `var` itself.

Now you have the basic concepts of pointers, and can understand the following program and its output.

```go
package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Println("An integer: %d; Its location in memory", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("THe value at memory location %p is %d\n", intP, *intP)
}
```

In the above code, a variable `i1` is declared and is set to the value `5`.
In the next line, the value is printed along with its memory address with the `&` operator, as `&i1`.
Then a pointer variable `intP` is declared, which is meant to have a memory address value.
The memory address value of `i1` is given to `intP` using `&i1`.
Finally, the memory address is printed by using the `*intP` syntax to dereference the `intP` pointer variable and retrieve its associated value.

The following code example is another program that deals with strings and shows that assigning a new value to the pointer variable changes the value of the variable itself.

#### Example: Changing values with pointers

```go
package main

import "fmt"

func main() {
	s := "Goodbye"
	var p *string = &s
	*p = "Ciao"
	
	fmt.Printf("Here's the pointer p: %p\n", p) // Prints memory address
	fmt.Printf("Here's the string *p: %s\n", *p) // Prints string
	fmt.Printf("Here's the string s: %s\n", s) // Prints same string
}
```

A pointer variable `p` is declared and assigned the address of a string variable `s`.
By dereferencing `p`, the value of `s` is changed to `"Ciao"`.
The change is confirmed through `Printf` statements, demonstrating how pointer dereferencing can modify the value of the variable it points to.

In programming, a literal is a fixed value directly represented in the code (like `5`, `"hello"`, etc.), and a constant is a variable whose value cannot be changed once it has been assigned.
Taking the address of a variable means accessing the memory location where the variable is stored.
However, literals and constants are treated differently by the compiler.

- **Literals**: Since they are fixed values, they don't have a specific memory address that can be accessed or modified.
  They are often stored in read-only sections of memory.
- **Constants**: While constants do have memory addresses, the idea of taking the address of a constant can be problematic in some contexts because constants are meant to be immutable values.
  Some languages or compilers might restrict taking the address of a constant to prevent unintended modifications or to optimize the storage of constants.

```go
package main

func main() {
	const i = 5
	ptr1 := &i // Error: Cannot take the address of i
	ptr2 := // Error: Cannot take the address of 10
}
```

Go supports pointers, similar to other system languages like C, C++, and D, but restricts pointer arithmetic to prevent memory access errors, enhancing its memory safety.

Pointers offer a memory-efficient way to handle variables in programs by allowing the passing of references instead of copies, which saves memory and enhances performance, especially with large or numerous variables.
They also extend the lifetime of variables beyond their creation scope as long as they are referenced.

Using pointers can lead to performance issues due to indirection, which is the process of redirecting to different memory addresses.
While pointers can point to other pointers, creating multiple levels of indirection, this often reduces code clarity.
However, Go simplifies programming by automatically handling some aspects of indirection, such as dereferencing, to improve the developer experience.

Run the following program to see that `nil` pointer dereferencing is not allowed.

#### `nil` pointer dereference

```go
package main

func main() {
	var p *int = nil // Making a nil pointer
	*p = 0
}
// panic: runtime error: invalid memory address or nil pointer dereference
```

The program crashes because a `nil` pointer is created and dereferenced, which is illegal and causes a crash.
This highlights the importance of avoiding `nil` pointer dereference in programming.
