package main

import "fmt"

const name = "radha" //we can decleare outside as well for both constant and normal varibale.
var name1 string = "krishna"

func main() {
	const age = 30
	fmt.Println((age))
	fmt.Println((name))
	fmt.Println((name1))

	//we can use constant grouping for declearing multiple constants.
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

}
