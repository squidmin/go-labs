# Times and Dates

In this section, you'll study the time package and the functions supported by it.

## Topics

- The `time` Package

THe `time` package in Go provides a `Time` data type for handling dates and times, allowing for the display and measurement of time.
It offers the `time.Now()` function to get the current time, from which specific components like day, minute, etc., can be extracted using methods like `t.Day()` and `t.Minute()`.

The below code demonstrates creating a custom time format:

```go
package main

import (
	"fmt"
	"time"
)

t := time.Now()
fmt.Printf("%03d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
```

There are three key concepts related to time handling in Golang:

1. The `Duration` data type: A data type that quantifies the time difference between two moments, measured in nanoseconds as an `int64` value.
2. The `Since(t Time)` function: This function calculates the time elapsed from a specified instant `t` to the current time.
3. The `Location` data type: This type associates time points with their respective time zones.
   The text specifically mentioned `UTC` as the Universal Coordinated Time standard.

There is a pre-defined function `Format` with the following syntax:

```
func (t Time) Format(s string) string
```

This formats the time `t` into a string according to a string `s`, with some pre-defined formats like `time.ANSIC` or time `time.RFC822`.
The general layout defines the format by showing the presentation of a standard time, which is then used to describe the time to be formatted.

```
t := time.Now().UTC()
fmt.Println(t.Format("02 Jan 2006 15:04)) // e.g.: 29 Oct 2019 11:00
```

Run the following program to see how the above-discussed functionalities work.

```go
package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.$02d.%04d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	
	// Calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // Must be in nanoseconds
	weekFromNow := t.Add(week)
	fmt.Println(weekFromNow)
	
	// Formatting times:
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	
	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)
}
```

In the above code, we declared time `t` and initialized it with the present time using `time.Now()`.
Printing the `t` prints complete time with a default format.
We then specify the format of `t` as `dd.mm.yyyy`.
Next, we decide to add a _week_ to the stored time value.
We declare the `week` variable and initialize it with a value of `60 * 60 * 24 * 7 * 1e9`.
We multiply a factor of `1e9` because we have to keep it in nanoseconds.
We add the value of `week` to the time `t`.
Now after printing `t`, change is noticeable, as weeks are added into `t`.
Time is printed in `RFC822` format, e.g., `30 Oct 19 11:34 UTC`.
At **line 22**, time is printed in `ANSIC` format, e.g., `Wed Oct 30 11:34:03 2019`.
At **line 23** time is printed in `02 Jan 2006 15:04` format, e.g., `30 Oct 2019 11:34`.
At **line 24** we create a `format(yyyy.mm.dd)` and store it in `s`.
In the next line, we print the time `t` in default format and then in the format specified above.

For more information on Golang's `time` package, visit [this page](https://pkg.go.dev/time).
