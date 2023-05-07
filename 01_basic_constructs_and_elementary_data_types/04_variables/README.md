# Variables


Names must start with a letter and may contain letters, numbers or the underscore (`_`) symbol.

#### Code example

```go
package main

import "fmt"

func main() {
	var someInt1 int // Declaration
	fmt.Println("someInt1 == ", someInt1)
    var someInt2 int = 3  // Declaration + initialization
    fmt.Println("someInt2 == ", someInt2)
	
	// If the declaration and initialization occur on one line, the type can be omitted:
	var someInt3 = 5
	fmt.Println("someInt3 == ", someInt3)
}
```