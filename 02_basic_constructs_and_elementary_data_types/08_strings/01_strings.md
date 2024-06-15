# Strings

This section discusses the datatype string in detail.

## Topics

- Introduction
- Strings in Go
  - Advantages of strings
  - Types of string literals
    - Interpreted strings
    - Raw strings
  - Length of a string
  - Concatenation of strings

## Introduction

Strings in Go are sequences of UTF-8 characters, utilizing 1-byte for ASCII and 2-4 bytes for other characters as needed.
UTF-8 is the predominant encoding for various text formats.
Go optimizes storage by using only 1 byte for ASCII characters within strings.

## Strings

Contrary to strings in other languages like C++, Java, or Python that are fixed-width (Java always uses 2 bytes), a Go string is a sequence of variable-width characters (each 1 to 4 bytes).

### Advantages of strings

The advantages of strings in Go include:

- Go strings and text files occupy less memory / disk space (because of variable-width characters).
- Since UTF-8 is the standard, Go doesn't need to encode and decode strings like other languages have to do.

Strings are value types and immutable.
This means that once a string is created, its content cannot be changed.
If you attempt to modify a string, what actually happens is the creation of a new string with the desired changes, leaving the original string unaltered.
This behavior is due to strings being treated as immutable arrays of bytes; each character in the string is stored as one or more bytes, and the sequence of these bytes does not change once the string is created.

This immutability property of strings in Go is crucial for understanding how strings behave in memory and how operations on strings might affect performance, especially in large-scale systems where efficiency and memory management are critical.

### Types of string literals

Two kinds of string literals exist in Go:

- interpreted strings
- raw strings

#### Interpreted strings

Interpreted strings are surrounded by double quotes (`""`). For example, escape sequences are interpreted:

```
"\n" // Represents a newline
"\r" // Represents a carriage return
"\t" // Represents a tab
"\u" // Represents Unicode characters
"\U" // Also represents Unicode characters
```

#### Raw strings

Raw strings are surrounded by back quotes (<code>`</code>) and are not interpreted.
For example in the following string:

```
`This is a raw string \n`
```

`\n` is not interpreted but taken literally. They can span multiple lines.

### Length of a string

In Go, strings are not ended with a special character like they are in C / C++.
Instead, an empty string (`""`) is considered the default value.
When comparing strings, Go uses the comparison operators to compare them byte by byte.

The `len()` function in Go returns the number of bytes in a string, not the number of characters.
You can access individual bytes in a string using standard indexing, which starts at **0**.

For a string `str`:

- The first byte is given by: `str[0]`
- The `i`-th byte is given by: `str[i]`
- The last byte is given by: `str[len[str] - 1]`

> **Note**: Taking the address of a character in a string is not valid syntax.

### Concatenation of strings

Two strings `s1` and `s2` can be made into one string `s` with:

```
s := s1 + s2
```

`s2` is appended to `s1` to form a new string `s`. Multi-line strings can be constructed as follows:

```
str := "Beginning of the string " +
"second part of the string"
```

The `+` operator has to be on the _first line_ due to the insertion of `;` by the compiler.
The append shorthand `+=` can also be used for strings.

```go
package main

import "fmt"

func main() {
	s := "Hel" + "lo"
	s += "World"
	fmt.Println(s)
}
```

The two strings `Hel` and `lo` are concatenated, and the result of this concatenation is further concatenated with the string `World`.
The final string `Hello World` is subsequently printed.
