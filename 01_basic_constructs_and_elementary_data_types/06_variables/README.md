# Variables

## Syntax

Names must start with a letter and may contain letters, numbers or the underscore (`_`) symbol.

#### Code example

```go
package main

import "fmt"

func main() {
    var someInt1 int // Declaration
    fmt.Println("someInt1 == ", someInt1)
    var someInt2 int = 3 // Declaration + initialization
    fmt.Println("someInt2 == ", someInt2)

    // If declaration and initialization occur on one line, the type can be omitted with the 'var' keyword:
    var someInt3 = 5
    fmt.Println("someInt3 == ", someInt3)
    var someInt4 = 7.8
    fmt.Println("someInt4 == ", someInt4)

    // Variables can also be initialized on one line using the ':=' operator:
    x := 5
    fmt.Println(x)
    // Go is able to infer the type based on the supplied literal value.
}
```


### Multivalue assignment

#### Code example

```go
package main

import "fmt"

func main() {
    var (
        a = 5
        b = 10
        c = 15
    )
    fmt.Println("a ==", a)
    fmt.Println("b ==", b)
    fmt.Println("c ==", c)
}
```
