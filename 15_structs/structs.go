package main

import (
	"fmt"
	"time"
)

// struct is a user-defined data type that allows us to group related data together.
// Struct is a collection of fields, where each field has a name and a type.
// Structs are used to create complex data types that represent real-world entities. In Go, we can define a struct using the "type" keyword followed by the struct name and the "struct" keyword.
// We can then create instances of the struct and access its fields using the dot notation.
// structs are like classes in other programming languages.
// we can use object-oriented programming concepts like encapsulation, inheritance, and polymorphism with structs in Go.
type order struct {
	id        string
	amount    int
	status    string
	createdAt time.Time
}

//constructor concept
//In golang , we don't have constructors like other programming languages but we can create a function that returns an instance of the struct and we can use that function to create an instance of the struct. This function is called a constructor function.
//We can define a constructor function for our struct by using the "func" keyword followed by the function name and the return type as the struct type.

// myOrder  →  type: order    (the struct instance)
// &myOrder →  type: *order   (address of that instance)

// In go , constructor function is a function that returns an instance of the struct. We can use this function to create an instance of the struct and we can also use it to initialize the fields of the struct. We can define a constructor function for our struct by using the "func" keyword followed by the function name and the return type as the struct type. The constructor function can take parameters that are used to initialize the fields of the struct. We can also use pointer receiver in constructor function to return pointer of struct instance.
// In constructor function syntax function name start with "new" and then struct name. for example if we have a struct named "order" then we can define a constructor function named "newOrder" that returns an instance of the order struct.
// Every value in Go has a type. The instance myOrder was created from the order blueprint, so its type is order.
func newOrder(id string, amount int, status string) *order { // we return pointer of order struct because we want to return the address of the struct instance and we can use that address to access the fields of the struct. if we return value of order struct then it will create a copy of the struct instance and we can't access the fields of the original struct instance.
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder //you're handing back instance address where it lives in memory (something like 0xc000010030). An address pointing to an order has type *order, so the signature must say *order.

}

// method concept
// In classes we attached methods to the class but in golang we can attach methods to struct as well. we can define a method for a struct by defining a function with a receiver of the struct type. The receiver is specified before the function name and is used to access the fields of the struct within the method. We can define multiple methods for a struct, each with its own receiver and functionality.
// we use methods to define behavior for our structs. Methods are functions that are associated with a specific struct type and can access its fields and perform operations on them. We can define methods for our structs by using the "func" keyword followed by the receiver type and the method name. The receiver type is specified in parentheses before the method name and is used to access the fields of the struct within the method.
// we use reciever to connect the method to the struct. reciever type has syntax of (variableName structName) and we can use variableName to access the fields of the struct within the method and variableName is the first letter of the struct name. We can define multiple methods for a struct, each with its own receiver and functionality..
// if we don't use pointer receiver then we can't change the value of the struct field inside the method because the method will receive a copy of the struct, and any changes made to the copy will not affect the original struct. If we want to modify the fields of the struct inside the method, we need to use a pointer receiver, which allows us to access and modify the original struct.
func (o *order) changeStatus(status string) {

	o.status = status // we don't use dereference operator here  to access the fields of the struct because we can access the fields of the struct directly using the dot notation. The dot notation automatically dereferences the pointer and accesses the field of the struct.
	fmt.Println("status changed to", o.status)
}

func (o order) getAmount() int { // we don't use pointer receiver here because we don't want to change the value of the struct field inside the method. we just want to access the value of the struct field.
	return o.amount
}
func main() {

	//1st way to create instance of struct
	// var myOrder order //myOrder is an instance of the order struct and we can access its fields using the dot notation.memory allocated: {id: "", amount: 0, status: "", createdAt: zero time}
	// myOrder.id = "1"
	// myOrder.amount = 29
	// myOrder.status = "recived"
	// myOrder.createdAt = time.Now()

	//2nd way to create instance of struct
	myOrder := order{ //myOrder is an instance of the order struct
		id:     "1",
		amount: 29,
		status: "recived",
	}

	myOrder1 := order{
		id:     "2",
		amount: 30,
		// status:    "shipped",
		createdAt: time.Now(),
	}

	// myOrder.createdAt = time.Now()        //we can assign value to field after creating instance of struct as well.
	myOrder.amount = 30                   //we can update value of field after creating instance of struct as well.
	fmt.Println("status", myOrder.status) //we can access value of field using dot notation.
	//Both myOrder and myOrder1 are instances of the order struct, but they have different values for their fields. We can create as many instances of a struct as we need, each with its own set of values for the fields.
	fmt.Println("myOrder", myOrder)
	fmt.Println("myOrder1", myOrder1)

	myOrder3 := order{
		id:        "3",
		amount:    40,
		status:    "shipped",
		createdAt: time.Now(),
	}
	myOrder3.changeStatus("confirmed")
	fmt.Println("Amount:", myOrder3.getAmount())
	fmt.Println("myOrder3", myOrder3)

	//if we don't set any field value while creating instance of struct then golang will set default value for that field based on its type. for example if we don't set value for amount field then golang will set default value for amount field as 0 because amount is of type int and default value for int is 0. similarly if we don't set value for status field then golang will set default value for status field as "" because status is of type string and default value for string is "".
	//Default values :- int=>0, string="", bool=false, float=0.0, time.Time=zero time.
	myOrder4 := order{
		id:     "4",
		status: "shipped",
	}
	fmt.Println("myOrder4", myOrder4)

	//COnstructor function
	myOrders5 := newOrder("5", 50, "recived") //we can use constructor function to create instance of struct and initialize its fields. we can also use pointer receiver in constructor function to return pointer of struct instance.
	myOrders5.amount = 10

	//myOrders5 holds the address. fmt knows raw hex addresses like 0xc000010030 are useless to humans, so for a pointer to a struct it follows the pointer and prints the contents, prefixed with & to tell you "this is a pointer to this data.
	//a variable of type *order that stores just that address. It's a signpost pointing at the instance, not the instance itself.
	fmt.Println("myOrders5", myOrders5, myOrders5.amount) //we don't need to use dereference operator here to access the fields of the struct because we can access the fields of the struct directly using the dot notation. The dot notation automatically dereferences the pointer and accesses the field of the struct.

	//Below language struct is an anonymous struct because it doesn't have a name.  we can define an anonymous struct by using the "struct" keyword followed by the fields of the struct.
	//if we want to create struct only once and don't want to create instance again and again like for configuration so we can use anonymous struct.
	//When i write that  "language struct is an anonymous struct", it's using "language struct" as shorthand for "the struct value stored in the language variable" — not saying "language" is the struct's type name.This is a common (if slightly loose) way people talk in casual comments/conversation
	language := struct {
		name   string
		isGood bool
	}{
		"golang",
		true,
	}
	fmt.Println("language", language)
}
