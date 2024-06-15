# Constants

This section discusses how constants are used to store data values in Go.


The following topics are covered:
- Introduction
- Explicit and implicit typing
- Typed and untyped constants
- Compilation
- Overflow
- Multiple assignments
- Enumerations


---


### Introduction

A value that _cannot_ be changed by a program is called a **constant**.
This data can only be of type _boolean_, _number_ (integer, float, or complex), or _string_.


---


### Explicit and implicit typing

In Go, a constant can be defined using the keyword `const` as:

```go
const identifier [type] = value
```

Here, `identifier` is the name, and `type` is the type of constant.
Below is an example of a declaration:

```go
const PI = 3.14159
```

You may have noticed that we didn't specify the type of constant `PI` here.
It's perfectly fine because the type specifier [type] is _optional_ because the compiler can implicitly derive the type from the value.
Let's look at another example of implicit typing:

```go
const B = "hello"
```

The compiler knows that the constant `B` is a string by looking at its value.
However, you can also write the above declaration with explicit typing as:

```go
const B string = "hello"
```

> **Remark**: There is a convention to name constant identifiers with all uppercase letters, e.g., `const INCHTOCM = 2.54`. This improves readability.


---


### Typed and untyped constants

Constants declared through explicit typing are called _typed constants_, and constants declared through implicit typing are _untyped constants_.
A value derived from an untyped constant becomes typed when it is used within a context that requires a typed value. For example:

```go
var n int
f(b + 5)  // untyped numeric constant "5" becomes typed as int, because n was int.
```


---


### Compilation

Constants must be evaluated at compile-time.
A const can be defined as a calculation, but all the values necessary for the calculation must be available at compile time.
See the case below:

```go
const C1 = 2/3  // OK
```

Here, the value of `C1` was available at compile time. But the following will give an error:

```go
const C2 = getNumber()  // Not OK
```

Because the function `getNumber()` can't provide the value at compile-time.
A constant's value should be known at compile time according to the design principles where the function's value is computed at run time.
So, it will give the build error: `getNumber() used as value`.


---


### Overflow

_Numeric constants_ have no size or sign.
They can be of _arbitrarily high precision_ and do not overflow:

```go
const Ln2= 0.693147180559945309417232121458\
176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const BILLION = 1e9 // float constant
const HARD_EIGHT = (1 << 100) >> 97
```

We used `\` (backslash) in declaring constant `Ln2`.
It can be used as a _continuation character_ in a constant.


---


### Multiple assignments

The assignments made in one single assignment statement are called multiple assignments.
Go allows different ways of doing multiple assignments. Let's start with a simple example:

```go
const BEEF, TWO, C = "meat", 2, "veg"
```

As you can see, we made 3 constants.
All of them are _untyped_ constants.
Let's look at another method where all the constants are named first, and then their values are written if needed. For example:

```go
const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
```

As you can see, the constants `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY`, and `SATURDAY` are _typed_ constants because their type (`int`) is mentioned explicitly, and they have the values 1, 2, 3, 4, 5, and 6 respectively.


---


### Enumerations

Listing of all elements of a set is called _enumeration_.
Constants can be used for enumerations. For example:

```go
const {
	ZERO = 0
	ONE = 1
	TWO = 2
}
```

`ZERO`, `ONE`, and `TWO` are now aliases for the `int` values 0, 1, and 2.
Interestingly, the value `iota` can be used to enumerate the values.
Let's enumerate the above example with `iota`:

```go
const {
	ZERO = iota
	ONE = iota
	TWO = opta
}
```

The first use of `iota` gives 0.
Whenever `iota` is used again on a new line, its value is incremented by 1; so `ZERO` gets 0, `ONE` gets 1, and `TWO` gets 2.

Remember that a new `const` block or declaration initializes `iota` back to 0.
The above notation can be shortened, making no difference as:

```go
const {
	ZERO = iota
	ONE
	TWO
}
```

You can give an enumeration a type name.
For example, `ZERO`, `ONE`, and `TWO` are each an **English digit**.
Let's give them `EnglishDigit` as the type name:

```go
type EnglishDigit int
const {
	ZERO EnglishDigit = iota
	ONE
	TWO
}
```
