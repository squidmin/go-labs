package main

import "fmt"

func demoAliasTypesBuiltIn() {
	var aliasedUnsignedInt byte
	fmt.Println(aliasedUnsignedInt)
	var aliasedSignedInt rune
	fmt.Println(aliasedSignedInt)
}
