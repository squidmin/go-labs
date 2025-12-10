# Functions and packages

```go
// functions_and_packages.go
package main

import (
	"fmt"
	"math"
)

// add demonstrates a simple function with parameters and a single return value.
func add(a int, b int) int {
	return a + b
}

// multiply demonstrates multiple parameters of the same type.
func multiply(a, b int) int {
	return a * b
}

// divide demonstrates multiple return values.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Rectangle demonstrates methods attached to a type.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area is a method with a value receiver.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale demonstrates a method with a pointer receiver (modifies the receiver).
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	// Calling basic functions
	sum := add(3, 5)
	product := multiply(4, 6)
	fmt.Println("sum:", sum)
	fmt.Println("product:", product)

	// Multiple return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error dividing:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Handling an error (division by zero)
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Expected error:", err)
	}

	// Using a method on a struct
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle area:", rect.Area())

	rect.Scale(2)
	fmt.Println("Scaled rectangle:", rect)
	fmt.Println("Scaled rectangle area:", rect.Area())

	// ----- Using other packages -----
	// math.Sqrt comes from the standard library `math` package.
	value := 16.0
	root := math.Sqrt(value)
	fmt.Printf("Square root of %.0f is %.2f\n", value, root)

	// Note:
	// - This file itself is in package "main".
	// - We imported fmt and math packages at the top.
	// - Functions that start with a capital letter are exported from their package
	//   (e.g., math.Sqrt), while lowercase names are not exported (e.g., add).
}
```
