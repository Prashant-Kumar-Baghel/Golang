package main

import (
	"fmt"
	"maps"
)

// maps:- Assosative data structure (hash, object)
func main() {
	//creating maps . we use make function for maps as well.
	m := make(map[string]string) //key and value both are strings
	n := make(map[string]int)
	s := make(map[int]bool)
	//setting elements
	m["name"] = "golang"
	m["area"] = "bacend"
	n["value"] = 6
	s[7] = true
	//get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["hello"])             //string case:- we get empty value for the key which doesnot exist in map.
	fmt.Println(n["value"], n["daata"]) //integer case:- we get zero value for the key which doesnot exist in map.
	fmt.Println(s[7], s[6])             //boolean case:- we get false value for the key which doesnot exist in map.

	fmt.Println(len(m))

	//way to delete the key from map.
	fmt.Println(m)
	delete(m, "area")
	fmt.Println(m)
	//way to clear the map
	clear(m)
	fmt.Println(m)
	//way to write map without make function
	p := map[string]int{"data": 5, "data1": 6}
	fmt.Println(p)

	//This is a Go (Golang) code snippet that checks whether the key "data" exists in a map p
	//  _, ok := p["data"]
	// In Go, when accessing a map value, you can get two return values:

	// value, ok := map[key]
	// value → actual value stored in the map
	// ok → boolean (true or false)
	// true = key exists
	// false = key does not exist
	//_ means "ignore the value"
	value, ok := p["data"]
	fmt.Println(value)

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//This code compares two maps (m1 and m2) in Go and prints the result.
	m1 := map[string]int{"data": 5, "data1": 6}
	m2 := map[string]int{"data": 5, "data1": 6}

	// maps.Equal() is a function from the Go maps package.
	// It checks whether:
	// both maps have the same keys
	// and the same values for those keys

	// It returns:

	// true → maps are equal
	// false → maps are different
	fmt.Println(maps.Equal(m1, m2))

}
