# Testing for Errors on Functions

In this section, you'll learn about errors that originate from functions and how to resolve them.

## Topics

- Testing support
- Tracking an error in a function

## Testing support

In Go, functions can return both a value and an execution status, where the status is indicated by either a boolean or an error variable.
Successful executions return a value with `true` or `nil` error, while failures return a potentially `nil` value with `false` or an error.
This allows for concise error checking using the "`comma, ok`" pattern.

Consider the following example:

```
v, ok = sampleFunction(parameter)
```

The "comma, ok" pattern in Go programming is used to check for errors.
The function returns two values: the result (`v`) and a boolean (`ok`) indicating success or failure.
If `ok` is true, the operation succeeded without errors; otherwise, and error occurred.

## Tracking an error in a function

The `strconv.Atoi` function from Go's standard library can be used to attempt converting a string to an integer.
The function returns two values: the converted integer and an error.
In previous examples of `strconv.Atoi` usage, the error is intentionally ignored using the blank identifier (`_`).
This means that if a string cannot be successfully converted to an integer (e.g., if it contains non-numeric characters), the error that `strconv.Atoi` returns is discarded, and the program does not handle or acknowledge the failure.
This approach can lead to issues in the program if the error is critical to subsequent operations.

A program must handle all potential errors by notifying the user about the issue, exiting the function, or stopping the program to ensure reliability and user awareness of issues.

Here's another example program to demonstrate testing for errors:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var an int
	var err error
	
	// Storing an integer and error information
	an, err = strconv.Atoi(orig)
	if err != nil { // If it was an error, discontinue
		fmt.Printf("'orig' %s is not an integer. Exiting with error", orig)
		return
    }
	fmt.Printf("The integer is %d\n", an)
	// Rest of the code
}
```

Two variables, `an` for storing a converted integer and `err` for any error information, are declared.
The `strconv.Atoi` function attempts to convert a string `orig` to an integer.
If an error occurs (`err != nil`), the program exits early.
Otherwise, it continues, but will still terminate with an error if `orig` is not an integer, demonstrating basic error checking and handling in Go.

There is an alternative approach available for handling the errors where, instead of immediately terminating the program upon encountering an error (using `return` or `os.Exit`), the function returns the error object (`return err`).
This enables the caller of the function to handle the error approximately based on its context or logic, providing more flexibility in error management.

```go
value, err := pack1.Function1(param1)
if err != nil {
	fmt.Printf("An error occurred in pack1.Function1 with parameter %v", param1)
	return err
}
// Normal case, continue execution...
```

If an error occurs within the `main()` function of a Go program, using `os.Exit(1)` instead of `return` will immediately terminate the program.
This approach is used when it's necessary to stop the program execution due to an error.

```go
if err != nil {
	fmt.Printf("Program stopping with error %v", err)
	os.Exit(1)
}
```

Look at the following code where we attempt to open a file:

```go
f, err := os.Open(fname)
if err != nil {
	return err
}
doSomething(f) // In case of no error, the file 'f' is passed to a function doSomething
```

Sometimes, the above idiom is repeated a number of times in succession.
No `else` branch is written.
If there is no error-condition, the code simply continues execution after the `if (...) { ... }`

> **Remark**: More information on the `os` package can be found [here](https://pkg.go.dev/os).
