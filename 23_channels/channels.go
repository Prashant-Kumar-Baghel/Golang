package main

// Channels are used to pass messages (data) between goroutines safely which runs concurrently .

//	make(chan string)  -> create a channel that carries string values
//	ch <- value        -> SEND: put a value into the channel
//	msg := <-ch        -> RECEIVE: take a value out of the channel
//
// The arrow always points in the direction the data flows.

/*An unbuffered channel like  (made without a size) has no storage — a send blocks until
another goroutine is ready to receive, and a receive blocks until another goroutine
sends. So a channel passes data AND synchronizes the two goroutines at the same time.*/

// messageChan := make(chan string)     // unbuffered — no size given
// messageChan := make(chan string, 5)  // buffered — can hold 5 values

// processNum is the RECEIVER goroutine. It gets the same channel main created
// (channels are references — no pointer needed, both sides share one pipe).
// <-numChan blocks until main sends a value, then evaluates to that value.
// func processNum(numChan chan int) {
// 	fmt.Println("Processing Number", <-numChan)
// }

//sending from main and recieve in this function.
// func processNumsValue(numsValueChan chan int) {
// 	for value := range numsValueChan { //range in Go isn't just for slices/arrays/maps — it works on anything Go defines "ranging" behavior for, and channels are one of those types.
// 		fmt.Println("Processing NumberValue", value) //we don't need to use arrow (<-) when we use range.
// 		time.Sleep(time.Second)                      // we using time here so we get wait between every iteration of loop.
// 	}
// }

// sending from sum and recieving in main
//
//	func sum(result chan int, a int, b int) {
//		finalResult := a + b
//		result <- finalResult
//	}

// func task(done chan bool) {

// 	defer func() { // defer in Go can only be applied to a function or method call — nothing else. done <- true is a send statement, not a function call, so defer done <- true is a compile error
// 		done <- true // we use true as a signal that task() is done. The value itself doesn't matter, we just need to send something.
// 	}()
// 	fmt.Println("processing")
// }

