package main

import (
	"fmt"
)

// Goroutines are lightweight threads managed by the Go runtime. When we need to run
// tasks concurrently (multithreading-style work), we use goroutines.

//Above statement is also true but bellow defination is best
//Goroutines are lightweight units of execution managed by the Go runtime. We use goroutines when we want to run tasks concurrently, similar to multithreading

// What is a Thread in Go?
// A thread is a unit of execution provided by the operating system. It is used to execute code concurrently.
// In Go, however, you usually don't create OS threads directly. Instead, you use goroutines.
//The Go scheduler runs many goroutines on a smaller number of OS threads, which makes concurrent programming efficient

// 1000 Goroutines
//        ↓
//     Go Runtime
//        ↓
//    4 OS Threads
//        ↓
//    CPU Cores

// Blocking: a piece of code stops and waits until something finishes before moving on.
// In normal Go code, every call is blocking.
//
// Multithreading: the OS-level way to avoid blocking — the operating system runs multiple
// threads at once, potentially on different CPU cores. Threads are powerful but heavy:
// each one costs ~1 MB of stack memory, and switching between them (context switching)
// is expensive. Creating 100,000 OS threads would kill most machines.
//
// Concurrency: structuring your program as multiple independent tasks that make progress
// without waiting on each other. Goroutines are Go's unit of concurrency — they start with
// only ~2 KB of stack, so we can create thousands or even millions of them.

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {
	for i := 0; i <= 10; i++ { // Blocking version: calls run one after another, in order.
		task(i) // Each call must finish before the next one starts.
	}

	// We want the task function to run concurrently. To achieve that, we put the "go"
	// keyword in front of the function call.
	// for i := 0; i <= 10; i++ {
	// 	//Now we donot get result in order.
	// 	go task(i) // Go runs task() inside a new goroutine (lightweight thread).
	// 	// All 11 goroutines run concurrently, Now there is no blocking here.
	// }

	// If we don't make main sleep, we get no output. Here's why: main itself runs in a
	// goroutine (the "main goroutine"). The "go" keyword hands each task function to the Go
	// scheduler, which manages and runs goroutines. The loop only *schedules* the 11
	// goroutines in a non-blocking way — it doesn't wait for them. So main reaches the
	// end and exits before the scheduler gets a chance to run them, and when the main
	// goroutine exits, the whole program exits and every other goroutine is killed.
	// That's why we make main sleep.

	// for i := 0; i <= 10; i++ {
	// 	// we can use 'go' keyword  on anonymous function as well
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()

	// }

	// time.Sleep(time.Second * 2) // time.Sleep(time.Second * 2) pauses (blocks) the current goroutine — here, the main goroutine — for 2 seconds before letting it continue.

}
