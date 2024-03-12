package maps

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

func shorterMapInitialization() {
	fmt.Println("Example 2: Shorter map initialization")
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func compositeMapping() {
	fmt.Println("Example 3: Composite mapping")
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

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
