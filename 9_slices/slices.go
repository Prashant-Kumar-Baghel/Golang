package main

import (
	"fmt"
	"slices"
)

//slices are dynamic array that is we don't need to increase size of array.
//its size automatically increased so when we don't know that how many element we are going to put then we can use slices.
//most used construct in go
//we get the useful methods as well

func main() {
	//way to declear the slices
	var nums []int //This is the unintialized(doesnot provide the values) slice and unintialized slice is nil.
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	//  2 = initial length (and capacity by default)
	var nums1 = make([]int, 2, 5) //make is the inbuilt function . make([]int, 2) creates a slice of integers with initial length  2, initialized with zero values. we can add more then two element as well but its initial size is two. here we are not limiting the size. in this third perameter is initial capacity.if we don't put capacity then it automatically become equal to inital length.
	fmt.Println(cap(nums1))       //This prints the capacity of the slice nums1. Capacity (cap) = maximum number of elements the slice can hold without resizing. it doesnot mean we can fit maximum only 2 numbers in 'num1' slice . the slice automatically resize because its dynamic array so capacity and length change dynamically.
	fmt.Println(nums1 == nil)     //num1 is not nil because it has 2 element .
	fmt.Println(len(nums1))

	//nums1 has capacity 5 but as we add elements and capacity is exceeded then its capacity automatically get resize(i.e capacity get double)
	//append adds elements to a slice in last index and returns a new slice; if capacity is exceeded, it allocates a new underlying array.
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 2)
	nums1 = append(nums1, 3)
	nums1 = append(nums1, 4)
	fmt.Println(nums1)
	fmt.Println(cap(nums1))
	fmt.Println(len(nums1))

	var nums3 = make([]int, 0, 5) //we generally  put initially length as 0 while creating the slices because we don't want starting elements as zero.
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)
	fmt.Println(nums3)
	fmt.Println(cap(nums3))
	fmt.Println(len(nums3))
	//we can put element inside slice with the help of index as well.
	// nums3[1] = 9

	//other way of creating slice
	nums4 := []int{} //empty slice
	nums4 = append(nums4, 1)
	fmt.Println(nums4)
	fmt.Println(cap(nums4))
	fmt.Println(len(nums4))

	//copy function (use to copy slice)
	// var nums5 = make([]int, 0, 5)
	nums5 := make([]int, 0, 5)
	nums5 = append(nums5, 1)
	nums5 = append(nums5, 2)
	var num6 = make([]int, len(nums5))

	copy(num6, nums5)
	fmt.Println(num6, nums5)

	//slice operator(colon(:) is slice operator) is used to take a portion of an array or slice.
	// slice[start:end]
	// start(index) → included
	// end(index) → excluded
	nums6 := []int{1, 2, 3}
	fmt.Println("num6:", nums6)
	fmt.Println(nums6[0:2])
	fmt.Println(nums6[:2]) //(by default it take 0 as starting index)
	fmt.Println(nums6[0:]) //(now it doesnot exclude last index)

	//slice package
	nums7 := []int{1, 2, 3}
	nums8 := []int{1, 2, 5}
	fmt.Println(slices.Equal(nums7, nums8))
	//2D array
	var nums9 = [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(nums9)

}
