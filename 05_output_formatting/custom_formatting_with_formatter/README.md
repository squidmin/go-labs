# Custom Formatting with `fmt.Formatter`

In Go, the `fmt` package provides a flexible way to format text, including predefined formats for basic data types.
For types where default formatting is not sufficient, Go allows you to define custom formatting by implementing the `fmt.Formatter` interface.
This document explains an implementation of custom formatting for a Person struct using the `fmt.Formatter` interface.

## Understanding the `fmt.Formatter` Interface

The `fmt.Formatter` interface has a single method, `Format`, which is used to format the value of a type.
The `Format` method has the following signature:

```go
Format(f fmt.State, c rune)
```

- `f fmt.State`: Holds flags, width, and precision details for the current formatting directive.
- `c rune`: The formatting verb (e.g., 'v', 's', 'd', etc.) that triggered the call to `Format`.

Implementing this interface allows you to control how your type is formatted by the `fmt` package's printing functions (`Printf`, `Println`, etc.).

## Example: Formatting a `Person` Type

### Defining the `Person` Struct

Consider a simple `Person` struct with `Name` and `Age` fields:

```go
type Person struct {
    Name string
    Age  int
}
```

Our goal is to customize the formatting of instances of this type when using the `fmt` package's printing functions.
The different formats will be based on the formatting verb and flags used in the format string.

### Implementing `fmt.Formatter` for `Person`

We implement custom formatting logic inside the `Format` method:

```go
func (p Person) Format(f fmt.State, c rune) {
	switch c {
	case 'v': // 'v' acts as a generic verb
		if f.Flag('+') {
			fmt.Fprintf(f, "%s (%d years old)", p.Name, p.Age)
		} else {
			fmt.Fprintf(f, "%s", p.Name)
		}
	}
}
```

- When the formatting verb is 'v', we check for the '+' flag.
- If the flag is present, we format the `Person` instance to include the person's name and age.
- Without the `+` flag, we format the `Person` instance to include only the person's name.

### Demonstrating Custom Formatting

The `FormatterExample` function demonstrates using the custom formatting:

```go
func FormatterExample() {
    person := Person{Name: "Test User", Age: 30}
    fmt.Printf("Name: %v\n", person)           // Default format
    fmt.Printf("Custom: %+v\n", person)        // Custom format triggered by '+'
}
```

- `fmt.Printf("Name: %v\n", person)` prints the `Person` instance using the default format (just the `name`).
- `fmt.Printf("Custom: %+v\n", person)` uses the custom format to include both the `name` and the `age`.

### Running the Example

When `FormatterExample` is called, the output demonstrates the two different formats:

```
Name: Test User
Custom: Test User (30 years old)
```

This implementation showcases the flexibility of Go's formatting system, allowing types to have sophisticated and context-sensitive string representations that can be easily printed using standard functions from the `fmt` package.
By implementing the `fmt.Formatter` interface, you can ensure that your types are always represented in a way that is meaningful for your application's logging, debugging, and output needs.
