package main

import "fmt"

func demoStrings() {
	fmt.Println(len("Hello World")) // len() usage
	fmt.Println("Hello World"[1])   // String indexing
	fmt.Println("Hello " + "World") // String concatenation
}

func main() {
	demoStrings()
}
