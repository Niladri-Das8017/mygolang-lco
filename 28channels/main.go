package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channeels in golang")

	myChan := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// myChan <- 5
	// fmt.Println(<-myChan)

	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myChan
		fmt.Println(isChanelOpen) //true/false
		fmt.Println(val)

		wg.Done()
	}(myChan, wg)

	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {

		//listning the closeed channels are allowed but not sending
		myChan <- 5
		myChan <- 6
		close(myChan)

		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
