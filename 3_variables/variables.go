package main

import "fmt"

func main() {
	//method1:- way to declear variable
	var name string = "golang" //we need to use the variable otherwise we get error "declared and not used"
	fmt.Println(name)
	//method2:- we can use sort way to decleare variable as well.if we are provide direct value and not provide type then golang automatically infer type .
	var name1 = "golang"
	var name2 = true
	fmt.Println(name1)
	fmt.Println(name2)

	//method3:-shorthand syntax(assigning and declearing value to variable at same time )
	name4 := "guru"
	fmt.Println(name4)

	//we can't use method2 and method3 every time as some time we need to assign value after declaration.
	var name5 string
	name5 = "jobs"
	fmt.Println((name5))
}

//var name int = 7
