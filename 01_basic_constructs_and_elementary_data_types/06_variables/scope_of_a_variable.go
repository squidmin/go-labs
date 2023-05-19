package main

import "fmt"

var number int = 5 // Number declared outside (global scope).

func scopeOfAVariable() {
	fmt.Println("Demo: Scope of a variable")
	var decision bool = true // Decision declared inside function (local scope).
	fmt.Println("Original value of number:", number)
	number = 10
	fmt.Println("New value of number:", number)
	fmt.Println("Value of decision:", decision)
}
