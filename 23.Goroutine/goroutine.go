package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.simple go routine exmaple
func SaySomething() {
	fmt.Println("This will be printed after few moments")
}

// 2realworld example
func count(n string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(n, ": ", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

// 3.waitgropup exmaple
func SayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from say hello function which is a memeber og waitgroup")
	for i := 1; i <= 5; i++ {
		fmt.Println("value of i is :", i)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	//go SaySomething()
	fmt.Println("Hello world from goroutine")
	//time.Sleep(1 * time.Second) // sleep for 1 seconds to allow goroutine to finish

	//2.example
	// go count("A")
	// go count("B")
	// count("C")

	//3.example-waitgropu scenen ek
	var wg sync.WaitGroup

	wg.Add(1) //tell the gropu we are starting 1 goroutine
	go SayHello(&wg)
	wg.Wait() //wait for all go routines
	fmt.Println("Main function is done")

}
