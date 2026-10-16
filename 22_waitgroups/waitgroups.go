package main

import (
	"fmt"
	"sync"
)

// WaitGroup (from the "sync" package) is the proper way to wait for goroutines to finish, instead of guessing with time.Sleep. It is basically a counter.

//	wg.Add(n)  -> increase the counter by n (n goroutines are about to start)
//	wg.Done()  -> decrease the counter by 1 (this goroutine has finished)
//	wg.Wait()  -> block until the counter reaches 0 (all goroutines are done)

/*
The goroutine receives a *pointer* to the WaitGroup. If we passed it by value,
each goroutine would get its own copy, call Done() on the copy, and the original
counter in main would never reach 0 — Wait() would block forever (deadlock).
*/
func task(i int, wg *sync.WaitGroup) {
	defer wg.Done() // defer guarantees Done() runs when the function finish.
	fmt.Println(i)
}

func main() {
	/*Yes — in Go, "zero value" means the value a variable gets automatically when you declare it without
	initializing it. For a struct like sync.WaitGroup, that means all its internal fields are
	zeroed out, and since the counter is one of those fields, the counter starts at 0.*/
	var wg sync.WaitGroup // zero value is ready to use, counter starts at 0

	for i := 0; i <= 10; i++ {
		wg.Add(1)       // register one goroutine BEFORE starting it — counter goes up by 1
		go task(i, &wg) // launch the goroutine, passing the WaitGroup's address
	}

	/*Wait() blocks the main goroutine here until every task has called Done()
	and the counter is back to 0. Unlike time.Sleep, this waits exactly as long
	as needed — not too short (lost output), not too long (wasted time).*/
	wg.Wait()
	fmt.Println("All tasks finished.")

}
