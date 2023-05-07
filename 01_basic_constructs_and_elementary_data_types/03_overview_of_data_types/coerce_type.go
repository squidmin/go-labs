package main

import "fmt"

func main() {
	coerceType()
}

func coerceType() { // Data type conversions
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number)) // Print the result of the type cast
}
