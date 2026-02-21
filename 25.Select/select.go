package main

import (
	"fmt"
	"time"
)

func main() {

	//1.normal example with channles created
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(20 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	go func() {
		time.Sleep(45 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println("Recieved from channel 1: ", msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println("Recieved from channel 2: ", msg2)
	// case <-time.After(300 * time.Second):
	// 	fmt.Println("Timeout after 3 seconds")
	// 	// default:
	// 	// 	fmt.Println("Just nothing, this part is actually wont necessary")
	// }

	//in this case it actively looking for ch1 and ch2 or any channel to give smething, this will wait forever since it doesnt have any stopping condition
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Got from ch1:", msg1)

		case msg2 := <-ch2:
			fmt.Println("got from ch2:", msg2)

		default:
			fmt.Println("Waiting fro somethin to happen.......")
			time.Sleep(time.Second)
		}
	}

}
