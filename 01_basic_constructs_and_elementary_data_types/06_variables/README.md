# Variables

## Introduction

The general form for declaring a variable uses the keyword `var` as:

```go
var identifier type
```

Here, `identifier` is the name of the variable and `type` is the type of the variable.
`type` is written _after_ the `identifier` of the variable, unlike most older programming languages.
When a variable is declared, _memory in Go is initialized_, which means it contains the default zero or null value (depending on its type) automatically.

For example:
- `0` for `int`
- `0.0` for `float`
- `false` for `bool`
- empty string (`""`) for `string`
- `nil` for pointer
- zero-ed `struct`
- etc.

Names must start with a letter and may contain letters, numbers or the underscore (`_`) symbol.


---


### Example 1: Declaring a variable

Run the following program to see how declaring a variable works:

```go
package main

import "fmt"

func main() {
	var number int         // Declaring an integer variable
	fmt.Println(number)    // Printing its value
	var decision bool      // Declaring a boolean variable
	fmt.Println(decision)  // Printing its value
}
```

You can see that in the above code, we declare a variable `number` of type `int` in the first line of the `main` function.
As memory is initialized, the default value for `number` is printed, which is `0`.
Similarly, a variable `decision` of type `bool` is declared, and `false` is printed as its value.

> **Note**: The naming of identifiers for variables follows the _camelCasing_ rules (start with a small letter, and each subsequent word in the variable name starts with a capital letter).
> But if the variable has to be exported, it must start with a capital letter.


---


### Example 2: Assigning values

A variable is assigned a value using the assignment operator (`=`) at compile time.
A value can also be computed or changed during runtime.
Declaration and assignment (initialization) can be combined in the general format:

```go
var identifier type = value
```

Here, `value` can be:
- a literal value of type `type`
- a variable of type `type`
- an expression that evaluates to type `type`

Run the following program to see how the assignment operator works on variables in Go:

```go
package main

import "fmt"

func main() {
	var number int = 5        // Declaring and initializing an integer variable
	fmt.Println(number)       // Printing its value
	var decision bool = true  // Declaring and intializing a boolean variable
	fmt.Println(decision)     // Printing its value
}
```

In the above code, we declared a variable `number` of type `int` and initialized it with the value `5`.
Similarly, a variable `decision` of type `bool` was declared and initialized with the value `true`.
These initialized values are printed.


### Example 3: Assigning values

The Go compiler is intelligent enough to derive the type of a variable from its value dynamically, also called _automatic type inference_ at runtime, so omitting the type of a variable is also a correct syntax.

Here's a program demonstrating automatic type inference:

```go
package main

import "fmt"

func main() {
	var number = 5         // Declaring and initializing an integer variable without stating its type
	fmt.Println(number)    // Printing its value
	var decision = true    // Declaring and initializing an integer variable without stating its type
	fmt.Println(decision)  // Printing its value
}
```

We declared a variable `number` and a variable `decision` without stating their types explicitly.
The compiler infers the types, and the result is the same as the previous program in which the types were declared explicitly.
`5` and `true` are printed.


---


### Example 4: Short form initialization with `:=` assignment operator

With the type omitted, the keyword `var` is pretty superfluous.
So we can also write it as:

```go
package main

import "fmt"

func main() {
	number := 5            // Declaring and initializing an integer variable without stating its type
	fmt.Println(number)    // Printing its value
	decision := true       // Declaring and initializing an integer variable without stating its type
	fmt.Println(decision)  // fmt.Println(number)  // Printing its value
}
```

Again, the types of `number` and `decision` (`int` and `bool`) are inferred by the compiler.
This is the preferred form, but it can only be used inside functions, _not_ in package scope.
The short-form assignment operator (`:=`) effectively makes a new variable; it is also called an _initializing declaration_.

If after the lines above in the same code block, we declare:

```go
number := 20
```

This is not allowed. The compiler gives an error: `no new variables`


### Example 5: Short form initialization with `:=` assignment operator

```go
package main

import "fmt"

func main() {
    var someInt1 int // Declaration
    fmt.Println("someInt1 == ", someInt1)
    var someInt2 int = 3 // Declaration + initialization
    fmt.Println("someInt2 == ", someInt2)

    // If declaration and initialization occur on one line, the type can be omitted with the 'var' keyword:
    var someInt3 = 5
    fmt.Println("someInt3 == ", someInt3)
    var someInt4 = 7.8
    fmt.Println("someInt4 == ", someInt4)

    // Variables can also be initialized on one line using the ':=' operator:
    x := 5
    fmt.Println(x)
    // Go is able to infer the type based on the supplied literal value.
}
```


### Example 6: Multivalue assignment

#### Code example

```go
package main

import "fmt"

func main() {
    var (
        a = 5
        b = 10
        c = 15
    )
    fmt.Println("a ==", a)
    fmt.Println("b ==", b)
    fmt.Println("c ==", c)
}
```


---


## Scope of variables


This section gives the details on scope of a variable and the difference between value and reference types.

The following topics are covered:
- Scope of a variable
  - Global scope
  - Local scope
- Printing
- Value types and reference types


### Scope of a variable

<details>
<summary>Expand</summary>

A variable of any type is only known within a certain range of a program, called the variable's _scope_.
In a programming language, there are two main types of scopes:
- Global scope
- Local scope

#### Global vs. local scope

The scope of the variables declared inside a function is called _local scope_.
They are only known in that function or "range of code"; the same goes for parameters and return variables of a function.

Mostly, you can think of a scope as a code block (surrounded by `{ }`) in which the variable is declared.
Run the following program to visualize the concept of scope:

<details>
<summary>Expand code</summary>

```go
package main

import "fmt"

var number int = 5  // Number declared outside (global scope).

func main() {
  fmt.Println("Demo: Scope of a variable")
  var decision bool = true // Decision declared inside function (local scope).
  fmt.Println("Original value of number:", number)
  number = 10
  fmt.Println("New value of number:", number)
  fmt.Println("Value of decision:", decision)
}
```

**Output**

```
Demo: Scope of a variable
Original value of number: 5
New value of number: 10
Value of decision: true
```

</details>


Variables with _global scope_ are declared outside any function.

Variables with _local scope_ are declared within an enclosing code block.

Although identifiers have to be unique, an identifier declared in a block may be re-declared in an inner block, but then the redeclared variable takes priority and _shadows_ the outer variable with the same name; if used, care must be taken to avoid subtle errors.

</details>


### Printing

<details>
<summary>Expand</summary>



</details>
