package main

import (
	"fmt"
	"sync"
)

type post struct {
	//we prefer to decleare mutex in struct where variable is present.
	views int
	mu    sync.Mutex // here we lock this variable and using this all the other goroutiens wait until first goroutien modify the variable. This is How mutex help to away from race condition.
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	//Always write mutex before the modification and below "p.mu.Lock()" line all code get lock .but its not good way to write like this beacuse let say we get any error in modification code then we never hit line "p.mu.Unlock()" then variable never unlock. so better way to move unlock in defer due to that unlock run everytime.
	p.mu.Lock()
	p.views = p.views + 1

	// p.views = p.views + 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{
		views: 0,
	}
	//  for i:=0;i<100;i++{//By using this views increases synchronously as the using for loop , as in real life views increases in concurrently.
	//      myPost.inc()
	//  }

	for i := 0; i < 100; i++ { //By using this views increases concurrently as the using goroutine , as in real life views increases in concurrently.
		//Through this loop when we run the program we can see that we are getting the different output some time because the goroutines are running concurrently and they are trying to access the same variable "views" at the same time. This is called race condition.
		//race condition is a situation where the output of a program depends on the sequence or timing of uncontrollable events like goroutines. In this case, the goroutines are trying to access the same variable "views" at the same time and they are not synchronized. So, we are getting different output some time.

		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("data", myPost.views)
}
