# The `if-else` Construct

This section discusses the `if-else` construct in detail.

## Topics

- Introduction
- The `if-else` Construct

## Introduction

Go programs begin execution with the `main()` function, executing statements in sequence.
To enable decision-making based on conditions, Go offers conditional or branching structures such as:

- the `if-else` construct
- the `switch-case` construct
- the `select` construct

In Go, repetitive tasks re handled using the `for` loop, which can iterate over ranges.
The loop's flow can be modified using `break` and `continue` keywords.
Additionally, `return` can exit a function, and `goto` can jump to a specific code label.
Unlike Java, C++, or C#, Go simplifies syntax by not requiring parentheses around conditions in `if`, `switch`, and `for` loops, reducing visual clutter.

```
for (range) construct
```

Some other keywords like `break` and `continue` can also alter the behavior of the loop.

## The `if-else` Construct

An `if` clause tests a conditional statement that can be logical or boolean.
If the statement evaluates to `true`, the body of statements between the `{` and `}` characters after the `if` is executed, and if it is `false`, these statements are ignored and the statement following the `if` after `}` is executed.
This creates two distinct paths in the code, but only one path is taken based on the condition's outcome.

```
if condition {
    // Do something
}
```

```
if condition {
    // Do something
} else {
    // Do something else
}
```

A second conditional check (`else-if`) is added after an initial `if` condition and before the final `else` block.
This creates a flow with three distinct paths, each taken based on different conditions being true.
The first `if` checks for the first condition, the `else if` checks for another condition if the first isn't met, and the `else` serves as a default action if none of the specified conditions are true.
This structure allows for more nuanced decision-making in code, enabling different actions to be taken based on various conditions.

While you can technically use unlimited `else if` branches in code, it's best to limit their use for clarity.
Prioritize placing the most probable condition first for efficiency.

In Go programming, using curly braces (`{}`) is compulsory even for single-statement blocks within conditional structures like `if` and `else`.
Additionally, the opening brace `{` must be on the same line as the `if` or `else` statement, and `else if` or `else` must follow the closing brace `}` of the preceding block on the same line.
These syntax rules are strictly enforced by the compiler to maintain consistency and adhere to standard software engineering practices.

The below code example is invalid:

```go
if x {
} else { // Invalid
}
```

Go enforces consistent indentation and alignment; specifically, branches should be indented with spaces or a tab, and the closing brace aligns if the `if` statement.
`gofmt`, Go's formatting tool, enforces these stylistic rules.
Additionally, it highlights that conditions in `if` statements can be complex, utilizing logical operators (`&&`, `||`, `!`) and parentheses to clarify precedence and enhance code readability.

In Go, parentheses around conditionals are optional but can enhance clarity for complex conditions, aiding in readability and understanding, especially when evaluating multiple variable values for different outcomes.

```go
package main

import "fmt"

func main() {
	n := 42
	// Use of a control structure if and else to check whether number is even
	if n % 2 == 0 {
		fmt.Printf("The value is even\n")
    } else {
		fmt.Printf("The value is odd\n")
    }
}
```

In the above code, a _conditional statement_ (`if-else`) is used to check if a variable is even or odd.

If `n` is even (`n % 2 == 0`), it prints `"The value is even"`.
If not, it executes the `else` part and prints `"The value is odd"`.
This demonstrates basic conditional logic in programming.

> A _conditional statement_ after the `if` keyword in programming evaluates to either `true` or `false`.
> This statement can be a boolean expression which involves variables and operators that the program evaluates to determine its truth value.
> Alternatively, it can directly be a boolean value (`true` or `false`), determining the flow of execution without the need for evaluation.

A boolean variable `bool1` is declared and set to `true`.
It's used in an `if` statement without comparing it to `true` explicitly because it's already a boolean.
Consequently, the program will always execute the code within the `if` block, printing `"The value is true"`.

For special cases, not that you could write something like:

