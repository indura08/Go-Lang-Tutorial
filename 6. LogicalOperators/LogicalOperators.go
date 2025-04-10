package main

import "fmt"

func main() {

	var isManager bool = false
	var isFather bool = true
	if isManager {
		fmt.Println("He is a manager")
	} else if !isManager || isFather {
		fmt.Println("he is a father and a manager also")
	}

	fmt.Println("hello'")
}
