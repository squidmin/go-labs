# `for` construct

This section discusses the `for` construct in detail.

## Topics

- Introduction
- Types of `for` loops
  - Counter-controlled iteration
  - Condition-controlled iteration

## Introduction

Up to this point, the module has covered the `if-else` and `switch` constructs.
It introduces the `for` construct as another crucial and commonly used programming control structure.

The "`for` statement" in Go is a control structure used to execute a block of code multiple times.
Each execution of the block is termed an "iteration."
This means that if you have a set of instructions that you need to repeat, you use a `for` loop to iterate over those instructions the specified number of times.
The fundamental purpose of the `for` loop in Go programming is to facilitate the repetition of code execution, with each cycle through the loop being one iteration.

> **Remark**: There is no `for` equivalent for the `do-while` statement found in many other languages.
> This design decision suggests that the creators of Go did not find the specific use case for a `do-while` loop compelling enough to include it in the language, possibly because the same functionality can be achieved through other looping constructs in Go, such as the `for` loop.

## Types of `for` loops

There are _two_ methods to control iteration.

- _Counter-controlled_ iteration
- _Condition-controlled_ iteration

### Counter-controlled iteration

The simplest form of iteration is counter-controlled iteration. The general format is:

```
for initialization; condition; modification { }
```

For example:

```
for i := 0; i < 10; i++ { }
```

A counter-controlled iteration using a `for` loop is a common control structure in programming for repeating a set of statements a fixed number of times.
This process involves three key steps:

1. **Initialization**: This step is executed only once at the beginning of the loop.
   It typically involves setting an initial value for the loop control variable, in this case, `i`.
   For example, `i := 0` initializes `i` to `0`.
2. **Condition Checking**: Before each iteration (including the first), the loop checks a condition involving the loop control variable (`i < 10` in the example).
   If `false`, the loop terminates, and control passes to the statements following the loop.
3. **Modification**: After the body of the loop executes, the loop control variable is modified (e.g., `i++` to increment `i` by `1`).
   This modification affects the result of the condition check in the next iteration.

This structure ensures that the loop's body executes a precise number of times determined by the initial value, condition, and modification step.
It's a powerful tool for tasks that require repetition, such as processing items in an array or generating a sequence of numbers.
The simplicity and predictability of counter-controlled loops make them a staple in programming for tasks with a known number of iterations.

To see how a `for` loop is implemented in Go, run the following program.

```go
package main

import "fmt"

func main() {
	// For loop with 5 iterations
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration:%d\n", i)
    }
}
```

A loop variable `i` is initialized to `0` and incremented by `1` in each iteration until it reaches `4`, resulting in a total of `5` iterations.
During each iteration, the current value of `i` is printed, demonstrating the loop's execution from `0` to `4`.

In general, a `for` loop's header typically consists of three parts:

- the initialization statement (`i := 0`)
- the condition (`i < 5`)
- the post statement (`i++`)

separated by semicolons (`;`).
This structure sets up the loop's starting condition, its continuation condition, and what happens after each iteration.

In Go, a `for` loop's header does not require parentheses around these three statements, unlike in some other languages like C or Java.
It is a common mistake in Go to add parentheses around these statements, which results in invalid code.

```
for (i = 0; i < 10; i++) { } // Invalid code
```

The opening `{` has to be on the same line as the `for` keyword.
The counter variable ceases to exist after the `}` of the `for`.
Always use short names for the counter variable, like `i`, `j`, `z`, or `x`.
It is considered bad practice to change the counter variable in a `for` loop.

More than one counter can also be used, as in:

```
for i, j := 0, N; i < j; i, j = i + 1, j - 1 { }
```

In Go programming, using parallel assignment within a `for` loop, with two counters that move in opposite directions (one increments and the other decrements), is a common and preferred approach for iteration control.

For loops can be nested, like this:

```go
for i := 0; i < 5; i++ {
    for (j := 0; j < 10; j++ {
        println(j)
    }
}
```

Here, for every iteration of `i`, `j` will run `10` times.
This means that the outer loop

```
for i := 0; i < 5; i++
```

will run `5` times and the inner loop:

```
for j := 0; j < 10; j++
```

will run `50` times in total.

What happens if we use a `for` loop for a general Unicode string?
Let's run a program and find out.

```go
package main

import "fmt"

func main() {
	str := "Demo: String length"
	fmt.Printf("The length of the string: %d\n", len(str))
	
	// For loop to find the character at each position
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character at position %d: %c\n", ix, str[ix])
    }
	
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2: %d\n", len(str2))
	
	// For loop to find character at each position
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character at position %d: %c\n", ix, str2[ix])
    }
}
```

A string variable `str` is declared with the value `"Demo: String length""` and using a `for` loop to iterate over each character of the string.
The loop uses an index `ix` to access and print each character of `str` sequentially, running for the length of `str`.

The `main` function processes a string `str2` containing both ASCII and non-ASCII characters (Chinese: 日本語).
`str2` is initialized, its length is calculated using `len(str2)`, and each character is iterated over using a `for` loop.

The loop uses an index `ix` to access and print each character of `str2` line by line, demonstrating how characters are handled differently based on their byte size in memory.

The length of `str` is `18`, while the length of `str2` is `18`.
ASCII characters, which use 1 byte, are accurately indexed, but non-ASCII characters, requiring 2 to 4 bytes, are not correctly indexed using the traditional method.
This issue can be addressed using the `for range` loop, which will be discussed later.

## Condition-controlled iteration

The second form of the `for` loop is a simplified version used in Go programming for iterating as long as a condition is `true`, similar to the `while` loop found in many other programming languages.
It omits the **initialization** and **modification** sections, focusing solely on the condition that controls the loop's execution.

This form demonstrates an alternative looping mechanism in Go, where instead of specifying a loop variable's start value, incrementation, and condition within the loop's header, only the condition is provided.
This is particularly useful for scenarios where the number of iterations is not known beforehand but is determined by a condition evaluated at each iteration.

Run the following program to see how a `for` loop does the required task by only using a condition.

```go
package main

import "fmt"

func main() {
	var i int = 0
	// Condition controlled for a loop with 5 iterations
	for i < 5 {
		fmt.Printf("Iteration: %d\n", i)
		i = i + 1
    }
}
```

The above code demonstrates the transition from using counter-controlled iterations to condition-controlled iterations in a program.
Unlike initializing the loop counter in the loop's initialization section, a variable `i` is declared beforehand and used in a condition (`i < 5`) to control the loop's execution.
The loop iterates until `i` is less than `5`, with the counter `i` being manually incremented within the loop body.
This results in the loop running exactly five times, demonstrating that condition-controlled iterations can achieve the same outcome as counter-controlled iterations.
