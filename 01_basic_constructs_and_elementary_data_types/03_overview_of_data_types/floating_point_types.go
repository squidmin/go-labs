package main

import "fmt"

func demoFloatingPointTypes() {
	var floatExample1 float32
	fmt.Println(floatExample1)
	var floatExample2 float64
	fmt.Println(floatExample2)
	var complexExample1 complex64
	fmt.Println(complexExample1)
	var complexExample2 complex128
	fmt.Println(complexExample2)

	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)
	fmt.Println("1 + 1.0 = ", 1+1.0)
}
