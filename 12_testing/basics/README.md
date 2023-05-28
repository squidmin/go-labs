# Testing: Basics

Contents:
- The `testing` package
- Writing tests
  - Writing simple unit tests
  - Writing table-driven tests


---


## The `testing` package

The `testing` package defines types and functions that facilitate writing automated tests for Go packages.
The `go test` command fully supports the `testing` package and is designed to run your tests and capture the results.

Here are some of the built-in capabilities of the `testing` package:
- Self-discovery of tests and benchmarks
- Nested tests and table-driven tests
- Skipping tests
- Verified example code
- Package-level setup and teardown


---


## Writing tests

Go tests use special types defined in the `testing` package and must implement a special signature:

```go
func TestFoo(t *testing.T) {
}
```

The `testing.T` types that is passed in to your test has various methods that you call to let the testing infrastructure know if the test failed, including:
- format and log error messages
- skip tests
- run sub-tests

For benchmarks there is a different signature:

```go
func BenchmarkFoo(*testing.B) 
}
```

The rolw of the `testing.B` type is similar to that of the `testing.T` type.
It is passed into every benchmark, and it has a slew of functionality that is useful for benchmarking such as:
- reporting memory allocations
- reporting metrics
- working with timers
- running benchmarks in parallel

The `T` and `B` types share a common interface called `TB`:

```go
type TB interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
	Helper()
}
```


---


## Writing simple unit tests

Unit tests go in the same package that they test in a file with the naming convention:

`_test.go`

This lets the `go test` command easily discover all the tests that need to be run.

### Example

`12_testing/basics/example.go`

```go
package basics

import (
	"errors"
)

var errorMessage = "width and height must be positive"

// CalcArea Calculate the area of a rectangle
func CalcArea(w, h int) (int, error) {
	if w < 1 || h < 1 {
		return 0, errors.New(errorMessage)
	}
	return w * h, nil
}
```

`12_testing/basics/example_test.go`

```go
package basics

import "testing"

func TestCalcAreaSuccess(t *testing.T) {
	result, err := CalcArea(3, 5)
	if err != nil {
		t.Error("Expected CalcArea(3, 5) to succeed")
	} else if result != 15 {
		t.Errorf("CalcArea(3, 5) returned %d. Expected 15", result)
	}
}

func TestCalcAreaFail(t *testing.T) {
	_, err := CalcArea(-3, 6)
	if err == nil {
		t.Error("Expected CalcArea(-3, 6) to return an error")
	}

	if err.Error() != errorMessage {
		t.Error("Expected error to be: " + errorMessage)
	}
}

func TestCalcAreaViaTable(t *testing.T) {
	var tests = []struct {
		width    int
		height   int
		expected int
	}{
		{1, 1, 1},
		{5, 6, 30},
		{1, 99, 99},
		{7, 6, 42},
	}

	for _, test := range tests {
		w := test.width
		h := test.height
		r, err := CalcArea(w, h)
		if err != nil {
			t.Errorf("CalcArea(%d, %d) returned an error", w, h)
		} else if r != test.expected {
			t.Errorf("CalcArea(%d, %d) returned %d. Expected %d",
				w, h, r, test.expected)
		}
	}
}
```
