package main

import (
	"fmt"
	"time"
)

func main() {
	//use switch When we have many conditions based on the same variable, writing multiple if-else becomes messy.
	//Notes :- 1) switch {} → expects boolean conditions. Each case returns true/false

	// |Type             | Syntax      | Use Case              |
	// | ---------------- | ----------- | --------------------- |
	// | Value switch     | `switch j`  | exact match (`1,2,3`) |
	// | Condition switch | `switch {}` | ranges (`<, >, <=`)   |

	i := 4
	switch i {
	case 1: //here switch i → compares value of i with each case. here it cheking i == case value.
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("other")
	}

	j := 9
	switch {
	case j < 5:
		fmt.Println("less than 5")
	case j < 10:
		fmt.Println("less than 10")
	default:
		fmt.Println("10 or more")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend", time.Saturday)
	default:
		fmt.Println("it's workday")
	}

	num := 2

	switch num {
	case 1, 2, 3:
		fmt.Println("Between 1 and 3")
	case 4, 5:
		fmt.Println("Between 4 and 5")
	default:
		fmt.Println("Other number")
	}

	//type switch(Type switch only works on interface values)

	whoAmi := func(i interface{}) { //storing function inside variable.An interface{} stores both (value + type).interface with empty bracket mean any  type  that is i variable can have value of any type. So, we can pass any type of value to this function.
		switch t := i.(type) { //i.(type) always return the type of variable like string,int etc .. and type of varable store in "t".
		//i is of type interface{} (generic),
		// t is of the actual concrete type (like string, int) → so it’s safer and directly usable.
		case int:
			fmt.Println("integers")
		case string:
			fmt.Println("string", t)
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmi(5)
	whoAmi("7")
	whoAmi(false)
	whoAmi(1.2)

}
