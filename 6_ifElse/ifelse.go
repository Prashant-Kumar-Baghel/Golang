package main

import "fmt"

func main() {
	age := 10
	if age >= 18 {
		fmt.Println("person is the adult")
	} else if age >= 12 {
		fmt.Println("person is the teeneger")
	} else {
		fmt.Println("person is the kid")
	}

	//logical operator
	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//we can decleare variable inside if.after declearing provide semicolon and write condition
	if value := 20; value >= 18 { //we can use the value variable inside if condition.
		fmt.Println("value is the legal", value)
	}

	//go doesnot has ternary operator so we use normal if-else condition.
}
