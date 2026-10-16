package main

import "fmt"

//use for the iterating over data structure.
func main() {
	nums := []int{9, 7, 6, 5}
	fmt.Println(nums)

	//normal for loop.
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	sum := 0
	//for-range loop used to iterate over elements of a slice, array, or map.
	//_ Ignores the index. basically this is index
	//num Stores current element value
	//Normally range returns index, value.
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	//map
	m := map[string]int{"team1": 1, "team2": 2}
	for k, v := range m { //here it return key and value.if we write only single value then we get only key
		fmt.Println(k, v)
	}

	//This is a Go range loop over a string.
	//The loop iterates through each character of the string "golang"
	// 	i → stores the index (position)
	// c → stores the Unicode value (ASCII value for normal English letters)
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c)) //string(c) → converts Unicode value (rune) into actual character
	}

}
