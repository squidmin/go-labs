# The `for range` Construct

This section discusses another variation for running loops in Go; i.e., the `for range` construct.

## Topics

- Infinite loop
- Use of `for range`

## Infinite loop

The condition can be absent, as in:

```
for i := 0; ; i++
```

or

```
for { }
```

This is the same as `for ;; { }`, but the two semicolons are removed by `gofmt`.
These are **infinite loops**.
The latter could also be written as:

```
for true { }
```

but the normal format is

```
for { }
```

In Go, if a `for` loop lacks a condition, it creates an **infinite loop**, necessitating an exit mechanism within the loop body to prevent endless execution.
This is typically achieved through a `break` statement to exit the loop or a `return` statement to exit the entire function.
It's crucial to ensure an exit condition will eventually be met to avoid infinite loops.

## Use of `for range`

The `for range` loop in Go is a powerful and elegant iteration mechanism that allows looping through each item in a collection, such as arrays, slices, or maps, making it highly versatile for various programming scenarios.

```
for ix, val := range coll { }
```

The `range` keyword iterates over an indexed collection, returning an index `ix` and a copy of the value `val` at that index for each iteration.
Initially, it returns the first element (`ix == 0 ` and `val == coll[0]`).
This mechanism is akin to the `foreach` loop in other languages, with the added benefit of accessing the index `ix` at each step.

When using `for range` in Go to iterate over a collection, the variable `val` represents a copy of the item at the current index `ix`.
This means `val` can only be used to read the item's value and cannot be used to modify the item in the original collection.

Let's solve the issue we encountered when printing characters of a Unicode string.

```go
package main

import "fmt"

func main() {
	str := "Demo: Printing characters"
	
	// for range
	for pos, char := range str {
		fmt.Printf("Character on position %d: %c\n", pos, char)
    }
	fmt.Println()
	str2 := "Japanese: 日本語"
	
	// for range
	for pos, char := range str2 {
		fmt.Printf("Character %c starts at byte position %d\n", char, pos)
    }
}
```

The above code demonstrates the use of the `for range` loop in Go to iterate over strings.
Two strings, `str`, and `str2`, are declared and initialized with `"Demo: Printing characters"`, and `"Japanese: 日本語"` respectively.
The `for range` loop is used to iterate over each character of these strings, printing the character along with its position.
It highlights that English characters occupy ` byte`, while Japanese characters occupy 3 bytes.
This example illustrates how Go handles strings iteration and character encoding.
