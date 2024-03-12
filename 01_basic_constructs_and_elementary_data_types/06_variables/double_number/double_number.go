package main

import "fmt"

func doubleNumber() {
	fmt.Print("Enter a number: ")
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return
	}

	output := input * 2
	fmt.Println(output)
}

func main() {
	doubleNumber()
}