//	func emailSender(emailChain chan string, done chan bool) {
//		defer func() { //
//			done <- true
//		}()
//		for email := range emailChain {
//			fmt.Println("Sending email to", email)
//			time.Sleep(time.Millisecond * 100) // Simulate time taken to send an email
//		}
//	}
func main() {
	// case 1 : -
	/*messageChan := make(chan string)
	//DEADLOCK Error — this does NOT work with only the main goroutine:

	messageChan <- "ping" // ← deadlock happens HERE (channels are always blocking)
	msg := <-messageChan

	fmt.Println("message", msg)*/
	// Why the deadlock:
	/*The step-by-step of the crash
	1. There is only ONE goroutine. No go keyword anywhere, so the main goroutine is the only worker in the whole program. Every line runs one after another, in order — that's the "every call is blocking" rule from your notes.
	2. Main reaches the send and must wait. messageChan is unbuffered — no storage. A send can't just "drop the value and move on"; it's a hand-off that needs a receiver standing on the other side at the same moment. No goroutine is currently receiving, so main stops on this line and waits.
	3. The receiver can never arrive. Here's the trap: the receiver is the next line (msg := <-messageChan). But main can only run that line after it finishes the send — and the send only finishes when someone receives. Main is waiting for a line that main itself is supposed to execute. It's waiting for itself:
	send (line 1) is waiting for → receive (line 2)
	receive (line 2) is waiting for → main to get past line 1
	        └────────── circular wait = deadlock ──────────┘
	Like trying to shake your own hand, or calling your own phone number and waiting for yourself to pick up.
	4. Go detects it and kills the program. The runtime keeps track of all goroutines. At this moment, the only goroutine that exists (main) is asleep waiting on a channel, and no other goroutine exists that could ever wake it. Nothing can ever happen again — so instead of hanging forever, Go crashes deliberately.
	“asleep” = blocked/waiting, not literally sleeping.
	fatal error: all goroutines are asleep - deadlock!*/

	// case 2 :- WORKING version — send and receive happen in two DIFFERENT goroutines.
	// Goroutine 1 = main (the sender). Goroutine 2 = processNum (the receiver).

	// numChan := make(chan int) // create one unbuffered channel; both goroutines will share it

	// go processNum(numChan) // launch goroutine 2; it reaches <-numChan and BLOCKS, waiting for a value

	// numChan <- 5 // main SENDS 5. Both sides are now at the channel → hand-off happens:
	// // 5 travels main → numChan → processNum, and both goroutines continue.
	// // (Unlike case 1, main is not waiting for itself — a second goroutine is receiving.)

	// // The send only guarantees processNum RECEIVED the value, not that it finished
	// // PRINTING it. Without this sleep, main could exit first and kill goroutine 2
	// // before its Println runs — so we keep main alive for 2 seconds.
	// time.Sleep(time.Second * 2)

	//case 3
	/*Note:- A send on an unbuffered channel (numsValueChan <- rand.Intn(100)) blocks until some other goroutine is ready to receive from that exact channel at that exact moment.
	with unbuffered channel, always make sure a receiver is already running (or about to run) before you block on a send — otherwise you're just talking to an empty room.*/
	// //we use channel as a queue to send the data because we send lot of data.
	// numsValueChan := make(chan int)
	// go processNumsValue(numsValueChan)

	// //use infinite loop and we don't put sleep in main function because it is infinite loop as program willnot stop.
	// for {
	// 	numsValueChan <- rand.Intn(100) //sending random value in a range (max is 100 )
	// }

	/*Note :- Yes — on an unbuffered channel (like make(chan int)), both operations block:
		1) Receive (<-ch) blocks the goroutine doing the receiving until some other goroutine sends a value.
	    2) Send (ch <-) blocks the goroutine doing the sending until some other goroutine is ready to receive it.*/
	//case 4 (sending from goroutrin and recieving in main)
	// result := make(chan int)
	// go sum(result, 9, 5)
	// finalSum := <-result //we doesnot use time.sleep in this case because recieve blocks the code.
	// fmt.Println(finalSum)

	//case 5 ( we use channels for gorouteins sychronisation as well)
	// note : - we can do the same work using channels for which waitgroups use .
	// The value "true" doesn't matter here, we just use it as a signal that task() is done.
	// task() first prints "processing", then (because of defer) sends true on done as it finishes.
	// go task(done) starts task() in the background — main does not wait on this line.
	// <-done is the line that actually waits. Main pauses here until task() sends, so main only
	// moves forward once task() (and its print) has fully finished.
	//
	// How this is same as a WaitGroup:
	// A WaitGroup counter and this channel are doing the same job — making main wait for a goroutine
	// to finish instead of guessing with time.Sleep. With WaitGroup: wg.Add(1) before starting,
	// wg.Done() when the goroutine finishes, wg.Wait() blocks main until the count is 0.
	// With a channel: done := make(chan bool) is like the counter, task() sending on done is like
	// using defer, and <-done blocking main is like wg.Wait(). The difference is that a WaitGroup has a built-in counter, while a channel does not.
	//When we use channel or waitgrups ?
	// Ans : - If I need to wait for a single goroutine to finish, I can use a channel. If I need to wait for multiple goroutines to complete, I generally use a WaitGroup .
	// done := make(chan bool)
	// go task(done)
	// <-done //blocking

	//case 6 (Buffered channel)
	//Buffered channel :- A buffered channel has storage, so a send and receive can happen at different times. A send only blocks when the buffer is full, and a receive only blocks when the buffer is empty. The buffer size is set when the channel is created (make(chan int, 5) creates a channel that can hold 5 ints).
	//In unbuffered channels, we can send only when a receiver is ready. In buffered channels, we can send multiple values before a receiver is ready, up to the buffer's capacity. Once the buffer is full, sends block until some goroutine receives from the channel.
	//To implement like queue we can use buffered channel not unbuffered channel because unbuffered channel can send only when receiver is ready but buffered channel can send multiple values before a receiver is ready, up to the buffer's capacity.

	//Below is the example of buffered channel with capacity of 100 and we send two values in it and recieve them in main function. and we donot get deadlock error because we have buffered channel with capacity of 100 and we send only two values in it and recieve them in main function. we can send 100 values in it before it gets full and blocks the sender goroutine.
	// emailChain := make(chan string, 100)//buffered channel with capacity of 100
	// emailChain <- "10@example.com"
	// emailChain <- "11@example.com"
	// fmt.Print(<-emailChain)
	// fmt.Print(<-emailChain)

	//we have recieved lot of mail from database and we want to send mail to them in banch using queue and we can use buffered channel for that because we can send multiple values before a receiver is ready, up to the buffer's capacity. Once the buffer is full, sends block until some goroutine receives from the channel.
	//Using channel we can run workers in backgroud , process the data in background.

	// emailChain := make(chan string, 100)
	// done := make(chan bool)
	// //here we can call emailSender function after sending the emails to the channel because we use buffered channel and send is non blocking until the channel is full and we have sent only 100 values in it and channel has capacity of 100 so it will not block the sender goroutine and we can call emailSender function afters sending the emails to the channel.
	// go emailSender(emailChain, done)
	// for i := 0; i < 100; i++ {
	// 	email := fmt.Sprintf("%d@example.com", i) //we can use fmt.Sprintf to format the string and send it to the channel.
	// 	emailChain <- email
	// } // Start the email sender goroutine

	// fmt.Println("All emails sent to the channel. Waiting for sender to finish...") //as it print first because we have sent all the emails to the channel and now we are waiting for the sender to finish sending them.as send is non blocking because we have buffered channel with capacity of 100 and we have sent only 100 values in it.
	/*we get a deadlock without using close keyword :-
	1. Main sends 100 emails into the buffered channel, then hits <-done.
	2. emailSender's range loop receives all 100, then tries to receive a 101st. Nothing's there, and nothing more will ever be sent — but the loop doesn't know that, so it blocks waiting.
	3. Since the loop never exits, the defer func(){ done <- true }() never runs.
	4. Main is stuck at <-done forever, waiting for a signal that will never come.
	5. So range just blocks on the next receive, forever, unless something tells it "no more values are coming." That signal is close().
	 only the sender should close a channel, never the receiver.
	*/
	//close() works the same way on unbuffered and buffered channels. What matters is whether the receiver needs to know that the sender has finished

	/*
			//Where to close a channel ?
			Does the receiver use `range` or check `ok`?
		├── NO  → don't close (receiver doesn't need it)
		└── YES → Who knows when ALL sending is done?
		          ├── One sender → that sender closes, after its last send
		          └── Many senders → a separate closer goroutine + wg.Wait()
				  when we have multiple senders then we need to write wg.Wait() first otherwise close() on a channel while a sender is still sending causes a panic.
				  Just call close() after all sending is done. If you have multiple senders, use a WaitGroup to wait for them all to finish before closing the channel.
	*/

	// close(emailChain) // we Close the channel
	// <-done            // we use done for blocking because we have sent all the emails to the channel and now we are waiting for the sender to finish sending them.

	/*// case 7 //Below we are lishen on multiple channels using select statement and we are sending data to both channels in different goroutines and we are receiving data from both channels in main function using select statement.
	//select can handle both receive and send operations on channels.
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { //goroutine 1
		chan1 <- 5
	}()
	go func() { //goroutine 2
		chan2 <- "raja"
	}()

	//Below is the example of select statement which is used to wait on multiple channel operations. The select statement blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
	for i := 0; i < 4; i++ { //we use for loop to receive data from both channels because we have sent data to both channels in different goroutines and we want to receive data from both channels in main function using select statement.
		select {
		case chanVal1 := <-chan1:
			fmt.Println("channel1 Data", chanVal1)

		case chanVal2 := <-chan2:
			fmt.Println("channel2 Data", chanVal2)

		}
	}*/
	//New concept :- Below we make restriction as to which channel can be used for sending and which channel can be used for receiving. This is called directional channels. We can use directional channels to make restriction as to which channel can be used for sending and which channel can be used for receiving.
	//To achieve this we use syntax like this :- func emailSender(emailChain <-chan string, done chan<- bool) { //here we use <-chan string to make restriction as to which channel can be used for receiving and chan<- bool to make restriction as to which channel can be used for sending.
	//	func emailSender(emailChain <-chan string, done chan<- bool) {
	//		defer func() { //
	//			done <- true
	//		}()
	//		for email := range emailChain {
	//			fmt.Println("Sending email to", email)
	//			time.Sleep(time.Millisecond * 100) // Simulate time taken to send an email
	//		}
	//	}
}
