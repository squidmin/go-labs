package main

import "fmt"

func basicArraysDemo1() {
	fmt.Println("Example 1: Basic arrays")
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println()
}

func basicArraysDemo2() {
	fmt.Println("Example 2: Basic arrays")

	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	//for i := 0; i < len(x); i++ {  // Classic for loop syntax
	//	total += x[i]
	//}
	//for i, value := range x {      // Simplified for loop syntax
	//	total += value
	//}
	// The Go compiler doesn't allow unused variables, so the 'i' variable needs to be changed to '_' in the simplified for loop syntax.
	for _, value := range x {
		total += value
	}
	// A single underscore (_) is used to tell the compiler that the variable isn't needed.
	// In this case, the iterator variable 'i' is being omitted.
	fmt.Println(total / float64(len(x)))
	fmt.Println()
}

func basicArraysDemo3() {
	fmt.Println("Example 3: Shorter array instantiation syntax")

	// Go provides a shorter syntax for creating arrays
	x := [5]float64{98, 93, 77, 82, 83}
	fmt.Println("x ==", x)

	// Can be split into multiple lines
	y := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println("y ==", y)

	fmt.Println()
}

func main() {
	basicArraysDemo1()
	basicArraysDemo2()
	basicArraysDemo3()
}
