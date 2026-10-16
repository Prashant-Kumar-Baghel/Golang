package main

import "fmt" //we import the fmt package inside main package.fmt package coming from golang standard lbrary .

func main() { // we use func to create the function. main functon is the entry point for any programe.from there every go programe start to execute.
	fmt.Println("helloworld") //fmt is the package of standard library.Println method is present in fmt.
}

//Notes
// To execute code we have two ways :-
// 1)When we run command 'go build .\1_hello_world\main.go' then we get main.exe file which is a machine code. the size of file is very less so when we create big appliction that size wll also less. to execute main.exe file we use ./ then write the path .
// 2)we can use direct way for execution by command 'go run .\1_hello_world\main.go'

// package main:-
//package main is used to create a runnable Go program and define where execution starts.
// Only files inside package main can be compiled into a runnable program. If either package main or func main() is missing ❌ → program won’t run.

// fmt:-
// fmt stands for format. it is used for formatting.
// It is a package from Go's standard library used for:
// Printing output to the console
// Formatting strings
// Reading user input

/*Note-
When you run go run ./1_hello_world, the Go toolchain does two things:
1. Finds the entry package — it looks for a package named main. Any other name is treated as a library.

2. Finds the entry function — inside that package, it looks for a function with the exact signature func main() (no parameters, no return value). The runtime calls this function first, automatically.

So "defines where execution starts" means the Go runtime is hardcoded to always call main.main — the main function inside the main package — as the very first line of your program.


package main        // step 1: runtime finds this package

func main() {       // step 2: runtime calls THIS function first
    // your code runs from here, top to bottom
}
Nothing in your code explicitly calls main() — the Go runtime does it for you, automatically, before anything else runs. You have no control over what runs before it (except init(), which is a separate concept).*/
