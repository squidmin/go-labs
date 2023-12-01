package main

import "fmt"

/**
 * Integer types:
 * uint8, uint16, uint32, uint64
 * int8, int16, int32, int64
 */
func demoIntegerTypes() {
	var unsignedIntExample1 uint8
	fmt.Println(unsignedIntExample1)
	var unsignedIntExample2 uint16
	fmt.Println(unsignedIntExample2)
	var unsignedIntExample3 uint32
	fmt.Println(unsignedIntExample3)
	var unsignedIntExample4 uint64
	fmt.Println(unsignedIntExample4)

	var signedIntExample1 int8
	fmt.Println(signedIntExample1)
	var signedIntExample2 int16
	fmt.Println(signedIntExample2)
	var signedIntExample3 int32
	fmt.Println(signedIntExample3)
	var signedIntExample4 int64
	fmt.Println(signedIntExample4)
}

/**
 * Generally if you are working with integers you should just use the 'int' type.
 */

func main() {
	demoIntegerTypes()
}
