# Variables, constants, and operators

```go
// variables_constants_operators.go
package main

import "fmt"

func main() {
	// ----- Variables -----
	// Explicit type
	var a int = 10
	// Type inference
	b := 3

	// Multiple variables
	var x, y int = 4, 5
	z := 7.5

	fmt.Println("a:", a, "b:", b, "x:", x, "y:", y, "z:", z)

	// ----- Constants -----
	const pi = 3.14159           // untyped constant
	const maxConnections int = 5 // typed constant

	// iota for enumerations
	const (
		StatusOK = iota
		StatusWarning
		StatusError
	)

	fmt.Println("pi:", pi, "maxConnections:", maxConnections)
	fmt.Println("StatusOK:", StatusOK, "StatusWarning:", StatusWarning, "StatusError:", StatusError)

	// ----- Arithmetic operators -----
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("sum:", sum, "diff:", diff, "product:", product, "quotient:", quotient, "remainder:", remainder)

	// ----- Assignment operators -----
	a += 5  // a = a + 5
	b *= 2  // b = b * 2
	x -= 1  // x = x - 1
	y /= 2  // y = y / 2
	fmt.Println("After assignment ops:", "a:", a, "b:", b, "x:", x, "y:", y)

	// ----- Comparison operators -----
	isEqual := a == b
	isNotEqual := a != b
	less := x < y
	greaterOrEqual := z >= 7.5

	fmt.Println("isEqual:", isEqual)
	fmt.Println("isNotEqual:", isNotEqual)
	fmt.Println("less:", less)
	fmt.Println("greaterOrEqual:", greaterOrEqual)

	// ----- Logical operators -----
	logicalAnd := (a > 0) && (b > 0)
	logicalOr := (a < 0) || (b > 0)
	logicalNot := !(a > 0)

	fmt.Println("logicalAnd:", logicalAnd)
	fmt.Println("logicalOr:", logicalOr)
	fmt.Println("logicalNot:", logicalNot)

	// ----- Bitwise operators -----
	// Use small ints for clarity
	var m uint8 = 0b0011 // 3
	var n uint8 = 0b0101 // 5

	bitAnd := m & n  // AND
	bitOr := m | n   // OR
	bitXor := m ^ n  // XOR
	shiftLeft := m << 1
	shiftRight := n >> 1

	fmt.Printf("m: %08b, n: %08b\n", m, n)
	fmt.Printf("AND: %08b\n", bitAnd)
	fmt.Printf("OR : %08b\n", bitOr)
	fmt.Printf("XOR: %08b\n", bitXor)
	fmt.Printf("<< : %08b\n", shiftLeft)
	fmt.Printf(">> : %08b\n", shiftRight)
}
```
