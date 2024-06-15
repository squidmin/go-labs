package main

import "fmt"

func main() {
	// Creating a slice using make
	slice1 := make([]int, 5)
	for i := range slice1 {
		slice1[i] = i * 10
	}

	// Creating a slice using a slice literal
	slice2 := []string{"hello", "world"}

	// Creating a slice from an array
	arr := [5]int{0, 1, 2, 3, 4}
	slice3 := arr[1:4] // includes elements 1, 2, and 3

	// Printing slices
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
