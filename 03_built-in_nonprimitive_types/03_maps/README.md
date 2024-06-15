## Maps


### Declaration syntax

```go
var x map[string]int
```

### Initialization syntax

<details>
<summary>Example 1</summary>

```go
x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])
```

**Output**

```
10
```

</details>


<details>
<summary>Example 2</summary>

```go
x := make(map[int]int)
x[1] = 10
fmt.Println(x[1])
```

**Output**

```
10
```

</details>


---


### `delete` function

```go
delete(x, 1)
```


---


### Examples

<details>
<summary>Example 1: "map" data type</summary>

```go
package main

import "fmt"

func mapExample() {
	fmt.Println("Example 1: \"map\" data type")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	
	fmt.Println(elements["Li"])
}
```

**Output**

```
Example 1: "map" data type
Lithium
```

</details>


<details>
<summary>Example 2: Shorter map initialization</summary>

```go
package main

import "fmt"

func shorterMapInitialization() {
	fmt.Println("Example 2: Shorter map initialization")
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
    }
	fmt.Println(elements)
}
```

**Output**

```
Example 2: Shorter map initialization
map[B:Boron Be:Beryllium C:Carbon F:Fluorine H:Hydrogen He:Helium Li:Lithium N:Nitrogen Ne:Neon O:Oxygen]
```

</details>


<details>
<summary>Example 3: Composite mapping</summary>

```go
package main

import "fmt"

func compositeMapping() {
	fmt.Println("Example 3: Composite mapping")
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
        },
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
        },
		"Li": map[string]string{
			"name": "Lithium",
			"state": "solid",
        },
    }
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
    }
}
```

**Output**

```
Example 3: Composite mapping
Lithium solid
```

</details>


<details>
<summary>Example 4: Map lookup verification</summary>

```go
package main

import "fmt"

func mapLookupVerificationExample() {
	fmt.Println("Example 4: Map lookup verification")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	
	// Accessing an element of a map can return two values instead of just one.
	// The first value is the result of the lookup; the second tells whether the lookup was successful,
	//name, ok := elements["Un"]
	//fmt.Println(name, ok)
	if name, ok := elements["Un"]; ok {
		fmt.Println(name)
    }
}
```

**Output**

Since there is no entry in the map with the value `"Un"`, the output of this code is empty.

</details>
