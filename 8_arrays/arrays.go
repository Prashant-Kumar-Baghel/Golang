package main

import "fmt"

//array is number sequence of specific length .whenever we create array we already decide its length.
//in golang index start from zero.
func main() {
	var nums [5]int        //way to decleare array
	fmt.Println(len(nums)) //"len" keyword use to find length of array
	nums[0] = 5            //way to add element in array
	fmt.Println(nums[0])   //way to print array element.
	fmt.Println(nums)      //way to print array .

	//when we create array then zero value added by default(when int type array). for boolean type array "false" value added by defalt.
	var vals [4]bool
	fmt.Println(vals)

	//when array is string type then empty string added by default
	var names [4]string
	names[0] = "golang"
	fmt.Println(names)

	// Way to add element at the time of array declaration.
	var elem = [4]int{4, 6, 7, 8}
	// elem := [4]int{4, 6, 7, 8} //other way to decleare
	fmt.Println(elem)

	//2d array
	twoDArray := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDArray)

	// in golang, we generally use slices instead of array. we generally use array when we already know the size of array so that we can easily optimize memory .

	//beneft of the array
	// fixed size , that is predictable
	// Memory optimazation
	// Constant time access

}
