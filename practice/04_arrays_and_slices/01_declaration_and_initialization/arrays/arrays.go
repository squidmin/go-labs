package main

import "fmt"

// Arrays
func main() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	for i, value := range arr1 {
		fmt.Printf("Value at index %d is %d\n", i, value)
	}
}
