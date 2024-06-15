# Strings and `strconv` Package

In this section, you'll study strings, the `strconv` package, and the functions supported by them.

## Topics

- Prefixes and suffixes
- Testing whether a string contains a substring
- Indicating the index of a substring or character in a string
- Replacing a substring
- Counting occurrences of a substring
- Repeating a string
- Changing the case of a string
- Trimming a string
- Splitting a string
  - Splitting a string on whitespaces
  - Splitting a string on a separator token
- Joining over a slice
- Reading from a string
- Conversion to and from a string

Strings are a basic data structure, and every language has a number of predefined functions for manipulating strings.
In Go, these are gathered in a package, `strings`.
We'll discuss some very useful functions.

## Prefixes and suffixes

`HasPrefix` tests whether a string `s` begins with a _prefix_ `prefix`:

```
strings.HasPrefix(s, prefix string) bool
```

`HasSuffix` tests whether a string `s` ends with a _suffix_ `suffix`:

```
strings.HasSuffix(s, suffix string) bool
```

The following program implements these functions:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T / F?\nThe string \"%s\" has the prefix %s:", str, "Th")
	fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th")) // Finding prefix
	
	fmt.Printf("T / F?\nThe string \"%s\" has the suffix %s:", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting")) // Finding suffix
}
```

In the above code, a string `str` is declared and initialized with `This is an example of a string`.
The function `HasPrefix` is used to find a prefix `Th` in string `str`.
The function returns `true` because `str` does start with `Th`.
Similarly, the function `HasSuffix` is used to find the suffix `ting` in the string `str`.
The function returns `false` because `str` does not end with `ting`.
This also illustrates the use of the escape character <code>&#92;</code> to output a literal `"` with <code>&#92;"</code>, and the use of 2 substitutions in a format string.

## Testing whether a string contains a substring

The function `Contains` returns true if a string `subStr` is within a string `s`:

```
strings.Contains(s, substr string) bool
```

## Indicating the index of a substring or character in a string

`Index` returns the index of the last instance of `str` in `s`, or `-1` if `str` is not present in `s`:

```
strings.LastIndex(s, str string) int
```

If `ch` is a _non-ASCII_ character, use:

```
strings.IndexRune(s string, ch int) int
```

