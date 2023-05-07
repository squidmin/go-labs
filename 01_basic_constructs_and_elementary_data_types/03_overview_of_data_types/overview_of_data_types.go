package main

import "fmt"

func main() {
	coerceType()
}

/**
 * Overview of Data Types
 *
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

/**
 * Integer types:
 * uint8, uint16, uint32, uint64
 * int8, int16, int32, int64
 */
var unsignedIntExample1 uint8
var unsignedIntExample2 uint16
var unsignedIntExample3 uint32
var unsignedIntExample4 uint64
var signedIntExample1 int8
var signedIntExample2 int16
var signedIntExample3 int32
var signedIntExample4 int64

/**
 * 3 machine-dependent (size occupied depends on architecture) integer types:
 * uint
 * int
 * uintptr
 */
var machineDependentUnsignedInt uint
var machineDependentSignedInt int
var machineDependentUnsignedIntPtr uintptr

/**
 * Alias types:
 * byte - Same as uint8
 * rune - Same as int32
 */
var aliasedUnsignedInt1 byte
var aliasedUnsignedInt2 rune

/**
 * Go has two floating-point types: float32 and float 64 (also often referred to as single precision and double precision respectively.)
 */

/**
 * Generally if you are working with integers you should just use the 'int' type.
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
 * Data type conversions
 */

func coerceType() {
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number)) // Print the result of the type cast
}

/**
 * Basic string operations
 */
func stringOperations() {
	fmt.Println(len("Hello World")) // len() usage
	fmt.Println("Hello World"[1])   // String indexing
	fmt.Println("Hello " + "World") // String concatenation
}
