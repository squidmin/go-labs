# Challenge: Season of a Month

## Topics

- Problem statement
- Input
- Output
- Sample input
- Sample output

## Problem statement

Create a function named `Season` that takes a month number as input and returns the season name associated with that month, without considering the specific day.

- January: 1
- February: 2
- March: 3
- April: 4
- May: 5
- June: 6
- July: 7
- August: 8
- September: 9
- October: 10
- November: 11
- December: 12

Another thing to note is the seasons of the months. Look at the following mapping:

- Winter: 1, 2, and 12
- Spring: 3, 4, and 5
- Summer: 6, 7, and 8
- Autumn: 9, 10, and 11

## Input

An `int` variable.

## Output

A `string` variable.

## Sample input

```
3
```

## Sample output

```
Spring
```

Handle invalid month inputs (values outside the 1-12 range) in a function by returning the string "Season unknown".
Attempt to write the function before consulting the solution.

## Code

```go
package main

func Season(month int) string {
	return ""
}
```

## Solution

### Using a `switch-case` construct

```go
package main

import "fmt"

func main() {
	fmt.Printf(Season(3))
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Season unknown"
	}
}
```

### Using an `if-else` construct

```go
package main

import "fmt"

func main() {
	fmt.Printf(Season(3))
}

func Season(month int) string {
	winter := month == 1 || month == 2 || month == 12
	spring := month >= 3 && month <= 5
	summer := month >= 6 && month <= 8
	autumn := month >= 9 && month <= 11
	if winter {
		return "Winter"
    } else if spring {
		return "Spring"
    } else if summer {
		return "Summer"
    } else if autumn {
		return "Autumn"
    } else {
		return "Season unknown"
    }
}
```
