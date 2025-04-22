package main

import (
	"fmt"
	"sync"
)

// 1.exmaple normal channel := passing and value and recieiving an value
func SendData(ch chan<- int) {
	ch <- 20
}

// 2.example - buffered channels i created
func ChannelFunc1(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello form 1st channel function"
}

func ChannelFunc2(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello form 2nd channel function"
}

func ChannelFunc3(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello form 3rd channel function"
}

func main() {
	// ch := make(chan int)
	// go SendData(ch)

	// value := <-ch
	// fmt.Println("Value received from channel is:", value)

	//2nd example buffered channels

	ch := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	ChannelFunc1(ch, &wg)
	ChannelFunc2(ch, &wg)
	ChannelFunc3(ch, &wg)

	close(ch)

	//dan me exmaple eke hatiyata channelfunctions run wenne go routine walin nowena nisa mehema range function ekn display krgnna puluwan
	//namuth go rputine walin run unoth thawa go routine function ekk one wenama wg.wait ekatai close(ch) ekatai pahatha example ek wage:
	//and me function eka thiyennath one for-range ekat kalin

	// func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	for message := range ch {
		fmt.Println("Value received from channel is:", message)
	}
}
