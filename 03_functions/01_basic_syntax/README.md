# Functions

## Introduction

In Go (Golang), functions are fundamental building blocks of the language.
They allow you to encapsulate code for performing specific tasks, which can be reused throughout your program.
Understanding functions in Go is crucial for effective programming.

## Basic Syntax

A function in Go is declared with the `func` keyword, followed by the function's name, parameters (if any), and the return type.

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // Function body
}
```

## Parameters

Functions can take zero or more parameters.
Parameters are specified in the function declaration and are used to pass data into a function.

```go
func add(x int, y int) int {
    return x + y
}
```

In this example, `add` takes two parameters of type `int`.

## Return Values

Function in Go can return one or more values.
The return type is specified after the parameters.

```go
func divide(a float64, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

Here, `divide` returns two values: a `float64` and an `error`.

## Named Return Values

Go allows you to name return values.
This can improve readability, especially when returning multiple values.

```go
func stats(numbers []int) (min int, max int) {
    min, max = numbers[0], numbers[0]
    for _, n := range numbers {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return
}
```

Named return values are automatically declared as variables in the function.

## Variadic Functions

Variadic functions can be called with any number of trailing arguments.
Use an ellipsis `...` before the type name to specify a variadic parameter.

```go
func sum(nums ...int) {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Anonymous Functions and Closures

Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.

```go
func sequence() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

This `sequence` function returns another function, which closes over the variable `i`.

## Conclusion

Functions in Go are versatile and powerful, allowing for clean and efficient code design.
By understanding the basics of function declaration, parameters, return values, variadic parameters, and closures, you can utilize functions effectively in your Go programs.

---

## Exercises

To reinforce your understanding of functions in Go, here are some simple coding exercises that cover various aspects like

- basic syntax
- parameters
- return values
- named return values
- variadic functions
- closures

These exercises are designed to be increasingly challenging as you progress.

### Exercise 1: Basic Function

**Objective**: Create a function `greet` that takes a `string` name and prints a greeting message.

```
// Example usage:
// greet("Alice") should print "Hello, Alice!"
```

### Exercise 2: Function with Return Value

**Objective**: Write a function `multiply` that takes two `int` values and returns their product.

```
// Example usage:
// result := multiply(3, 4) // result should be 12
```

### Exercise 3: Function with Return Value

**Objective**: Write a function `average` that takes an array of `float64` (e.g., `[]float64`) and returns the average of all numbers in the array.

```
// Example usage:
// result := average([]float64{98,93,77,82,83}) // result should be 86.6
```

### Exercise 4: Multiple Return Values

**Objective**: Implement a function `swap` that takes two `int` values and returns them in reverse order.

```
// Example usage:
// a, b := swap(5, 10) // a should be 10, b should be 5
```

### Exercise 5: Named Return Values

**Objective**: Create a function `rectProperties` that calculates and returns the area and perimeter of a rectangle given its length and width.

```
// Example usage:
// area, perimeter := rectProperties(4, 5) // area should be 20, perimeter should be 18
```

### Exercise 6: Variadic Function

**Objective**: Write a variadic function `sum` that sums an arbitrary number of `int` values.

```
// Example usage:
// total := sum(1, 2, 3, 4) // total should be 10
```

### Exercise 7: Closure

**Objective**: Implement a function `counter` that returns a closure. Each time the closure is called, it should return an incremented value.

```
// Example usage:
// count := counter()
// count() // returns 1
// count() // returns 2
```

### Exercise 8: Recursive Function

**Objective**: Write a recursive function `fibonacci` that returns the **n**th Fibonacci number.

```
// Example usage:
// fib := fibonacci(5) // fib should be 5 (sequence: 0, 1, 1, 2, 3, 5)
```

### Exercise 9: Error Handling in Functions

**Objective**: Create a function `sqrt` that calculates the square root of a number. If the input is negative, the function should return an error.

```
// Example usage:
// result, err := sqrt(16) // result should be 4, err should be nil
// result, err := sqrt(-1) // result should be 0, err should be non-nil
```

Work through these exercises at your own pace.
Each one targets a specific concept related to functions in Go, helping you to gradually build your proficiency.
The key to learning Golang is practice, so don't hesitate to experiment and modify these exercises as you learn!
