package main

import (
	"fmt"
	"sync"
	"time"
)

//we can't return value while run code in goroutine. so to send data from one goroutine to another goroutine we can use channel. channel is a way to communicate between goroutines.
//channel work in FIFO(First In First Out) manner.
// func worker(url string) string {
// 	time.Sleep(time.Millisecond * 50)

// 	fmt.Printf("image processed: %s\n", url)
// 	return url
// }

func worker(url string, wg *sync.WaitGroup, resultChain chan string) {
	defer func() {
		wg.Done()
	}()
	time.Sleep(time.Millisecond * 50)
	fmt.Printf("image processed: %s\n", url)

	resultChain <- url // here we send the data to channel. so we can send data from one goroutine to another goroutine using channel.

}

func main() {

	//Below code is sequential and will take 150ms to complete, because it is running synchronously, one after another.
	/*startTime := time.Now()

	result1 := worker("image_1.png")
	result2 := worker("image_2.png")
	worker("image_3.png")
	fmt.Println("result1", result	1)
	fmt.Println("result2", result2)

	fmt.Printf("it took %s ms.\n", time.Since(startTime))*/

	//Below code is concurrent and will take 50ms to complete, because it is running concurrently, all at once.
	//Below is the concurrency model of golang which is also called fork and join model as well as parallel and join model. we called this lauch and wait model .
	//Here main function run on main thread and worker function which is goroutien run on light weight thread. so code will run asynchronously and concurrently.
	//it is called launch and wait model because we launch the goroutine and wait for it to complete. To wait for the goroutine to complete, we use wait group.
	//Below pattern is called fan-out and fan-in pattern. we called fan-out because we are fanning out the work to multiple goroutines.we called fan-in because we are waiting for all the goroutines to complete and then we are collecting the results.if we are sending the data then we need to aggregate that data then that aggreation(collection) of results is called fan-in pattern. we achieve fan-in using channel.
	var wg sync.WaitGroup
	resultChain := make(chan string, 3)

	startTime := time.Now()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(fmt.Sprintf("image_%d.png", i), &wg, resultChain)
	}
	go func() { // we need to close the channel after all the goroutines are done, otherwise it will cause deadlock.if we call close after range loop then it will cause panic.
		wg.Wait()
		close(resultChain)
	}()
	for result := range resultChain {
		fmt.Println(result)
	}

	//wait for all goroutines to complete, otherwise main function will exit before goroutines complete.

	fmt.Printf("it took %s ms.\n", time.Since(startTime))
}
