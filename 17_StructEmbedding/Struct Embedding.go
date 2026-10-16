package main

import "fmt"

type customer struct { //it is a struct that represents a customer with a name and age.
	name string
	age  int
}

type order struct { //Here we are embedding the customer struct into the order struct. embedding means that the fields of the customer struct will be accessible directly from the order struct.
	id     int
	amount float64
	status string
	customer
}

func main() {
	// Create a new order with embedded customer struct (1st way)
	newOrder := order{ //Here we are creating a new order struct and initializing it with values. We are also initializing the embedded customer struct with values.
		id:     1,
		amount: 100.50,
		status: "pending",
		customer: customer{
			name: "John Doe",
			age:  30,
		},
	}
	//2nd way
	customer1 := customer{
		name: "Jane Smith",
		age:  25,
	}
	newOrder2 := order{
		id:       2,
		amount:   200.75,
		status:   "shipped",
		customer: customer1,
	}
	fmt.Println("new order:", newOrder)
	fmt.Println("new order 2:", newOrder2)
	newOrder2.customer.name = "Jane Doe"
	fmt.Println("new order 2 after name change:", newOrder2)
}
