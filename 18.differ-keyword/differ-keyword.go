package main

import (
	"fmt"
)

// 3. The arguments in a deferred function are evaluated immediately, but the function call is delayed.
func sumation(a int) int {
	x := 10
	fmt.Println("X value of the sumation function at first was:", x)
	x = 35
	fmt.Println("X value of the sumation function at last was:", x)

	fmt.Println("and your sum is :", a+x)
	return a + x
}

func main() {

	//1.simple example
	fmt.Println("starting fmt execustions")
	defer fmt.Println("This is a defer keyword usage , check this will printed lastly")
	fmt.Println("Middle of the fmt execution")

	//2.multiple defers , diferes act as LIFO
	func() {
		defer fmt.Println("This will printed lastly in the anonyous function execution")
		defer fmt.Println("This will printed secondly in the anonyous function execution")
		defer fmt.Println("This will printed first in the anonyous function execution")
	}()

	//3.
	fmt.Println(sumation(10))

}
