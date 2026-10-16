// ENUMS = Enumerated types: a fixed set of named values (an order can ONLY be Recieved, Confirmed, Delieverd or Prepared).
// Go has NO built-in "enum" keyword (unlike Java/C#). We simulate enums with: a custom type + a const block of that type.

//we use enums because if we pass simple string in changeOrderStatus function then same Confirmed used in many places in code if we need to change Confirmed status in future then we have to change in many places that is not good way , so it is like an group of status so its better to put all in seprate place.
//we donot face string typo as well using enums as we using const variables and code become more readable. In real world they place in other file and we create them also for multiple roles etc.
// Bruit-force way
// package main

// import "fmt"

// func changeOrderStatus(status string) {
// 	fmt.Println(status)
// }

// func main() {
// 	changeOrderStatus("Confirmed")
// }

// ----------------- VERSION 1: int-based enum (using iota) -----------------
// package main

// import "fmt"

// // Step 1: a new custom type based on int. Gives type-safety , we can create own custom types.
// // a function that accepts orderType will not accept a plain int variable.

//we create custom type like orderType as we can also write simple string because a plain string says nothing about what values are meant to go there — a custom type lets the compiler and the reader enforce your intent. second If tomorrow you decide statuses should be int codes instead of strings (like your iota version), you change type orderType string to type orderType int and the constant block — every function signature using orderType stays the same. With raw string sprinkled through 40 signatures, that's a painful hunt.

// type orderType int

// // Step 2: the allowed values as constants of that type.
// // iota is a built-in counter: starts at 0, increases by 1 on each line of the const block.
// // With iota, the type AND the "= iota" repeat automatically for the lines below.
//iota can ONLY be used inside a const block.
// const (
// 	Recieved  orderType = iota // 0// Generally start with Capital letter
// 	Confirmed                  // 1
// 	Delieverd                  // 2
// 	Prepared                   // 3
// )

// func changeOrderStatus(status orderType) {
// 	fmt.Println(status) //drawback of int enums: prints the number (1), NOT the name "Confirmed"
// }

// func main() {
// 	changeOrderStatus(Confirmed)
// }

//----------------- VERSION 2: string-based enum (active code) -----------------
//Advantage over int enums: printing/logging shows a readable name ("Confirmed") instead of a number (1).

package main

import "fmt"

// Custom type based on string — the "enum type". Only its declared constants should be used as values.
type orderType string

// The fixed set of allowed values.
// Unlike the iota version, the type does NOT carry down automatically when every line
// has its own "= value" — so write orderType explicitly on each line.
// (Otherwise Confirmed/Delieverd/Prepared would be plain untyped string constants, not orderType.)
const (
	Recieved  orderType = "Recieved"
	Confirmed orderType = "Confirmed"
	Delieverd orderType = "Delieverd"
	Prepared  orderType = "Prepared"
)

// Accepts only orderType. The type-safety has a LIMIT though:
//   - a string VARIABLE will not compile:  var s string = "x"; changeOrderStatus(s) → ERROR
//   - but a raw LITERAL still compiles:    changeOrderStatus("cnfirmed") → works, even with a typo!
//
// Reason: literals are "untyped constants" and Go converts them to orderType automatically.
// So Go enums are not fully closed — the compiler protects against wrong variables, not wrong literals.
// A literal is a value written directly in the source code, rather than stored in a variable — you're typing the actual data itself at that spot.
//An untyped constant (like "confirmed" written directly in code) has no fixed type until it's used — Go auto-converts it to whatever type the context needs (here, orderType).
//A variable (like var s string = "x") is already locked to a type (string), so it won't auto-convert to a different named type like orderType.

func changeOrderStatus(status orderType) {
	fmt.Println(status) //prints the readable name, e.g. "Confirmed"
}

func main() {

	changeOrderStatus(Confirmed) //always pass one of the declared constants, never a raw string literal
}
