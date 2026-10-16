package main

import "fmt"

// in golang we have only for loop . we don't have while loop as well. we implement while loop wth the help of for loop.
func main() {
	//while loop
	// i := 1
	// for i <= 5 { //no round bracket used. here it work like while loop
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	//classic for loop .
	// for j := 0; j < 4; j++ {
	// 	// break //used to break the loop
	// 	if j == 2 {
	// 		continue//skip the iteration

	// 	}
	// 	fmt.Println(j)
	// }

	//1.22 range

	for i := range 10 { //using range we print till 10 and 10 is exclude. so we print 0 to 9.
		fmt.Println(i)
	}

}