The following program finds the index of the substrings in a string:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Example: Finding the index of substrings in a string"
	
	fmt.Printf("The position of the first instance of \"Example\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Example"))
	
	fmt.Printf("The position of the first instance of \"index\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "index"))
	
	fmt.Printf("The position of the last instance of \"index\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "index"))
	
	fmt.Printf("The position of the first instance of \"substrings\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "substrings"))
	
	fmt.Printf("The position of the first occurrence of \"Golang\" is:")
  fmt.Printf("%d\n", strings.Index(str, "Golang"))
}
```

In the above code, a string `str` is declared and initialized with `Example: Finding the index of substrings in a string`.
Subsequently, the index of the first occurrence of `Example` in `str` is found using the `Index` function.
Similarly, we find the index of the first occurrence of `index` in `str` using the `Index` function, as well as its last occurrence by using the `LastIndex` function.
Finally, the last occurrence of `substrings` in `str` is found, and the call to `Index` to find the first occurrence of the string `Golang` in `str` returns `-1` because it doesn't exist in `str`..

## Replacing substring

We can replace an old string with a new string like:

```
strings.Replace(str, old, new string, n int)
```

We can replace the first `n` occurrences of `old` in `str` with `new`.
A copy of `str` is returned, and if `n` is `-1`, all occurrences are replaced.

## Counting occurrences of a substring

`Count` counts the number of non-overlapping instance of a substring `str` in `s` with:

```
strings.Count(s, str string) int
```

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, here's a simple Go program that counts the occurrences of a substring"
	var manyG = "gggggggggg"
	fmt.Printf("Number of H's in %s: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H")) // Count occurrences
	fmt.Printf("Number of double g's in %s: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) // Count occurrences
}
```

In the above code, we declare the two strings `str` and `manyG` and initialize them with

- `Hi, here's a simple Go program that counts the occurrences of a substring`
- `gggggggggg`

respectively.
We find the count of `H` in `str` using the `Count` function.
Similarly, we find the count of `gg` in `manyG` using the `Count` function.

## Repeating a string

The `Repeat` function returns a new string consisting of `coutn` copies of the string `s`:

```
strings.Repeat(s, count int) string
```

The following program demonstrates an implementation of the `Repeat` function:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Original string"
	var newS string
	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s", newS)
}
```

In the above code, we declare a string `origS` and initialize it with `Original string`.
We repeat `origS` **3** times using the `Repeat` function and store the new string in `newS`.
The result is printed afterwards.

## Changing the case of a string

The `ToLower` function returns a _copy_ of the string `s` with all Unicode letters mapped to their lower case:

```
strings.ToLower(s) string
```

All uppercase is obtained with:

```
strings.ToUpper(s) string
```

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Here's another example string"
	var lower string
	var upper string
	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig) // Changing to lower case
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig) // Changing to upper case
	fmt.Printf("The uppercase string is: %s\n", upper)
}
```

In the above code, we declare a string `orig` and initialize it with `Here's another example string`.
We change the case of `orig` to lowercase using the function `ToLower` and the case of `orig` to uppercase using the function `ToUpper`.
The new strings are then printed to verify the result.

## Trimming a string

`TrimSpace` can be used to remove all leading and trailing whitespaces:

```
strings.TrimSpace(s)
```

If you want to trim a specific string `str` from a string `s`, use:

```
strings.Trim(s, str)
```

For example:

```
strings.Trim(s, "\r\n")
```

The above statement will remove all _leading_ and _trailing_ `\r` and `\n` characters from the string `s`.
The 2nd string parameter can contain any characters, which are all removed from the left and right side of `s`.

If you want to remove only the leading or only the trailing characters or strings, use `TrimLeft` or `TrimRight` independently.

## Splitting a string

### On whitespaces

The `strings.Fields(s)` function divides a string `s` into a slice of substrings based on whitespace, returning an empty slice if `s` is only whitespace.

### On a separator

The `strings.Split(s, sep)` function takes a string `s` and a separator `sep`, then divides `s` into a slice of substrings based on the occurrences of `sep`.
Unlike `strings.Fields(s)` which automatically splits a string based on whitespace, `strings.Split` allows for custom separators.
These separators can be single characters like commas, colons, semicolons, or even strings of multiple characters.
This function is useful for parsing data that is delimited in a specific way, such as CSV (comma-separated values) or log files with custom delimiters.

### Joining over a slice

The `strings.Join(sl []string, sep string)` function in Go's `strings` package concatenates the elements of the slice `sl` into a single string, with each element separated by the string `sep`.
This is useful for combining multiple strings together with a specific separator, such as commas or spaces.

## Reading from a string

The `strings` package also has a function called `strings.NewReader(str)`.
This produces a _pointer_ to a _Reader_ value that provides amongst others the following functions to operate on `str`:

- `Read()` to read a `[]byte`
- `ReadByte()` to read the next byte from a string
- `ReadRune()` to read the next rune from a string

## Conversion to and from a string

The `strconv` package contains a few variables to calculate the size in bits of the `int` of a given platform on which the program runs:

```
strconv.IntSize
```

Converting a variable of a certain type to a string will always succeed.
For converting from numbers, we have the following functions:

```
strconv.Itoa(i int) string
```

This returns the decimal string representation of `i`. Next, we have:

```
strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
```

This converts the 64-bit floating-point number `f` to a string, according to the format `fmt` (`b`, `e`, `f`, or `g`); precision `prec`, with `bitSize` being 32 bits for `float32` and 64 bits for `float64`.

For converting to numbers,, we have the following functions:

```
strconv.Atoi(s string) (i int, err error)
```

It converts to an `int`. Second, we have:

```
strconv.ParseFloat(s string, bitSize int) (f float64, err error)
```

It converts to a 64-bit floating-point number.

As can be seen from the return type, these functions will return **2** values:

- the converted value (if possible)
- the possible error

When calling such a function, the multiple assignment form will be used:

```
val, err = strconv.Atoi(s)
```

Converting a string to another type will not always be possible (as in the case of the functions `Atoi` and `ParseFloat`).
Hence, a _runtime_ error is thrown: `parsing "...": invalid argument`.

```go
package main
import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "whelp"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig) // Converting to a number
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an) // Converting to a string
	fmt.Printf("The new string is: %s\n", newS)
}
```

In the above code, we declare a string `orig` and initialize it with `whelp`.
We calculate the size of ints using `strconv.IntSize`, and convert `orig` to an integer using the function `strconv.Atoi(orig)`.
The result of this call to `strconv.Atoi` is stored in a new variable `an`.
We modify `an` and convert it into a string using `strconv.Itoa(an)` and then store it in a new variable `newS`.
`an` and `newS` are printed to verify the results.

Documentation for other functions in this package can be found [here](https://pkg.go.dev/strconv).
