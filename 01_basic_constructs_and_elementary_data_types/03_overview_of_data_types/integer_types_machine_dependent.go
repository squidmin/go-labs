package main

import "fmt"

func demoMachineDependentIntegerTypes() {
	var unsignedIntExample uint
	fmt.Println(unsignedIntExample)
	var signedIntExample int
	fmt.Println(signedIntExample)
	var unsignedIntPtrExample uintptr
	fmt.Println(unsignedIntPtrExample)
}
