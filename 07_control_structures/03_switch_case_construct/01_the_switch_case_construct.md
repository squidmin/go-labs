# The `switch-case` Construct

This section discusses the `switch-case` construct in detail.

## Topics

- Introduction
- `switch` statement with values
- `switch` statement with conditions
- Initialization within the `switch` statement

## Introduction

In the last two sections, we studied the `if-else` construct.
There is another control structure known as the `switch-case` structure.

The `switch` keyword simplifies code by replacing lengthy `if-else` statements that compare a variable against multiple values.
It directs the execution flow to different code sections based on the variable's value, as illustrated by a figure explaining the `switch-case` structure.

## `switch` statement with valeus

Compared to the C and Java languages, `switch` in Go is considerably more flexible.
It takes the general form:

```go
switch var1 {
case val1:
	// ...
case v2:
	// ...
default:
	// ...
}
```

In Go's `switch` statement, the case values can be of any type, not limited to constants or integers, but all must be of the same type or evaluate to that type.
Additionally, the syntax requires the opening brace `{` to be on the same line as the `switch` keyword.
Multiple statements can follow a case without needing braces, although braces are permitted.

In Go's `switch` statements, if a case has only one statement, it can be written on the same line as the case itself.
If a case ends with a `return` statement, a `return` statement must also follow the closing brace of the `switch` block.
This ensures clarity and proper flow control within the `switch` structure.

More than one value can be tested in a case. They are checked sequentially from left to right.
We should place the most probable values first, to save the time of computation.
The first branch that is correct is executed, and then the `switch` statement is complete.

Compare the following two cases:

```go
package main

import "fmt"

func main() {
	var i int = 10
	switch i {
	case 0: // Empty case body. Nothing is executed when i == 0.
	case 1:
		f() // The function 'f' is not called when i == 0.
	}
}

func f() {
	fmt.Printf("Test")
}
```

And:

```go
package main

import "fmt"

func main() {
	var i int = 0
	switch i {
	case 0: fallthrough
	case 1:
		f() // 'f' is now also called when i == 0.
    }
}

func f() {
	fmt.Println("Test")
}
```

In Go, if a `case` in a `switch` statement has no code, like `case 0`, nothing executes for that case, and the `switch` ends.
To execute the next case's code (`case 1`) when `case 0` is matched, use `fallthrough`.
This continues execution to subsequent cases until a case without `fallthrough` is reached, unlike in C, where a `break` is needed to prevent falling through by default.

The `fallthrough` keyword in Go allows the execution of sequential case blocks in a `switch` statement, enabling additional actions at each level.
The `default` case acts like an `else` clause, executing when no other case matches.
It's optional and can be placed anywhere, but it's typically at the end for clarity.

Here's a simple program to see how the `switch` statement works.

```go
package main

import "fmt"

func main() {
	var num1 int = 100
	// Adding switch on 'num1'
	switch num1 {
	case 98, 99: // First case: num1 = 98 or 99
		fmt.Println("It's equal to 98")
	case 100: // Second case: num1 = 100
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
    }
}
```

In the above code, ee declare a variable `num1` and initialize it to a value `100`.
A `switch` statement is used against the `num` value, which means different cases will be written on `num1`.
A total of **3** cases are created, including the `default` case.

### `switch` statement with conditions

There is a variant of the `switch` statement in Go where instead of matching a variable against various values, it evaluates different conditions directly.
This first condition that evaluates to `true` triggers its associated case to execute.
This approach simplifies code, making it more readable, especially when dealing with multiple conditional branches, and functions similarly to an `if-else` chain but with a syntax that may be easier to follow.

```go
package main

func main() {
	switch {
	case condition1:
		// ...
	case condition2:
		// ...
	default:
		// ...
	}
}
```

For example:

```go
package main

func main() {
	switch {
	case i < 0:
		f() // Function call
	case i == 0:
		f2() // Function call
	case i > 0:
		f3() // Function call
    }
}
```

In a `switch` statement, conditions can be based on any data type that allows for equality checks, including integers, strings, and pointers.
Examine the below program to understand how `switch` statements can evaluate rather than just matching specific values.

```go
package main

import "fmt"

func main() {
	var num1 int = 100
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 1 and 10")
	default:
		fmt.Println("Number is 10 or greater")
    }
}
```

- **Variable Initialization**: The variable `num1` is initialized with a value of 100, which is then used within the `switch` statement to evaluate the conditions defined in each case.
- **Cases without a Value**: Instead of switching on `num1` directly, the cases are defined by conditions on (`num1 < 0`, `num1 > 0 && num1 < 10`).
  This allows for more complex and flexible decision-kaking with the switch.
- **Default Case**: The default case acts as a catch-all, executing if none of the specified conditions are met.
  In this scenario, since `num1` is 100, it doesn't satisfy the conditions of the first two cases, leading to the execution of the default case.
- **Outcome**: The execution of the default case results in printing `"Number is 10 or greater."` which is the intended behavior given the value of `num1`.

In the above code, the `switch` statement is used in a slightly unconventional manner, where instead of switching on a specific variable's value, the `switch` operates on conditions directly.
This approach allows for more complex decision-making processes within a single `switch` statement, akin to chaining `if-else` statements but with potentially cleaner syntax and readability.

This approach to using `switch` statements is particularly useful for simplifying complex conditional logic, making the code easier to read and maintain by grouping related conditions in a structured format.

## Initialization within the `switch` statement

A `switch` can contain an _initialization statement_, like the `if` construct.

```
switch initialization; {
case val1:
    // ...
case val2:
    // ...
default:
    // ...
}
```

After writing the `switch` keyword, we can do initialization and add a `;` at the end:

```
switch a, b := x[i], y[j]; {
case a < b: t = -1
case a == b: t = 0
case a > b: t = 1
}
```

In a `switch` statement, variables `a` and `b` are initialized simultaneously (`a` with `x[i]` and `b` with `y[j]`) before evaluating the `switch` cases, which are based on conditions comparing `a` and `b`.
Write a short program to demonstrate this initialization feature within a `switch` statement.

```go
package main

import "fmt"

func main() {
	// Initialization within a 'switch' block
	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 1 and 10")
	default:
		fmt.Println("Number is 10 or greater")
    }
}
```

This is the same program written previously, but with a little modification.
Previously, `num1` was declared separately. But now, it's declared in a `switch` block and initialized on the same line.
The rest of the program is the same and will produce the same output.
