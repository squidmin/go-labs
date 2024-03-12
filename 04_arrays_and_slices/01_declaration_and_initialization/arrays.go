package main

import "fmt"

func main() {
	var arr1 [5]int // Declaration of an array
	// Initialization of array elements
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	// Iteration over array elements
	for i, value := range arr1 {
		fmt.Printf("Item at index %d is %d\n", i, value)
	}
}
