package main

import "fmt"

func add(a int, b int) int { //we write type of variables as well as after variables. If you do NOT write the return type (int) before {, then Go thinks the function returns nothing.But inside function you are returning a value thatswhy we get error " returning too many values"
	return a + b
}

// func add(a, b int) int { //If multiple parameters have the same type, then in Go we can write the type only once.
// 	return a + b
// }

//
func getLanguage() (string, int, string, string) { //function can return multiple values as well
	return "go", 1, "c", "python" // Returns multiple values in same order as return types
}

//In Golang, functions are first-class citizens. In Go, saying “functions are first-class citizens” means:
//A function can be stored in a variable.
// A function can be passed as an argument to another function.
// A function can be returned from another function.

func processit(fn func(a int) int) int { //processit function recive the function .argument fn is a function that takes an int and returns an int. we can pass any function which has same signature as argument to processit function.
	return fn(2)
}

func main() { //main functon is entry point for us and golang call it automatically.

	result := add(5, 7)
	fmt.Println(result)
	lang1, lang2, lang3, lang4 := getLanguage()
	// lang1, lang2, lang3, _ := getLanguage()//we can use _ to skip value
	// fmt.Println(getLanguage())
	fmt.Println(lang1, lang2, lang3, lang4)

	//
	fn := func(a int) int { //this is anonymous function because it doesnot have name and we can store this function in variable and we can pass this function as argument to another function.
		return a
	}
	fmt.Println(processit(fn)) // this is how we can pass function as argument to another function.

}
