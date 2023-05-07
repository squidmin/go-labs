package main

import "fmt"

func main() {
	printMachineDependentIntegerTypes()
}

func printMachineDependentIntegerTypes() {
	var unsignedIntExample uint
	fmt.Println(unsignedIntExample)
	var signedIntExample int
	fmt.Println(signedIntExample)
	var unsignedIntPtrExample uintptr
	fmt.Println(unsignedIntPtrExample)
}
