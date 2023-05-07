package main

import "fmt"

func main() {
	stringOperations()
}

func stringOperations() {
	fmt.Println(len("Hello World")) // len() usage
	fmt.Println("Hello World"[1])   // String indexing
	fmt.Println("Hello " + "World") // String concatenation
}
