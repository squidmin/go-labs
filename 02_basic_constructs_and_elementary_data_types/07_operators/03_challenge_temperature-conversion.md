# Challenge: Temperature Conversion

## Topics

- Problem statement
- Input
- Output
- Sample input
- Sample output

## Problem statement

Implement a function `toFahrenheit(t Celsius)` to convert from Celsius to Fahrenheit.
The data type for the temperatures should be `float32`.

### Input

A variable of type `Celsius`.

### Output

A variable of type `Fahrenheit`.

## Sample input

```
100
```

## Sample output

```
212
```

Try to implement the function below.

```go
package main

import "fmt"
import "strconv"
import "encoding/json"

// Aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert Celsius to Fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	return temp
}
```

## Solution

```go
package main

import "fmt"
import "strconv"
import "encoding/json"

// Aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert Celsius to Fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t * 9 / 5) + 32)
	return temp
}

func main() {
	var celsiusTemp Celsius = 100
	fahrTemp := toFahrenheit(celsiusTemp)
	fmt.Printf("%f (degrees celsius is equal to %f degrees fahrenheit)", celsiusTemp, fahrTemp)
}
```

In the above code, two new types, `Celsius` and `Fahrenheit`, are created by aliasing the `float32` type to represent temperature units.
Subsequently, a function is defined to convert temperatures from Celsius to Fahrenheit using these aliased types.

The function `toFahrenheit` describes a function named `toFahrenheit` that converts a temperature from Celsius to Fahrenheit.
It takes a parameter `t` of type `Celsius`, applies a conversion formula, and returns the result as type `Fahrenheit`.
The conversion is done by multiplying `t` by `9 / 5`, adding `32`, and then casting the result to `Fahrenheit`.

In the main function, a variable `celsiusTemp` of type `Celsius` is declared. The `toFahrenheit` function is then called with `celsiusTemp` as an argument, converting it to `Fahrenheit`.
The result is stored in `fahrTemp`, and finally, the converted temperature is printed.
