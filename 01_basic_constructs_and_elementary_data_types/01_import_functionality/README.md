# Import functionality

### Single package

```go
import "fmt"
```

### Multiple packages

```go
import (
    "fmt"
    "rsc.io/quote"
)
```


---


### Code example

```go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println("Console output using 'fmt' package...")
    fmt.Println(quote.Go())
}
```
