# Arithmetic Operators

The common binary operators `+`, `-`, `*`, `/`, and `/` that exist for both integers and floats in Golang are:

- `+`: Addition operator
- `-`: Subtraction operator
- `/`: Division operator
- `%`: Modulus operator
- `*`: Multiplication operator

The `+` operator also exists for strings.
The `/` operator for integers is (floored) integral division.
For example, `9 / 4` will give you `2`, and `9 / 10` will give you `0`.
The modulus operator `%` is only defined for integers.
For example, `9 % 4` gives `1`.

Integer division by `0` causes the program to crash, and a run-time panic occurs (in many cases the compiler can detect this condition).
Division by `0.0` with floating-point numbers gives an infinite result: **+Inf.**

There are shortcuts for some operations. For example, the statement:

```
b = b + a
```

can be shortened as

```
b += a
```

The same goes for `-=`, `*=`, `/=`, and `%=`.

The unary operators for integers and floats are:

- `++`: Increment operator
- `--`: Decrement operator

However, these operators can only be used after the number, which means: `i++` is short for `i += 1`, which is, in turn, short for `i = i + 1`.
Similarly, `i==` is short for `i -= 1`, which is short for `i = i - 1`.

Moreover, `++` and `--` may only be used as statements, not expressions; so `n = i++` is invalid, and subtler expressions like `f(i++)` or `a[i] = b[i++]`, which are accepted in C, C++, and Java, _cannot_ be used in Go.

No error is generated when an _overflow_ occurs during an operation because high bits are simply _discarded_.
Constants can be of help here.
If we need integers or rational numbers of unbounded size (only limited by available memory), we can use the `math` / `big` package from the standard library, which provides the types `big.Int` and `big.Rat`.

## Logical operators

Following are logical operators present in Go:

- `==`: Boolean equality operator
- `!=`: Boolean non-equality operator
- `<`: Boolean less-than operator
- `>`: Boolean greater-than operator
- `<=`: Boolean less-than or equal-to operator
- `>=`: Boolean greater-than or equal-to operator

In Go, for comparison operations, both operand values must be of the same type. If a constant is involved, it should be compatible with the other value's type.
Otherwise, type conversion is required before comparison.

The operators `<` `<=`, `>`, `>=`, `==`, and `!=` are used for comparison and work with both numbers and strings in Go, producing a boolean result (`true` or `false`).
These operators are considered logical because they evaluate the relationship between their operands and return a boolean value indicating the outcome of the comparison.

```go
package main
import "fmt"

func main() {
	b3 := 10 > 5
	fmt.Println(b3)
	
	b3 = 10 < 5
	fmt.Println(b3)
	
	b3 = 5 <= 10
	fmt.Println(b3)
	
	b3 = 10 != 10
	fmt.Println(b3)
}
```

The comparison operators shown in the above code example (`>`, `<`, `<=`, `!=`) are used to modify a variable's value at specific lines of code and assess the outcomes, which are boolean values (true or false) depending on the comparisons made.

Boolean constants and variables can also be combined using logical operators (`&&`, `||`, `!`) to yield a boolean result.
However, these logical expressions alone do not constitute a complete statement in Go.

**AND** (`&&`) and **OR** (`||`) are binary operators requiring two operands, while **NOT** (`!`) is a unary operator needing one operand.
The `&&` and `||` operators use shortcut evaluation: if the outcome can be determined from the left operand alone (`false` for `&&`, `true` for `||`), the right operand is not evaluated.
This suggests placing longer computations on the right side to avoid unnecessary processing.

```go
package main
import "fmt"

func main() {
	b3 := 10 > 5 && 7 < 15;
	fmt.Println(b3)
	
	b3 = 10 < 5 || 2 < 7
	fmt.Println(b3)
	
	b3 = !b3
	fmt.Println(b3)
}
```

In the above example, logical operators are used to manipulate and evaluate the value of a variable (`b3`), resulting in boolean outputs (true or false) based on the comparisons performed.

## Bitwise operators

Bitwise operators, including

- **AND** (`&`)
- **OR** (`|`)
- **XOR** (`^`)
- **CLEAR** (`&^`)
- **COMPLEMENT** (`^`)

operate exclusively on integer variables with equal-length bit patterns.
The `%b` format string is used for displaying bit representations.
These operators manipulate bits directly, performing logical operations at the bit level.

Bitwise **AND**, **OR**, **XOR**, and **CLEAR** are _binary_ operators, which means they require _two_ operands to work on.
However, the COMPLEMENT operator is a unary operator.

There two other major bitwise operators used for _shifting_.

- Left shift operator `<<`
- Right shift operator `>>`

Left shift by `n` causes multiplication by 2^n, and right shifting causes division by 2^n.
For example, if `n` is **2**, the number is _multiplied_ by **4** for left shifting and _divided_ by **4** for right shifting.

## Operators and precedence

Operators have a defined order of precedence, meaning some are evaluated before others in expressions.
Operators with the same precedence level are processed from left to right.

| Precedence | Operator(s)                          |
|------------|--------------------------------------|
| 7          | `^`, `!`                             |
| 6          | `*`, `/`, `%`, `<<`, `>>`, `&`, `&^` |
| 5          | `*`, `/`, `&#124;`                   |
| 4          | `==`, `!=`, `<`, `<=`, `>=`, `>`     |
| 3          | `<-`                                 |
| 2          | `&&`                                 |
| 1          | `&#124;&#124;`                       |

Using `()` is allowed for clarifying expressions,, to indicate priority in operations as expressions contained in `()` are always evaluated first.
