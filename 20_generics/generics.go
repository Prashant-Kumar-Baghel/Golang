//Case 1:-
// package main

// import "fmt"

// //As we can see there is lot of duplicasy here as printSlice and printStringSlice have the same logic but only name and parameter types changes which is not better way as we need to follow DOY(Donot repeat yourself). To solve this issue we use generics.

//function
// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// // func printStringSlice(items []string) {
// // 	for _, item := range items {
// // 		fmt.Println(item)
// // 	}
// // }

//struct (In this case also dulicasy let say we want to do some operations with string struct then we need to create seprate string struct with same logic . so we also use interface in this case. )
// type stack struct{
// 	elements []int
// }

// func main() {

//for function

// 	items := []int{1, 2, 3, 4, 5}
// 	printSlice(items)
// 	// itemsString := []string{"golang", "cs"}
// 	// printStringSlice(itemsString)

//for struct

// mystack := stack{
// 	elements : []int{1,2,3,4}
// }
// fmt.Println(mystack)

// }

// Case2:-(Generics) // GENERIC FUNCTION — one function that works for many types.

package main

import "fmt"

// func printSlice[T any](items []T) {//any means any type and we can write any other alfphabet as well inplace of "T" . we can also write interface{} instead  of "any" keyword.
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//we are using generics on functions
// func printSlice[T int | string | bool](items []T) { //using "any" keyword is not good practise as we can pass any argument so let say we want only int,string slices we can use | .
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types).
func printSlice[T comparable, V string](items []T, name V) { // we can use this comparable keyword also instead of writing "int | string | bool" as we can also pass multiple parameter also
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println(name)
}

////we can also use generics on struct. we are creating stack(LIFO:- last in first out)
type stack[T any] struct {
	elements []T
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	// itemsString := []string{"golang", "cs"}
	// itemsBool := []bool{true, false}
	printSlice(items, "jsks")
	// printSlice(itemsString)
	// printSlice(itemsBool)

	// struct
	// stack[int] and stack[string] are two completely different, incompatible types.
	myStack := stack[int]{ // we need to mention the type of T in struct here by using brackets
		elements: []int{1, 2, 3, 4, 5}, // must match: T is int, so the slice must be []int — []string here would not compile
	}
	fmt.Println(myStack)
	//int
	myStack1 := stack[string]{ // we need to mention the type of T in struct here by using brackets
		elements: []string{"golang", "c++++"}, // must match: T is strint, so the slice must be []string — []int here would not compile
	}
	fmt.Println(myStack1)

}
