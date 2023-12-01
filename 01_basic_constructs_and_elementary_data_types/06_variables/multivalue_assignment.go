package main

import "fmt"

func demoMultivalueAssignment() {
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println("a ==", a)
	fmt.Println("b ==", b)
	fmt.Println("c ==", c)
}

func main() {
	demoMultivalueAssignment()
}
