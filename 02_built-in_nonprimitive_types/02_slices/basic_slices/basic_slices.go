package main

import "fmt"

func slices1() {
	fmt.Println("Example 1: append function")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func slices2() {
	fmt.Println("Example 2: copy function")
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func main() {
	slices1()
	slices2()
}