```go
if true {
	fmt.Printf("I'm always executing this branch")
}
```

Here, the code block in the curly braces (`{}`) is always executed, so this is not very useful. Inversely, you could write:

```go
if false {
	fmt.Printf("I will never execute this code")
}
```

Now, the code block in the curly braces (`{}`) will never be executed.

When writing conditional statements, it's generally more straightforward and readable to test for conditions that evaluate to `truw`.
For instance `if condition` is preferred over `if !condition` because it's generally easier to understand conditions that are being met rather than conditions that are not being met.
However, the `!` operator allows for testing the opposite of a condition, enabling the execution of code blocks based on a condition being `false`.

The principle is below above with the use of the `!` operator to reverse a condition's truth value.

```go
if !bool1
```

For example, if you have a boolean variable `bool1` set to `true`, writing `if !bool1` allows you to execute a block of code if `bool1` is `false`.
This demonstrates the flexibility in handling conditions but emphasizes that direct, positive condition checks (`if condition`) are usually more intuitive.

Parentheses (`()`) around the condition are often necessary. For example:

```go
if !(var1 == var2)
```

which can be rewritten as the shorter:

```go
if var1 != var2
```

In Go programming, it's common to skip the `else` clause when the block inside an `if` statement ends with a control statement like `break`, `continue`, `goto`, or `return`.
This practice makes the code cleaner and easier to read by reducing unnecessary nesting.

For example, when an even value is returned from code, you would write:

```go
if n % 2 == 0 {
	return n
}
fmt.Printf("Continuing with an odd value\n")
```

In Go, it is generally recommended to structure your code to directly return one value if a condition is met and another value if it is not, without using an `else` statement.
This rule of thumb should be used in situations where different outcomes are returned from a function or a block of code based on a condition.

When returning different values `x` and `y`, regardless of whether a condition evaluates to `true`, use the following pattern:

```go
if condition {
    return x
}
    return y
```

Here is an example:

```go
if n % 2 == 0 {
	return n // Return an even value
}
return (n + 1) // Return an odd value
```

In Go, an `if` statement can include an initialization statement before the condition, separated by a semicolon (`;`).
This allows for declaring and initializing a variable that will only be accessible within the `if`, `else-if`, and `else` blocks.
This pattern is useful for concise code where a variable is needed only for the scope of the conditional logic.

```go
if initialization; condition {
	// Do something
}
```

For example, instead of:

```go
val := 10
if val > max {
	// Do something
}
```

you can write:

```go
if val := 10; val > max {
	// Do something
}
```

The scope of a variable declared in an `if` statement (like `val`) is limited to the blocks of that `if` statement, including any associated `else-if` blocks and `else` blocks.
Attempting to use it outside these blocks results in a compiler error due to the variable being undefined.
This concept allows for concise and localized variable initialization and usage within conditional structures.

```go
if value := process(data); value > max {
    // ...
}
```

The result of the function `process()` can be retrieved in the `if`, and action is taken according to the `value`.

Here is another complete example using variants of the `if` construct.

```go
package main

import "fmt"

func main() {
	var first int = 10
	var cond int
	if first <= 0 {
		fmt.Printf("First is less than or equal to 0\n")
    } else if first > 0 && first < 5 {
		fmt.Printf("First is between 0 and 5\n")
    } else {
		fmt.Printf("First is 5 or greater\n")
    }
	if cond = 5; cond > 10 {
		fmt.Printf("'cond' is greater than 10\n")
    } else {
		fmt.Printf("'cond' is not greater than `0\n")
    }
}
```

The flow of control in the above code is based on conditional statements.
It evaluates conditions associated with the `first` variable, directing the flow through `if`, `else if`, and `else` blocks.
The execution path is determined by the value of `first`, leading to different blocks of code being executed.
Additionally, the `cond` variable is both initialized and evaluated in a conditional statement, affecting the flow of control based on the outcome of the condition.
