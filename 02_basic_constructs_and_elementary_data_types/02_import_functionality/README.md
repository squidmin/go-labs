# Import functionality

This section discusses a basic component of a Go program, e.g., _packages_.


The following topics are covered:
- Packages
  - Package dependencies
- Import keyword
- Visibility
  - Visibility rule


---


## Packages

A library, module, or namespace in any other language is called a **package**.
Packages are a way to structure code.
A program is constructed as a _package_ which may use facilities from other packages.
A package is often abbreviated as "_pkg_".

Every Go file belongs to only _one_ package whereas one package can comprise many different Go files.
Hence, the filename(s) and the package names are generally _not_ the same.
The package to which the code-file belongs must be indicated on the _first_ line.
A package name is written in _lowercase_ letters.
For example, if your code-file belongs to a package called **main**, do the following:

```go
package main
```

A standalone executable belongs to _main_. Each Go application contains one _main_.

An application can consist of different packages.
But even if you use only package _main_, you don't have to place all code in one big file.
You can make a number of smaller files, each having `package main` as the first line of code.
If you compile a source file with a package name other than main, e.g., **`pack1`**, the object file is stored in `pack1.a`.


---


## Import keyword

A Go program is created by linking sets of packages together, with the `import` keyword.
For example, if you want to import a package, sat `fmt`, then you do:

```go
package main
import "fmt"
```

`import "fmt""` tells Go that this program needs functions, or other elements from the package `fmt`, which implements a functionality for formatted IO.
The package names are enclosed within `""` (double-quotes).

Import loads the public declarations from the compiled package; it does not insert the source code.
If multiple packages are needed, they can each be imported by a separate statement.
For example, if you want to import two packages, `fmt` and `os` in one code file, the following indicates some ways to do so:

```go
import "fmt"
import "os"
```

or you can do:

```go
import "fmt"; import "os"
```

Factoring means calling a _keyword_ once on multiple instances.
You may have notices that we imported two packages using a single `import` keyword.
It is also applicable to keywords like `const`, `var` and `type`.


---


## Visibility

Packages contain all other code objects apart from the blank identifier (`_`).
Also, identifiers of code-objects in a package have to be unique, which means that there can be no naming conflicts.
However, the same identifier can be used in different packages.
The package name qualifies a package to be different.

### Visibility rule

Packages expose their code objects to code outside of the package according to the following rule enforced by the compiler:

When the identifier (of a constant, variable, type, function, struct field, ...) starts with an uppercase letter, like `Group1`, then the "object" with this identifier is visible in code outside the package (thus available to client-programs, or "importers" of the package), and it is said to be exported (like `public` identifiers/variables in OO languages).
Identifiers that start with a lowercase letter are not visible outside the package, but they are visible and usable in the whole package (like `private` identifiers/variables).

> **Note**: Capital letters can come from the entire _Unicode-range_, like Greek; not only ASCII letters are allowed.

Importing a package gives access only to the exported objects in that package.
Suppose we have an instance of a variable or a function called `Object` (starts with capital `O` so it is exported) in a package `pack1`.
When `pack1` is imported in the current package, `Object` can be called with the usual dot-notation from OO-languages.

```
pack1.Object
```

Packages also serve as namespaces and can help us avoid name-conflicts.
For example, variables with the _same_ name in two packages are differentiated by their package name, like `pack1.Object` and `pack2.Object`.

A package can also be given another name called an _alias_.
If you name a package then its alias will be used throughout the code, rather than its original name. For example:

```go
import fm "fmt"
```

Now in the code, whenever you want to use `fmt`, use its alias: `fm` (not `fmt`).

> **Note**: Go has a motto known as _"No unnecessary code!"_. So importing a package which is not used in the rest of the code is a _build-error_.


---


## Syntax

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
