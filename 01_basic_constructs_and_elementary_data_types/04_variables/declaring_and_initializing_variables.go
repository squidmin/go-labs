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

	/*
	 * Generally, the shorter form of declaration & initialization given by type inference should be used whenever possible,
	 *   whether with 'var' or the ':=' operator.
	 */
}
