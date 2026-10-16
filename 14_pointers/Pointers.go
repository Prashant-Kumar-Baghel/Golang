package main

import "fmt"

//Pass by value
//In go, we do pass by value. It means when we pass a variable to a function, it creates a copy of the variable and passes it to the function. So when we change the value of the variable inside the function, it does not affect the original variable outside the function. This is called pass by value.
func changeNum(num int) { //num is a copy of the value of num in main function. so when we change the value of num in changeNum function it will not affect the value of num in main function.
	num = 10
	fmt.Println(num)
}

//Pass by reference is the process of passing the memory address of a variable to a function. In go, we can use pointer to achieve pass by reference. Pointer is a variable that stores the memory address (reference) of another variable.
//we need to use pointer to change the value of the variable inside the function. Pointer is a variable that stores the memory address (reference) of another variable.

func changeNumPtr(num *int) { // num is a pointer: an address of an int . num *int means the parameter num is a pointer to an int — it doesn't hold an int value itself, it holds the memory address of an int.
	fmt.Println("kdk", num) //this will print the memory address of num in main function. because num is a pointer and it holds the memory address of num in main function.
	*num = 10               //* is called dereference operator. it is used to get the value of the variable from its memory address. here we are changing the value of num in main function by using its memory address.
	fmt.Println(*num)
}

func main() {
	fmt.Println("Start small. Ship something.")
	num := 2
	changeNum(num)
	fmt.Println(&num)
	fmt.Println(num)   //num is still 2 because we passed a copy of the value of num to changeNum function. so when we change the value of num in changeNum function it will not affect the value of num in main function.
	changeNumPtr(&num) //we pass the memory address of num to changeNumPtr function using & operator. & is called amperstand operator. it is used to get the memory address of a variable.
	fmt.Println(num)   //num is now 10 because we passed the memory address of num to changeNumPtr function and changed the value at that memory address.
}
