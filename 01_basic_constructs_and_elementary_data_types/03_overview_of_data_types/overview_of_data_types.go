package main

import "fmt"

func main() {
	coerceType()
}

/**
 * Overview of Data Types
 * This lesson tells us about the type of data that Go can handle.
 * The following topics are covered:
 * - Types
 *   - Conversions
 */

/**
 * Types                                Examples
 *
 * elementary (or primitive)            int, float, bool, string
 * structured (or composite)            struct, array, slice, map, channel
 * interfaces                           They describe the behavior of a type.
 */

var var1 int // Declare a variable "var1" of type "int".

func functionName(a int, b bool) (c float32) { // Function types
	return 2.0
}

// type IZ int  // Data type aliases
// var a IZ = 5

// If you have more than one type to define, you can use the factored keyword form, as in:
type (
	IZ  int
	FZ  float32
	STR string
)

/**
 * Conversions
 */

func coerceType() {
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number)) // Print the result of the type cast
}
