package main

import (
	"fmt"
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
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	//go SaySomething()
	fmt.Println("Hello world from goroutine")
	//time.Sleep(1 * time.Second) // sleep for 1 seconds to allow goroutine to finish

	//2.example
	go count("A")
	go count("B")
	count("C")

}

//sync.wait heta code krnna one , waitgroups b;nna one
