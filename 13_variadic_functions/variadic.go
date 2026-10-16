package main

import "fmt"

// In the definition (func sum(nums ...int)) — it collects many arguments into a slice.
// In the call (sum(nums...)) — it spreads a slice into many arguments.

//sum is a variadic function because it can accept n numbers of values. Variadic functions are functions that can take a variable number of arguments. In Go, you can define a variadic function by using an ellipsis (...) before the type of the last parameter. In this case, the sum function takes a variable number of int arguments and returns their sum.
func sum(nums ...int) int { //nums is a slice of int type. we can pass n numbers of values to this function and it will accept all values as slice.

	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

func getNumber(nums ...interface{}) { //nums is a slice of interface{} type. we can pass n numbers of values to this function and it will accept all values as slice. interface{} is a type that can hold any value. it is used to accept any type of values.
	for _, v := range nums {
		fmt.Println(v)
	}
}

func main() {
	//variadic function

	fmt.Println(1, 2, 3, 4, "Hello", "World") // fmt.Println is an variadic function beacuse it can accept n numbers of values.
	sum1 := sum(1, 2, 3, 4, 5)                //sum is an variadic function because it can accept n numbers of values.

	nums := []int{1, 2, 3, 4, 5}
	sum2 := sum(nums...) //we can pass slice to variadic function by using ... after slice name. this is called unpacking the slice. it will pass all values of slice as individual values to the function. nums... is the spread (unpack) operator. It tells Go: "take each element of this slice and pass them as separate arguments," so sum(nums...) behaves exactly like sum(1, 2, 3).
	fmt.Println(sum1)
	fmt.Println(sum2)
	getNumber(1, 2, 3, 4, "Hello") //getNumber is an variadic function because it can accept n numbers of values.
}
